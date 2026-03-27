package service

import (
	"context"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
)

type (
	AdminProfileGetInput struct {
		AdminID int64
	}

	AdminProfileGetOutput struct {
		ID        int64
		Username  string
		RealName  string
		Phone     string
		Status    int
		CreatedAt string
	}
)

type IAdminProfile interface {
	GetProfile(ctx context.Context, in AdminProfileGetInput) (*AdminProfileGetOutput, error)
}

var AdminProfile IAdminProfile = adminProfileService{}

type adminProfileService struct{}

func (s adminProfileService) GetProfile(ctx context.Context, in AdminProfileGetInput) (*AdminProfileGetOutput, error) {
	record, err := dao.Admin.Ctx(ctx).
		Where(dao.Admin.Columns().Id, in.AdminID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询管理员资料失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("管理员资料不存在")
	}

	return &AdminProfileGetOutput{
		ID:        record[dao.Admin.Columns().Id].Int64(),
		Username:  record[dao.Admin.Columns().Username].String(),
		RealName:  record[dao.Admin.Columns().RealName].String(),
		Phone:     record[dao.Admin.Columns().Phone].String(),
		Status:    record[dao.Admin.Columns().Status].Int(),
		CreatedAt: record[dao.Admin.Columns().CreatedAt].Time().Format("2006-01-02 15:04:05"),
	}, nil
}
