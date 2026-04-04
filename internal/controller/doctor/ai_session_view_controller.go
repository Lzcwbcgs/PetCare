package doctor

import (
	"context"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *AIController) SessionList(ctx context.Context, req *v1.AISessionListReq) (res *v1.AISessionListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.AI.ListDoctorSessions(ctx, service.AIDoctorSessionListInput{
		DoctorID: claims.UserID,
		Page:     derefPage(req.Page, 1),
		Size:     derefPage(req.PageSize, 10),
		PetID:    req.PetID,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.AISessionListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.AISessionListItem{
			ID:        item.ID,
			SessionNo: item.SessionNo,
			PetID:     item.PetID,
			PetName:   item.PetName,
			ModelName: item.ModelName,
			Status:    item.Status,
			CreatedAt: item.CreatedAt,
		})
	}

	return &v1.AISessionListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}
