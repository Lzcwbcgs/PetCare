package v1

import "github.com/gogf/gf/v2/frame/g"

type PetCreateReq struct {
	g.Meta     `path:"/" method:"post" tags:"宠物管理" summary:"新增宠物档案"`
	PetName    string  `json:"pet_name" v:"required|length:1,64#请输入宠物名称|宠物名称长度需在1到64位之间" dc:"宠物名称"`
	PetType    string  `json:"pet_type" v:"length:0,32#宠物类型长度不能超过32位" dc:"宠物类型"`
	AvatarURL  string  `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"头像地址"`
	Gender     int     `json:"gender" v:"in:0,1,2#性别取值不合法" dc:"性别：1公，2母，0未知"`
	Age        int     `json:"age" v:"min:0#年龄不能为负" dc:"年龄"`
	AgeUnit    *string `json:"age_unit" v:"in:month,year#年龄单位不合法" dc:"年龄单位：month/月，year/岁"`
	Breed      string  `json:"breed" v:"max-length:64#品种长度不能超过64位" dc:"品种"`
	Weight     float64 `json:"weight" v:"min:0#体重不能为负" dc:"体重(kg)"`
	Sterilized int     `json:"sterilized" v:"in:0,1#绝育状态不合法" dc:"是否绝育：1是，0否"`
	Remark     string  `json:"remark" v:"max-length:255#备注长度不能超过255位" dc:"备注"`
}

type PetCreateRes struct {
	PetID int64 `json:"pet_id" dc:"新建宠物ID"`
}

type PetListReq struct {
	g.Meta   `path:"/" method:"get" tags:"宠物管理" summary:"获取宠物列表"`
	Page     int `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize int `json:"page_size" v:"min:1|max:100#每页数量必须大于0|每页数量不能超过100" dc:"每页数量"`
}

type PetListItem struct {
	ID         int64   `json:"id"`
	PetName    string  `json:"pet_name"`
	PetType    string  `json:"pet_type"`
	AvatarURL  string  `json:"avatar_url"`
	Gender     int     `json:"gender"`
	Age        int     `json:"age"`
	AgeUnit    string  `json:"age_unit"`
	Breed      string  `json:"breed"`
	Weight     float64 `json:"weight"`
	Sterilized int     `json:"sterilized"`
	Remark     string  `json:"remark"`
	CreatedAt  string  `json:"created_at"`
}

type PetListRes struct {
	List     []PetListItem `json:"list"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
}

type PetDetailReq struct {
	g.Meta `path:"/{pet_id}" method:"get" tags:"宠物管理" summary:"获取宠物详情"`
	PetID  int64 `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
}

type PetDetailRes struct {
	ID         int64   `json:"id"`
	UserID     int64   `json:"user_id"`
	PetName    string  `json:"pet_name"`
	PetType    string  `json:"pet_type"`
	AvatarURL  string  `json:"avatar_url"`
	Gender     int     `json:"gender"`
	Age        int     `json:"age"`
	AgeUnit    string  `json:"age_unit"`
	Breed      string  `json:"breed"`
	Weight     float64 `json:"weight"`
	Sterilized int     `json:"sterilized"`
	Remark     string  `json:"remark"`
	Status     int     `json:"status"`
	CreatedAt  string  `json:"created_at"`
}

type PetUpdateReq struct {
	g.Meta     `path:"/{pet_id}" method:"put" tags:"宠物管理" summary:"修改宠物档案"`
	PetID      int64    `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
	PetName    *string  `json:"pet_name" v:"length:1,64#宠物名称长度需在1到64位之间" dc:"宠物名称"`
	PetType    *string  `json:"pet_type" v:"length:0,32#宠物类型长度不能超过32位" dc:"宠物类型"`
	AvatarURL  *string  `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"头像地址"`
	Gender     *int     `json:"gender" v:"in:0,1,2#性别取值不合法" dc:"性别：1公，2母，0未知"`
	Age        *int     `json:"age" v:"min:0#年龄不能为负" dc:"年龄"`
	AgeUnit    *string  `json:"age_unit" v:"in:month,year#年龄单位不合法" dc:"年龄单位：month/月，year/岁"`
	Breed      *string  `json:"breed" v:"max-length:64#品种长度不能超过64位" dc:"品种"`
	Weight     *float64 `json:"weight" v:"min:0#体重不能为负" dc:"体重(kg)"`
	Sterilized *int     `json:"sterilized" v:"in:0,1#绝育状态不合法" dc:"是否绝育：1是，0否"`
	Remark     *string  `json:"remark" v:"max-length:255#备注长度不能超过255位" dc:"备注"`
}

type PetUpdateRes struct{}

type PetDeleteReq struct {
	g.Meta `path:"/{pet_id}" method:"delete" tags:"宠物管理" summary:"删除宠物档案"`
	PetID  int64 `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
}

type PetDeleteRes struct{}

// --- 病史 ---

type MedicalHistoryCreateReq struct {
	g.Meta      `path:"/{pet_id}/medical-histories" method:"post" tags:"宠物健康" summary:"新增病史记录"`
	PetID       int64  `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
	HistoryType string `json:"history_type" v:"required|length:1,64#请输入病史类型|病史类型长度需在1到64位之间" dc:"病史类型"`
	Description string `json:"description" v:"required|max-length:65535#请输入病史描述|病史描述过长" dc:"病史描述"`
	DiagnosedAt string `json:"diagnosed_at" v:"required|date-format:Y-m-d#请输入确诊日期|确诊日期格式应为YYYY-MM-DD" dc:"确诊日期"`
	IsCurrent   int    `json:"is_current" v:"in:0,1#是否当前仍存在取值不合法" dc:"是否当前仍存在：1是，0否"`
}

type MedicalHistoryCreateRes struct {
	RecordID int64 `json:"record_id" dc:"病史记录ID"`
}

type MedicalHistoryListReq struct {
	g.Meta   `path:"/{pet_id}/medical-histories" method:"get" tags:"宠物健康" summary:"获取病史列表"`
	PetID    int64 `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
	Page     int   `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize int   `json:"page_size" v:"min:1|max:100#每页数量必须大于0|每页数量不能超过100" dc:"每页数量"`
}

type MedicalHistoryItem struct {
	ID          int64  `json:"id"`
	HistoryType string `json:"history_type"`
	Description string `json:"description"`
	DiagnosedAt string `json:"diagnosed_at"`
	IsCurrent   int    `json:"is_current"`
	CreatedAt   string `json:"created_at"`
}

type MedicalHistoryListRes struct {
	List     []MedicalHistoryItem `json:"list"`
	Total    int                  `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"page_size"`
}

// --- 疫苗 ---

type VaccinationCreateReq struct {
	g.Meta          `path:"/{pet_id}/vaccinations" method:"post" tags:"宠物健康" summary:"新增疫苗记录"`
	PetID           int64  `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
	VaccineName     string `json:"vaccine_name" v:"required|length:1,128#请输入疫苗名称|疫苗名称长度需在1到128位之间" dc:"疫苗名称"`
	VaccinationDate string `json:"vaccination_date" v:"required|date-format:Y-m-d#请输入接种日期|接种日期格式应为YYYY-MM-DD" dc:"接种日期"`
	NextDueDate     string `json:"next_due_date" v:"date-format:Y-m-d#下次接种日期格式应为YYYY-MM-DD" dc:"下次应接种日期"`
	HospitalName    string `json:"hospital_name" v:"max-length:128#接种机构长度不能超过128位" dc:"接种机构"`
	Remark          string `json:"remark" v:"max-length:255#备注长度不能超过255位" dc:"备注"`
}

type VaccinationCreateRes struct {
	RecordID int64 `json:"record_id" dc:"疫苗记录ID"`
}

type VaccinationListReq struct {
	g.Meta   `path:"/{pet_id}/vaccinations" method:"get" tags:"宠物健康" summary:"获取疫苗记录列表"`
	PetID    int64 `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
	Page     int   `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize int   `json:"page_size" v:"min:1|max:100#每页数量必须大于0|每页数量不能超过100" dc:"每页数量"`
}

type VaccinationItem struct {
	ID              int64  `json:"id"`
	VaccineName     string `json:"vaccine_name"`
	VaccinationDate string `json:"vaccination_date"`
	NextDueDate     string `json:"next_due_date"`
	HospitalName    string `json:"hospital_name"`
	Remark          string `json:"remark"`
	CreatedAt       string `json:"created_at"`
}

type VaccinationListRes struct {
	List     []VaccinationItem `json:"list"`
	Total    int               `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
}

// --- 过敏 ---

type AllergyCreateReq struct {
	g.Meta             `path:"/{pet_id}/allergies" method:"post" tags:"宠物健康" summary:"新增过敏记录"`
	PetID              int64  `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
	Allergen           string `json:"allergen" v:"required|length:1,128#请输入过敏源|过敏源长度需在1到128位之间" dc:"过敏源"`
	SymptomDescription string `json:"symptom_description" v:"required|max-length:1024#请输入症状描述|症状描述过长" dc:"症状描述"`
	SeverityLevel      int    `json:"severity_level" v:"in:1,2,3#严重程度取值不合法" dc:"严重程度：1轻微，2中等，3严重"`
	Remark             string `json:"remark" v:"max-length:255#备注长度不能超过255位" dc:"备注"`
}

type AllergyCreateRes struct {
	RecordID int64 `json:"record_id" dc:"过敏记录ID"`
}

type AllergyListReq struct {
	g.Meta   `path:"/{pet_id}/allergies" method:"get" tags:"宠物健康" summary:"获取过敏记录列表"`
	PetID    int64 `json:"pet_id" v:"required|min:1#宠物ID必须大于0" dc:"宠物ID"`
	Page     int   `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize int   `json:"page_size" v:"min:1|max:100#每页数量必须大于0|每页数量不能超过100" dc:"每页数量"`
}

type AllergyItem struct {
	ID                 int64  `json:"id"`
	Allergen           string `json:"allergen"`
	SymptomDescription string `json:"symptom_description"`
	SeverityLevel      int    `json:"severity_level"`
	Remark             string `json:"remark"`
	CreatedAt          string `json:"created_at"`
}

type AllergyListRes struct {
	List     []AllergyItem `json:"list"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
}
