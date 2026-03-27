package v1

import "github.com/gogf/gf/v2/frame/g"

type DoctorProfileReq struct {
	g.Meta `path:"/profile" method:"get" tags:"医生资料" summary:"获取医生个人信息"`
}

type DoctorProfileRes struct {
	ID           int64  `json:"id" dc:"医生ID"`
	HospitalID   int64  `json:"hospital_id" dc:"医院ID"`
	HospitalName string `json:"hospital_name" dc:"医院名称"`
	Username     string `json:"username" dc:"用户名"`
	DoctorName   string `json:"doctor_name" dc:"医生姓名"`
	Gender       int    `json:"gender" dc:"性别"`
	Phone        string `json:"phone" dc:"手机号"`
	Email        string `json:"email" dc:"邮箱"`
	Title        string `json:"title" dc:"职称"`
	Specialty    string `json:"specialty" dc:"擅长领域"`
	AvatarURL    string `json:"avatar_url" dc:"头像地址"`
	Intro        string `json:"intro" dc:"简介"`
	Status       int    `json:"status" dc:"状态"`
}

type DoctorUpdateProfileReq struct {
	g.Meta    `path:"/profile" method:"put" tags:"医生资料" summary:"更新医生个人信息"`
	Phone     *string `json:"phone" v:"phone#手机号格式不正确" dc:"手机号"`
	Email     *string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	AvatarURL *string `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"头像地址"`
	Intro     *string `json:"intro" v:"max-length:65535#简介长度不能超过65535位" dc:"简介"`
}

type DoctorUpdateProfileRes struct{}

type DoctorUpdatePasswordReq struct {
	g.Meta      `path:"/password" method:"put" tags:"医生资料" summary:"修改医生密码"`
	OldPassword string `json:"old_password" v:"required|length:6,64#请输入旧密码|旧密码长度需在6到64位之间" dc:"旧密码"`
	NewPassword string `json:"new_password" v:"required|length:6,64#请输入新密码|新密码长度需在6到64位之间" dc:"新密码"`
}

type DoctorUpdatePasswordRes struct{}
