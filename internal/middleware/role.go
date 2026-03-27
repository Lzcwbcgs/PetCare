package middleware

import (
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Role checks whether the current authenticated subject belongs to one of the allowed roles.
func Role(roles ...string) func(r *ghttp.Request) {
	var allowedRoles = make([]string, 0, len(roles))
	for _, role := range roles {
		if normalizedRole := service.NormalizeRole(role); normalizedRole != "" {
			allowedRoles = append(allowedRoles, normalizedRole)
		}
	}

	return func(r *ghttp.Request) {
		if len(allowedRoles) == 0 {
			r.Middleware.Next()
			return
		}

		claims, ok := r.GetCtxVar(consts.CtxKeyAuthClaims).Val().(*service.AuthClaims)
		if !ok || claims == nil {
			r.SetError(consts.NewUnauthorizedError(""))
			return
		}
		if !claims.HasRole(allowedRoles...) {
			r.SetError(consts.NewForbiddenError(""))
			return
		}

		r.Middleware.Next()
	}
}
