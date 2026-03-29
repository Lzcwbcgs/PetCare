package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	AllergyCreateInput struct {
		UserID             int64
		PetID              int64
		Allergen           string
		SymptomDescription *string
		SeverityLevel      *int
		Remark             *string
	}
	AllergyListInput struct {
		UserID        int64
		PetID         int64
		Page          int
		Size          int
		SeverityLevel *int
	}
	AllergyCreateOutput struct {
		ID int64
	}
	AllergyItem struct {
		ID                 int64
		PetID              int64
		Allergen           string
		SymptomDescription string
		SeverityLevel      int
		Remark             string
		CreatedAt          string
		UpdatedAt          string
	}
	AllergyListOutput struct {
		Items []AllergyItem
		Total int
		Page  int
		Size  int
	}
)

type IPetAllergy interface {
	Create(ctx context.Context, in AllergyCreateInput) (*AllergyCreateOutput, error)
	List(ctx context.Context, in AllergyListInput) (*AllergyListOutput, error)
}

var PetAllergy IPetAllergy = petAllergyService{}

type petAllergyService struct{}

func (s petAllergyService) Create(ctx context.Context, in AllergyCreateInput) (*AllergyCreateOutput, error) {
	if err := checkPetOwner(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	data := do.PetAllergyRecord{
		PetId:     in.PetID,
		Allergen:  in.Allergen,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if in.SymptomDescription != nil {
		data.SymptomDescription = *in.SymptomDescription
	}
	if in.SeverityLevel != nil {
		data.SeverityLevel = *in.SeverityLevel
	}
	if in.Remark != nil {
		data.Remark = *in.Remark
	}

	result, err := dao.PetAllergyRecord.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增过敏记录失败")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增过敏记录失败")
	}
	return &AllergyCreateOutput{ID: id}, nil
}

func (s petAllergyService) List(ctx context.Context, in AllergyListInput) (*AllergyListOutput, error) {
	if err := checkPetOwner(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	page, size := in.Page, in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	if size > 100 {
		size = 100
	}

	cols := dao.PetAllergyRecord.Columns()
	model := dao.PetAllergyRecord.Ctx(ctx).Where(cols.PetId, in.PetID)
	if in.SeverityLevel != nil {
		model = model.Where(cols.SeverityLevel, *in.SeverityLevel)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询过敏记录失败")
	}
	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询过敏记录失败")
	}

	items := make([]AllergyItem, 0, len(records))
	for _, r := range records {
		items = append(items, AllergyItem{
			ID:                 r[cols.Id].Int64(),
			PetID:              r[cols.PetId].Int64(),
			Allergen:           r[cols.Allergen].String(),
			SymptomDescription: r[cols.SymptomDescription].String(),
			SeverityLevel:      r[cols.SeverityLevel].Int(),
			Remark:             r[cols.Remark].String(),
			CreatedAt:          r[cols.CreatedAt].GTime().Format(petTimeLayout),
			UpdatedAt:			r[cols.UpdatedAt].GTime().Format(petTimeLayout),
		})
	}
	return &AllergyListOutput{Items: items, Total: total, Page: page, Size: size}, nil
}
