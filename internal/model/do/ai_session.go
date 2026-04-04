// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AiSession is the golang structure of table ai_session for DAO operations like Where/Data.
type AiSession struct {
	g.Meta         `orm:"table:ai_session, do:true"`
	Id             any // AI会话ID
	SessionNo      any // 会话编号
	UserId         any // 用户ID
	PetId          any // 宠物ID
	HospitalId     any // 关联医院ID，可为空
	DoctorId       any // 关联医生ID，可为空
	SourceType     any // 来源：1用户端发起，2医生端发起
	ModelType      any // 模型类型，如api/local
	ModelName      any // 模型名称
	ProviderName   any // 模型提供方名称
	SessionSummary any // AI会话总结
	RagEnabled     any // 是否启用RAG：1是，0否
	RetrievalCount any // 最近一次召回片段数
	SyncToAdmin    any // 是否同步给管理端：1是，0否
	Status         any // 状态：1进行中，2已结束，3已归档
	LastMessageAt  any // 最后消息时间
	ExtraMetadata  any // 扩展元数据
	CreatedAt      any // 创建时间
	UpdatedAt      any // 更新时间
}
