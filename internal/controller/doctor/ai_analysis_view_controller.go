package doctor

import (
	"context"

	v1 "PetCare/api/doctor/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

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
