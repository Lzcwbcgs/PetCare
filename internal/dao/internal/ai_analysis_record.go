// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AiAnalysisRecordDao is the data access object for the table ai_analysis_record.
type AiAnalysisRecordDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  AiAnalysisRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// AiAnalysisRecordColumns defines and stores column names for the table ai_analysis_record.
type AiAnalysisRecordColumns struct {
	Id               string // AI分析记录ID
	PetId            string // 宠物ID
	SessionId        string // 关联AI会话ID
	MedicalRecordId  string // 关联病历ID
	AnalysisType     string // 分析类型：1病历总结，2症状归纳，3风险提示，4健康建议
	InputSource      string // 输入来源：1AI对话，2病历记录，3体检数据，4综合数据
	AnalysisResult   string // 分析结果
	RuleBasedResult  string // 规则分析结果
	LlmBasedResult   string // 大模型分析结果
	RiskLevel        string // 风险等级：1低，2中，3高
	ReviewedByDoctor string // 医生是否已审核：1是，0否
	CreatedAt        string // 创建时间
}

// aiAnalysisRecordColumns holds the columns for the table ai_analysis_record.
var aiAnalysisRecordColumns = AiAnalysisRecordColumns{
	Id:               "id",
	PetId:            "pet_id",
	SessionId:        "session_id",
	MedicalRecordId:  "medical_record_id",
	AnalysisType:     "analysis_type",
	InputSource:      "input_source",
	AnalysisResult:   "analysis_result",
	RuleBasedResult:  "rule_based_result",
	LlmBasedResult:   "llm_based_result",
	RiskLevel:        "risk_level",
	ReviewedByDoctor: "reviewed_by_doctor",
	CreatedAt:        "created_at",
}

// NewAiAnalysisRecordDao creates and returns a new DAO object for table data access.
func NewAiAnalysisRecordDao(handlers ...gdb.ModelHandler) *AiAnalysisRecordDao {
	return &AiAnalysisRecordDao{
		group:    "default",
		table:    "ai_analysis_record",
		columns:  aiAnalysisRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AiAnalysisRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AiAnalysisRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AiAnalysisRecordDao) Columns() AiAnalysisRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AiAnalysisRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AiAnalysisRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AiAnalysisRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
