// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Hospital is the golang structure for table hospital.
type Hospital struct {
	Id           int64     `json:"id"           orm:"id"            description:"医院主键ID"`     // 医院主键ID
	HospitalName string    `json:"hospitalName" orm:"hospital_name" description:"医院名称"`       // 医院名称
	Address      string    `json:"address"      orm:"address"       description:"医院地址"`       // 医院地址
	Phone        string    `json:"phone"        orm:"phone"         description:"医院联系电话"`     // 医院联系电话
	Description  string    `json:"description"  orm:"description"   description:"医院简介"`       // 医院简介
	Status       int       `json:"status"       orm:"status"        description:"状态：1启用，0停用"` // 状态：1启用，0停用
	CreatedAt    time.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`       // 创建时间
	UpdatedAt    time.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`       // 更新时间
}
