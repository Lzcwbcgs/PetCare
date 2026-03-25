// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AiMessageDao is the data access object for the table ai_message.
type AiMessageDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AiMessageColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AiMessageColumns defines and stores column names for the table ai_message.
type AiMessageColumns struct {
	Id             string // AI消息ID
	SessionId      string // AI会话ID
	SenderType     string // 发送者类型：1用户，2AI，3医生，4管理员
	SenderId       string // 发送者ID，可为空
	MessageContent string // 消息内容
	MessageType    string // 消息类型：1文本
	CreatedAt      string // 发送时间
}

// aiMessageColumns holds the columns for the table ai_message.
var aiMessageColumns = AiMessageColumns{
	Id:             "id",
	SessionId:      "session_id",
	SenderType:     "sender_type",
	SenderId:       "sender_id",
	MessageContent: "message_content",
	MessageType:    "message_type",
	CreatedAt:      "created_at",
}

// NewAiMessageDao creates and returns a new DAO object for table data access.
func NewAiMessageDao(handlers ...gdb.ModelHandler) *AiMessageDao {
	return &AiMessageDao{
		group:    "default",
		table:    "ai_message",
		columns:  aiMessageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AiMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AiMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AiMessageDao) Columns() AiMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AiMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AiMessageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AiMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
