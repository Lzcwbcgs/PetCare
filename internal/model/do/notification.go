// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Notification is the golang structure of table notification for DAO operations like Where/Data.
type Notification struct {
	g.Meta           `orm:"table:notification, do:true"`
	Id               any // 通知ID
	UserId           any // 接收用户ID
	DoctorId         any // 接收医生ID
	AppointmentId    any // 关联预约ID
	NotificationType any // 通知类型：1预约提醒，2系统通知，3AI分析提醒
	Title            any // 通知标题
	Content          any // 通知内容
	SendTime         any // 计划发送时间
	Status           any // 状态：0待发送，1已发送，2发送失败，3已读
	CreatedAt        any // 创建时间
	UpdatedAt        any // 更新时间
}
