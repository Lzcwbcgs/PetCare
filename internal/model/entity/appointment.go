// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Appointment is the golang structure for table appointment.
type Appointment struct {
	Id                 int64     `json:"id"                 orm:"id"                  description:"预约主键ID"`                 // 预约主键ID
	AppointmentNo      string    `json:"appointmentNo"      orm:"appointment_no"      description:"预约单号"`                   // 预约单号
	UserId             int64     `json:"userId"             orm:"user_id"             description:"用户ID"`                   // 用户ID
	PetId              int64     `json:"petId"              orm:"pet_id"              description:"宠物ID"`                   // 宠物ID
	HospitalId         int64     `json:"hospitalId"         orm:"hospital_id"         description:"医院ID"`                   // 医院ID
	DoctorId           int64     `json:"doctorId"           orm:"doctor_id"           description:"医生ID，可为空"`               // 医生ID，可为空
	AppointmentType    int       `json:"appointmentType"    orm:"appointment_type"    description:"预约类型：1体检预约，2看病预约"`       // 预约类型：1体检预约，2看病预约
	SymptomDescription string    `json:"symptomDescription" orm:"symptom_description" description:"症状描述（看病预约时填写）"`          // 症状描述（看病预约时填写）
	AppointmentTime    time.Time `json:"appointmentTime"    orm:"appointment_time"    description:"预约时间"`                   // 预约时间
	ReminderTime       time.Time `json:"reminderTime"       orm:"reminder_time"       description:"提醒时间，一般为预约前1小时"`         // 提醒时间，一般为预约前1小时
	Status             int       `json:"status"             orm:"status"              description:"状态：1待就诊，2已完成，3已取消，4已过期"` // 状态：1待就诊，2已完成，3已取消，4已过期
	Source             int       `json:"source"             orm:"source"              description:"来源：1用户端预约，2医生代录入，3后台创建"` // 来源：1用户端预约，2医生代录入，3后台创建
	CreatedAt          time.Time `json:"createdAt"          orm:"created_at"          description:"创建时间"`                   // 创建时间
	UpdatedAt          time.Time `json:"updatedAt"          orm:"updated_at"          description:"更新时间"`                   // 更新时间
}
