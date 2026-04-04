package ai

import (
	"context"
	"net/http"

	v1 "PetCare/api/ai/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *Controller) SendMessage(ctx context.Context, req *v1.SessionSendMessageReq) (res *v1.SessionSendMessageRes, err error) {
	claims, err := requireAIRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("X-Accel-Buffering", "no")
	r.Response.WriteStatus(http.StatusOK)

	err = service.AI.SendMessageStream(ctx, service.AISendMessageInput{
		SessionID:      req.SessionID,
		SenderUserID:   claims.UserID,
		SenderRole:     claims.Role,
		MessageContent: req.MessageContent,
		MessageType:    req.MessageType,
	}, func(event service.AIStreamEvent) {
		writeSSEEvent(r, event.Type, event.Payload)
	})
	if err != nil {
		if r.Response.BufferLength() > 0 || r.Response.BytesWritten() > 0 {
			return nil, nil
		}
		return nil, err
	}
	return nil, nil
}

func (c *Controller) ListMessages(ctx context.Context, req *v1.SessionMessageListReq) (res *v1.SessionMessageListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireAIRoles(ctx, consts.RoleUser, consts.RoleDoctor)
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

	items := make([]v1.SessionMessageItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.SessionMessageItem{
			ID:             item.ID,
			SenderType:     item.SenderType,
			SenderID:       item.SenderID,
			MessageContent: item.MessageContent,
			MessageType:    item.MessageType,
			ProviderType:   item.ProviderType,
			ProviderName:   item.ProviderName,
			FinishReason:   item.FinishReason,
			CreatedAt:      item.CreatedAt,
		})
	}

	return &v1.SessionMessageListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}
