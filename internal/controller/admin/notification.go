package admin

import (
	"context"

	v1 "PetCare/api/admin/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type NotificationController struct{}

func NewNotification() *NotificationController {
	return &NotificationController{}
}

func (c *NotificationController) Create(ctx context.Context, req *v1.NotificationCreateReq) (res *v1.NotificationCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "发布成功")

	output, err := service.AdminNotification.Create(ctx, service.AdminNotificationCreateInput{
		ReceiverType:     req.ReceiverType,
		ReceiverIDs:      req.ReceiverIDs,
		NotificationType: req.NotificationType,
		Title:            req.Title,
		Content:          req.Content,
		SendTime:         req.SendTime,
	})
	if err != nil {
		return nil, err
	}
	return &v1.NotificationCreateRes{Count: output.Count}, nil
}

func (c *NotificationController) List(ctx context.Context, req *v1.NotificationListReq) (res *v1.NotificationListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	var (
		page int
		size int
	)
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		size = *req.PageSize
	}

	output, err := service.AdminNotification.List(ctx, service.AdminNotificationListInput{
		Page:             page,
		Size:             size,
		NotificationType: req.NotificationType,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.NotificationListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.NotificationListItem{
			ID:               item.ID,
			NotificationType: item.NotificationType,
			Title:            item.Title,
			Content:          item.Content,
			SendTime:         item.SendTime,
			Status:           item.Status,
			CreatedAt:        item.CreatedAt,
		})
	}

	return &v1.NotificationListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *NotificationController) Delete(ctx context.Context, req *v1.NotificationDeleteReq) (res *v1.NotificationDeleteRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "撤回成功")

	err = service.AdminNotification.Delete(ctx, service.AdminNotificationDeleteInput{
		NotificationID: req.NotificationID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.NotificationDeleteRes{}, nil
}
