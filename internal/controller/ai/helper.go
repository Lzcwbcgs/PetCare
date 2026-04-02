package ai

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

func requireAIRoles(ctx context.Context, roles ...string) (*service.AuthClaims, error) {
	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	if !claims.HasRole(roles...) {
		return nil, consts.NewForbiddenError("")
	}
	return claims, nil
}

func derefPage(value *int, fallback int) int {
	if value == nil {
		return fallback
	}
	return *value
}
