// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// AiMessage is the golang structure for table ai_message.
type AiMessage struct {
	Id             int64     `json:"id"             orm:"id"              description:"AI消息ID"`                 // AI消息ID
	SessionId      int64     `json:"sessionId"      orm:"session_id"      description:"AI会话ID"`                 // AI会话ID
	SenderType     int       `json:"senderType"     orm:"sender_type"     description:"发送者类型：1用户，2AI，3医生，4管理员"` // 发送者类型：1用户，2AI，3医生，4管理员
	SenderId       int64     `json:"senderId"       orm:"sender_id"       description:"发送者ID，可为空"`              // 发送者ID，可为空
	MessageContent string    `json:"messageContent" orm:"message_content" description:"消息内容"`                   // 消息内容
	MessageType    int       `json:"messageType"    orm:"message_type"    description:"消息类型：1文本"`               // 消息类型：1文本
	CreatedAt      time.Time `json:"createdAt"      orm:"created_at"      description:"发送时间"`                   // 发送时间
}
