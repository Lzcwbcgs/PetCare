package pet

import (
	"context"

	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func authClaimsFromCtx(ctx context.Context) (*service.AuthClaims, error) {
	claims, ok := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxKeyAuthClaims).Val().(*service.AuthClaims)
	if !ok || claims == nil {
		return nil, consts.NewUnauthorizedError("")
	}
	return claims, nil
}
