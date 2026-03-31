package service

import (
	"context"
	"net/url"
	"path"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

const (
	doctorMedicalReportUploadDir = "resource/public/uploads/reports"
	doctorMedicalReportURLPrefix = "/reports"
)

type (
	DoctorMedicalReportUploadInput struct {
		DoctorID        int64
		MedicalRecordID int64
		ReportTitle     string
		ReportType      string
		ReportContent   string
		File            *ghttp.UploadFile
	}

	DoctorMedicalReportUploadOutput struct {
		ID      int64
		FileURL string
	}

	DoctorMedicalReportListInput struct {
		DoctorID        int64
		MedicalRecordID int64
	}

	DoctorMedicalReportItem struct {
		ID            int64
		ReportTitle   string
		ReportType    string
		FileURL       string
		ReportContent string
		UploadedAt    string
	}

	DoctorMedicalReportDeleteInput struct {
		DoctorID int64
		ReportID int64
	}
)

type IDoctorMedicalReport interface {
	Upload(ctx context.Context, in DoctorMedicalReportUploadInput) (*DoctorMedicalReportUploadOutput, error)
	List(ctx context.Context, in DoctorMedicalReportListInput) ([]DoctorMedicalReportItem, error)
	Delete(ctx context.Context, in DoctorMedicalReportDeleteInput) error
}

var DoctorMedicalReport IDoctorMedicalReport = doctorMedicalReportService{}

type doctorMedicalReportService struct{}

func (s doctorMedicalReportService) Upload(ctx context.Context, in DoctorMedicalReportUploadInput) (*DoctorMedicalReportUploadOutput, error) {
	if err := ensureDoctorMedicalRecord(ctx, in.DoctorID, in.MedicalRecordID); err != nil {
		return nil, err
	}

	var (
		err     error
		fileURL string
	)
	if in.File != nil {
		var savedName string
		savedName, err = in.File.Save(doctorMedicalReportUploadDir, true)
		if err != nil {
			return nil, consts.WrapInternalError(err, "上传报告附件失败")
		}
		fileURL = doctorMedicalReportURLPrefix + "/" + savedName
	}

	now := time.Now()
	result, err := dao.MedicalReport.Ctx(ctx).Data(do.MedicalReport{
		MedicalRecordId: in.MedicalRecordID,
		DoctorId:        in.DoctorID,
		ReportTitle:     in.ReportTitle,
		ReportType:      in.ReportType,
		FileUrl:         fileURL,
		ReportContent:   in.ReportContent,
		UploadedAt:      now,
	}).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "上传医疗报告失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取报告ID失败")
	}
	return &DoctorMedicalReportUploadOutput{
		ID:      id,
		FileURL: fileURL,
	}, nil
}

func (s doctorMedicalReportService) List(ctx context.Context, in DoctorMedicalReportListInput) ([]DoctorMedicalReportItem, error) {
	if err := ensureDoctorMedicalRecord(ctx, in.DoctorID, in.MedicalRecordID); err != nil {
		return nil, err
	}

	records, err := dao.MedicalReport.Ctx(ctx).
		Where(dao.MedicalReport.Columns().MedicalRecordId, in.MedicalRecordID).
		Where(dao.MedicalReport.Columns().DoctorId, in.DoctorID).
		OrderDesc(dao.MedicalReport.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医疗报告列表失败")
	}

	items := make([]DoctorMedicalReportItem, 0, len(records))
	for _, record := range records {
		items = append(items, DoctorMedicalReportItem{
			ID:            record[dao.MedicalReport.Columns().Id].Int64(),
			ReportTitle:   record[dao.MedicalReport.Columns().ReportTitle].String(),
			ReportType:    record[dao.MedicalReport.Columns().ReportType].String(),
			FileURL:       record[dao.MedicalReport.Columns().FileUrl].String(),
			ReportContent: record[dao.MedicalReport.Columns().ReportContent].String(),
			UploadedAt:    record[dao.MedicalReport.Columns().UploadedAt].GTime().Format(appointmentTimeLayout),
		})
	}
	return items, nil
}

func (s doctorMedicalReportService) Delete(ctx context.Context, in DoctorMedicalReportDeleteInput) error {
	record, err := dao.MedicalReport.Ctx(ctx).
		Where(dao.MedicalReport.Columns().Id, in.ReportID).
		Where(dao.MedicalReport.Columns().DoctorId, in.DoctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询医疗报告失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("医疗报告不存在")
	}

	result, err := dao.MedicalReport.Ctx(ctx).
		Where(dao.MedicalReport.Columns().Id, in.ReportID).
		Where(dao.MedicalReport.Columns().DoctorId, in.DoctorID).
		Delete()
	if err != nil {
		return consts.WrapInternalError(err, "删除医疗报告失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取删除结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("医疗报告不存在")
	}

	filePath := doctorMedicalReportFilePath(record[dao.MedicalReport.Columns().FileUrl].String())
	if filePath != "" && gfile.Exists(filePath) {
		if err = gfile.Remove(filePath); err != nil {
			g.Log().Warningf(ctx, "remove medical report file failed: %v", err)
		}
	}
	return nil
}

func ensureDoctorMedicalRecord(ctx context.Context, doctorID int64, medicalRecordID int64) error {
	record, err := dao.MedicalRecord.Ctx(ctx).
		Where(dao.MedicalRecord.Columns().Id, medicalRecordID).
		Where(dao.MedicalRecord.Columns().DoctorId, doctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询病历信息失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("病历不存在")
	}
	return nil
}

func doctorMedicalReportFilePath(fileURL string) string {
	if strings.TrimSpace(fileURL) == "" {
		return ""
	}

	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return ""
	}

	reportPath := parsedURL.Path
	if !strings.HasPrefix(reportPath, doctorMedicalReportURLPrefix+"/") {
		return ""
	}
	return gfile.Join(doctorMedicalReportUploadDir, strings.TrimPrefix(reportPath, doctorMedicalReportURLPrefix+"/"))
}

func BuildDoctorMedicalReportFileURL(baseURL string, fileURL string) string {
	if strings.TrimSpace(fileURL) == "" {
		return ""
	}
	if strings.HasPrefix(fileURL, "http://") || strings.HasPrefix(fileURL, "https://") {
		return fileURL
	}
	if strings.TrimSpace(baseURL) == "" {
		return fileURL
	}
	return strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(path.Clean(fileURL), "/")
}
