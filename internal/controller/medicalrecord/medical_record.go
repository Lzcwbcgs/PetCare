package medicalrecord

import (
	"context"
	"strings"

	v1 "PetCare/api/medicalrecord/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) List(ctx context.Context, req *v1.MedicalRecordListReq) (res *v1.MedicalRecordListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireMedicalRecordRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	output, err := service.MedicalRecordView.List(ctx, service.MedicalRecordViewListInput{
		UserID: claims.UserID,
		Page:   derefPage(req.Page, 1),
		Size:   derefPage(req.PageSize, 10),
		PetID:  req.PetID,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.MedicalRecordListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.MedicalRecordListItem{
			ID:                   item.ID,
			AppointmentID:        item.AppointmentID,
			PetID:                item.PetID,
			DoctorID:             item.DoctorID,
			DoctorName:           item.DoctorName,
			PreliminaryDiagnosis: item.PreliminaryDiagnosis,
			VisitTime:            item.VisitTime,
			Status:               item.Status,
		})
	}

	return &v1.MedicalRecordListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *Controller) Detail(ctx context.Context, req *v1.MedicalRecordDetailReq) (res *v1.MedicalRecordDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireMedicalRecordRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	item, err := service.MedicalRecordView.Detail(ctx, service.MedicalRecordViewDetailInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		MedicalRecordID: req.MedicalRecordID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.MedicalRecordDetailRes{
		ID:                   item.ID,
		AppointmentID:        item.AppointmentID,
		PetID:                item.PetID,
		UserID:               item.UserID,
		DoctorID:             item.DoctorID,
		ChiefComplaint:       item.ChiefComplaint,
		PresentHistory:       item.PresentHistory,
		PhysicalExamination:  item.PhysicalExamination,
		PreliminaryDiagnosis: item.PreliminaryDiagnosis,
		TreatmentPlan:        item.TreatmentPlan,
		Prescription:         item.Prescription,
		DoctorAdvice:         item.DoctorAdvice,
		VisitTime:            item.VisitTime,
		Status:               item.Status,
	}, nil
}

func (c *Controller) ReportList(ctx context.Context, req *v1.MedicalRecordReportListReq) (res *v1.MedicalRecordReportListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireMedicalRecordRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	items, err := service.MedicalRecordView.ReportList(ctx, service.MedicalRecordViewReportListInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		MedicalRecordID: req.MedicalRecordID,
	})
	if err != nil {
		return nil, err
	}

	baseURL := medicalRecordBaseURL(ctx)
	list := make([]v1.MedicalRecordReportListItem, 0, len(items))
	for _, item := range items {
		list = append(list, v1.MedicalRecordReportListItem{
			ID:            item.ID,
			ReportTitle:   item.ReportTitle,
			ReportType:    item.ReportType,
			FileURL:       service.BuildDoctorMedicalReportFileURL(baseURL, item.FileURL),
			ReportContent: item.ReportContent,
			UploadedAt:    item.UploadedAt,
		})
	}

	return &v1.MedicalRecordReportListRes{List: list}, nil
}

func authClaimsFromCtx(ctx context.Context) (*service.AuthClaims, error) {
	claims, ok := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxKeyAuthClaims).Val().(*service.AuthClaims)
	if !ok || claims == nil {
		return nil, consts.NewUnauthorizedError("")
	}
	return claims, nil
}

func requireMedicalRecordRoles(ctx context.Context, roles ...string) (*service.AuthClaims, error) {
	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	if !claims.HasRole(roles...) {
		return nil, consts.NewForbiddenError("")
	}
	return claims, nil
}

func derefPage(value *int, fallback int) int {
	if value == nil {
		return fallback
	}
	return *value
}

func medicalRecordBaseURL(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}
	return strings.TrimRight(r.GetSchema()+"://"+r.Host, "/")
}
