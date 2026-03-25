// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AiAnalysisRecord is the golang structure of table ai_analysis_record for DAO operations like Where/Data.
type AiAnalysisRecord struct {
	g.Meta           `orm:"table:ai_analysis_record, do:true"`
	Id               any // AI分析记录ID
	PetId            any // 宠物ID
	SessionId        any // 关联AI会话ID
	MedicalRecordId  any // 关联病历ID
	AnalysisType     any // 分析类型：1病历总结，2症状归纳，3风险提示，4健康建议
	InputSource      any // 输入来源：1AI对话，2病历记录，3体检数据，4综合数据
	AnalysisResult   any // 分析结果
	RuleBasedResult  any // 规则分析结果
	LlmBasedResult   any // 大模型分析结果
	RiskLevel        any // 风险等级：1低，2中，3高
	ReviewedByDoctor any // 医生是否已审核：1是，0否
	CreatedAt        any // 创建时间
}
