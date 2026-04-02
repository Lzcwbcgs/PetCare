package doctor

import (
	"context"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type NotificationController struct{}

func NewNotification() *NotificationController {
	return &NotificationController{}
}

func (c *NotificationController) List(ctx context.Context, req *v1.NotificationListReq) (res *v1.NotificationListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.DoctorNotification.List(ctx, service.DoctorNotificationListInput{
		DoctorID: claims.UserID,
		Page:     derefDoctorPage(req.Page, 1),
		Size:     derefDoctorPage(req.PageSize, 10),
	})
	if err != nil {
		return nil, err
	}

	list := make([]v1.NotificationListItem, 0, len(output.Items))
	for _, item := range output.Items {
		list = append(list, v1.NotificationListItem{
			ID:               item.ID,
			NotificationType: item.NotificationType,
			Title:            item.Title,
			Content:          item.Content,
			AppointmentID:    item.AppointmentID,
			SendTime:         item.SendTime,
			Status:           item.Status,
		})
	}

	return &v1.NotificationListRes{
		List: list,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *NotificationController) Read(ctx context.Context, req *v1.NotificationReadReq) (res *v1.NotificationReadRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "已标记为已读")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = service.DoctorNotification.Read(ctx, service.DoctorNotificationReadInput{
		DoctorID:       claims.UserID,
		NotificationID: req.NotificationID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.NotificationReadRes{}, nil
}

func derefDoctorPage(value *int, fallback int) int {
	if value == nil {
		return fallback
	}
	return *value
}
