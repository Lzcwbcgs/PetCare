// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Notification is the golang structure for table notification.
type Notification struct {
	Id               int64     `json:"id"               orm:"id"                description:"通知ID"`                     // 通知ID
	UserId           int64     `json:"userId"           orm:"user_id"           description:"接收用户ID"`                   // 接收用户ID
	DoctorId         int64     `json:"doctorId"         orm:"doctor_id"         description:"接收医生ID"`                   // 接收医生ID
	AppointmentId    int64     `json:"appointmentId"    orm:"appointment_id"    description:"关联预约ID"`                   // 关联预约ID
	NotificationType int       `json:"notificationType" orm:"notification_type" description:"通知类型：1预约提醒，2系统通知，3AI分析提醒"` // 通知类型：1预约提醒，2系统通知，3AI分析提醒
	Title            string    `json:"title"            orm:"title"             description:"通知标题"`                     // 通知标题
	Content          string    `json:"content"          orm:"content"           description:"通知内容"`                     // 通知内容
	SendTime         time.Time `json:"sendTime"         orm:"send_time"         description:"计划发送时间"`                   // 计划发送时间
	Status           int       `json:"status"           orm:"status"            description:"状态：0待发送，1已发送，2发送失败，3已读"`   // 状态：0待发送，1已发送，2发送失败，3已读
	CreatedAt        time.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`                     // 创建时间
	UpdatedAt        time.Time `json:"updatedAt"        orm:"updated_at"        description:"更新时间"`                     // 更新时间
}
