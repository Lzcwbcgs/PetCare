// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OperationLogDao is the data access object for the table operation_log.
type OperationLogDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  OperationLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// OperationLogColumns defines and stores column names for the table operation_log.
type OperationLogColumns struct {
	Id              string // 日志ID
	OperatorType    string // 操作人类型：1管理员，2医生，3用户
	OperatorId      string // 操作人ID
	OperationModule string // 操作模块
	OperationType   string // 操作类型
	OperationDesc   string // 操作描述
	IpAddress       string // IP地址
	CreatedAt       string // 操作时间
}

// operationLogColumns holds the columns for the table operation_log.
var operationLogColumns = OperationLogColumns{
	Id:              "id",
	OperatorType:    "operator_type",
	OperatorId:      "operator_id",
	OperationModule: "operation_module",
	OperationType:   "operation_type",
	OperationDesc:   "operation_desc",
	IpAddress:       "ip_address",
	CreatedAt:       "created_at",
}

// NewOperationLogDao creates and returns a new DAO object for table data access.
func NewOperationLogDao(handlers ...gdb.ModelHandler) *OperationLogDao {
	return &OperationLogDao{
		group:    "default",
		table:    "operation_log",
		columns:  operationLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OperationLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OperationLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OperationLogDao) Columns() OperationLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OperationLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OperationLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OperationLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
