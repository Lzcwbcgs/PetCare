package v1

import "github.com/gogf/gf/v2/frame/g"

type UserProfileReq struct {
	g.Meta `path:"/profile" method:"get" tags:"用户资料" summary:"获取用户个人信息"`
}

type UserProfileRes struct {
	ID        int64  `json:"id" dc:"用户ID"`
	Username  string `json:"username" dc:"用户名"`
	Nickname  string `json:"nickname" dc:"昵称"`
	Phone     string `json:"phone" dc:"手机号"`
	Email     string `json:"email" dc:"邮箱"`
	AvatarURL string `json:"avatar_url" dc:"头像地址"`
	Status    int    `json:"status" dc:"状态"`
	CreatedAt string `json:"created_at" dc:"创建时间"`
}

type UserUpdateProfileReq struct {
	g.Meta    `path:"/profile" method:"put" tags:"用户资料" summary:"更新用户个人信息"`
	Nickname  *string `json:"nickname" v:"length:1,32#昵称长度需在1到32位之间" dc:"昵称"`
	Phone     *string `json:"phone" v:"phone#手机号格式不正确" dc:"手机号"`
	Email     *string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	AvatarURL *string `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"头像地址"`
}

type UserUpdateProfileRes struct{}

type UserUpdatePasswordReq struct {
	g.Meta      `path:"/password" method:"put" tags:"用户资料" summary:"修改用户密码"`
	OldPassword string `json:"old_password" v:"required|length:6,64#请输入旧密码|旧密码长度需在6到64位之间" dc:"旧密码"`
	NewPassword string `json:"new_password" v:"required|length:6,64#请输入新密码|新密码长度需在6到64位之间" dc:"新密码"`
}

type UserUpdatePasswordRes struct{}
