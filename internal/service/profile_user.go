package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	UserProfileGetInput struct {
		UserID int64
	}

	UserProfileGetOutput struct {
		ID        int64
		Username  string
		Nickname  string
		Phone     string
		Email     string
		AvatarURL string
		Status    int
		CreatedAt string
	}

	UserProfileUpdateInput struct {
		UserID    int64
		Nickname  *string
		Phone     *string
		Email     *string
		AvatarURL *string
	}

	UserProfileUpdatePasswordInput struct {
		UserID      int64
		OldPassword string
		NewPassword string
	}
)

type IUserProfile interface {
	GetProfile(ctx context.Context, in UserProfileGetInput) (*UserProfileGetOutput, error)
	UpdateProfile(ctx context.Context, in UserProfileUpdateInput) error
	UpdatePassword(ctx context.Context, in UserProfileUpdatePasswordInput) error
}

var UserProfile IUserProfile = userProfileService{}

type userProfileService struct{}

func (s userProfileService) GetProfile(ctx context.Context, in UserProfileGetInput) (*UserProfileGetOutput, error) {
	record, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Id, in.UserID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询用户资料失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("用户资料不存在")
	}

	return &UserProfileGetOutput{
		ID:        record[dao.User.Columns().Id].Int64(),
		Username:  record[dao.User.Columns().Username].String(),
		Nickname:  record[dao.User.Columns().Nickname].String(),
		Phone:     record[dao.User.Columns().Phone].String(),
		Email:     record[dao.User.Columns().Email].String(),
		AvatarURL: record[dao.User.Columns().AvatarUrl].String(),
		Status:    record[dao.User.Columns().Status].Int(),
		CreatedAt: record[dao.User.Columns().CreatedAt].Time().Format("2006-01-02 15:04:05"),
	}, nil
}

func (s userProfileService) UpdateProfile(ctx context.Context, in UserProfileUpdateInput) error {
	data := do.User{}
	if in.Nickname != nil {
		data.Nickname = *in.Nickname
	}
	if in.Phone != nil {
		data.Phone = *in.Phone
	}
	if in.Email != nil {
		data.Email = *in.Email
	}
	if in.AvatarURL != nil {
		data.AvatarUrl = *in.AvatarURL
	}
	if !hasUserProfileUpdates(in) {
		return consts.NewBadRequestError("至少提供一个更新字段")
	}
	data.UpdatedAt = time.Now()

	result, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Id, in.UserID).
		Data(data).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "更新用户资料失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("用户资料不存在")
	}
	return nil
}

func (s userProfileService) UpdatePassword(ctx context.Context, in UserProfileUpdatePasswordInput) error {
	record, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Id, in.UserID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询用户资料失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("用户资料不存在")
	}
	if !verifyPassword(in.OldPassword, record[dao.User.Columns().PasswordHash].String()) {
		return consts.NewBadRequestError("旧密码错误")
	}

	result, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Id, in.UserID).
		Data(do.User{
			PasswordHash: hashPassword(in.NewPassword),
			UpdatedAt:    time.Now(),
		}).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "修改用户密码失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("用户资料不存在")
	}
	return nil
}

func hasUserProfileUpdates(in UserProfileUpdateInput) bool {
	return in.Nickname != nil || in.Phone != nil || in.Email != nil || in.AvatarURL != nil
}
