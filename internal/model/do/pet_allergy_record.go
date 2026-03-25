// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PetAllergyRecord is the golang structure of table pet_allergy_record for DAO operations like Where/Data.
type PetAllergyRecord struct {
	g.Meta             `orm:"table:pet_allergy_record, do:true"`
	Id                 any // 过敏记录ID
	PetId              any // 宠物ID
	Allergen           any // 过敏源
	SymptomDescription any // 症状描述
	SeverityLevel      any // 严重程度：1轻微，2中等，3严重
	Remark             any // 备注
	CreatedAt          any // 创建时间
	UpdatedAt          any // 更新时间
}
