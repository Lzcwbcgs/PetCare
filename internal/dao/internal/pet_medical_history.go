// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PetMedicalHistoryDao is the data access object for the table pet_medical_history.
type PetMedicalHistoryDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  PetMedicalHistoryColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// PetMedicalHistoryColumns defines and stores column names for the table pet_medical_history.
type PetMedicalHistoryColumns struct {
	Id          string // 病史记录ID
	PetId       string // 宠物ID
	HistoryType string // 病史类型
	Description string // 病史描述
	DiagnosedAt string // 确诊时间
	IsCurrent   string // 是否当前仍存在：1是，0否
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// petMedicalHistoryColumns holds the columns for the table pet_medical_history.
var petMedicalHistoryColumns = PetMedicalHistoryColumns{
	Id:          "id",
	PetId:       "pet_id",
	HistoryType: "history_type",
	Description: "description",
	DiagnosedAt: "diagnosed_at",
	IsCurrent:   "is_current",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewPetMedicalHistoryDao creates and returns a new DAO object for table data access.
func NewPetMedicalHistoryDao(handlers ...gdb.ModelHandler) *PetMedicalHistoryDao {
	return &PetMedicalHistoryDao{
		group:    "default",
		table:    "pet_medical_history",
		columns:  petMedicalHistoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PetMedicalHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PetMedicalHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PetMedicalHistoryDao) Columns() PetMedicalHistoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PetMedicalHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PetMedicalHistoryDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PetMedicalHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
