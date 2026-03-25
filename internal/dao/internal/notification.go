// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotificationDao is the data access object for the table notification.
type NotificationDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  NotificationColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// NotificationColumns defines and stores column names for the table notification.
type NotificationColumns struct {
	Id               string // 通知ID
	UserId           string // 接收用户ID
	DoctorId         string // 接收医生ID
	AppointmentId    string // 关联预约ID
	NotificationType string // 通知类型：1预约提醒，2系统通知，3AI分析提醒
	Title            string // 通知标题
	Content          string // 通知内容
	SendTime         string // 计划发送时间
	Status           string // 状态：0待发送，1已发送，2发送失败，3已读
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// notificationColumns holds the columns for the table notification.
var notificationColumns = NotificationColumns{
	Id:               "id",
	UserId:           "user_id",
	DoctorId:         "doctor_id",
	AppointmentId:    "appointment_id",
	NotificationType: "notification_type",
	Title:            "title",
	Content:          "content",
	SendTime:         "send_time",
	Status:           "status",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewNotificationDao creates and returns a new DAO object for table data access.
func NewNotificationDao(handlers ...gdb.ModelHandler) *NotificationDao {
	return &NotificationDao{
		group:    "default",
		table:    "notification",
		columns:  notificationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NotificationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NotificationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NotificationDao) Columns() NotificationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NotificationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NotificationDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *NotificationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
