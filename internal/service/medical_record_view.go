package service

import (
	"context"

	"PetCare/internal/consts"
	"PetCare/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	MedicalRecordViewListInput struct {
		UserID int64
		Page   int
		Size   int
		PetID  *int64
	}

	MedicalRecordViewListItem struct {
		ID                   int64
		AppointmentID        int64
		PetID                int64
		DoctorID             int64
		DoctorName           string
		PreliminaryDiagnosis string
		VisitTime            string
		Status               int
	}

	MedicalRecordViewListOutput struct {
		Items []MedicalRecordViewListItem
		Total int
		Page  int
		Size  int
	}

	MedicalRecordViewDetailInput struct {
		RequesterUserID int64
		RequesterRole   string
		MedicalRecordID int64
	}

	MedicalRecordViewDetailItem struct {
		ID                   int64
		AppointmentID        int64
		PetID                int64
		UserID               int64
		DoctorID             int64
		ChiefComplaint       string
		PresentHistory       string
		PhysicalExamination  string
		PreliminaryDiagnosis string
		TreatmentPlan        string
		Prescription         string
		DoctorAdvice         string
		VisitTime            string
		Status               int
	}

	MedicalRecordViewReportListInput struct {
		RequesterUserID int64
		RequesterRole   string
		MedicalRecordID int64
	}
)

type IMedicalRecordView interface {
	List(ctx context.Context, in MedicalRecordViewListInput) (*MedicalRecordViewListOutput, error)
	Detail(ctx context.Context, in MedicalRecordViewDetailInput) (*MedicalRecordViewDetailItem, error)
	ReportList(ctx context.Context, in MedicalRecordViewReportListInput) ([]DoctorMedicalReportItem, error)
}

var MedicalRecordView IMedicalRecordView = medicalRecordViewService{}

type medicalRecordViewService struct{}

func (s medicalRecordViewService) List(ctx context.Context, in MedicalRecordViewListInput) (*MedicalRecordViewListOutput, error) {
	var (
		page = in.Page
		size = in.Size
		cols = dao.MedicalRecord.Columns()
	)

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	model := dao.MedicalRecord.Ctx(ctx).Where(cols.UserId, in.UserID)
	if in.PetID != nil {
		if err := checkPetOwner(ctx, *in.PetID, in.UserID); err != nil {
			return nil, err
		}
		model = model.Where(cols.PetId, *in.PetID)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病历列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病历列表失败")
	}

	doctorNameMap, err := loadDoctorNameMap(ctx, collectMedicalRecordDoctorIDs(records))
	if err != nil {
		return nil, err
	}

	items := make([]MedicalRecordViewListItem, 0, len(records))
	for _, record := range records {
		items = append(items, MedicalRecordViewListItem{
			ID:                   record[cols.Id].Int64(),
			AppointmentID:        record[cols.AppointmentId].Int64(),
			PetID:                record[cols.PetId].Int64(),
			DoctorID:             record[cols.DoctorId].Int64(),
			DoctorName:           doctorNameMap[record[cols.DoctorId].Int64()],
			PreliminaryDiagnosis: record[cols.PreliminaryDiagnosis].String(),
			VisitTime:            record[cols.VisitTime].GTime().Format(appointmentTimeLayout),
			Status:               record[cols.Status].Int(),
		})
	}

	return &MedicalRecordViewListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s medicalRecordViewService) Detail(ctx context.Context, in MedicalRecordViewDetailInput) (*MedicalRecordViewDetailItem, error) {
	record, err := getReadableMedicalRecord(ctx, in.RequesterUserID, in.RequesterRole, in.MedicalRecordID)
	if err != nil {
		return nil, err
	}

	cols := dao.MedicalRecord.Columns()
	return &MedicalRecordViewDetailItem{
		ID:                   record[cols.Id].Int64(),
		AppointmentID:        record[cols.AppointmentId].Int64(),
		PetID:                record[cols.PetId].Int64(),
		UserID:               record[cols.UserId].Int64(),
		DoctorID:             record[cols.DoctorId].Int64(),
		ChiefComplaint:       record[cols.ChiefComplaint].String(),
		PresentHistory:       record[cols.PresentHistory].String(),
		PhysicalExamination:  record[cols.PhysicalExamination].String(),
		PreliminaryDiagnosis: record[cols.PreliminaryDiagnosis].String(),
		TreatmentPlan:        record[cols.TreatmentPlan].String(),
		Prescription:         record[cols.Prescription].String(),
		DoctorAdvice:         record[cols.DoctorAdvice].String(),
		VisitTime:            record[cols.VisitTime].GTime().Format(appointmentTimeLayout),
		Status:               record[cols.Status].Int(),
	}, nil
}

func (s medicalRecordViewService) ReportList(ctx context.Context, in MedicalRecordViewReportListInput) ([]DoctorMedicalReportItem, error) {
	if _, err := getReadableMedicalRecord(ctx, in.RequesterUserID, in.RequesterRole, in.MedicalRecordID); err != nil {
		return nil, err
	}

	records, err := dao.MedicalReport.Ctx(ctx).
		Where(dao.MedicalReport.Columns().MedicalRecordId, in.MedicalRecordID).
		OrderDesc(dao.MedicalReport.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病历报告列表失败")
	}

	items := make([]DoctorMedicalReportItem, 0, len(records))
	for _, record := range records {
		items = append(items, DoctorMedicalReportItem{
			ID:            record[dao.MedicalReport.Columns().Id].Int64(),
			ReportTitle:   record[dao.MedicalReport.Columns().ReportTitle].String(),
			ReportType:    record[dao.MedicalReport.Columns().ReportType].String(),
			FileURL:       record[dao.MedicalReport.Columns().FileUrl].String(),
			ReportContent: record[dao.MedicalReport.Columns().ReportContent].String(),
			UploadedAt:    record[dao.MedicalReport.Columns().UploadedAt].GTime().Format(appointmentTimeLayout),
		})
	}
	return items, nil
}

func getReadableMedicalRecord(ctx context.Context, requesterUserID int64, requesterRole string, medicalRecordID int64) (gdb.Record, error) {
	record, err := dao.MedicalRecord.Ctx(ctx).
		Where(dao.MedicalRecord.Columns().Id, medicalRecordID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病历信息失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("病历不存在")
	}

	switch NormalizeRole(requesterRole) {
	case consts.RoleUser:
		if record[dao.MedicalRecord.Columns().UserId].Int64() != requesterUserID {
			return nil, consts.NewForbiddenError("")
		}
	case consts.RoleDoctor:
		if record[dao.MedicalRecord.Columns().DoctorId].Int64() != requesterUserID {
			return nil, consts.NewForbiddenError("")
		}
	default:
		return nil, consts.NewForbiddenError("")
	}

	return record, nil
}

func collectMedicalRecordDoctorIDs(records gdb.Result) []int64 {
	return collectUniqueInt64(records, func(record gdb.Record) int64 {
		return record[dao.MedicalRecord.Columns().DoctorId].Int64()
	})
}
