package service

import (
	"context"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	MedicalHistoryCreateInput struct {
		UserID      int64
		PetID       int64
		HistoryType *string
		Description string
		DiagnosedAt *string
		IsCurrent   *int
	}
	MedicalHistoryListInput struct {
		UserID    int64
		PetID     int64
		Page      int
		Size      int
		IsCurrent *int
	}
	MedicalHistoryCreateOutput struct {
		ID int64
	}
	MedicalHistoryItem struct {
		ID          int64
		PetID       int64
		HistoryType string
		Description string
		DiagnosedAt string
		IsCurrent   int
		CreatedAt   string
		UpdatedAt   string
	}
	MedicalHistoryListOutput struct {
		Items []MedicalHistoryItem
		Total int
		Page  int
		Size  int
	}
)

type IPetMedicalHistory interface {
	Create(ctx context.Context, in MedicalHistoryCreateInput) (*MedicalHistoryCreateOutput, error)
	List(ctx context.Context, in MedicalHistoryListInput) (*MedicalHistoryListOutput, error)
}

var PetMedicalHistory IPetMedicalHistory = petMedicalHistoryService{}

type petMedicalHistoryService struct{}

func (s petMedicalHistoryService) Create(ctx context.Context, in MedicalHistoryCreateInput) (*MedicalHistoryCreateOutput, error) {
	if err := checkPetOwner(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	data := do.PetMedicalHistory{
		PetId:       in.PetID,
		Description: in.Description,
		IsCurrent:   0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if in.HistoryType != nil {
		data.HistoryType = *in.HistoryType
	}
	if in.IsCurrent != nil {
		data.IsCurrent = *in.IsCurrent
	}
	if in.DiagnosedAt != nil && strings.TrimSpace(*in.DiagnosedAt) != "" {
		t, err := time.ParseInLocation(petTimeLayout, *in.DiagnosedAt, time.Local)
		if err != nil {
			return nil, consts.NewBadRequestError("确诊时间格式不正确")
		}
		data.DiagnosedAt = t
	}

	result, err := dao.PetMedicalHistory.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增病史记录失败")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增病史记录失败")
	}
	return &MedicalHistoryCreateOutput{ID: id}, nil
}

func (s petMedicalHistoryService) List(ctx context.Context, in MedicalHistoryListInput) (*MedicalHistoryListOutput, error) {
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

	cols := dao.PetMedicalHistory.Columns()
	model := dao.PetMedicalHistory.Ctx(ctx).Where(cols.PetId, in.PetID)
	if in.IsCurrent != nil {
		model = model.Where(cols.IsCurrent, *in.IsCurrent)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病史列表失败")
	}
	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病史列表失败")
	}

	items := make([]MedicalHistoryItem, 0, len(records))
	for _, r := range records {
		diagnosedAt := ""
		if !r[cols.DiagnosedAt].IsNil() {
			diagnosedAt = r[cols.DiagnosedAt].GTime().Format(petTimeLayout)
		}
		items = append(items, MedicalHistoryItem{
			ID:          r[cols.Id].Int64(),
			PetID:       r[cols.PetId].Int64(),
			HistoryType: r[cols.HistoryType].String(),
			Description: r[cols.Description].String(),
			DiagnosedAt: diagnosedAt,
			IsCurrent:   r[cols.IsCurrent].Int(),
			CreatedAt:   r[cols.CreatedAt].GTime().Format(petTimeLayout),
			UpdatedAt:   r[cols.UpdatedAt].GTime().Format(petTimeLayout),
		})
	}
	return &MedicalHistoryListOutput{Items: items, Total: total, Page: page, Size: size}, nil
}

