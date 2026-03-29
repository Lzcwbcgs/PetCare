package pet

import (
	"context"

	v1 "PetCare/api/auth/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type UserController struct{}

func NewUser() *UserController {
	return &UserController{}
}

func (c *UserController) Create(ctx context.Context, req *v1.PetCreateReq) (res *v1.PetCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.Pet.Create(ctx, service.PetCreateInput{
		UserID:     claims.UserID,
		PetName:    req.PetName,
		PetType:    req.PetType,
		AvatarURL:  req.AvatarURL,
		Gender:     req.Gender,
		Age:        req.Age,
		AgeUnit:    req.AgeUnit,
		Breed:      req.Breed,
		Weight:     req.Weight,
		Sterilized: req.Sterilized,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PetCreateRes{PetID: output.PetID}, nil
}

func (c *UserController) List(ctx context.Context, req *v1.PetListReq) (res *v1.PetListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.Pet.List(ctx, service.PetListInput{
		UserID:   claims.UserID,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	list := make([]v1.PetListItem, 0, len(output.List))
	for _, item := range output.List {
		list = append(list, v1.PetListItem{
			ID:         item.ID,
			PetName:    item.PetName,
			PetType:    item.PetType,
			AvatarURL:  item.AvatarURL,
			Gender:     item.Gender,
			Age:        item.Age,
			AgeUnit:    item.AgeUnit,
			Breed:      item.Breed,
			Weight:     item.Weight,
			Sterilized: item.Sterilized,
			Remark:     item.Remark,
			CreatedAt:  item.CreatedAt,
		})
	}
	return &v1.PetListRes{
		List:     list,
		Total:    output.Total,
		Page:     output.Page,
		PageSize: output.PageSize,
	}, nil
}

func (c *UserController) Detail(ctx context.Context, req *v1.PetDetailReq) (res *v1.PetDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.Pet.Detail(ctx, service.PetDetailInput{
		UserID: claims.UserID,
		PetID:  req.PetID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PetDetailRes{
		ID:         output.ID,
		UserID:     output.UserID,
		PetName:    output.PetName,
		PetType:    output.PetType,
		AvatarURL:  output.AvatarURL,
		Gender:     output.Gender,
		Age:        output.Age,
		AgeUnit:    output.AgeUnit,
		Breed:      output.Breed,
		Weight:     output.Weight,
		Sterilized: output.Sterilized,
		Remark:     output.Remark,
		Status:     output.Status,
		CreatedAt:  output.CreatedAt,
	}, nil
}

func (c *UserController) Update(ctx context.Context, req *v1.PetUpdateReq) (res *v1.PetUpdateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "更新成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = service.Pet.Update(ctx, service.PetUpdateInput{
		UserID:     claims.UserID,
		PetID:      req.PetID,
		PetName:    req.PetName,
		PetType:    req.PetType,
		AvatarURL:  req.AvatarURL,
		Gender:     req.Gender,
		Age:        req.Age,
		AgeUnit:    req.AgeUnit,
		Breed:      req.Breed,
		Weight:     req.Weight,
		Sterilized: req.Sterilized,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PetUpdateRes{}, nil
}

func (c *UserController) Delete(ctx context.Context, req *v1.PetDeleteReq) (res *v1.PetDeleteRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "删除成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = service.Pet.Delete(ctx, service.PetDeleteInput{
		UserID: claims.UserID,
		PetID:  req.PetID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PetDeleteRes{}, nil
}

func (c *UserController) CreateMedicalHistory(ctx context.Context, req *v1.MedicalHistoryCreateReq) (res *v1.MedicalHistoryCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.PetHealth.CreateMedicalHistory(ctx, service.PetMedicalHistoryCreateInput{
		UserID:      claims.UserID,
		PetID:       req.PetID,
		HistoryType: req.HistoryType,
		Description: req.Description,
		DiagnosedAt: req.DiagnosedAt,
		IsCurrent:   req.IsCurrent,
	})
	if err != nil {
		return nil, err
	}
	return &v1.MedicalHistoryCreateRes{RecordID: output.RecordID}, nil
}

func (c *UserController) CreateVaccination(ctx context.Context, req *v1.VaccinationCreateReq) (res *v1.VaccinationCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.PetHealth.CreateVaccination(ctx, service.PetVaccinationCreateInput{
		UserID:          claims.UserID,
		PetID:           req.PetID,
		VaccineName:     req.VaccineName,
		VaccinationDate: req.VaccinationDate,
		NextDueDate:     req.NextDueDate,
		HospitalName:    req.HospitalName,
		Remark:          req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.VaccinationCreateRes{RecordID: output.RecordID}, nil
}

func (c *UserController) CreateAllergy(ctx context.Context, req *v1.AllergyCreateReq) (res *v1.AllergyCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.PetHealth.CreateAllergy(ctx, service.PetAllergyCreateInput{
		UserID:             claims.UserID,
		PetID:              req.PetID,
		Allergen:           req.Allergen,
		SymptomDescription: req.SymptomDescription,
		SeverityLevel:      req.SeverityLevel,
		Remark:             req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AllergyCreateRes{RecordID: output.RecordID}, nil
}
