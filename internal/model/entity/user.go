// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// User is the golang structure for table user.
type User struct {
	Id           int64     `json:"id"           orm:"id"            description:"用户主键ID"`     // 用户主键ID
	Username     string    `json:"username"     orm:"username"      description:"登录用户名"`      // 登录用户名
	PasswordHash string    `json:"passwordHash" orm:"password_hash" description:"加密后的密码"`     // 加密后的密码
	Nickname     string    `json:"nickname"     orm:"nickname"      description:"昵称"`         // 昵称
	Phone        string    `json:"phone"        orm:"phone"         description:"手机号"`        // 手机号
	Email        string    `json:"email"        orm:"email"         description:"邮箱"`         // 邮箱
	AvatarUrl    string    `json:"avatarUrl"    orm:"avatar_url"    description:"用户头像URL"`    // 用户头像URL
	Status       int       `json:"status"       orm:"status"        description:"状态：1正常，0禁用"` // 状态：1正常，0禁用
	CreatedAt    time.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`       // 创建时间
	UpdatedAt    time.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`       // 更新时间
}
