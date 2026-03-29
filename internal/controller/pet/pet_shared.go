package pet

import (
	"context"

	v1 "PetCare/api/auth/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type SharedController struct{}

func NewShared() *SharedController {
	return &SharedController{}
}

func (c *SharedController) ListMedicalHistory(ctx context.Context, req *v1.MedicalHistoryListReq) (res *v1.MedicalHistoryListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	userID := claims.UserID
	if claims.Role == consts.RoleDoctor {
		userID = 0
	}
	output, err := service.PetHealth.ListMedicalHistory(ctx, service.PetMedicalHistoryListInput{
		UserID:   userID,
		PetID:    req.PetID,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	list := make([]v1.MedicalHistoryItem, 0, len(output.List))
	for _, item := range output.List {
		list = append(list, v1.MedicalHistoryItem{
			ID:          item.ID,
			HistoryType: item.HistoryType,
			Description: item.Description,
			DiagnosedAt: item.DiagnosedAt,
			IsCurrent:   item.IsCurrent,
			CreatedAt:   item.CreatedAt,
		})
	}
	return &v1.MedicalHistoryListRes{
		List:     list,
		Total:    output.Total,
		Page:     output.Page,
		PageSize: output.PageSize,
	}, nil
}

func (c *SharedController) ListVaccinations(ctx context.Context, req *v1.VaccinationListReq) (res *v1.VaccinationListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	userID := claims.UserID
	if claims.Role == consts.RoleDoctor {
		userID = 0
	}
	output, err := service.PetHealth.ListVaccinations(ctx, service.PetVaccinationListInput{
		UserID:   userID,
		PetID:    req.PetID,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	list := make([]v1.VaccinationItem, 0, len(output.List))
	for _, item := range output.List {
		list = append(list, v1.VaccinationItem{
			ID:              item.ID,
			VaccineName:     item.VaccineName,
			VaccinationDate: item.VaccinationDate,
			NextDueDate:     item.NextDueDate,
			HospitalName:    item.HospitalName,
			Remark:          item.Remark,
			CreatedAt:       item.CreatedAt,
		})
	}
	return &v1.VaccinationListRes{
		List:     list,
		Total:    output.Total,
		Page:     output.Page,
		PageSize: output.PageSize,
	}, nil
}

func (c *SharedController) ListAllergies(ctx context.Context, req *v1.AllergyListReq) (res *v1.AllergyListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	userID := claims.UserID
	if claims.Role == consts.RoleDoctor {
		userID = 0
	}
	output, err := service.PetHealth.ListAllergies(ctx, service.PetAllergyListInput{
		UserID:   userID,
		PetID:    req.PetID,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	list := make([]v1.AllergyItem, 0, len(output.List))
	for _, item := range output.List {
		list = append(list, v1.AllergyItem{
			ID:                 item.ID,
			Allergen:           item.Allergen,
			SymptomDescription: item.SymptomDescription,
			SeverityLevel:      item.SeverityLevel,
			Remark:             item.Remark,
			CreatedAt:          item.CreatedAt,
		})
	}
	return &v1.AllergyListRes{
		List:     list,
		Total:    output.Total,
		Page:     output.Page,
		PageSize: output.PageSize,
	}, nil
}
