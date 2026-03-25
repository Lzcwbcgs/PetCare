// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AiSessionDao is the data access object for the table ai_session.
type AiSessionDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AiSessionColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AiSessionColumns defines and stores column names for the table ai_session.
type AiSessionColumns struct {
	Id             string // AI会话ID
	SessionNo      string // 会话编号
	UserId         string // 用户ID
	PetId          string // 宠物ID
	HospitalId     string // 关联医院ID，可为空
	DoctorId       string // 关联医生ID，可为空
	SourceType     string // 来源：1用户端发起，2医生端发起
	ModelType      string // 模型类型，如本地模型/API模型
	ModelName      string // 模型名称
	SessionSummary string // AI会话总结
	SyncToAdmin    string // 是否同步给管理端：1是，0否
	Status         string // 状态：1进行中，2已结束，3已归档
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// aiSessionColumns holds the columns for the table ai_session.
var aiSessionColumns = AiSessionColumns{
	Id:             "id",
	SessionNo:      "session_no",
	UserId:         "user_id",
	PetId:          "pet_id",
	HospitalId:     "hospital_id",
	DoctorId:       "doctor_id",
	SourceType:     "source_type",
	ModelType:      "model_type",
	ModelName:      "model_name",
	SessionSummary: "session_summary",
	SyncToAdmin:    "sync_to_admin",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewAiSessionDao creates and returns a new DAO object for table data access.
func NewAiSessionDao(handlers ...gdb.ModelHandler) *AiSessionDao {
	return &AiSessionDao{
		group:    "default",
		table:    "ai_session",
		columns:  aiSessionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AiSessionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AiSessionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AiSessionDao) Columns() AiSessionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AiSessionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AiSessionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AiSessionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
