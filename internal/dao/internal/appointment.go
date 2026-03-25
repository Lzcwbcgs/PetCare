// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppointmentDao is the data access object for the table appointment.
type AppointmentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppointmentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppointmentColumns defines and stores column names for the table appointment.
type AppointmentColumns struct {
	Id                 string // 预约主键ID
	AppointmentNo      string // 预约单号
	UserId             string // 用户ID
	PetId              string // 宠物ID
	HospitalId         string // 医院ID
	DoctorId           string // 医生ID，可为空
	AppointmentType    string // 预约类型：1体检预约，2看病预约
	SymptomDescription string // 症状描述（看病预约时填写）
	AppointmentTime    string // 预约时间
	ReminderTime       string // 提醒时间，一般为预约前1小时
	Status             string // 状态：1待就诊，2已完成，3已取消，4已过期
	Source             string // 来源：1用户端预约，2医生代录入，3后台创建
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
}

// appointmentColumns holds the columns for the table appointment.
var appointmentColumns = AppointmentColumns{
	Id:                 "id",
	AppointmentNo:      "appointment_no",
	UserId:             "user_id",
	PetId:              "pet_id",
	HospitalId:         "hospital_id",
	DoctorId:           "doctor_id",
	AppointmentType:    "appointment_type",
	SymptomDescription: "symptom_description",
	AppointmentTime:    "appointment_time",
	ReminderTime:       "reminder_time",
	Status:             "status",
	Source:             "source",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewAppointmentDao creates and returns a new DAO object for table data access.
func NewAppointmentDao(handlers ...gdb.ModelHandler) *AppointmentDao {
	return &AppointmentDao{
		group:    "default",
		table:    "appointment",
		columns:  appointmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppointmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppointmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppointmentDao) Columns() AppointmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppointmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppointmentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppointmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
