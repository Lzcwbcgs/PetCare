// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id           int64     `json:"id"           orm:"id"            description:"管理员主键ID"`    // 管理员主键ID
	Username     string    `json:"username"     orm:"username"      description:"管理员账号"`      // 管理员账号
	PasswordHash string    `json:"passwordHash" orm:"password_hash" description:"加密后的密码"`     // 加密后的密码
	RealName     string    `json:"realName"     orm:"real_name"     description:"管理员姓名"`      // 管理员姓名
	Phone        string    `json:"phone"        orm:"phone"         description:"联系电话"`       // 联系电话
	Status       int       `json:"status"       orm:"status"        description:"状态：1正常，0禁用"` // 状态：1正常，0禁用
	CreatedAt    time.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`       // 创建时间
	UpdatedAt    time.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`       // 更新时间
}
