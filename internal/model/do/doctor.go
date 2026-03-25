// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Doctor is the golang structure of table doctor for DAO operations like Where/Data.
type Doctor struct {
	g.Meta       `orm:"table:doctor, do:true"`
	Id           any // 医生主键ID
	HospitalId   any // 所属医院ID
	Username     any // 医生登录账号
	PasswordHash any // 加密后的密码
	DoctorName   any // 医生姓名
	Gender       any // 性别：1男，2女，0未知
	Phone        any // 手机号
	Email        any // 邮箱
	Title        any // 职称
	Specialty    any // 擅长领域
	AvatarUrl    any // 头像URL
	Intro        any // 医生简介
	Status       any // 状态：1在职可接诊，0停用
	CreatedAt    any // 创建时间
	UpdatedAt    any // 更新时间
}
