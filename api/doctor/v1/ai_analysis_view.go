package v1

import "github.com/gogf/gf/v2/frame/g"

type AISessionAnalysisListReq struct {
	g.Meta    `path:"/sessions/{session_id}/analysis-records" method:"get" tags:"医生AI查看" summary:"获取AI分析结果"`
	SessionID int64 `json:"session_id" p:"session_id" v:"required|min:1#会话ID不能为空|会话ID不合法" dc:"会话ID"`
}

type AISessionAnalysisItem struct {
	ID               int64  `json:"id" dc:"分析ID"`
	AnalysisType     int    `json:"analysis_type" dc:"分析类型"`
	InputSource      int    `json:"input_source" dc:"输入来源"`
	AnalysisResult   string `json:"analysis_result" dc:"分析结果"`
	RuleBasedResult  string `json:"rule_based_result" dc:"规则分析结果"`
	LlmBasedResult   string `json:"llm_based_result" dc:"模型分析结果"`
	RiskLevel        *int   `json:"risk_level" dc:"风险等级"`
	ReviewedByDoctor int    `json:"reviewed_by_doctor" dc:"是否已审核"`
	CreatedAt        string `json:"created_at" dc:"创建时间"`
}

type AISessionAnalysisListRes struct {
	List []AISessionAnalysisItem `json:"list" dc:"分析记录列表"`
}
