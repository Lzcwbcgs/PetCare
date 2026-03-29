package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	PetCreateInput struct {
		UserID     int64
		PetName    string
		PetType    string
		AvatarURL  string
		Gender     int
		Age        int
		AgeUnit    *string
		Breed      string
		Weight     float64
		Sterilized int
		Remark     string
	}

	PetCreateOutput struct {
		PetID int64
	}

	PetListInput struct {
		UserID   int64
		Page     int
		PageSize int
	}

	PetListItem struct {
		ID         int64
		PetName    string
		PetType    string
		AvatarURL  string
		Gender     int
		Age        int
		AgeUnit    string
		Breed      string
		Weight     float64
		Sterilized int
		Remark     string
		CreatedAt  string
	}

	PetListOutput struct {
		List     []PetListItem
		Total    int
		Page     int
		PageSize int
	}

	PetDetailInput struct {
		UserID int64
		PetID  int64
	}

	PetDetailOutput struct {
		ID         int64
		UserID     int64
		PetName    string
		PetType    string
		AvatarURL  string
		Gender     int
		Age        int
		AgeUnit    string
		Breed      string
		Weight     float64
		Sterilized int
		Remark     string
		Status     int
		CreatedAt  string
	}

	PetUpdateInput struct {
		UserID     int64
		PetID      int64
		PetName    *string
		PetType    *string
		AvatarURL  *string
		Gender     *int
		Age        *int
		AgeUnit    *string
		Breed      *string
		Weight     *float64
		Sterilized *int
		Remark     *string
	}

	PetDeleteInput struct {
		UserID int64
		PetID  int64
	}
)

type IPet interface {
	Create(ctx context.Context, in PetCreateInput) (*PetCreateOutput, error)
	List(ctx context.Context, in PetListInput) (*PetListOutput, error)
	Detail(ctx context.Context, in PetDetailInput) (*PetDetailOutput, error)
	Update(ctx context.Context, in PetUpdateInput) error
	Delete(ctx context.Context, in PetDeleteInput) error
}

var Pet IPet = petService{}

type petService struct{}

func (s petService) Create(ctx context.Context, in PetCreateInput) (*PetCreateOutput, error) {
	now := time.Now()
	ageUnit := ""
	if in.AgeUnit != nil {
		ageUnit = *in.AgeUnit
	}
	result, err := dao.Pet.Ctx(ctx).Data(do.Pet{
		UserId:     in.UserID,
		PetName:    in.PetName,
		PetType:    in.PetType,
		AvatarUrl:  in.AvatarURL,
		Gender:     in.Gender,
		Age:        in.Age,
		AgeUnit:    ageUnit,
		Breed:      in.Breed,
		Weight:     in.Weight,
		Sterilized: in.Sterilized,
		Remark:     in.Remark,
		Status:     1,
		CreatedAt:  now,
		UpdatedAt:  now,
	}).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增宠物档案失败")
	}
	petID, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取宠物ID失败")
	}
	return &PetCreateOutput{PetID: petID}, nil
}

func (s petService) List(ctx context.Context, in PetListInput) (*PetListOutput, error) {
	cols := dao.Pet.Columns()
	total, err := dao.Pet.Ctx(ctx).
		Where(cols.UserId, in.UserID).
		Where(cols.Status, 1).
		Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物列表失败")
	}

	offset := (in.Page - 1) * in.PageSize
	records, err := dao.Pet.Ctx(ctx).
		Where(cols.UserId, in.UserID).
		Where(cols.Status, 1).
		OrderDesc(cols.CreatedAt).
		Limit(offset, in.PageSize).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物列表失败")
	}

	list := make([]PetListItem, 0, len(records))
	for _, record := range records {
		list = append(list, PetListItem{
			ID:         record[cols.Id].Int64(),
			PetName:    record[cols.PetName].String(),
			PetType:    record[cols.PetType].String(),
			AvatarURL:  record[cols.AvatarUrl].String(),
			Gender:     record[cols.Gender].Int(),
			Age:        record[cols.Age].Int(),
			AgeUnit:    record[cols.AgeUnit].String(),
			Breed:      record[cols.Breed].String(),
			Weight:     record[cols.Weight].Float64(),
			Sterilized: record[cols.Sterilized].Int(),
			Remark:     record[cols.Remark].String(),
			CreatedAt:  formatDateTime(record[cols.CreatedAt].Time()),
		})
	}

	return &PetListOutput{
		List:     list,
		Total:    total,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}

func (s petService) Detail(ctx context.Context, in PetDetailInput) (*PetDetailOutput, error) {
	cols := dao.Pet.Columns()
	record, err := dao.Pet.Ctx(ctx).
		Where(cols.Id, in.PetID).
		Where(cols.UserId, in.UserID).
		Where(cols.Status, 1).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物详情失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("宠物档案不存在")
	}

	return &PetDetailOutput{
		ID:         record[cols.Id].Int64(),
		UserID:     record[cols.UserId].Int64(),
		PetName:    record[cols.PetName].String(),
		PetType:    record[cols.PetType].String(),
		AvatarURL:  record[cols.AvatarUrl].String(),
		Gender:     record[cols.Gender].Int(),
		Age:        record[cols.Age].Int(),
		AgeUnit:    record[cols.AgeUnit].String(),
		Breed:      record[cols.Breed].String(),
		Weight:     record[cols.Weight].Float64(),
		Sterilized: record[cols.Sterilized].Int(),
		Remark:     record[cols.Remark].String(),
		Status:     record[cols.Status].Int(),
		CreatedAt:  formatDateTime(record[cols.CreatedAt].Time()),
	}, nil
}

func (s petService) Update(ctx context.Context, in PetUpdateInput) error {
	if !hasPetUpdates(in) {
		return consts.NewBadRequestError("至少提供一个更新字段")
	}

	data := do.Pet{}
	if in.PetName != nil {
		data.PetName = *in.PetName
	}
	if in.PetType != nil {
		data.PetType = *in.PetType
	}
	if in.AvatarURL != nil {
		data.AvatarUrl = *in.AvatarURL
	}
	if in.Gender != nil {
		data.Gender = *in.Gender
	}
	if in.Age != nil {
		data.Age = *in.Age
	}
	if in.AgeUnit != nil {
		data.AgeUnit = *in.AgeUnit
	}
	if in.Breed != nil {
		data.Breed = *in.Breed
	}
	if in.Weight != nil {
		data.Weight = *in.Weight
	}
	if in.Sterilized != nil {
		data.Sterilized = *in.Sterilized
	}
	if in.Remark != nil {
		data.Remark = *in.Remark
	}
	data.UpdatedAt = time.Now()

	cols := dao.Pet.Columns()
	result, err := dao.Pet.Ctx(ctx).
		Where(cols.Id, in.PetID).
		Where(cols.UserId, in.UserID).
		Where(cols.Status, 1).
		Data(data).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "更新宠物档案失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("宠物档案不存在")
	}
	return nil
}

func (s petService) Delete(ctx context.Context, in PetDeleteInput) error {
	cols := dao.Pet.Columns()
	result, err := dao.Pet.Ctx(ctx).
		Where(cols.Id, in.PetID).
		Where(cols.UserId, in.UserID).
		Where(cols.Status, 1).
		Data(do.Pet{
			Status:    0,
			UpdatedAt: time.Now(),
		}).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "删除宠物档案失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取删除结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("宠物档案不存在")
	}
	return nil
}

func hasPetUpdates(in PetUpdateInput) bool {
	return in.PetName != nil ||
		in.PetType != nil ||
		in.AvatarURL != nil ||
		in.Gender != nil ||
		in.Age != nil ||
		in.AgeUnit != nil ||
		in.Breed != nil ||
		in.Weight != nil ||
		in.Sterilized != nil ||
		in.Remark != nil
}
