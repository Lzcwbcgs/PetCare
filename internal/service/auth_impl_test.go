package service

import (
	"testing"

	"PetCare/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
)

func TestValidateClaimsRecord(t *testing.T) {
	t.Run("valid record", func(t *testing.T) {
		err := validateClaimsRecord(
			AuthClaims{UserID: 1, Username: "user001", Role: consts.RoleUser},
			&authRecord{ID: 1, Username: "user001", Status: 1},
		)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	})

	t.Run("missing record", func(t *testing.T) {
		err := validateClaimsRecord(AuthClaims{UserID: 1}, nil)
		if err == nil {
			t.Fatalf("expected unauthorized error")
		}
		if code := gerror.Code(err); code == nil || code.Code() != 401 {
			t.Fatalf("expected 401 code, got: %v", code)
		}
	})

	t.Run("disabled record", func(t *testing.T) {
		err := validateClaimsRecord(
			AuthClaims{UserID: 1, Username: "user001", Role: consts.RoleUser},
			&authRecord{ID: 1, Username: "user001", Status: 0},
		)
		if err == nil {
			t.Fatalf("expected forbidden error")
		}
		if code := gerror.Code(err); code == nil || code.Code() != 403 {
			t.Fatalf("expected 403 code, got: %v", code)
		}
	})
}
