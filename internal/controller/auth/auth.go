package auth

import (
	"context"

	authv1 "PetCare/api/auth/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
)

type PublicController struct{}
type PrivateController struct{}

func NewPublic() *PublicController {
	return &PublicController{}
}

func NewPrivate() *PrivateController {
	return &PrivateController{}
}

func (c *PublicController) Register(ctx context.Context, req *authv1.RegisterReq) (res *authv1.RegisterRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "注册成功")
	output, err := service.Auth.Register(ctx, service.RegisterInput{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &authv1.RegisterRes{UserID: output.UserID}, nil
}

func (c *PublicController) Login(ctx context.Context, req *authv1.LoginReq) (res *authv1.LoginRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "登录成功")
	output, err := service.Auth.Login(ctx, service.LoginInput{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		return nil, err
	}
	return &authv1.LoginRes{
		Token:    output.Token,
		ExpireAt: output.ExpireAt,
		UserID:   output.UserID,
		Role:     output.Role,
	}, nil
}

func (c *PrivateController) Me(ctx context.Context, req *authv1.MeReq) (res *authv1.MeRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")
	claims, ok := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxKeyAuthClaims).Val().(*service.AuthClaims)
	if !ok || claims == nil {
		return nil, gerror.NewCode(gcode.New(401, "", nil), "未登录或 token 无效")
	}

	output, err := service.Auth.Me(ctx, *claims)
	if err != nil {
		return nil, err
	}
	return &authv1.MeRes{
		ID:         output.ID,
		Username:   output.Username,
		Role:       output.Role,
		Nickname:   output.Nickname,
		DoctorName: output.DoctorName,
		RealName:   output.RealName,
		AvatarURL:  output.AvatarURL,
	}, nil
}

func (c *PrivateController) Logout(ctx context.Context, req *authv1.LogoutReq) (res *authv1.LogoutRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "退出成功")
	return &authv1.LogoutRes{}, nil
}
