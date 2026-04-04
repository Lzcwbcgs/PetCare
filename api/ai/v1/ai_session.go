package v1

import "github.com/gogf/gf/v2/frame/g"

type SessionCreateReq struct {
	g.Meta     `path:"/sessions" method:"post" tags:"AI Consultation" summary:"Create AI session"`
	PetID      int64  `json:"pet_id" v:"required|min:1#Pet ID is required|Invalid pet ID" dc:"Pet ID"`
	HospitalID *int64 `json:"hospital_id" v:"min:1#Invalid hospital ID" dc:"Hospital ID (optional)"`
	DoctorID   *int64 `json:"doctor_id" v:"min:1#Invalid doctor ID" dc:"Doctor ID (optional)"`
	ModelType  string `json:"model_type" v:"in:api,local#Invalid model_type" dc:"Model provider type: api/local"`
	ModelName  string `json:"model_name" v:"max-length:100#Model name is too long" dc:"Model name"`
}

type SessionCreateRes struct {
	SessionID int64  `json:"session_id" dc:"Session ID"`
	SessionNo string `json:"session_no" dc:"Session number"`
	Status    int    `json:"status" dc:"Session status"`
}

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
	SessionSummary string `json:"session_summary" dc:"Session summary"`
	Status         int    `json:"status" dc:"Session status"`
	CreatedAt      string `json:"created_at" dc:"Created time"`
	UpdatedAt      string `json:"updated_at" dc:"Updated time"`
}
