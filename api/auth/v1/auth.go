package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post" tags:"认证模块" summary:"用户注册"`
	Username string `json:"username" v:"required|length:3,32#请输入用户名|用户名长度需在3到32位之间" dc:"登录用户名"`
	Password string `json:"password" v:"required|length:6,64#请输入密码|密码长度需在6到64位之间" dc:"登录密码"`
	Nickname string `json:"nickname" v:"required|length:1,32#请输入昵称|昵称长度需在1到32位之间" dc:"用户昵称"`
	Phone    string `json:"phone" v:"required|phone#请输入手机号|手机号格式不正确" dc:"手机号"`
	Email    string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
}

type RegisterRes struct {
	UserID int64 `json:"user_id" dc:"新注册用户ID"`
}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"认证模块" summary:"统一登录"`
	Username string `json:"username" v:"required|length:3,32#请输入用户名|用户名长度需在3到32位之间" dc:"登录用户名"`
	Password string `json:"password" v:"required|length:6,64#请输入密码|密码长度需在6到64位之间" dc:"登录密码"`
	Role     string `json:"role" v:"required|in:user,doctor,admin#请选择登录角色|登录角色不合法" dc:"登录角色: user/doctor/admin"`
}

type LoginRes struct {
	Token    string `json:"token" dc:"JWT访问令牌"`
	ExpireAt string `json:"expire_at" dc:"令牌过期时间"`
	UserID   int64  `json:"user_id" dc:"当前登录用户ID"`
	Role     string `json:"role" dc:"当前登录角色"`
}

type MeReq struct {
	g.Meta `path:"/me" method:"get" tags:"认证模块" summary:"获取当前登录信息"`
}

type MeRes struct {
	ID         int64  `json:"id" dc:"当前登录主体ID"`
	Username   string `json:"username" dc:"用户名"`
	Role       string `json:"role" dc:"角色"`
	Nickname   string `json:"nickname,omitempty" dc:"用户昵称"`
	DoctorName string `json:"doctor_name,omitempty" dc:"医生姓名"`
	RealName   string `json:"real_name,omitempty" dc:"管理员姓名"`
	AvatarURL  string `json:"avatar_url,omitempty" dc:"头像地址"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"认证模块" summary:"退出登录"`
}

type LogoutRes struct{}
