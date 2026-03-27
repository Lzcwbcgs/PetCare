package service

import "context"

type (
	RegisterInput struct {
		Username string
		Password string
		Nickname string
		Phone    string
		Email    string
	}

	RegisterOutput struct {
		UserID int64
	}

	LoginInput struct {
		Username string
		Password string
		Role     string
	}

	LoginOutput struct {
		Token    string
		ExpireAt string
		UserID   int64
		Role     string
	}

	MeOutput struct {
		ID         int64
		Username   string
		Role       string
		Nickname   string
		DoctorName string
		RealName   string
		AvatarURL  string
	}

	AuthClaims struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
		Role     string `json:"role"`
		Issuer   string `json:"iss"`
		IssuedAt int64  `json:"iat"`
		Exp      int64  `json:"exp"`
	}
)

type IAuth interface {
	Register(ctx context.Context, in RegisterInput) (*RegisterOutput, error)
	Login(ctx context.Context, in LoginInput) (*LoginOutput, error)
	Me(ctx context.Context, claims AuthClaims) (*MeOutput, error)
	Logout(ctx context.Context, token string) error
	VerifyToken(ctx context.Context, token string) (*AuthClaims, error)
}

var Auth IAuth = authService{}
