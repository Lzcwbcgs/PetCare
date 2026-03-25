// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AiMessage is the golang structure of table ai_message for DAO operations like Where/Data.
type AiMessage struct {
	g.Meta         `orm:"table:ai_message, do:true"`
	Id             any // AI消息ID
	SessionId      any // AI会话ID
	SenderType     any // 发送者类型：1用户，2AI，3医生，4管理员
	SenderId       any // 发送者ID，可为空
	MessageContent any // 消息内容
	MessageType    any // 消息类型：1文本
	CreatedAt      any // 发送时间
}
