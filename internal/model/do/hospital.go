// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Hospital is the golang structure of table hospital for DAO operations like Where/Data.
type Hospital struct {
	g.Meta       `orm:"table:hospital, do:true"`
	Id           any // 医院主键ID
	HospitalName any // 医院名称
	Address      any // 医院地址
	Phone        any // 医院联系电话
	Description  any // 医院简介
	Status       any // 状态：1启用，0停用
	CreatedAt    any // 创建时间
	UpdatedAt    any // 更新时间
}
