package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/middleware"
	"PetCare/internal/service"
	"PetCare/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

type controllerMockAuthService struct {
	registerFunc    func(ctx context.Context, in service.RegisterInput) (*service.RegisterOutput, error)
	loginFunc       func(ctx context.Context, in service.LoginInput) (*service.LoginOutput, error)
	meFunc          func(ctx context.Context, claims service.AuthClaims) (*service.MeOutput, error)
	logoutFunc      func(ctx context.Context, token string) error
	verifyTokenFunc func(ctx context.Context, token string) (*service.AuthClaims, error)
}

func (m controllerMockAuthService) Register(ctx context.Context, in service.RegisterInput) (*service.RegisterOutput, error) {
	return m.registerFunc(ctx, in)
}

func (m controllerMockAuthService) Login(ctx context.Context, in service.LoginInput) (*service.LoginOutput, error) {
	return m.loginFunc(ctx, in)
}

func (m controllerMockAuthService) Me(ctx context.Context, claims service.AuthClaims) (*service.MeOutput, error) {
	return m.meFunc(ctx, claims)
}

func (m controllerMockAuthService) Logout(ctx context.Context, token string) error {
	return m.logoutFunc(ctx, token)
}

func (m controllerMockAuthService) VerifyToken(ctx context.Context, token string) (*service.AuthClaims, error) {
	return m.verifyTokenFunc(ctx, token)
}

type controllerEnvelope struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type registerResponse struct {
	UserID int64 `json:"user_id"`
}

type loginResponse struct {
	Token    string `json:"token"`
	ExpireAt string `json:"expire_at"`
	UserID   int64  `json:"user_id"`
	Role     string `json:"role"`
}

type meResponse struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	DoctorName string `json:"doctor_name"`
}

func TestAuthEndpoints(t *testing.T) {
	utility.ConfigureTestConfig(t)

	oldAuthService := service.Auth
	service.Auth = controllerMockAuthService{
		registerFunc: func(ctx context.Context, in service.RegisterInput) (*service.RegisterOutput, error) {
			if in.Username != "user001" || in.Nickname != "small-cat-owner" {
				t.Fatalf("unexpected register input: %+v", in)
			}
			return &service.RegisterOutput{UserID: 101}, nil
		},
		loginFunc: func(ctx context.Context, in service.LoginInput) (*service.LoginOutput, error) {
			if in.Username != "doctor_zhang" || in.Role != consts.RoleDoctor {
				t.Fatalf("unexpected login input: %+v", in)
			}
			return &service.LoginOutput{
				Token:    "doctor-token",
				ExpireAt: "2026-03-26 10:00:00",
				UserID:   2,
				Role:     consts.RoleDoctor,
			}, nil
		},
		meFunc: func(ctx context.Context, claims service.AuthClaims) (*service.MeOutput, error) {
			if claims.Role != consts.RoleDoctor || claims.UserID != 2 {
				t.Fatalf("unexpected claims: %+v", claims)
			}
			return &service.MeOutput{
				ID:         2,
				Username:   "doctor_zhang",
				Role:       consts.RoleDoctor,
				DoctorName: "Dr. Zhang",
			}, nil
		},
		logoutFunc: func(ctx context.Context, token string) error {
			if token != "doctor-token" {
				t.Fatalf("unexpected logout token: %s", token)
			}
			return nil
		},
		verifyTokenFunc: func(ctx context.Context, token string) (*service.AuthClaims, error) {
			if token != "doctor-token" {
				return nil, consts.NewUnauthorizedError("")
			}
			return &service.AuthClaims{
				UserID:   2,
				Username: "doctor_zhang",
				Role:     consts.RoleDoctor,
			}, nil
		},
	}
	t.Cleanup(func() {
		service.Auth = oldAuthService
	})

	server := g.Server(guid.S())
	server.SetAddr("127.0.0.1:0")
	server.SetDumpRouterMap(false)
	server.Use(middleware.ResponseHandler)
	server.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/auth", func(group *ghttp.RouterGroup) {
			group.Bind(NewPublic())
			group.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.Auth)
				group.Bind(NewPrivate())
			})
		})
	})
	if err := server.Start(); err != nil {
		t.Fatalf("start test server: %v", err)
	}
	t.Cleanup(func() {
		_ = server.Shutdown()
	})

	time.Sleep(100 * time.Millisecond)

	client := g.Client()
	client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", server.GetListenedPort()))
	client.SetContentType("application/json")

	t.Run("register", func(t *testing.T) {
		content := client.PostContent(context.Background(), "/api/auth/register", map[string]any{
			"username": "user001",
			"password": "123456",
			"nickname": "small-cat-owner",
			"phone":    "13800000000",
			"email":    "user001@example.com",
		})
		var envelope controllerEnvelope
		if err := json.Unmarshal([]byte(content), &envelope); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if envelope.Code != 200 || envelope.Message != "注册成功" {
			t.Fatalf("unexpected response: %+v", envelope)
		}

		var data registerResponse
		if err := json.Unmarshal(envelope.Data, &data); err != nil {
			t.Fatalf("decode data: %v", err)
		}
		if data.UserID != 101 {
			t.Fatalf("unexpected user id: %+v", data)
		}
	})

	t.Run("login", func(t *testing.T) {
		content := client.PostContent(context.Background(), "/api/auth/login", map[string]any{
			"username": "doctor_zhang",
			"password": "123456",
			"role":     "doctor",
		})
		var envelope controllerEnvelope
		if err := json.Unmarshal([]byte(content), &envelope); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if envelope.Code != 200 || envelope.Message != "登录成功" {
			t.Fatalf("unexpected response: %+v", envelope)
		}

		var data loginResponse
		if err := json.Unmarshal(envelope.Data, &data); err != nil {
			t.Fatalf("decode data: %v", err)
		}
		if data.Token != "doctor-token" || data.Role != consts.RoleDoctor {
			t.Fatalf("unexpected login response: %+v", data)
		}
	})

	t.Run("me", func(t *testing.T) {
		meClient := g.Client()
		meClient.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", server.GetListenedPort()))
		meClient.SetHeader("Authorization", "Bearer doctor-token")
		content := meClient.GetContent(context.Background(), "/api/auth/me")
		var envelope controllerEnvelope
		if err := json.Unmarshal([]byte(content), &envelope); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if envelope.Code != 200 || envelope.Message != "success" {
			t.Fatalf("unexpected response: %+v", envelope)
		}

		var data meResponse
		if err := json.Unmarshal(envelope.Data, &data); err != nil {
			t.Fatalf("decode data: %v", err)
		}
		if data.Role != consts.RoleDoctor || data.DoctorName != "Dr. Zhang" {
			t.Fatalf("unexpected me response: %+v", data)
		}
	})

	t.Run("logout", func(t *testing.T) {
		logoutClient := g.Client()
		logoutClient.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", server.GetListenedPort()))
		logoutClient.SetHeader("Authorization", "Bearer doctor-token")
		content := logoutClient.PostContent(context.Background(), "/api/auth/logout")
		var envelope controllerEnvelope
		if err := json.Unmarshal([]byte(content), &envelope); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if envelope.Code != 200 || envelope.Message != "退出成功" {
			t.Fatalf("unexpected response: %+v", envelope)
		}
	})
}
