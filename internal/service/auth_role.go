package service

import (
	"strings"

	"PetCare/internal/consts"
)

var supportedRoles = map[string]struct{}{
	consts.RoleUser:   {},
	consts.RoleDoctor: {},
	consts.RoleAdmin:  {},
}

func NormalizeRole(role string) string {
	return strings.ToLower(strings.TrimSpace(role))
}

func IsSupportedRole(role string) bool {
	_, ok := supportedRoles[NormalizeRole(role)]
	return ok
}

func RoleAllowed(role string, roles ...string) bool {
	var normalizedRole = NormalizeRole(role)
	if normalizedRole == "" {
		return false
	}
	for _, candidate := range roles {
		if normalizedRole == NormalizeRole(candidate) {
			return true
		}
	}
	return false
}

func (c AuthClaims) HasRole(roles ...string) bool {
	return RoleAllowed(c.Role, roles...)
}
