// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Appointment is the golang structure of table appointment for DAO operations like Where/Data.
type Appointment struct {
	g.Meta             `orm:"table:appointment, do:true"`
	Id                 any // 预约主键ID
	AppointmentNo      any // 预约单号
	UserId             any // 用户ID
	PetId              any // 宠物ID
	HospitalId         any // 医院ID
	DoctorId           any // 医生ID，可为空
	AppointmentType    any // 预约类型：1体检预约，2看病预约
	SymptomDescription any // 症状描述（看病预约时填写）
	AppointmentTime    any // 预约时间
	ReminderTime       any // 提醒时间，一般为预约前1小时
	Status             any // 状态：1待就诊，2已完成，3已取消，4已过期
	Source             any // 来源：1用户端预约，2医生代录入，3后台创建
	CreatedAt          any // 创建时间
	UpdatedAt          any // 更新时间
}
