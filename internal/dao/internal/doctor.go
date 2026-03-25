// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DoctorDao is the data access object for the table doctor.
type DoctorDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DoctorColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DoctorColumns defines and stores column names for the table doctor.
type DoctorColumns struct {
	Id           string // 医生主键ID
	HospitalId   string // 所属医院ID
	Username     string // 医生登录账号
	PasswordHash string // 加密后的密码
	DoctorName   string // 医生姓名
	Gender       string // 性别：1男，2女，0未知
	Phone        string // 手机号
	Email        string // 邮箱
	Title        string // 职称
	Specialty    string // 擅长领域
	AvatarUrl    string // 头像URL
	Intro        string // 医生简介
	Status       string // 状态：1在职可接诊，0停用
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// doctorColumns holds the columns for the table doctor.
var doctorColumns = DoctorColumns{
	Id:           "id",
	HospitalId:   "hospital_id",
	Username:     "username",
	PasswordHash: "password_hash",
	DoctorName:   "doctor_name",
	Gender:       "gender",
	Phone:        "phone",
	Email:        "email",
	Title:        "title",
	Specialty:    "specialty",
	AvatarUrl:    "avatar_url",
	Intro:        "intro",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewDoctorDao creates and returns a new DAO object for table data access.
func NewDoctorDao(handlers ...gdb.ModelHandler) *DoctorDao {
	return &DoctorDao{
		group:    "default",
		table:    "doctor",
		columns:  doctorColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DoctorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DoctorDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DoctorDao) Columns() DoctorColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DoctorDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DoctorDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DoctorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
