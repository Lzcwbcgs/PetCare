// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// PetAllergyRecord is the golang structure for table pet_allergy_record.
type PetAllergyRecord struct {
	Id                 int64     `json:"id"                 orm:"id"                  description:"过敏记录ID"`           // 过敏记录ID
	PetId              int64     `json:"petId"              orm:"pet_id"              description:"宠物ID"`             // 宠物ID
	Allergen           string    `json:"allergen"           orm:"allergen"            description:"过敏源"`              // 过敏源
	SymptomDescription string    `json:"symptomDescription" orm:"symptom_description" description:"症状描述"`             // 症状描述
	SeverityLevel      int       `json:"severityLevel"      orm:"severity_level"      description:"严重程度：1轻微，2中等，3严重"` // 严重程度：1轻微，2中等，3严重
	Remark             string    `json:"remark"             orm:"remark"              description:"备注"`               // 备注
	CreatedAt          time.Time `json:"createdAt"          orm:"created_at"          description:"创建时间"`             // 创建时间
	UpdatedAt          time.Time `json:"updatedAt"          orm:"updated_at"          description:"更新时间"`             // 更新时间
}
