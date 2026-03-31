package doctor

import (
	"context"
	"strings"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type MedicalReportController struct{}

func NewReport() *MedicalReportController {
	return &MedicalReportController{}
}

func (c *MedicalReportController) Upload(ctx context.Context, req *v1.MedicalReportUploadReq) (res *v1.MedicalReportUploadRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "上传成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.DoctorMedicalReport.Upload(ctx, service.DoctorMedicalReportUploadInput{
		DoctorID:        claims.UserID,
		MedicalRecordID: req.MedicalRecordID,
		ReportTitle:     req.ReportTitle,
		ReportType:      req.ReportType,
		ReportContent:   req.ReportContent,
		File:            req.File,
	})
	if err != nil {
		return nil, err
	}

	baseURL := doctorRequestBaseURL(ctx)

	return &v1.MedicalReportUploadRes{
		ReportID: output.ID,
		FileURL:  service.BuildDoctorMedicalReportFileURL(baseURL, output.FileURL),
	}, nil
}

func (c *MedicalReportController) List(ctx context.Context, req *v1.MedicalReportListReq) (res *v1.MedicalReportListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	items, err := service.DoctorMedicalReport.List(ctx, service.DoctorMedicalReportListInput{
		DoctorID:        claims.UserID,
		MedicalRecordID: req.MedicalRecordID,
	})
	if err != nil {
		return nil, err
	}

	baseURL := doctorRequestBaseURL(ctx)
	list := make([]v1.MedicalReportListItem, 0, len(items))
	for _, item := range items {
		list = append(list, v1.MedicalReportListItem{
			ID:            item.ID,
			ReportTitle:   item.ReportTitle,
			ReportType:    item.ReportType,
			FileURL:       service.BuildDoctorMedicalReportFileURL(baseURL, item.FileURL),
			ReportContent: item.ReportContent,
			UploadedAt:    item.UploadedAt,
		})
	}
	return &v1.MedicalReportListRes{List: list}, nil
}

func (c *MedicalReportController) Delete(ctx context.Context, req *v1.MedicalReportDeleteReq) (res *v1.MedicalReportDeleteRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "删除成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = service.DoctorMedicalReport.Delete(ctx, service.DoctorMedicalReportDeleteInput{
		DoctorID: claims.UserID,
		ReportID: req.ReportID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.MedicalReportDeleteRes{}, nil
}

func doctorRequestBaseURL(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}
	return strings.TrimRight(r.GetSchema()+"://"+r.Host, "/")
}
