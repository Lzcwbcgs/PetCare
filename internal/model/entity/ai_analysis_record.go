// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// AiAnalysisRecord is the golang structure for table ai_analysis_record.
type AiAnalysisRecord struct {
	Id               int64     `json:"id"               orm:"id"                  description:"AI分析记录ID"`                     // AI分析记录ID
	PetId            int64     `json:"petId"            orm:"pet_id"              description:"宠物ID"`                         // 宠物ID
	SessionId        int64     `json:"sessionId"        orm:"session_id"          description:"关联AI会话ID"`                     // 关联AI会话ID
	MedicalRecordId  int64     `json:"medicalRecordId"  orm:"medical_record_id"   description:"关联病历ID"`                       // 关联病历ID
	AnalysisType     int       `json:"analysisType"     orm:"analysis_type"       description:"分析类型：1病历总结，2症状归纳，3风险提示，4健康建议"` // 分析类型：1病历总结，2症状归纳，3风险提示，4健康建议
	InputSource      int       `json:"inputSource"      orm:"input_source"        description:"输入来源：1AI对话，2病历记录，3体检数据，4综合数据"` // 输入来源：1AI对话，2病历记录，3体检数据，4综合数据
	SummaryTitle     string    `json:"summaryTitle"     orm:"summary_title"       description:"分析标题"`                         // 分析标题
	AnalysisResult   string    `json:"analysisResult"   orm:"analysis_result"     description:"分析结果"`                         // 分析结果
	RuleBasedResult  string    `json:"ruleBasedResult"  orm:"rule_based_result"   description:"规则分析结果"`                       // 规则分析结果
	LlmBasedResult   string    `json:"llmBasedResult"   orm:"llm_based_result"    description:"大模型分析结果"`                      // 大模型分析结果
	RiskLevel        int       `json:"riskLevel"        orm:"risk_level"          description:"风险等级：1低，2中，3高"`                // 风险等级：1低，2中，3高
	ConfidenceScore  float64   `json:"confidenceScore"  orm:"confidence_score"    description:"置信度分数，0-100"`                  // 置信度分数，0-100
	ReferenceChunks  string    `json:"referenceChunks"  orm:"reference_chunks"    description:"引用知识片段信息"`                    // 引用知识片段信息
	ExtraMetadata    string    `json:"extraMetadata"    orm:"extra_metadata"      description:"扩展元数据"`                       // 扩展元数据
	ReviewedByDoctor int       `json:"reviewedByDoctor" orm:"reviewed_by_doctor"  description:"医生是否已审核：1是，0否"`                // 医生是否已审核：1是，0否
	CreatedAt        time.Time `json:"createdAt"        orm:"created_at"          description:"创建时间"`                         // 创建时间
}
