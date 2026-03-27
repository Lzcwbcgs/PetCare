package auth

import (
	"context"

	v1 "PetCare/api/auth/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type DoctorController struct{}

func NewDoctor() *DoctorController {
	return &DoctorController{}
}

func (c *DoctorController) Profile(ctx context.Context, req *v1.DoctorProfileReq) (res *v1.DoctorProfileRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.DoctorProfile.GetProfile(ctx, service.DoctorProfileGetInput{
		DoctorID: claims.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DoctorProfileRes{
		ID:           output.ID,
		HospitalID:   output.HospitalID,
		HospitalName: output.HospitalName,
		Username:     output.Username,
		DoctorName:   output.DoctorName,
		Gender:       output.Gender,
		Phone:        output.Phone,
		Email:        output.Email,
		Title:        output.Title,
		Specialty:    output.Specialty,
		AvatarURL:    output.AvatarURL,
		Intro:        output.Intro,
		Status:       output.Status,
	}, nil
}

func (c *DoctorController) UpdateProfile(ctx context.Context, req *v1.DoctorUpdateProfileReq) (res *v1.DoctorUpdateProfileRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "更新成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = service.DoctorProfile.UpdateProfile(ctx, service.DoctorProfileUpdateInput{
		DoctorID:  claims.UserID,
		Phone:     req.Phone,
		Email:     req.Email,
		AvatarURL: req.AvatarURL,
		Intro:     req.Intro,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DoctorUpdateProfileRes{}, nil
}

func (c *DoctorController) UpdatePassword(ctx context.Context, req *v1.DoctorUpdatePasswordReq) (res *v1.DoctorUpdatePasswordRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "密码修改成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = service.DoctorProfile.UpdatePassword(ctx, service.DoctorProfileUpdatePasswordInput{
		DoctorID:    claims.UserID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DoctorUpdatePasswordRes{}, nil
}
