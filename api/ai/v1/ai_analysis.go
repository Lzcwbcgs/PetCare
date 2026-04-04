package v1

import "github.com/gogf/gf/v2/frame/g"

type SessionAnalysisListReq struct {
	g.Meta    `path:"/sessions/{session_id}/analysis-records" method:"get" tags:"AI Consultation" summary:"Get session analysis records"`
	SessionID int64 `json:"session_id" p:"session_id" v:"required|min:1#Session ID is required|Invalid session ID" dc:"Session ID"`
}

type SessionAnalysisItem struct {
	ID               int64  `json:"id" dc:"Analysis record ID"`
	AnalysisType     int    `json:"analysis_type" dc:"Analysis type"`
	InputSource      int    `json:"input_source" dc:"Input source"`
	SummaryTitle     string `json:"summary_title" dc:"Summary title"`
	AnalysisResult   string `json:"analysis_result" dc:"Analysis result"`
	RiskLevel        *int   `json:"risk_level" dc:"Risk level"`
	ReviewedByDoctor int    `json:"reviewed_by_doctor" dc:"Reviewed by doctor"`
	CreatedAt        string `json:"created_at" dc:"Created time"`
}

type SessionAnalysisListRes struct {
	List []SessionAnalysisItem `json:"list" dc:"Analysis record list"`
}
