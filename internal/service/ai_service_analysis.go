package service

import (
	"context"
	"fmt"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
)

func (s aiService) ListAnalysisRecords(ctx context.Context, in AIAnalysisListInput) (*AIAnalysisListOutput, error) {
	_, err := aiLoadAccessibleSession(ctx, in.SessionID, in.RequesterUserID, in.RequesterRole)
	if err != nil {
		return nil, err
	}

	records, err := dao.AiAnalysisRecord.Ctx(ctx).
		Where(dao.AiAnalysisRecord.Columns().SessionId, in.SessionID).
		OrderDesc(dao.AiAnalysisRecord.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query analysis records failed")
	}

	items := make([]AIAnalysisItem, 0, len(records))
	for _, record := range records {
		var riskLevel *int
		if record[aiAnalysisColumnRiskLevel].Val() != nil {
			value := record[aiAnalysisColumnRiskLevel].Int()
			riskLevel = &value
		}

		items = append(items, AIAnalysisItem{
			ID:               record[dao.AiAnalysisRecord.Columns().Id].Int64(),
			AnalysisType:     record[dao.AiAnalysisRecord.Columns().AnalysisType].Int(),
			InputSource:      record[dao.AiAnalysisRecord.Columns().InputSource].Int(),
			SummaryTitle:     record[aiAnalysisColumnSummaryTitle].String(),
			AnalysisResult:   record[dao.AiAnalysisRecord.Columns().AnalysisResult].String(),
			RuleBasedResult:  record[aiAnalysisColumnRuleBased].String(),
			LlmBasedResult:   record[aiAnalysisColumnLlmBased].String(),
			RiskLevel:        riskLevel,
			ReviewedByDoctor: record[dao.AiAnalysisRecord.Columns().ReviewedByDoctor].Int(),
			CreatedAt:        formatGTime(record[dao.AiAnalysisRecord.Columns().CreatedAt].GTime()),
		})
	}

	RecordOperationLogByRole(
		ctx,
		in.RequesterRole,
		in.RequesterUserID,
		"ai_session",
		"list_analysis",
		fmt.Sprintf("list ai analysis session_id=%d", in.SessionID),
	)

	return &AIAnalysisListOutput{Items: items}, nil
}
