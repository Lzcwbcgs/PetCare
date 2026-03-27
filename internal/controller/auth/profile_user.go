package auth

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

func (c *UserController) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.UserProfile.GetProfile(ctx, service.UserProfileGetInput{
		UserID: claims.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserProfileRes{
		ID:        output.ID,
		Username:  output.Username,
		Nickname:  output.Nickname,
		Phone:     output.Phone,
		Email:     output.Email,
		AvatarURL: output.AvatarURL,
		Status:    output.Status,
		CreatedAt: output.CreatedAt,
	}, nil
}

func (c *UserController) UpdateProfile(ctx context.Context, req *v1.UserUpdateProfileReq) (res *v1.UserUpdateProfileRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "更新成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = service.UserProfile.UpdateProfile(ctx, service.UserProfileUpdateInput{
		UserID:    claims.UserID,
		Nickname:  req.Nickname,
		Phone:     req.Phone,
		Email:     req.Email,
		AvatarURL: req.AvatarURL,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserUpdateProfileRes{}, nil
}

func (c *UserController) UpdatePassword(ctx context.Context, req *v1.UserUpdatePasswordReq) (res *v1.UserUpdatePasswordRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "密码修改成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = service.UserProfile.UpdatePassword(ctx, service.UserProfileUpdatePasswordInput{
		UserID:      claims.UserID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserUpdatePasswordRes{}, nil
}
