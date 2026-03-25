package service

import (
	"context"
	"testing"
	"time"
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
