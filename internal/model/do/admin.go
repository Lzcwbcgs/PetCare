// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Admin is the golang structure of table admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta       `orm:"table:admin, do:true"`
	Id           any // 管理员主键ID
	Username     any // 管理员账号
	PasswordHash any // 加密后的密码
	RealName     any // 管理员姓名
	Phone        any // 联系电话
	Status       any // 状态：1正常，0禁用
	CreatedAt    any // 创建时间
	UpdatedAt    any // 更新时间
}
