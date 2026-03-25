// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:user, do:true"`
	Id           any // 用户主键ID
	Username     any // 登录用户名
	PasswordHash any // 加密后的密码
	Nickname     any // 昵称
	Phone        any // 手机号
	Email        any // 邮箱
	AvatarUrl    any // 用户头像URL
	Status       any // 状态：1正常，0禁用
	CreatedAt    any // 创建时间
	UpdatedAt    any // 更新时间
}
