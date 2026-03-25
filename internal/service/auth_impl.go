package service

import (
	"context"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type authService struct{}

type authRecord struct {
	ID           int64
	Username     string
	PasswordHash string
	Password     string
	Nickname     string
	DoctorName   string
	RealName     string
	AvatarURL    string
	Status       int
}

func (s authService) Register(ctx context.Context, in RegisterInput) (*RegisterOutput, error) {
	record, err := s.findByUsername(ctx, consts.RoleUser, in.Username)
	if err != nil {
		return nil, err
	}
	if record != nil {
		return nil, gerror.NewCode(gcode.New(409, "", nil), "用户名已存在")
	}

	hashedPassword := hashPassword(in.Password)
	now := time.Now()
	result, err := dao.User.Ctx(ctx).Data(do.User{
		Username:     in.Username,
		PasswordHash: hashedPassword,
		Nickname:     in.Nickname,
		Phone:        in.Phone,
		Email:        in.Email,
		Status:       1,
		CreatedAt:    now,
		UpdatedAt:    now,
	}).Insert()
	if err != nil {
		if isDuplicateErr(err) {
			return nil, gerror.NewCode(gcode.New(409, "", nil), "用户名已存在")
		}
		return nil, gerror.Wrap(err, "创建用户失败")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户ID失败")
	}

	return &RegisterOutput{UserID: userID}, nil
}

func (s authService) Login(ctx context.Context, in LoginInput) (*LoginOutput, error) {
	record, err := s.findByUsername(ctx, in.Role, in.Username)
	if err != nil {
		return nil, err
	}
	if record == nil || !verifyPassword(in.Password, record.PasswordHash, record.Password) {
		return nil, gerror.NewCode(gcode.New(401, "", nil), "用户名或密码错误")
	}
	if record.Status != 1 {
		return nil, gerror.NewCode(gcode.New(403, "", nil), "账号已被禁用")
	}

	expireHours := g.Cfg().MustGet(ctx, "auth.expireHours", 24).Int()
	now := time.Now()
	expireAt := time.Now().Add(time.Duration(expireHours) * time.Hour)
	claims := AuthClaims{
		UserID:   record.ID,
		Username: record.Username,
		Role:     in.Role,
		Issuer:   g.Cfg().MustGet(ctx, "auth.jwtIssuer", "petcare").String(),
		IssuedAt: now.Unix(),
		Exp:      expireAt.Unix(),
	}

	token, err := generateToken(ctx, claims)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Token:    token,
		ExpireAt: expireAt.Format("2006-01-02 15:04:05"),
		UserID:   record.ID,
		Role:     in.Role,
	}, nil
}

func (s authService) Me(ctx context.Context, claims AuthClaims) (*MeOutput, error) {
	record, err := s.findByUsername(ctx, claims.Role, claims.Username)
	if err != nil {
		return nil, err
	}
	if record == nil || record.ID != claims.UserID {
		return nil, gerror.NewCode(gcode.New(401, "", nil), "未登录或 token 无效")
	}

	return &MeOutput{
		ID:         record.ID,
		Username:   record.Username,
		Role:       claims.Role,
		Nickname:   record.Nickname,
		DoctorName: record.DoctorName,
		RealName:   record.RealName,
		AvatarURL:  record.AvatarURL,
	}, nil
}

func (s authService) VerifyToken(ctx context.Context, token string) (*AuthClaims, error) {
	claims, err := parseToken(ctx, token)
	if err != nil {
		return nil, gerror.NewCode(gcode.New(401, "", nil), "未登录或 token 无效")
	}
	return claims, nil
}

func (s authService) findByUsername(ctx context.Context, role string, username string) (*authRecord, error) {
	switch strings.ToLower(role) {
	case consts.RoleUser:
		record, err := dao.User.Ctx(ctx).
			Where(dao.User.Columns().Username, username).
			One()
		if err != nil {
			return nil, gerror.Wrap(err, "查询用户失败")
		}
		if record.IsEmpty() {
			return nil, nil
		}
		return &authRecord{
			ID:           record[dao.User.Columns().Id].Int64(),
			Username:     record[dao.User.Columns().Username].String(),
			PasswordHash: record[dao.User.Columns().PasswordHash].String(),
			Nickname:     record[dao.User.Columns().Nickname].String(),
			AvatarURL:    record[dao.User.Columns().AvatarUrl].String(),
			Status:       record[dao.User.Columns().Status].Int(),
		}, nil
	case consts.RoleDoctor:
		record, err := dao.Doctor.Ctx(ctx).
			Where(dao.Doctor.Columns().Username, username).
			One()
		if err != nil {
			return nil, gerror.Wrap(err, "查询医生失败")
		}
		if record.IsEmpty() {
			return nil, nil
		}
		return &authRecord{
			ID:           record[dao.Doctor.Columns().Id].Int64(),
			Username:     record[dao.Doctor.Columns().Username].String(),
			PasswordHash: record[dao.Doctor.Columns().PasswordHash].String(),
			DoctorName:   record[dao.Doctor.Columns().DoctorName].String(),
			AvatarURL:    record[dao.Doctor.Columns().AvatarUrl].String(),
			Status:       record[dao.Doctor.Columns().Status].Int(),
		}, nil
	case consts.RoleAdmin:
		record, err := dao.Admin.Ctx(ctx).
			Where(dao.Admin.Columns().Username, username).
			One()
		if err != nil {
			return nil, gerror.Wrap(err, "查询管理员失败")
		}
		if record.IsEmpty() {
			return nil, nil
		}
		return &authRecord{
			ID:           record[dao.Admin.Columns().Id].Int64(),
			Username:     record[dao.Admin.Columns().Username].String(),
			PasswordHash: record[dao.Admin.Columns().PasswordHash].String(),
			RealName:     record[dao.Admin.Columns().RealName].String(),
			Status:       record[dao.Admin.Columns().Status].Int(),
		}, nil
	default:
		return nil, gerror.NewCode(gcode.New(400, "", nil), "登录角色不合法")
	}
}

func isDuplicateErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), "duplicate")
}
