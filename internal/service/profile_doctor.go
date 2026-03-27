package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	DoctorProfileGetInput struct {
		DoctorID int64
	}

	DoctorProfileGetOutput struct {
		ID           int64
		HospitalID   int64
		HospitalName string
		Username     string
		DoctorName   string
		Gender       int
		Phone        string
		Email        string
		Title        string
		Specialty    string
		AvatarURL    string
		Intro        string
		Status       int
	}

	DoctorProfileUpdateInput struct {
		DoctorID  int64
		Phone     *string
		Email     *string
		AvatarURL *string
		Intro     *string
	}

	DoctorProfileUpdatePasswordInput struct {
		DoctorID    int64
		OldPassword string
		NewPassword string
	}
)

type IDoctorProfile interface {
	GetProfile(ctx context.Context, in DoctorProfileGetInput) (*DoctorProfileGetOutput, error)
	UpdateProfile(ctx context.Context, in DoctorProfileUpdateInput) error
	UpdatePassword(ctx context.Context, in DoctorProfileUpdatePasswordInput) error
}

var DoctorProfile IDoctorProfile = doctorProfileService{}

type doctorProfileService struct{}

func (s doctorProfileService) GetProfile(ctx context.Context, in DoctorProfileGetInput) (*DoctorProfileGetOutput, error) {
	record, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, in.DoctorID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生资料失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("医生资料不存在")
	}

	var hospitalName string
	if hospitalID := record[dao.Doctor.Columns().HospitalId].Int64(); hospitalID > 0 {
		hospitalRecord, hospitalErr := dao.Hospital.Ctx(ctx).
			Where(dao.Hospital.Columns().Id, hospitalID).
			One()
		if hospitalErr != nil {
			return nil, consts.WrapInternalError(hospitalErr, "查询医院资料失败")
		}
		if !hospitalRecord.IsEmpty() {
			hospitalName = hospitalRecord[dao.Hospital.Columns().HospitalName].String()
		}
	}

	return &DoctorProfileGetOutput{
		ID:           record[dao.Doctor.Columns().Id].Int64(),
		HospitalID:   record[dao.Doctor.Columns().HospitalId].Int64(),
		HospitalName: hospitalName,
		Username:     record[dao.Doctor.Columns().Username].String(),
		DoctorName:   record[dao.Doctor.Columns().DoctorName].String(),
		Gender:       record[dao.Doctor.Columns().Gender].Int(),
		Phone:        record[dao.Doctor.Columns().Phone].String(),
		Email:        record[dao.Doctor.Columns().Email].String(),
		Title:        record[dao.Doctor.Columns().Title].String(),
		Specialty:    record[dao.Doctor.Columns().Specialty].String(),
		AvatarURL:    record[dao.Doctor.Columns().AvatarUrl].String(),
		Intro:        record[dao.Doctor.Columns().Intro].String(),
		Status:       record[dao.Doctor.Columns().Status].Int(),
	}, nil
}

func (s doctorProfileService) UpdateProfile(ctx context.Context, in DoctorProfileUpdateInput) error {
	data := do.Doctor{}
	if in.Phone != nil {
		data.Phone = *in.Phone
	}
	if in.Email != nil {
		data.Email = *in.Email
	}
	if in.AvatarURL != nil {
		data.AvatarUrl = *in.AvatarURL
	}
	if in.Intro != nil {
		data.Intro = *in.Intro
	}
	if !hasDoctorProfileUpdates(in) {
		return consts.NewBadRequestError("至少提供一个更新字段")
	}
	data.UpdatedAt = time.Now()

	result, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, in.DoctorID).
		Data(data).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "更新医生资料失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("医生资料不存在")
	}
	return nil
}

func (s doctorProfileService) UpdatePassword(ctx context.Context, in DoctorProfileUpdatePasswordInput) error {
	record, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, in.DoctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询医生资料失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("医生资料不存在")
	}
	if !verifyPassword(in.OldPassword, record[dao.Doctor.Columns().PasswordHash].String()) {
		return consts.NewBadRequestError("旧密码错误")
	}

	result, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, in.DoctorID).
		Data(do.Doctor{
			PasswordHash: hashPassword(in.NewPassword),
			UpdatedAt:    time.Now(),
		}).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "修改医生密码失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("医生资料不存在")
	}
	return nil
}

func hasDoctorProfileUpdates(in DoctorProfileUpdateInput) bool {
	return in.Phone != nil || in.Email != nil || in.AvatarURL != nil || in.Intro != nil
}
