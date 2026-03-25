// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Doctor is the golang structure for table doctor.
type Doctor struct {
	Id           int64     `json:"id"           orm:"id"            description:"医生主键ID"`        // 医生主键ID
	HospitalId   int64     `json:"hospitalId"   orm:"hospital_id"   description:"所属医院ID"`        // 所属医院ID
	Username     string    `json:"username"     orm:"username"      description:"医生登录账号"`        // 医生登录账号
	PasswordHash string    `json:"passwordHash" orm:"password_hash" description:"加密后的密码"`        // 加密后的密码
	DoctorName   string    `json:"doctorName"   orm:"doctor_name"   description:"医生姓名"`          // 医生姓名
	Gender       int       `json:"gender"       orm:"gender"        description:"性别：1男，2女，0未知"`  // 性别：1男，2女，0未知
	Phone        string    `json:"phone"        orm:"phone"         description:"手机号"`           // 手机号
	Email        string    `json:"email"        orm:"email"         description:"邮箱"`            // 邮箱
	Title        string    `json:"title"        orm:"title"         description:"职称"`            // 职称
	Specialty    string    `json:"specialty"    orm:"specialty"     description:"擅长领域"`          // 擅长领域
	AvatarUrl    string    `json:"avatarUrl"    orm:"avatar_url"    description:"头像URL"`         // 头像URL
	Intro        string    `json:"intro"        orm:"intro"         description:"医生简介"`          // 医生简介
	Status       int       `json:"status"       orm:"status"        description:"状态：1在职可接诊，0停用"` // 状态：1在职可接诊，0停用
	CreatedAt    time.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`          // 创建时间
	UpdatedAt    time.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`          // 更新时间
}
