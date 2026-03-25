// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// PetMedicalHistory is the golang structure for table pet_medical_history.
type PetMedicalHistory struct {
	Id          int64     `json:"id"          orm:"id"           description:"病史记录ID"`        // 病史记录ID
	PetId       int64     `json:"petId"       orm:"pet_id"       description:"宠物ID"`          // 宠物ID
	HistoryType string    `json:"historyType" orm:"history_type" description:"病史类型"`          // 病史类型
	Description string    `json:"description" orm:"description"  description:"病史描述"`          // 病史描述
	DiagnosedAt time.Time `json:"diagnosedAt" orm:"diagnosed_at" description:"确诊时间"`          // 确诊时间
	IsCurrent   int       `json:"isCurrent"   orm:"is_current"   description:"是否当前仍存在：1是，0否"` // 是否当前仍存在：1是，0否
	CreatedAt   time.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`          // 创建时间
	UpdatedAt   time.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`          // 更新时间
}
