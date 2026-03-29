package v1

import "github.com/gogf/gf/v2/frame/g"

type MedicalHistoryCreateReq struct {
	g.Meta      `path:"/{pet_id}/medical-histories" method:"post" tags:"宠物病史" summary:"新增病史记录"`
	PetID       int64   `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物ID"`
	HistoryType *string `json:"history_type" v:"max-length:50#病史类型长度不能超过50" dc:"病史类型"`
	Description string  `json:"description" v:"required#请填写病史描述" dc:"病史描述"`
	DiagnosedAt *string `json:"diagnosed_at" v:"datetime#确诊时间格式不正确" dc:"确诊时间 (yyyy-MM-dd HH:mm:ss)"`
	IsCurrent   *int    `json:"is_current" v:"in:0,1#是否当前存在不合法" dc:"是否当前仍存在：1是，0否"`
}

type MedicalHistoryCreateRes struct {
	ID int64 `json:"id" dc:"新建病史记录ID"`
}

type MedicalHistoryListReq struct {
	g.Meta    `path:"/{pet_id}/medical-histories" method:"get" tags:"宠物病史" summary:"获取病史列表"`
	PetID     int64 `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物ID"`
	Page      *int  `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	Size      *int  `json:"size" v:"min:1|max:100#每页数量必须在1-100之间" dc:"每页数量"`
	IsCurrent *int  `json:"is_current" v:"in:0,1#状态不合法" dc:"是否当前仍存在：1是，0否"`
}

type MedicalHistoryItem struct {
	ID          int64  `json:"id" dc:"病史记录ID"`
	PetID       int64  `json:"pet_id" dc:"宠物ID"`
	HistoryType string `json:"history_type" dc:"病史类型"`
	Description string `json:"description" dc:"病史描述"`
	DiagnosedAt string `json:"diagnosed_at" dc:"确诊时间"`
	IsCurrent   int    `json:"is_current" dc:"是否当前仍存在：1是，0否"`
	CreatedAt   string `json:"created_at" dc:"创建时间"`
	UpdatedAt   string `json:"updated_at" dc:"更新时间"`
}

type MedicalHistoryListRes struct {
	List  []MedicalHistoryItem `json:"list" dc:"病史列表"`
	Total int                  `json:"total" dc:"总数"`
	Page  int                  `json:"page" dc:"当前页"`
	Size  int                  `json:"size" dc:"每页数量"`
}
