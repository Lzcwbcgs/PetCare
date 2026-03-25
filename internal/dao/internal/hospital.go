// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HospitalDao is the data access object for the table hospital.
type HospitalDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  HospitalColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// HospitalColumns defines and stores column names for the table hospital.
type HospitalColumns struct {
	Id           string // 医院主键ID
	HospitalName string // 医院名称
	Address      string // 医院地址
	Phone        string // 医院联系电话
	Description  string // 医院简介
	Status       string // 状态：1启用，0停用
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// hospitalColumns holds the columns for the table hospital.
var hospitalColumns = HospitalColumns{
	Id:           "id",
	HospitalName: "hospital_name",
	Address:      "address",
	Phone:        "phone",
	Description:  "description",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewHospitalDao creates and returns a new DAO object for table data access.
func NewHospitalDao(handlers ...gdb.ModelHandler) *HospitalDao {
	return &HospitalDao{
		group:    "default",
		table:    "hospital",
		columns:  hospitalColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HospitalDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HospitalDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HospitalDao) Columns() HospitalColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HospitalDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HospitalDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HospitalDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
