package v1

import "github.com/gogf/gf/v2/frame/g"

type VaccinationCreateReq struct {
	g.Meta          `path:"/{pet_id}/vaccinations" method:"post" tags:"宠物疫苗" summary:"新增疫苗记录"`
	PetID           int64   `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物ID"`
	VaccineName     string  `json:"vaccine_name" v:"required|max-length:100#请填写疫苗名称|疫苗名称长度不能超过100" dc:"疫苗名称"`
	VaccinationDate *string `json:"vaccination_date" v:"date#接种日期格式不正确" dc:"接种日期 (yyyy-MM-dd)"`
	NextDueDate     *string `json:"next_due_date" v:"date#下次接种日期格式不正确" dc:"下次应接种日期 (yyyy-MM-dd)"`
	HospitalName    *string `json:"hospital_name" v:"max-length:100#接种机构长度不能超过100" dc:"接种机构"`
	Remark          *string `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`
}

type VaccinationCreateRes struct {
	ID int64 `json:"id" dc:"新建疫苗记录ID"`
}

type VaccinationListReq struct {
	g.Meta `path:"/{pet_id}/vaccinations" method:"get" tags:"宠物疫苗" summary:"获取疫苗记录列表"`
	PetID  int64 `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物ID"`
	Page   *int  `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	Size   *int  `json:"size" v:"min:1|max:100#每页数量必须在1-100之间" dc:"每页数量"`
}

type VaccinationItem struct {
	ID              int64  `json:"id" dc:"疫苗记录ID"`
	PetID           int64  `json:"pet_id" dc:"宠物ID"`
	VaccineName     string `json:"vaccine_name" dc:"疫苗名称"`
	VaccinationDate string `json:"vaccination_date" dc:"接种日期"`
	NextDueDate     string `json:"next_due_date" dc:"下次应接种日期"`
	HospitalName    string `json:"hospital_name" dc:"接种机构"`
	Remark          string `json:"remark" dc:"备注"`
	CreatedAt       string `json:"created_at" dc:"创建时间"`
	UpdatedAt       string `json:"updated_at" dc:"更新时间"`
}

type VaccinationListRes struct {
	List  []VaccinationItem `json:"list" dc:"疫苗记录列表"`
	Total int               `json:"total" dc:"总数"`
	Page  int               `json:"page" dc:"当前页"`
	Size  int               `json:"size" dc:"每页数量"`
}
