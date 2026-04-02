package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	DoctorMedicalRecordCreateInput struct {
		DoctorID             int64
		AppointmentID        int64
		PetID                int64
		UserID               int64
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

	DoctorMedicalRecordCreateOutput struct {
		ID int64
	}

	DoctorMedicalRecordUpdateInput struct {
		DoctorID             int64
		MedicalRecordID      int64
		ChiefComplaint       *string
		PresentHistory       *string
		PhysicalExamination  *string
		PreliminaryDiagnosis *string
		TreatmentPlan        *string
		Prescription         *string
		DoctorAdvice         *string
		VisitTime            *string
		Status               *int
	}

	DoctorMedicalRecordDetailInput struct {
		DoctorID        int64
		MedicalRecordID int64
	}

	DoctorMedicalRecordItem struct {
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
)

type IDoctorMedicalRecord interface {
	Create(ctx context.Context, in DoctorMedicalRecordCreateInput) (*DoctorMedicalRecordCreateOutput, error)
	Update(ctx context.Context, in DoctorMedicalRecordUpdateInput) error
	Detail(ctx context.Context, in DoctorMedicalRecordDetailInput) (*DoctorMedicalRecordItem, error)
}

var DoctorMedicalRecord IDoctorMedicalRecord = doctorMedicalRecordService{}

type doctorMedicalRecordService struct{}

func (s doctorMedicalRecordService) Create(ctx context.Context, in DoctorMedicalRecordCreateInput) (*DoctorMedicalRecordCreateOutput, error) {
	appointmentRecord, err := dao.Appointment.Ctx(ctx).
		Where(dao.Appointment.Columns().Id, in.AppointmentID).
		Where(dao.Appointment.Columns().DoctorId, in.DoctorID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询预约信息失败")
	}
	if appointmentRecord.IsEmpty() {
		return nil, consts.NewNotFoundError("预约不存在")
	}
	if appointmentRecord[dao.Appointment.Columns().PetId].Int64() != in.PetID {
		return nil, consts.NewConflictError("预约宠物信息不匹配")
	}
	if appointmentRecord[dao.Appointment.Columns().UserId].Int64() != in.UserID {
		return nil, consts.NewConflictError("预约用户信息不匹配")
	}

	appointmentStatus := appointmentRecord[dao.Appointment.Columns().Status].Int()
	if appointmentStatus == 3 {
		return nil, consts.NewConflictError("预约已取消，不能创建病历")
	}
	if appointmentStatus == 4 {
		return nil, consts.NewConflictError("预约已过期，不能创建病历")
	}

	exists, err := dao.MedicalRecord.Ctx(ctx).
		Where(dao.MedicalRecord.Columns().AppointmentId, in.AppointmentID).
		Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病历信息失败")
	}
	if exists > 0 {
		return nil, consts.NewConflictError("该预约已创建病历")
	}

	visitTime, err := time.ParseInLocation(appointmentTimeLayout, in.VisitTime, time.Local)
	if err != nil {
		return nil, consts.NewBadRequestError("就诊时间格式不正确")
	}

	now := time.Now()
	result, err := dao.MedicalRecord.Ctx(ctx).Data(do.MedicalRecord{
		AppointmentId:        in.AppointmentID,
		PetId:                in.PetID,
		UserId:               in.UserID,
		DoctorId:             in.DoctorID,
		ChiefComplaint:       in.ChiefComplaint,
		PresentHistory:       in.PresentHistory,
		PhysicalExamination:  in.PhysicalExamination,
		PreliminaryDiagnosis: in.PreliminaryDiagnosis,
		TreatmentPlan:        in.TreatmentPlan,
		Prescription:         in.Prescription,
		DoctorAdvice:         in.DoctorAdvice,
		VisitTime:            visitTime,
		Status:               in.Status,
		CreatedAt:            now,
		UpdatedAt:            now,
	}).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "创建病历失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取病历ID失败")
	}
	return &DoctorMedicalRecordCreateOutput{ID: id}, nil
}

func (s doctorMedicalRecordService) Update(ctx context.Context, in DoctorMedicalRecordUpdateInput) error {
	if !hasDoctorMedicalRecordUpdates(in) {
		return consts.NewBadRequestError("至少提供一个更新字段")
	}

	record, err := dao.MedicalRecord.Ctx(ctx).
		Where(dao.MedicalRecord.Columns().Id, in.MedicalRecordID).
		Where(dao.MedicalRecord.Columns().DoctorId, in.DoctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询病历信息失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("病历不存在")
	}

	data := do.MedicalRecord{
		UpdatedAt: time.Now(),
	}
	if in.ChiefComplaint != nil {
		data.ChiefComplaint = *in.ChiefComplaint
	}
	if in.PresentHistory != nil {
		data.PresentHistory = *in.PresentHistory
	}
	if in.PhysicalExamination != nil {
		data.PhysicalExamination = *in.PhysicalExamination
	}
	if in.PreliminaryDiagnosis != nil {
		data.PreliminaryDiagnosis = *in.PreliminaryDiagnosis
	}
	if in.TreatmentPlan != nil {
		data.TreatmentPlan = *in.TreatmentPlan
	}
	if in.Prescription != nil {
		data.Prescription = *in.Prescription
	}
	if in.DoctorAdvice != nil {
		data.DoctorAdvice = *in.DoctorAdvice
	}
	if in.VisitTime != nil {
		visitTime, err := time.ParseInLocation(appointmentTimeLayout, *in.VisitTime, time.Local)
		if err != nil {
			return consts.NewBadRequestError("就诊时间格式不正确")
		}
		data.VisitTime = visitTime
	}
	if in.Status != nil {
		data.Status = *in.Status
	}

	result, err := dao.MedicalRecord.Ctx(ctx).
		Where(dao.MedicalRecord.Columns().Id, in.MedicalRecordID).
		Where(dao.MedicalRecord.Columns().DoctorId, in.DoctorID).
		Data(data).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "更新病历失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("病历不存在")
	}
	return nil
}

func (s doctorMedicalRecordService) Detail(ctx context.Context, in DoctorMedicalRecordDetailInput) (*DoctorMedicalRecordItem, error) {
	record, err := dao.MedicalRecord.Ctx(ctx).
		Where(dao.MedicalRecord.Columns().Id, in.MedicalRecordID).
		Where(dao.MedicalRecord.Columns().DoctorId, in.DoctorID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病历详情失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("病历不存在")
	}

	return &DoctorMedicalRecordItem{
		ID:                   record[dao.MedicalRecord.Columns().Id].Int64(),
		AppointmentID:        record[dao.MedicalRecord.Columns().AppointmentId].Int64(),
		PetID:                record[dao.MedicalRecord.Columns().PetId].Int64(),
		UserID:               record[dao.MedicalRecord.Columns().UserId].Int64(),
		DoctorID:             record[dao.MedicalRecord.Columns().DoctorId].Int64(),
		ChiefComplaint:       record[dao.MedicalRecord.Columns().ChiefComplaint].String(),
		PresentHistory:       record[dao.MedicalRecord.Columns().PresentHistory].String(),
		PhysicalExamination:  record[dao.MedicalRecord.Columns().PhysicalExamination].String(),
		PreliminaryDiagnosis: record[dao.MedicalRecord.Columns().PreliminaryDiagnosis].String(),
		TreatmentPlan:        record[dao.MedicalRecord.Columns().TreatmentPlan].String(),
		Prescription:         record[dao.MedicalRecord.Columns().Prescription].String(),
		DoctorAdvice:         record[dao.MedicalRecord.Columns().DoctorAdvice].String(),
		VisitTime:            record[dao.MedicalRecord.Columns().VisitTime].GTime().Format(appointmentTimeLayout),
		Status:               record[dao.MedicalRecord.Columns().Status].Int(),
	}, nil
}

func hasDoctorMedicalRecordUpdates(in DoctorMedicalRecordUpdateInput) bool {
	return in.ChiefComplaint != nil ||
		in.PresentHistory != nil ||
		in.PhysicalExamination != nil ||
		in.PreliminaryDiagnosis != nil ||
		in.TreatmentPlan != nil ||
		in.Prescription != nil ||
		in.DoctorAdvice != nil ||
		in.VisitTime != nil ||
		in.Status != nil
}
