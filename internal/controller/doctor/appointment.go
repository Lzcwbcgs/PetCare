package doctor

import (
	"context"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type AppointmentController struct{}

func NewAppointment() *AppointmentController {
	return &AppointmentController{}
}

func (c *AppointmentController) List(ctx context.Context, req *v1.AppointmentListReq) (res *v1.AppointmentListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	var (
		page int
		size int
	)
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		size = *req.PageSize
	}

	output, err := service.DoctorAppointment.List(ctx, service.DoctorAppointmentListInput{
		DoctorID: claims.UserID,
		Page:     page,
		Size:     size,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.AppointmentListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.AppointmentListItem{
			ID:              item.ID,
			AppointmentNo:   item.AppointmentNo,
			PetID:           item.PetID,
			PetName:         item.PetName,
			UserID:          item.UserID,
			UserNickname:    item.UserNickname,
			AppointmentType: item.AppointmentType,
			AppointmentTime: item.AppointmentTime,
			Status:          item.Status,
		})
	}

	return &v1.AppointmentListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *AppointmentController) Detail(ctx context.Context, req *v1.AppointmentDetailReq) (res *v1.AppointmentDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.DoctorAppointment.Detail(ctx, service.DoctorAppointmentDetailInput{
		DoctorID:      claims.UserID,
		AppointmentID: req.AppointmentID,
	})
	if err != nil {
		return nil, err
	}

	medicalHistories := make([]v1.AppointmentDetailMedicalHistory, 0, len(output.MedicalHistories))
	for _, item := range output.MedicalHistories {
		medicalHistories = append(medicalHistories, v1.AppointmentDetailMedicalHistory{
			ID:          item.ID,
			HistoryType: item.HistoryType,
			Description: item.Description,
			DiagnosedAt: item.DiagnosedAt,
			IsCurrent:   item.IsCurrent,
		})
	}

	vaccinations := make([]v1.AppointmentDetailVaccination, 0, len(output.Vaccinations))
	for _, item := range output.Vaccinations {
		vaccinations = append(vaccinations, v1.AppointmentDetailVaccination{
			ID:              item.ID,
			VaccineName:     item.VaccineName,
			VaccinationDate: item.VaccinationDate,
			NextDueDate:     item.NextDueDate,
		})
	}

	allergies := make([]v1.AppointmentDetailAllergy, 0, len(output.Allergies))
	for _, item := range output.Allergies {
		allergies = append(allergies, v1.AppointmentDetailAllergy{
			ID:                 item.ID,
			Allergen:           item.Allergen,
			SymptomDescription: item.SymptomDescription,
			SeverityLevel:      item.SeverityLevel,
		})
	}

	return &v1.AppointmentDetailRes{
		Appointment: v1.AppointmentDetailAppointment{
			ID:                 output.Appointment.ID,
			AppointmentNo:      output.Appointment.AppointmentNo,
			AppointmentType:    output.Appointment.AppointmentType,
			SymptomDescription: output.Appointment.SymptomDescription,
			AppointmentTime:    output.Appointment.AppointmentTime,
			Status:             output.Appointment.Status,
		},
		Pet: v1.AppointmentDetailPet{
			ID:         output.Pet.ID,
			PetName:    output.Pet.PetName,
			PetType:    output.Pet.PetType,
			Gender:     output.Pet.Gender,
			Age:        output.Pet.Age,
			AgeUnit:    output.Pet.AgeUnit,
			Breed:      output.Pet.Breed,
			Weight:     output.Pet.Weight,
			Sterilized: output.Pet.Sterilized,
			Remark:     output.Pet.Remark,
		},
		MedicalHistories: medicalHistories,
		Vaccinations:     vaccinations,
		Allergies:        allergies,
	}, nil
}

func (c *AppointmentController) UpdateStatus(ctx context.Context, req *v1.AppointmentUpdateStatusReq) (res *v1.AppointmentUpdateStatusRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "状态更新成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = service.DoctorAppointment.UpdateStatus(ctx, service.DoctorAppointmentUpdateStatusInput{
		DoctorID:      claims.UserID,
		AppointmentID: req.AppointmentID,
		Status:        req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AppointmentUpdateStatusRes{}, nil
}
