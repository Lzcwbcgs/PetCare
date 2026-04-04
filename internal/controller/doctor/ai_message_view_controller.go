package doctor

import (
	"context"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *AIController) MessageList(ctx context.Context, req *v1.AISessionMessageListReq) (res *v1.AISessionMessageListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.AI.ListMessages(ctx, service.AIMessageListInput{
		SessionID:       req.SessionID,
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		Page:            derefPage(req.Page, 1),
		Size:            derefPage(req.PageSize, 20),
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.AISessionMessageItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.AISessionMessageItem{
			ID:             item.ID,
			SenderType:     item.SenderType,
			SenderID:       item.SenderID,
			MessageContent: item.MessageContent,
			CreatedAt:      item.CreatedAt,
		})
	}

	return &v1.AISessionMessageListRes{List: items}, nil
}
