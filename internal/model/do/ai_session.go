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
	ModelType      any // 模型类型，如本地模型/API模型
	ModelName      any // 模型名称
	SessionSummary any // AI会话总结
	SyncToAdmin    any // 是否同步给管理端：1是，0否
	Status         any // 状态：1进行中，2已结束，3已归档
	CreatedAt      any // 创建时间
	UpdatedAt      any // 更新时间
}
