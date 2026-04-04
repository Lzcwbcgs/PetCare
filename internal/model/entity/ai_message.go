// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// AiMessage is the golang structure for table ai_message.
type AiMessage struct {
	Id               int64     `json:"id"               orm:"id"                description:"AI消息ID"`                 // AI消息ID
	SessionId        int64     `json:"sessionId"        orm:"session_id"        description:"AI会话ID"`                 // AI会话ID
	SenderType       int       `json:"senderType"       orm:"sender_type"       description:"发送者类型：1用户，2AI，3医生，4管理员"` // 发送者类型：1用户，2AI，3医生，4管理员
	SenderId         int64     `json:"senderId"         orm:"sender_id"         description:"发送者ID，可为空"`              // 发送者ID，可为空
	MessageContent   string    `json:"messageContent"   orm:"message_content"   description:"消息内容"`                   // 消息内容
	MessageType      int       `json:"messageType"      orm:"message_type"      description:"消息类型：1文本，2系统事件，3结构化结果"`  // 消息类型：1文本，2系统事件，3结构化结果
	ProviderType     string    `json:"providerType"     orm:"provider_type"     description:"模型类型，如api/local"`      // 模型类型，如api/local
	ProviderName     string    `json:"providerName"     orm:"provider_name"     description:"模型提供方名称"`              // 模型提供方名称
	PromptTokens     int       `json:"promptTokens"     orm:"prompt_tokens"     description:"提示词token数"`              // 提示词token数
	CompletionTokens int       `json:"completionTokens" orm:"completion_tokens" description:"回复token数"`               // 回复token数
	FinishReason     string    `json:"finishReason"     orm:"finish_reason"     description:"停止原因"`                  // 停止原因
	ExtraMetadata    string    `json:"extraMetadata"    orm:"extra_metadata"    description:"扩展元数据"`                 // 扩展元数据
	CreatedAt        time.Time `json:"createdAt"        orm:"created_at"        description:"发送时间"`                   // 发送时间
}
