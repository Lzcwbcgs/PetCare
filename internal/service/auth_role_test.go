package service

import (
	"testing"

	"PetCare/internal/consts"
)

func TestNormalizeRole(t *testing.T) {
	if got := NormalizeRole(" Admin "); got != consts.RoleAdmin {
		t.Fatalf("unexpected normalized role: %q", got)
	}
}

func TestIsSupportedRole(t *testing.T) {
	if !IsSupportedRole("doctor") {
		t.Fatalf("expected doctor to be supported")
	}
	if IsSupportedRole("guest") {
		t.Fatalf("expected guest to be unsupported")
	}
}

func TestAuthClaimsHasRole(t *testing.T) {
	claims := AuthClaims{Role: consts.RoleDoctor}
	if !claims.HasRole(consts.RoleDoctor, consts.RoleAdmin) {
		t.Fatalf("expected role match")
	}
	if claims.HasRole(consts.RoleUser) {
		t.Fatalf("expected role mismatch")
	}
}
