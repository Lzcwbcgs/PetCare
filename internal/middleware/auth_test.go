package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/service"
	"PetCare/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

type middlewareMockAuthService struct {
	verifyTokenFunc func(ctx context.Context, token string) (*service.AuthClaims, error)
}

func (m middlewareMockAuthService) Register(ctx context.Context, in service.RegisterInput) (*service.RegisterOutput, error) {
	return nil, nil
}

func (m middlewareMockAuthService) Login(ctx context.Context, in service.LoginInput) (*service.LoginOutput, error) {
	return nil, nil
}

func (m middlewareMockAuthService) Me(ctx context.Context, claims service.AuthClaims) (*service.MeOutput, error) {
	return nil, nil
}

func (m middlewareMockAuthService) Logout(ctx context.Context, token string) error {
	return nil
}

func (m middlewareMockAuthService) VerifyToken(ctx context.Context, token string) (*service.AuthClaims, error) {
	return m.verifyTokenFunc(ctx, token)
}

type middlewareEnvelope struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type adminDashboardReq struct {
	g.Meta `path:"/dashboard" method:"get"`
}

type adminDashboardRes struct {
	Scope string `json:"scope"`
}

type adminDashboardController struct{}

func (c *adminDashboardController) Dashboard(ctx context.Context, req *adminDashboardReq) (res *adminDashboardRes, err error) {
	return &adminDashboardRes{Scope: consts.RoleAdmin}, nil
}

func TestAuthAndRoleMiddleware(t *testing.T) {
	utility.ConfigureTestConfig(t)

	oldAuthService := service.Auth
	service.Auth = middlewareMockAuthService{
		verifyTokenFunc: func(ctx context.Context, token string) (*service.AuthClaims, error) {
			switch token {
			case "admin-token":
				return &service.AuthClaims{UserID: 1, Username: "admin", Role: consts.RoleAdmin}, nil
			case "user-token":
				return &service.AuthClaims{UserID: 2, Username: "user", Role: consts.RoleUser}, nil
			default:
				return nil, consts.NewUnauthorizedError("")
			}
		},
	}
	t.Cleanup(func() {
		service.Auth = oldAuthService
	})

	server := g.Server(guid.S())
	server.SetAddr("127.0.0.1:0")
	server.SetDumpRouterMap(false)
	server.Use(ResponseHandler)
	server.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/admin", func(group *ghttp.RouterGroup) {
			group.Middleware(Auth, Role(consts.RoleAdmin))
			group.Bind(new(adminDashboardController))
		})
	})
	if err := server.Start(); err != nil {
		t.Fatalf("start test server: %v", err)
	}
	t.Cleanup(func() {
		_ = server.Shutdown()
	})

	time.Sleep(100 * time.Millisecond)

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", server.GetListenedPort())

	t.Run("missing token", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		content := client.GetContent(context.Background(), "/api/admin/dashboard")
		var envelope middlewareEnvelope
		if err := json.Unmarshal([]byte(content), &envelope); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if envelope.Code != 401 || envelope.Message != consts.ErrorMessageUnauthorized {
			t.Fatalf("unexpected response: %+v", envelope)
		}
	})

	t.Run("forbidden role", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer user-token")
		content := client.GetContent(context.Background(), "/api/admin/dashboard")
		var envelope middlewareEnvelope
		if err := json.Unmarshal([]byte(content), &envelope); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if envelope.Code != 403 || envelope.Message != consts.ErrorMessageForbidden {
			t.Fatalf("unexpected response: %+v", envelope)
		}
	})

	t.Run("allowed role", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer admin-token")
		content := client.GetContent(context.Background(), "/api/admin/dashboard")
		var envelope middlewareEnvelope
		if err := json.Unmarshal([]byte(content), &envelope); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if envelope.Code != 200 || envelope.Message != "success" {
			t.Fatalf("unexpected response: %+v", envelope)
		}

		var data adminDashboardRes
		if err := json.Unmarshal(envelope.Data, &data); err != nil {
			t.Fatalf("decode data: %v", err)
		}
		if data.Scope != consts.RoleAdmin {
			t.Fatalf("unexpected scope: %+v", data)
		}
	})
}
