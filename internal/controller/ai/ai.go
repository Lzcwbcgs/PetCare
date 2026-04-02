package ai

import (
	"context"
	"encoding/json"
	"net/http"

	v1 "PetCare/api/ai/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) CreateSession(ctx context.Context, req *v1.SessionCreateReq) (res *v1.SessionCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "Session created successfully")

	claims, err := requireAIRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	output, err := service.AI.CreateSession(ctx, service.AISessionCreateInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		PetID:           req.PetID,
		HospitalID:      req.HospitalID,
		DoctorID:        req.DoctorID,
		ModelType:       req.ModelType,
		ModelName:       req.ModelName,
		RagEnabled:      req.RagEnabled,
	})
	if err != nil {
		return nil, err
	}

	return &v1.SessionCreateRes{
		SessionID: output.SessionID,
		SessionNo: output.SessionNo,
		Status:    output.Status,
	}, nil
}

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

func (c *Controller) Detail(ctx context.Context, req *v1.SessionDetailReq) (res *v1.SessionDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireAIRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	output, err := service.AI.Detail(ctx, service.AISessionDetailInput{
		SessionID:       req.SessionID,
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
	})
	if err != nil {
		return nil, err
	}

	return &v1.SessionDetailRes{
		ID:             output.ID,
		SessionNo:      output.SessionNo,
		PetID:          output.PetID,
		SourceType:     output.SourceType,
		ModelType:      output.ModelType,
		ModelName:      output.ModelName,
		ProviderName:   output.ProviderName,
		SessionSummary: output.SessionSummary,
		RagEnabled:     output.RagEnabled,
		Status:         output.Status,
		LastMessageAt:  output.LastMessageAt,
		CreatedAt:      output.CreatedAt,
		UpdatedAt:      output.UpdatedAt,
	}, nil
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

func (c *Controller) ListAnalysisRecords(ctx context.Context, req *v1.SessionAnalysisListReq) (res *v1.SessionAnalysisListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireAIRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	output, err := service.AI.ListAnalysisRecords(ctx, service.AIAnalysisListInput{
		SessionID:       req.SessionID,
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.SessionAnalysisItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.SessionAnalysisItem{
			ID:               item.ID,
			AnalysisType:     item.AnalysisType,
			InputSource:      item.InputSource,
			SummaryTitle:     item.SummaryTitle,
			AnalysisResult:   item.AnalysisResult,
			RiskLevel:        item.RiskLevel,
			ReviewedByDoctor: item.ReviewedByDoctor,
			CreatedAt:        item.CreatedAt,
		})
	}

	return &v1.SessionAnalysisListRes{
		List: items,
	}, nil
}

func writeSSEEvent(r *ghttp.Request, event string, payload any) {
	data, _ := json.Marshal(payload)
	r.Response.Writef("event: %s\n", event)
	r.Response.Writef("data: %s\n\n", string(data))
	r.Response.Flush()
}
