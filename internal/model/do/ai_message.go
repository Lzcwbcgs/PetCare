// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AiMessage is the golang structure of table ai_message for DAO operations like Where/Data.
type AiMessage struct {
	g.Meta           `orm:"table:ai_message, do:true"`
	Id               any // AI消息ID
	SessionId        any // AI会话ID
	SenderType       any // 发送者类型：1用户，2AI，3医生，4管理员
	SenderId         any // 发送者ID，可为空
	MessageContent   any // 消息内容
	MessageType      any // 消息类型：1文本，2系统事件，3结构化结果
	ProviderType     any // 模型类型，如api/local
	ProviderName     any // 模型提供方名称
	PromptTokens     any // 提示词token数
	CompletionTokens any // 回复token数
	FinishReason     any // 停止原因
	ExtraMetadata    any // 扩展元数据
	CreatedAt        any // 发送时间
}
