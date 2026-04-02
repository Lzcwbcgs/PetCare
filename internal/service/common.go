package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

const (
	commonUploadRootDir = "resource/public/uploads"
	commonUploadURLRoot = "/uploads"
)

type (
	CommonHospitalOptionItem struct {
		ID           int64
		HospitalName string
	}

	CommonHospitalDoctorOptionsInput struct {
		HospitalID int64
	}

	CommonHospitalDoctorOptionItem struct {
		ID         int64
		DoctorName string
		Title      string
	}

	CommonUploadInput struct {
		File    *ghttp.UploadFile
		BizType *string
	}

	CommonUploadOutput struct {
		FileURL  string
		FileName string
	}
)

type ICommon interface {
	HospitalOptions(ctx context.Context) ([]CommonHospitalOptionItem, error)
	HospitalDoctorOptions(ctx context.Context, in CommonHospitalDoctorOptionsInput) ([]CommonHospitalDoctorOptionItem, error)
	Upload(ctx context.Context, in CommonUploadInput) (*CommonUploadOutput, error)
}

var Common ICommon = commonService{}

type commonService struct{}

func (s commonService) HospitalOptions(ctx context.Context) ([]CommonHospitalOptionItem, error) {
	records, err := dao.Hospital.Ctx(ctx).
		Where(dao.Hospital.Columns().Status, 1).
		OrderAsc(dao.Hospital.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医院下拉数据失败")
	}

	items := make([]CommonHospitalOptionItem, 0, len(records))
	for _, record := range records {
		items = append(items, CommonHospitalOptionItem{
			ID:           record[dao.Hospital.Columns().Id].Int64(),
			HospitalName: record[dao.Hospital.Columns().HospitalName].String(),
		})
	}
	return items, nil
}

func (s commonService) HospitalDoctorOptions(ctx context.Context, in CommonHospitalDoctorOptionsInput) ([]CommonHospitalDoctorOptionItem, error) {
	records, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().HospitalId, in.HospitalID).
		Where(dao.Doctor.Columns().Status, 1).
		OrderAsc(dao.Doctor.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医院医生下拉数据失败")
	}

	items := make([]CommonHospitalDoctorOptionItem, 0, len(records))
	for _, record := range records {
		items = append(items, CommonHospitalDoctorOptionItem{
			ID:         record[dao.Doctor.Columns().Id].Int64(),
			DoctorName: record[dao.Doctor.Columns().DoctorName].String(),
			Title:      record[dao.Doctor.Columns().Title].String(),
		})
	}
	return items, nil
}

func (s commonService) Upload(ctx context.Context, in CommonUploadInput) (*CommonUploadOutput, error) {
	if in.File == nil {
		return nil, consts.NewBadRequestError("请上传文件")
	}
	if in.BizType != nil {
		_ = strings.TrimSpace(*in.BizType)
	}

	now := time.Now()
	relativeDir := fmt.Sprintf("%04d/%02d/%02d", now.Year(), now.Month(), now.Day())
	saveDir := gfile.Join(commonUploadRootDir, relativeDir)

	fileName, err := in.File.Save(saveDir, true)
	if err != nil {
		return nil, consts.WrapInternalError(err, "上传文件失败")
	}

	return &CommonUploadOutput{
		FileURL:  commonUploadURLRoot + "/" + strings.ReplaceAll(relativeDir, "\\", "/") + "/" + fileName,
		FileName: fileName,
	}, nil
}
