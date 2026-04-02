package notification

import (
	"context"

	v1 "PetCare/api/notification/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireNotificationRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	output, err := service.Notification.List(ctx, service.NotificationListInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		Page:            derefPage(req.Page, 1),
		Size:            derefPage(req.PageSize, 10),
		Status:          req.Status,
	})
	if err != nil {
		return nil, err
	}

	list := make([]v1.ListItem, 0, len(output.Items))
	for _, item := range output.Items {
		list = append(list, v1.ListItem{
			ID:               item.ID,
			NotificationType: item.NotificationType,
			Title:            item.Title,
			Content:          item.Content,
			AppointmentID:    item.AppointmentID,
			SendTime:         item.SendTime,
			Status:           item.Status,
		})
	}

	return &v1.ListRes{
		List: list,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *Controller) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireNotificationRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	item, err := service.Notification.Detail(ctx, service.NotificationDetailInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		NotificationID:  req.NotificationID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.DetailRes{
		ID:               item.ID,
		NotificationType: item.NotificationType,
		Title:            item.Title,
		Content:          item.Content,
		AppointmentID:    item.AppointmentID,
		SendTime:         item.SendTime,
		Status:           item.Status,
	}, nil
}

func (c *Controller) Read(ctx context.Context, req *v1.ReadReq) (res *v1.ReadRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "已标记为已读")

	claims, err := requireNotificationRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	err = service.Notification.Read(ctx, service.NotificationReadInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		NotificationID:  req.NotificationID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.ReadRes{}, nil
}

func authClaimsFromCtx(ctx context.Context) (*service.AuthClaims, error) {
	claims, ok := g.RequestFromCtx(ctx).GetCtxVar(consts.CtxKeyAuthClaims).Val().(*service.AuthClaims)
	if !ok || claims == nil {
		return nil, consts.NewUnauthorizedError("")
	}
	return claims, nil
}

func requireNotificationRoles(ctx context.Context, roles ...string) (*service.AuthClaims, error) {
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
