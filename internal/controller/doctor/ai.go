package doctor

import (
	"context"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type AIController struct{}

func NewAI() *AIController {
	return &AIController{}
}

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

func (c *AIController) AnalysisList(ctx context.Context, req *v1.AISessionAnalysisListReq) (res *v1.AISessionAnalysisListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
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

	items := make([]v1.AISessionAnalysisItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.AISessionAnalysisItem{
			ID:               item.ID,
			AnalysisType:     item.AnalysisType,
			InputSource:      item.InputSource,
			AnalysisResult:   item.AnalysisResult,
			RuleBasedResult:  item.RuleBasedResult,
			LlmBasedResult:   item.LlmBasedResult,
			RiskLevel:        item.RiskLevel,
			ReviewedByDoctor: item.ReviewedByDoctor,
			CreatedAt:        item.CreatedAt,
		})
	}

	return &v1.AISessionAnalysisListRes{List: items}, nil
}

func derefPage(value *int, fallback int) int {
	if value == nil {
		return fallback
	}
	return *value
}
