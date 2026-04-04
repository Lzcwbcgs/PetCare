package admin

import (
	"context"

	v1 "PetCare/api/admin/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type OperationLogController struct{}

func NewOperationLog() *OperationLogController {
	return &OperationLogController{}
}

func (c *OperationLogController) List(ctx context.Context, req *v1.OperationLogListReq) (res *v1.OperationLogListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	output, err := service.AdminOperationLog.List(ctx, service.AdminOperationLogListInput{
		Page:            derefPage(req.Page, 1),
		Size:            derefPage(req.PageSize, 10),
		OperatorType:    req.OperatorType,
		OperationModule: req.OperationModule,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.OperationLogListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.OperationLogListItem{
			ID:              item.ID,
			OperatorType:    item.OperatorType,
			OperatorID:      item.OperatorID,
			OperationModule: item.OperationModule,
			OperationType:   item.OperationType,
			OperationDesc:   item.OperationDesc,
			IPAddress:       item.IPAddress,
			CreatedAt:       item.CreatedAt,
		})
	}

	return &v1.OperationLogListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *OperationLogController) Detail(ctx context.Context, req *v1.OperationLogDetailReq) (res *v1.OperationLogDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	item, err := service.AdminOperationLog.Detail(ctx, service.AdminOperationLogDetailInput{LogID: req.LogID})
	if err != nil {
		return nil, err
	}

	return &v1.OperationLogDetailRes{
		ID:              item.ID,
		OperatorType:    item.OperatorType,
		OperatorID:      item.OperatorID,
		OperationModule: item.OperationModule,
		OperationType:   item.OperationType,
		OperationDesc:   item.OperationDesc,
		IPAddress:       item.IPAddress,
		CreatedAt:       item.CreatedAt,
	}, nil
}

func derefPage(value *int, fallback int) int {
	if value == nil {
		return fallback
	}
	return *value
}
