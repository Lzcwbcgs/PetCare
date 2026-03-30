package v1

import "github.com/gogf/gf/v2/frame/g"

type AllergyCreateReq struct {
	g.Meta             `path:"/{pet_id}/allergies" method:"post" tags:"宠物过敏" summary:"新增过敏记录"`
	PetID              int64   `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物ID"`
	Allergen           string  `json:"allergen" v:"required|max-length:100#请填写过敏源|过敏源长度不能超过100" dc:"过敏源"`
	SymptomDescription *string `json:"symptom_description" v:"max-length:255#症状描述长度不能超过255" dc:"症状描述"`
	SeverityLevel      *int    `json:"severity_level" v:"in:1,2,3#严重程度不合法" dc:"严重程度，1轻微，2中等，3严重"`
	Remark             *string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`
}

type AllergyCreateRes struct {
	ID int64 `json:"id" dc:"新建过敏记录ID"`
}

type AllergyListReq struct {
	g.Meta        `path:"/{pet_id}/allergies" method:"get" tags:"宠物过敏" summary:"获取过敏记录列表"`
	PetID         int64 `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物ID"`
	Page          *int  `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize      *int  `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1-100之间" dc:"每页数量"`
	SeverityLevel *int  `json:"severity_level" p:"severity_level" v:"in:1,2,3#严重程度不合法" dc:"严重程度，1轻微，2中等，3严重"`
}

type AllergyItem struct {
	ID                 int64  `json:"id" dc:"过敏记录ID"`
	Allergen           string `json:"allergen" dc:"过敏源"`
	SymptomDescription string `json:"symptom_description" dc:"症状描述"`
	SeverityLevel      int    `json:"severity_level" dc:"严重程度，1轻微，2中等，3严重"`
	Remark             string `json:"remark" dc:"备注"`
}

type AllergyListRes struct {
	List       []AllergyItem `json:"list" dc:"过敏记录列表"`
	Pagination Pagination    `json:"pagination" dc:"分页信息"`
}
