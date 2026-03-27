package service

import (
	"context"
	"testing"
	"time"

	"PetCare/internal/consts"
	"PetCare/utility"

	"github.com/gogf/gf/v2/errors/gerror"
)

func TestPasswordHashAndVerify(t *testing.T) {
	hashed := hashPassword("123456")
	if !verifyPassword("123456", hashed) {
		t.Fatalf("expected hashed password to verify")
	}
	if !verifyPassword("123456", "123456") {
		t.Fatalf("expected plain password compatibility to verify")
	}
	if verifyPassword("wrong", hashed) {
		t.Fatalf("expected wrong password to fail")
	}
}

func TestGenerateAndParseToken(t *testing.T) {
	utility.ConfigureTestConfig(t)

	claims := AuthClaims{
		UserID:   1,
		Username: "tester",
		Role:     "user",
		Issuer:   "petcare",
		IssuedAt: time.Now().Unix(),
		Exp:      time.Now().Add(time.Hour).Unix(),
	}

	token, err := generateToken(context.Background(), claims)
	if err != nil {
		t.Fatalf("generate token failed: %v", err)
	}

	parsed, err := parseToken(context.Background(), token)
	if err != nil {
		t.Fatalf("parse token failed: %v", err)
	}

	if parsed.UserID != claims.UserID || parsed.Username != claims.Username || parsed.Role != claims.Role {
		t.Fatalf("claims mismatch after parse")
	}
}

func TestParseTokenRejectsExpiredToken(t *testing.T) {
	token, err := generateToken(context.Background(), AuthClaims{
		UserID:   1,
		Username: "tester",
		Role:     consts.RoleUser,
		Issuer:   "petcare",
		IssuedAt: time.Now().Add(-2 * time.Hour).Unix(),
		Exp:      time.Now().Add(-1 * time.Hour).Unix(),
	})
	if err != nil {
		t.Fatalf("generate token failed: %v", err)
	}

	_, err = parseToken(context.Background(), token)
	if err == nil {
		t.Fatalf("expected expired token to fail")
	}
	if code := gerror.Code(err); code == nil || code.Code() != 401 {
		t.Fatalf("expected 401 code, got: %v", code)
	}
}

func TestParseTokenRejectsInvalidRole(t *testing.T) {
	token, err := generateToken(context.Background(), AuthClaims{
		UserID:   1,
		Username: "tester",
		Role:     "guest",
		Issuer:   "petcare",
		IssuedAt: time.Now().Unix(),
		Exp:      time.Now().Add(time.Hour).Unix(),
	})
	if err != nil {
		t.Fatalf("generate token failed: %v", err)
	}

	_, err = parseToken(context.Background(), token)
	if err == nil {
		t.Fatalf("expected invalid role token to fail")
	}
	if code := gerror.Code(err); code == nil || code.Code() != 401 {
		t.Fatalf("expected 401 code, got: %v", code)
	}
}

func TestRevokeToken(t *testing.T) {
	token := "revoked-token"
	revokeToken(token, time.Minute)
	if !isTokenRevoked(token) {
		t.Fatalf("expected token to be revoked")
	}
}

func TestRevokedTokenExpiresFromStore(t *testing.T) {
	token := "short-lived-token"
	revokeToken(token, 10*time.Millisecond)
	time.Sleep(30 * time.Millisecond)
	if isTokenRevoked(token) {
		t.Fatalf("expected expired revoked token to be cleaned up")
	}
}
