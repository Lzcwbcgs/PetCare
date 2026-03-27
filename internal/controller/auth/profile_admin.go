package auth

import (
	"context"

	v1 "PetCare/api/auth/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminController struct{}

func NewAdmin() *AdminController {
	return &AdminController{}
}

func (c *AdminController) Profile(ctx context.Context, req *v1.AdminProfileReq) (res *v1.AdminProfileRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	output, err := service.AdminProfile.GetProfile(ctx, service.AdminProfileGetInput{
		AdminID: claims.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AdminProfileRes{
		ID:        output.ID,
		Username:  output.Username,
		RealName:  output.RealName,
		Phone:     output.Phone,
		Status:    output.Status,
		CreatedAt: output.CreatedAt,
	}, nil
}
