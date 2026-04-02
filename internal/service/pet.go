package service

import (
	"context"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
)

const petTimeLayout = "2006-01-02 15:04:05"

type (
	PetListInput struct {
		UserID  int64
		Page    int
		Size    int
		PetName *string
		PetType *string
		Status  *int
	}

	PetDetailInput struct {
		UserID int64
		PetID  int64
	}

	PetCreateInput struct {
		UserID     int64
		PetName    string
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

	PetItem struct {
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
		UpdatedAt  string
	}

	PetListOutput struct {
		Items []PetItem
		Total int
		Page  int
		Size  int
	}

	PetCreateOutput struct {
		ID int64
	}
)

type IPet interface {
	List(ctx context.Context, in PetListInput) (*PetListOutput, error)
	Detail(ctx context.Context, in PetDetailInput) (*PetItem, error)
	Create(ctx context.Context, in PetCreateInput) (*PetCreateOutput, error)
	Update(ctx context.Context, in PetUpdateInput) error
	Delete(ctx context.Context, in PetDeleteInput) error
}

var Pet IPet = petService{}

type petService struct{}

func (s petService) List(ctx context.Context, in PetListInput) (*PetListOutput, error) {
	var (
		page = in.Page
		size = in.Size
		cols = dao.Pet.Columns()
	)

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	if size > 100 {
		size = 100
	}

	model := dao.Pet.Ctx(ctx).Where(cols.UserId, in.UserID)
	if in.Status != nil {
		model = model.Where(cols.Status, *in.Status)
	}
	if in.PetName != nil && strings.TrimSpace(*in.PetName) != "" {
		petName := strings.TrimSpace(*in.PetName)
		model = model.WhereLike(cols.PetName, "%"+petName+"%")
	}
	if in.PetType != nil && strings.TrimSpace(*in.PetType) != "" {
		petType := strings.TrimSpace(*in.PetType)
		model = model.Where(cols.PetType, petType)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物列表失败")
	}

	items := make([]PetItem, 0, len(records))
	for _, record := range records {
		items = append(items, petItemFromRecord(record))
	}

	return &PetListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s petService) Detail(ctx context.Context, in PetDetailInput) (*PetItem, error) {
	cols := dao.Pet.Columns()
	record, err := dao.Pet.Ctx(ctx).
		Where(cols.Id, in.PetID).
		Where(cols.UserId, in.UserID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物详情失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("宠物不存在")
	}
	item := petItemFromRecord(record)
	return &item, nil
}

func (s petService) Create(ctx context.Context, in PetCreateInput) (*PetCreateOutput, error) {
	data := do.Pet{
		UserId:    in.UserID,
		PetName:   in.PetName,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	result, err := dao.Pet.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增宠物档案失败")
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取宠物ID失败")
	}

	return &PetCreateOutput{ID: lastID}, nil
}

func (s petService) Update(ctx context.Context, in PetUpdateInput) error {
	if !hasPetUpdate(in) {
		return consts.NewBadRequestError("至少提供一个更新字段")
	}

	data := do.Pet{
		UpdatedAt: time.Now(),
	}
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

	cols := dao.Pet.Columns()
	result, err := dao.Pet.Ctx(ctx).
		Where(cols.Id, in.PetID).
		Where(cols.UserId, in.UserID).
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
		return consts.NewNotFoundError("宠物不存在")
	}
	return nil
}

func (s petService) Delete(ctx context.Context, in PetDeleteInput) error {
	cols := dao.Pet.Columns()
	result, err := dao.Pet.Ctx(ctx).
		Where(cols.Id, in.PetID).
		Where(cols.UserId, in.UserID).
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
		return consts.NewNotFoundError("宠物不存在")
	}
	return nil
}

func hasPetUpdate(in PetUpdateInput) bool {
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

func petItemFromRecord(record gdb.Record) PetItem {
	cols := dao.Pet.Columns()
	return PetItem{
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
		CreatedAt:  record[cols.CreatedAt].Time().Format(petTimeLayout),
		UpdatedAt:  record[cols.UpdatedAt].Time().Format(petTimeLayout),
	}
}
