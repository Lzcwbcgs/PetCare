// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PetDao is the data access object for the table pet.
type PetDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PetColumns         // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PetColumns defines and stores column names for the table pet.
type PetColumns struct {
	Id         string // 宠物主键ID
	UserId     string // 所属用户ID
	PetName    string // 宠物名字
	PetType    string // 宠物类型，当前默认猫
	AvatarUrl  string // 宠物头像URL
	Gender     string // 性别：1公，2母，0未知
	Age        string // 年龄
	AgeUnit    string // 年龄单位：month/月，year/岁
	Breed      string // 品种
	Weight     string // 体重（kg）
	Sterilized string // 是否绝育：1是，0否
	Remark     string // 备注
	Status     string // 状态：1正常，0停用/删除
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// petColumns holds the columns for the table pet.
var petColumns = PetColumns{
	Id:         "id",
	UserId:     "user_id",
	PetName:    "pet_name",
	PetType:    "pet_type",
	AvatarUrl:  "avatar_url",
	Gender:     "gender",
	Age:        "age",
	AgeUnit:    "age_unit",
	Breed:      "breed",
	Weight:     "weight",
	Sterilized: "sterilized",
	Remark:     "remark",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewPetDao creates and returns a new DAO object for table data access.
func NewPetDao(handlers ...gdb.ModelHandler) *PetDao {
	return &PetDao{
		group:    "default",
		table:    "pet",
		columns:  petColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PetDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PetDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PetDao) Columns() PetColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PetDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PetDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PetDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
