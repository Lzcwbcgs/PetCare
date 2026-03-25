// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// AiSession is the golang structure for table ai_session.
type AiSession struct {
	Id             int64     `json:"id"             orm:"id"              description:"AI会话ID"`            // AI会话ID
	SessionNo      string    `json:"sessionNo"      orm:"session_no"      description:"会话编号"`              // 会话编号
	UserId         int64     `json:"userId"         orm:"user_id"         description:"用户ID"`              // 用户ID
	PetId          int64     `json:"petId"          orm:"pet_id"          description:"宠物ID"`              // 宠物ID
	HospitalId     int64     `json:"hospitalId"     orm:"hospital_id"     description:"关联医院ID，可为空"`        // 关联医院ID，可为空
	DoctorId       int64     `json:"doctorId"       orm:"doctor_id"       description:"关联医生ID，可为空"`        // 关联医生ID，可为空
	SourceType     int       `json:"sourceType"     orm:"source_type"     description:"来源：1用户端发起，2医生端发起"`  // 来源：1用户端发起，2医生端发起
	ModelType      string    `json:"modelType"      orm:"model_type"      description:"模型类型，如本地模型/API模型"`  // 模型类型，如本地模型/API模型
	ModelName      string    `json:"modelName"      orm:"model_name"      description:"模型名称"`              // 模型名称
	SessionSummary string    `json:"sessionSummary" orm:"session_summary" description:"AI会话总结"`            // AI会话总结
	SyncToAdmin    int       `json:"syncToAdmin"    orm:"sync_to_admin"   description:"是否同步给管理端：1是，0否"`    // 是否同步给管理端：1是，0否
	Status         int       `json:"status"         orm:"status"          description:"状态：1进行中，2已结束，3已归档"` // 状态：1进行中，2已结束，3已归档
	CreatedAt      time.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`              // 创建时间
	UpdatedAt      time.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`              // 更新时间
}
