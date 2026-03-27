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

type profileMockAuthService struct {
	verifyTokenFunc func(ctx context.Context, token string) (*service.AuthClaims, error)
}

func (m profileMockAuthService) Register(ctx context.Context, in service.RegisterInput) (*service.RegisterOutput, error) {
	return nil, nil
}

func (m profileMockAuthService) Login(ctx context.Context, in service.LoginInput) (*service.LoginOutput, error) {
	return nil, nil
}

func (m profileMockAuthService) Me(ctx context.Context, claims service.AuthClaims) (*service.MeOutput, error) {
	return nil, nil
}

func (m profileMockAuthService) Logout(ctx context.Context, token string) error {
	return nil
}

func (m profileMockAuthService) VerifyToken(ctx context.Context, token string) (*service.AuthClaims, error) {
	return m.verifyTokenFunc(ctx, token)
}

type profileMockUserProfileService struct {
	getFunc            func(ctx context.Context, in service.UserProfileGetInput) (*service.UserProfileGetOutput, error)
	updateFunc         func(ctx context.Context, in service.UserProfileUpdateInput) error
	updatePasswordFunc func(ctx context.Context, in service.UserProfileUpdatePasswordInput) error
}

func (m profileMockUserProfileService) GetProfile(ctx context.Context, in service.UserProfileGetInput) (*service.UserProfileGetOutput, error) {
	return m.getFunc(ctx, in)
}

func (m profileMockUserProfileService) UpdateProfile(ctx context.Context, in service.UserProfileUpdateInput) error {
	return m.updateFunc(ctx, in)
}

func (m profileMockUserProfileService) UpdatePassword(ctx context.Context, in service.UserProfileUpdatePasswordInput) error {
	return m.updatePasswordFunc(ctx, in)
}

type profileMockDoctorProfileService struct {
	getFunc            func(ctx context.Context, in service.DoctorProfileGetInput) (*service.DoctorProfileGetOutput, error)
	updateFunc         func(ctx context.Context, in service.DoctorProfileUpdateInput) error
	updatePasswordFunc func(ctx context.Context, in service.DoctorProfileUpdatePasswordInput) error
}

func (m profileMockDoctorProfileService) GetProfile(ctx context.Context, in service.DoctorProfileGetInput) (*service.DoctorProfileGetOutput, error) {
	return m.getFunc(ctx, in)
}

func (m profileMockDoctorProfileService) UpdateProfile(ctx context.Context, in service.DoctorProfileUpdateInput) error {
	return m.updateFunc(ctx, in)
}

func (m profileMockDoctorProfileService) UpdatePassword(ctx context.Context, in service.DoctorProfileUpdatePasswordInput) error {
	return m.updatePasswordFunc(ctx, in)
}

type profileMockAdminProfileService struct {
	getFunc func(ctx context.Context, in service.AdminProfileGetInput) (*service.AdminProfileGetOutput, error)
}

func (m profileMockAdminProfileService) GetProfile(ctx context.Context, in service.AdminProfileGetInput) (*service.AdminProfileGetOutput, error) {
	return m.getFunc(ctx, in)
}

type profileEnvelope struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func TestProfileEndpoints(t *testing.T) {
	utility.ConfigureTestConfig(t)

	oldAuthService := service.Auth
	oldUserProfile := service.UserProfile
	oldDoctorProfile := service.DoctorProfile
	oldAdminProfile := service.AdminProfile

	service.Auth = profileMockAuthService{
		verifyTokenFunc: func(ctx context.Context, token string) (*service.AuthClaims, error) {
			switch token {
			case "user-token":
				return &service.AuthClaims{UserID: 11, Username: "user001", Role: consts.RoleUser}, nil
			case "doctor-token":
				return &service.AuthClaims{UserID: 22, Username: "doctor_zhang", Role: consts.RoleDoctor}, nil
			case "admin-token":
				return &service.AuthClaims{UserID: 33, Username: "admin001", Role: consts.RoleAdmin}, nil
			default:
				return nil, consts.NewUnauthorizedError("")
			}
		},
	}
	service.UserProfile = profileMockUserProfileService{
		getFunc: func(ctx context.Context, in service.UserProfileGetInput) (*service.UserProfileGetOutput, error) {
			if in.UserID != 11 {
				t.Fatalf("unexpected user profile get input: %+v", in)
			}
			return &service.UserProfileGetOutput{
				ID:        11,
				Username:  "user001",
				Nickname:  "小猫家长",
				Phone:     "13800000000",
				Email:     "user001@example.com",
				AvatarURL: "https://example.com/user.jpg",
				Status:    1,
				CreatedAt: "2026-03-25 10:00:00",
			}, nil
		},
		updateFunc: func(ctx context.Context, in service.UserProfileUpdateInput) error {
			if in.UserID != 11 || in.Nickname == nil || *in.Nickname != "元元" {
				t.Fatalf("unexpected user profile update input: %+v", in)
			}
			return nil
		},
		updatePasswordFunc: func(ctx context.Context, in service.UserProfileUpdatePasswordInput) error {
			if in.UserID != 11 || in.OldPassword != "123456" || in.NewPassword != "654321" {
				t.Fatalf("unexpected user password update input: %+v", in)
			}
			return nil
		},
	}
	service.DoctorProfile = profileMockDoctorProfileService{
		getFunc: func(ctx context.Context, in service.DoctorProfileGetInput) (*service.DoctorProfileGetOutput, error) {
			if in.DoctorID != 22 {
				t.Fatalf("unexpected doctor profile get input: %+v", in)
			}
			return &service.DoctorProfileGetOutput{
				ID:           22,
				HospitalID:   1,
				HospitalName: "爱宠动物医院",
				Username:     "doctor_zhang",
				DoctorName:   "张医生",
				Gender:       1,
				Phone:        "13900000000",
				Email:        "doctor@example.com",
				Title:        "主治医师",
				Specialty:    "猫科内科",
				AvatarURL:    "https://example.com/doctor.jpg",
				Intro:        "擅长猫科消化系统疾病诊疗",
				Status:       1,
			}, nil
		},
		updateFunc: func(ctx context.Context, in service.DoctorProfileUpdateInput) error {
			if in.DoctorID != 22 || in.Phone == nil || *in.Phone != "13911112222" {
				t.Fatalf("unexpected doctor profile update input: %+v", in)
			}
			return nil
		},
		updatePasswordFunc: func(ctx context.Context, in service.DoctorProfileUpdatePasswordInput) error {
			if in.DoctorID != 22 || in.OldPassword != "123456" || in.NewPassword != "654321" {
				t.Fatalf("unexpected doctor password update input: %+v", in)
			}
			return nil
		},
	}
	service.AdminProfile = profileMockAdminProfileService{
		getFunc: func(ctx context.Context, in service.AdminProfileGetInput) (*service.AdminProfileGetOutput, error) {
			if in.AdminID != 33 {
				t.Fatalf("unexpected admin profile get input: %+v", in)
			}
			return &service.AdminProfileGetOutput{
				ID:        33,
				Username:  "admin001",
				RealName:  "系统管理员",
				Phone:     "13700000000",
				Status:    1,
				CreatedAt: "2026-03-25 09:00:00",
			}, nil
		},
	}

	t.Cleanup(func() {
		service.Auth = oldAuthService
		service.UserProfile = oldUserProfile
		service.DoctorProfile = oldDoctorProfile
		service.AdminProfile = oldAdminProfile
	})

	server := g.Server(guid.S())
	server.SetAddr("127.0.0.1:0")
	server.SetDumpRouterMap(false)
	server.Use(middleware.ResponseHandler)
	server.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/users", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth, middleware.Role(consts.RoleUser))
			group.Bind(NewUser())
		})
		group.Group("/doctors", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth, middleware.Role(consts.RoleDoctor))
			group.Bind(NewDoctor())
		})
		group.Group("/admin", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth, middleware.Role(consts.RoleAdmin))
			group.Bind(NewAdmin())
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

	t.Run("user profile", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer user-token")
		content := client.GetContent(context.Background(), "/api/users/profile")
		assertEnvelopeMessage(t, content, 200, "success")
	})

	t.Run("user update profile", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer user-token")
		client.SetContentType("application/json")
		content := client.PutContent(context.Background(), "/api/users/profile", map[string]any{
			"nickname":   "元元",
			"phone":      "13812345678",
			"email":      "new@example.com",
			"avatar_url": "https://example.com/new-avatar.jpg",
		})
		assertEnvelopeMessage(t, content, 200, "更新成功")
	})

	t.Run("user update password", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer user-token")
		client.SetContentType("application/json")
		content := client.PutContent(context.Background(), "/api/users/password", map[string]any{
			"old_password": "123456",
			"new_password": "654321",
		})
		assertEnvelopeMessage(t, content, 200, "密码修改成功")
	})

	t.Run("doctor profile", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer doctor-token")
		content := client.GetContent(context.Background(), "/api/doctors/profile")
		assertEnvelopeMessage(t, content, 200, "success")
	})

	t.Run("doctor update profile", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer doctor-token")
		client.SetContentType("application/json")
		content := client.PutContent(context.Background(), "/api/doctors/profile", map[string]any{
			"phone":      "13911112222",
			"email":      "doctor_new@example.com",
			"avatar_url": "https://example.com/new-doctor.jpg",
			"intro":      "擅长猫科消化系统与呼吸系统疾病",
		})
		assertEnvelopeMessage(t, content, 200, "更新成功")
	})

	t.Run("doctor update password", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer doctor-token")
		client.SetContentType("application/json")
		content := client.PutContent(context.Background(), "/api/doctors/password", map[string]any{
			"old_password": "123456",
			"new_password": "654321",
		})
		assertEnvelopeMessage(t, content, 200, "密码修改成功")
	})

	t.Run("admin profile", func(t *testing.T) {
		client := g.Client()
		client.SetPrefix(baseURL)
		client.SetHeader("Authorization", "Bearer admin-token")
		content := client.GetContent(context.Background(), "/api/admin/profile")
		assertEnvelopeMessage(t, content, 200, "success")
	})
}

func assertEnvelopeMessage(t *testing.T, content string, code int, message string) {
	t.Helper()

	var envelope profileEnvelope
	if err := json.Unmarshal([]byte(content), &envelope); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if envelope.Code != code || envelope.Message != message {
		t.Fatalf("unexpected response: %+v", envelope)
	}
}
