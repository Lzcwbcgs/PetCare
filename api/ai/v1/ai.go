package v1

import "github.com/gogf/gf/v2/frame/g"

type Pagination struct {
	Page     int `json:"page" dc:"Current page number"`
	PageSize int `json:"page_size" dc:"Page size"`
	Total    int `json:"total" dc:"Total records"`
}

type SessionCreateReq struct {
	g.Meta     `path:"/sessions" method:"post" tags:"AI Consultation" summary:"Create AI session"`
	PetID      int64  `json:"pet_id" v:"required|min:1#Pet ID is required|Invalid pet ID" dc:"Pet ID"`
	HospitalID *int64 `json:"hospital_id" v:"min:1#Invalid hospital ID" dc:"Hospital ID (optional)"`
	DoctorID   *int64 `json:"doctor_id" v:"min:1#Invalid doctor ID" dc:"Doctor ID (optional)"`
	ModelType  string `json:"model_type" v:"in:api,local#Invalid model_type" dc:"Model provider type: api/local"`
	ModelName  string `json:"model_name" v:"max-length:100#Model name is too long" dc:"Model name"`
	RagEnabled *int   `json:"rag_enabled" v:"in:0,1#rag_enabled must be 0 or 1" dc:"Enable RAG: 1 yes, 0 no"`
}

type SessionCreateRes struct {
	SessionID int64  `json:"session_id" dc:"Session ID"`
	SessionNo string `json:"session_no" dc:"Session number"`
	Status    int    `json:"status" dc:"Session status"`
}

type SessionSendMessageReq struct {
	g.Meta         `path:"/sessions/{session_id}/messages" method:"post" tags:"AI Consultation" summary:"Send message (SSE stream)"`
	SessionID      int64  `json:"session_id" p:"session_id" v:"required|min:1#Session ID is required|Invalid session ID" dc:"Session ID"`
	MessageContent string `json:"message_content" v:"required|max-length:65535#Message content is required|Message content is too long" dc:"Message content"`
	MessageType    int    `json:"message_type" v:"required|in:1#message_type must be 1" dc:"Message type, first version supports only text(1)"`
}

type SessionSendMessageRes struct{}

type SessionDetailReq struct {
	g.Meta    `path:"/sessions/{session_id}" method:"get" tags:"AI Consultation" summary:"Get session detail"`
	SessionID int64 `json:"session_id" p:"session_id" v:"required|min:1#Session ID is required|Invalid session ID" dc:"Session ID"`
}

type SessionDetailRes struct {
	ID             int64  `json:"id" dc:"Session ID"`
	SessionNo      string `json:"session_no" dc:"Session number"`
	PetID          int64  `json:"pet_id" dc:"Pet ID"`
	SourceType     int    `json:"source_type" dc:"Source type"`
	ModelType      string `json:"model_type" dc:"Model type"`
	ModelName      string `json:"model_name" dc:"Model name"`
	ProviderName   string `json:"provider_name" dc:"Provider name"`
	SessionSummary string `json:"session_summary" dc:"Session summary"`
	RagEnabled     int    `json:"rag_enabled" dc:"Whether RAG is enabled"`
	Status         int    `json:"status" dc:"Session status"`
	LastMessageAt  string `json:"last_message_at" dc:"Last message time"`
	CreatedAt      string `json:"created_at" dc:"Created time"`
	UpdatedAt      string `json:"updated_at" dc:"Updated time"`
}

type SessionMessageListReq struct {
	g.Meta    `path:"/sessions/{session_id}/messages" method:"get" tags:"AI Consultation" summary:"Get session messages"`
	SessionID int64 `json:"session_id" p:"session_id" v:"required|min:1#Session ID is required|Invalid session ID" dc:"Session ID"`
	Page      *int  `json:"page" p:"page" v:"min:1#Page must be greater than 0" dc:"Page number"`
	PageSize  *int  `json:"page_size" p:"page_size" v:"min:1|max:100#Page size must be between 1 and 100" dc:"Page size"`
}

type SessionMessageItem struct {
	ID             int64  `json:"id" dc:"Message ID"`
	SenderType     int    `json:"sender_type" dc:"Sender type"`
	SenderID       *int64 `json:"sender_id" dc:"Sender ID"`
	MessageContent string `json:"message_content" dc:"Message content"`
	MessageType    int    `json:"message_type" dc:"Message type"`
	ProviderType   string `json:"provider_type" dc:"Provider type"`
	ProviderName   string `json:"provider_name" dc:"Provider name"`
	FinishReason   string `json:"finish_reason" dc:"Finish reason"`
	CreatedAt      string `json:"created_at" dc:"Created time"`
}

type SessionMessageListRes struct {
	List       []SessionMessageItem `json:"list" dc:"Message list"`
	Pagination Pagination           `json:"pagination" dc:"Pagination info"`
}

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
