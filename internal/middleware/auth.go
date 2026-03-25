package middleware

import (
	"strings"

	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Auth(r *ghttp.Request) {
	token := bearerToken(r.Header.Get("Authorization"))
	if token == "" {
		r.SetError(gerror.NewCode(gcode.New(401, "", nil), "未登录或 token 无效"))
		return
	}

	claims, err := service.Auth.VerifyToken(r.GetCtx(), token)
	if err != nil {
		r.SetError(err)
		return
	}

	r.SetCtxVar(consts.CtxKeyAuthClaims, claims)
	r.Middleware.Next()
}

func bearerToken(authorization string) string {
	if authorization == "" {
		return ""
	}
	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}
	return strings.TrimSpace(parts[1])
}
