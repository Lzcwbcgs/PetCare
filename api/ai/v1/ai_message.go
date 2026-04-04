package v1

import "github.com/gogf/gf/v2/frame/g"

type SessionSendMessageReq struct {
	g.Meta         `path:"/sessions/{session_id}/messages" method:"post" tags:"AI Consultation" summary:"Send message (SSE stream)"`
	SessionID      int64  `json:"session_id" p:"session_id" v:"required|min:1#Session ID is required|Invalid session ID" dc:"Session ID"`
	MessageContent string `json:"message_content" v:"required|max-length:65535#Message content is required|Message content is too long" dc:"Message content"`
	MessageType    int    `json:"message_type" v:"required|in:1#message_type must be 1" dc:"Message type, first version supports only text(1)"`
}

type SessionSendMessageRes struct{}

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
