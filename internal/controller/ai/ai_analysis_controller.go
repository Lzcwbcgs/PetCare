package ai

import (
	"context"

	v1 "PetCare/api/ai/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

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
