// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PetVaccinationRecordDao is the data access object for the table pet_vaccination_record.
type PetVaccinationRecordDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  PetVaccinationRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// PetVaccinationRecordColumns defines and stores column names for the table pet_vaccination_record.
type PetVaccinationRecordColumns struct {
	Id              string // 疫苗记录ID
	PetId           string // 宠物ID
	VaccineName     string // 疫苗名称
	VaccinationDate string // 接种日期
	NextDueDate     string // 下次应接种日期
	HospitalName    string // 接种机构
	Remark          string // 备注
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// petVaccinationRecordColumns holds the columns for the table pet_vaccination_record.
var petVaccinationRecordColumns = PetVaccinationRecordColumns{
	Id:              "id",
	PetId:           "pet_id",
	VaccineName:     "vaccine_name",
	VaccinationDate: "vaccination_date",
	NextDueDate:     "next_due_date",
	HospitalName:    "hospital_name",
	Remark:          "remark",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewPetVaccinationRecordDao creates and returns a new DAO object for table data access.
func NewPetVaccinationRecordDao(handlers ...gdb.ModelHandler) *PetVaccinationRecordDao {
	return &PetVaccinationRecordDao{
		group:    "default",
		table:    "pet_vaccination_record",
		columns:  petVaccinationRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PetVaccinationRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PetVaccinationRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PetVaccinationRecordDao) Columns() PetVaccinationRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PetVaccinationRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PetVaccinationRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PetVaccinationRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
