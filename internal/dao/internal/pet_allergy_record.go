// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PetAllergyRecordDao is the data access object for the table pet_allergy_record.
type PetAllergyRecordDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  PetAllergyRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// PetAllergyRecordColumns defines and stores column names for the table pet_allergy_record.
type PetAllergyRecordColumns struct {
	Id                 string // 过敏记录ID
	PetId              string // 宠物ID
	Allergen           string // 过敏源
	SymptomDescription string // 症状描述
	SeverityLevel      string // 严重程度：1轻微，2中等，3严重
	Remark             string // 备注
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
}

// petAllergyRecordColumns holds the columns for the table pet_allergy_record.
var petAllergyRecordColumns = PetAllergyRecordColumns{
	Id:                 "id",
	PetId:              "pet_id",
	Allergen:           "allergen",
	SymptomDescription: "symptom_description",
	SeverityLevel:      "severity_level",
	Remark:             "remark",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewPetAllergyRecordDao creates and returns a new DAO object for table data access.
func NewPetAllergyRecordDao(handlers ...gdb.ModelHandler) *PetAllergyRecordDao {
	return &PetAllergyRecordDao{
		group:    "default",
		table:    "pet_allergy_record",
		columns:  petAllergyRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PetAllergyRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PetAllergyRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PetAllergyRecordDao) Columns() PetAllergyRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PetAllergyRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PetAllergyRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PetAllergyRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
