// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PetMedicalHistory is the golang structure of table pet_medical_history for DAO operations like Where/Data.
type PetMedicalHistory struct {
	g.Meta      `orm:"table:pet_medical_history, do:true"`
	Id          any // 病史记录ID
	PetId       any // 宠物ID
	HistoryType any // 病史类型
	Description any // 病史描述
	DiagnosedAt any // 确诊时间
	IsCurrent   any // 是否当前仍存在：1是，0否
	CreatedAt   any // 创建时间
	UpdatedAt   any // 更新时间
}
