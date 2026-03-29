package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	PetMedicalHistoryCreateInput struct {
		UserID      int64
		PetID       int64
		HistoryType string
		Description string
		DiagnosedAt string
		IsCurrent   int
	}

	PetMedicalHistoryCreateOutput struct {
		RecordID int64
	}

	PetMedicalHistoryListInput struct {
		UserID   int64
		PetID    int64
		Page     int
		PageSize int
	}

	PetMedicalHistoryItem struct {
		ID          int64
		HistoryType string
		Description string
		DiagnosedAt string
		IsCurrent   int
		CreatedAt   string
	}

	PetMedicalHistoryListOutput struct {
		List     []PetMedicalHistoryItem
		Total    int
		Page     int
		PageSize int
	}

	PetVaccinationCreateInput struct {
		UserID          int64
		PetID           int64
		VaccineName     string
		VaccinationDate string
		NextDueDate     string
		HospitalName    string
		Remark          string
	}

	PetVaccinationCreateOutput struct {
		RecordID int64
	}

	PetVaccinationListInput struct {
		UserID   int64
		PetID    int64
		Page     int
		PageSize int
	}

	PetVaccinationItem struct {
		ID              int64
		VaccineName     string
		VaccinationDate string
		NextDueDate     string
		HospitalName    string
		Remark          string
		CreatedAt       string
	}

	PetVaccinationListOutput struct {
		List     []PetVaccinationItem
		Total    int
		Page     int
		PageSize int
	}

	PetAllergyCreateInput struct {
		UserID             int64
		PetID              int64
		Allergen           string
		SymptomDescription string
		SeverityLevel      int
		Remark             string
	}

	PetAllergyCreateOutput struct {
		RecordID int64
	}

	PetAllergyListInput struct {
		UserID   int64
		PetID    int64
		Page     int
		PageSize int
	}

	PetAllergyItem struct {
		ID                 int64
		Allergen           string
		SymptomDescription string
		SeverityLevel      int
		Remark             string
		CreatedAt          string
	}

	PetAllergyListOutput struct {
		List     []PetAllergyItem
		Total    int
		Page     int
		PageSize int
	}
)

type IPetHealth interface {
	CreateMedicalHistory(ctx context.Context, in PetMedicalHistoryCreateInput) (*PetMedicalHistoryCreateOutput, error)
	ListMedicalHistory(ctx context.Context, in PetMedicalHistoryListInput) (*PetMedicalHistoryListOutput, error)
	CreateVaccination(ctx context.Context, in PetVaccinationCreateInput) (*PetVaccinationCreateOutput, error)
	ListVaccinations(ctx context.Context, in PetVaccinationListInput) (*PetVaccinationListOutput, error)
	CreateAllergy(ctx context.Context, in PetAllergyCreateInput) (*PetAllergyCreateOutput, error)
	ListAllergies(ctx context.Context, in PetAllergyListInput) (*PetAllergyListOutput, error)
}

var PetHealth IPetHealth = petHealthService{}

type petHealthService struct{}

func (s petHealthService) CreateMedicalHistory(ctx context.Context, in PetMedicalHistoryCreateInput) (*PetMedicalHistoryCreateOutput, error) {
	if err := ensurePetAccessible(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}
	diagnosedAt, err := parseDate(in.DiagnosedAt)
	if err != nil {
		return nil, consts.NewBadRequestError("确诊日期格式不正确")
	}

	now := time.Now()
	result, err := dao.PetMedicalHistory.Ctx(ctx).Data(do.PetMedicalHistory{
		PetId:       in.PetID,
		HistoryType: in.HistoryType,
		Description: in.Description,
		DiagnosedAt: diagnosedAt,
		IsCurrent:   in.IsCurrent,
		CreatedAt:   now,
		UpdatedAt:   now,
	}).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增病史记录失败")
	}
	recordID, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取病史记录ID失败")
	}
	return &PetMedicalHistoryCreateOutput{RecordID: recordID}, nil
}

func (s petHealthService) ListMedicalHistory(ctx context.Context, in PetMedicalHistoryListInput) (*PetMedicalHistoryListOutput, error) {
	if err := ensurePetAccessible(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	cols := dao.PetMedicalHistory.Columns()
	total, err := dao.PetMedicalHistory.Ctx(ctx).
		Where(cols.PetId, in.PetID).
		Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病史列表失败")
	}

	offset := (in.Page - 1) * in.PageSize
	records, err := dao.PetMedicalHistory.Ctx(ctx).
		Where(cols.PetId, in.PetID).
		OrderDesc(cols.CreatedAt).
		Limit(offset, in.PageSize).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病史列表失败")
	}

	list := make([]PetMedicalHistoryItem, 0, len(records))
	for _, record := range records {
		list = append(list, PetMedicalHistoryItem{
			ID:          record[cols.Id].Int64(),
			HistoryType: record[cols.HistoryType].String(),
			Description: record[cols.Description].String(),
			DiagnosedAt: formatDate(record[cols.DiagnosedAt].Time()),
			IsCurrent:   record[cols.IsCurrent].Int(),
			CreatedAt:   formatDateTime(record[cols.CreatedAt].Time()),
		})
	}

	return &PetMedicalHistoryListOutput{
		List:     list,
		Total:    total,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}

func (s petHealthService) CreateVaccination(ctx context.Context, in PetVaccinationCreateInput) (*PetVaccinationCreateOutput, error) {
	if err := ensurePetAccessible(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}
	vaccinationDate, err := parseDate(in.VaccinationDate)
	if err != nil {
		return nil, consts.NewBadRequestError("接种日期格式不正确")
	}
	var nextDueDate time.Time
	if in.NextDueDate != "" {
		nextDueDate, err = parseDate(in.NextDueDate)
		if err != nil {
			return nil, consts.NewBadRequestError("下次接种日期格式不正确")
		}
	}

	now := time.Now()
	data := do.PetVaccinationRecord{
		PetId:           in.PetID,
		VaccineName:     in.VaccineName,
		VaccinationDate: vaccinationDate,
		HospitalName:    in.HospitalName,
		Remark:          in.Remark,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	if !nextDueDate.IsZero() {
		data.NextDueDate = nextDueDate
	}

	result, err := dao.PetVaccinationRecord.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增疫苗记录失败")
	}
	recordID, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取疫苗记录ID失败")
	}
	return &PetVaccinationCreateOutput{RecordID: recordID}, nil
}

func (s petHealthService) ListVaccinations(ctx context.Context, in PetVaccinationListInput) (*PetVaccinationListOutput, error) {
	if err := ensurePetAccessible(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	cols := dao.PetVaccinationRecord.Columns()
	total, err := dao.PetVaccinationRecord.Ctx(ctx).
		Where(cols.PetId, in.PetID).
		Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询疫苗记录列表失败")
	}

	offset := (in.Page - 1) * in.PageSize
	records, err := dao.PetVaccinationRecord.Ctx(ctx).
		Where(cols.PetId, in.PetID).
		OrderDesc(cols.CreatedAt).
		Limit(offset, in.PageSize).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询疫苗记录列表失败")
	}

	list := make([]PetVaccinationItem, 0, len(records))
	for _, record := range records {
		list = append(list, PetVaccinationItem{
			ID:              record[cols.Id].Int64(),
			VaccineName:     record[cols.VaccineName].String(),
			VaccinationDate: formatDate(record[cols.VaccinationDate].Time()),
			NextDueDate:     formatDate(record[cols.NextDueDate].Time()),
			HospitalName:    record[cols.HospitalName].String(),
			Remark:          record[cols.Remark].String(),
			CreatedAt:       formatDateTime(record[cols.CreatedAt].Time()),
		})
	}

	return &PetVaccinationListOutput{
		List:     list,
		Total:    total,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}

func (s petHealthService) CreateAllergy(ctx context.Context, in PetAllergyCreateInput) (*PetAllergyCreateOutput, error) {
	if err := ensurePetAccessible(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	now := time.Now()
	result, err := dao.PetAllergyRecord.Ctx(ctx).Data(do.PetAllergyRecord{
		PetId:              in.PetID,
		Allergen:           in.Allergen,
		SymptomDescription: in.SymptomDescription,
		SeverityLevel:      in.SeverityLevel,
		Remark:             in.Remark,
		CreatedAt:          now,
		UpdatedAt:          now,
	}).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增过敏记录失败")
	}
	recordID, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取过敏记录ID失败")
	}
	return &PetAllergyCreateOutput{RecordID: recordID}, nil
}

func (s petHealthService) ListAllergies(ctx context.Context, in PetAllergyListInput) (*PetAllergyListOutput, error) {
	if err := ensurePetAccessible(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	cols := dao.PetAllergyRecord.Columns()
	total, err := dao.PetAllergyRecord.Ctx(ctx).
		Where(cols.PetId, in.PetID).
		Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询过敏记录列表失败")
	}

	offset := (in.Page - 1) * in.PageSize
	records, err := dao.PetAllergyRecord.Ctx(ctx).
		Where(cols.PetId, in.PetID).
		OrderDesc(cols.CreatedAt).
		Limit(offset, in.PageSize).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询过敏记录列表失败")
	}

	list := make([]PetAllergyItem, 0, len(records))
	for _, record := range records {
		list = append(list, PetAllergyItem{
			ID:                 record[cols.Id].Int64(),
			Allergen:           record[cols.Allergen].String(),
			SymptomDescription: record[cols.SymptomDescription].String(),
			SeverityLevel:      record[cols.SeverityLevel].Int(),
			Remark:             record[cols.Remark].String(),
			CreatedAt:          formatDateTime(record[cols.CreatedAt].Time()),
		})
	}

	return &PetAllergyListOutput{
		List:     list,
		Total:    total,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}
