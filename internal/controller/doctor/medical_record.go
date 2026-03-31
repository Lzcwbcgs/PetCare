package doctor

import (
	"context"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type MedicalRecordController struct{}

func NewMedicalRecord() *MedicalRecordController {
	return &MedicalRecordController{}
}

func (c *MedicalRecordController) Create(ctx context.Context, req *v1.MedicalRecordCreateReq) (res *v1.MedicalRecordCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "病历创建成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.DoctorMedicalRecord.Create(ctx, service.DoctorMedicalRecordCreateInput{
		DoctorID:             claims.UserID,
		AppointmentID:        req.AppointmentID,
		PetID:                req.PetID,
		UserID:               req.UserID,
		ChiefComplaint:       req.ChiefComplaint,
		PresentHistory:       req.PresentHistory,
		PhysicalExamination:  req.PhysicalExamination,
		PreliminaryDiagnosis: req.PreliminaryDiagnosis,
		TreatmentPlan:        req.TreatmentPlan,
		Prescription:         req.Prescription,
		DoctorAdvice:         req.DoctorAdvice,
		VisitTime:            req.VisitTime,
		Status:               req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.MedicalRecordCreateRes{MedicalRecordID: output.ID}, nil
}

func (c *MedicalRecordController) Update(ctx context.Context, req *v1.MedicalRecordUpdateReq) (res *v1.MedicalRecordUpdateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "病历更新成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = service.DoctorMedicalRecord.Update(ctx, service.DoctorMedicalRecordUpdateInput{
		DoctorID:             claims.UserID,
		MedicalRecordID:      req.MedicalRecordID,
		ChiefComplaint:       req.ChiefComplaint,
		PresentHistory:       req.PresentHistory,
		PhysicalExamination:  req.PhysicalExamination,
		PreliminaryDiagnosis: req.PreliminaryDiagnosis,
		TreatmentPlan:        req.TreatmentPlan,
		Prescription:         req.Prescription,
		DoctorAdvice:         req.DoctorAdvice,
		VisitTime:            req.VisitTime,
		Status:               req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.MedicalRecordUpdateRes{}, nil
}

func (c *MedicalRecordController) Detail(ctx context.Context, req *v1.MedicalRecordDetailReq) (res *v1.MedicalRecordDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	item, err := service.DoctorMedicalRecord.Detail(ctx, service.DoctorMedicalRecordDetailInput{
		DoctorID:        claims.UserID,
		MedicalRecordID: req.MedicalRecordID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.MedicalRecordDetailRes{
		ID:                   item.ID,
		AppointmentID:        item.AppointmentID,
		PetID:                item.PetID,
		UserID:               item.UserID,
		DoctorID:             item.DoctorID,
		ChiefComplaint:       item.ChiefComplaint,
		PresentHistory:       item.PresentHistory,
		PhysicalExamination:  item.PhysicalExamination,
		PreliminaryDiagnosis: item.PreliminaryDiagnosis,
		TreatmentPlan:        item.TreatmentPlan,
		Prescription:         item.Prescription,
		DoctorAdvice:         item.DoctorAdvice,
		VisitTime:            item.VisitTime,
		Status:               item.Status,
	}, nil
}
