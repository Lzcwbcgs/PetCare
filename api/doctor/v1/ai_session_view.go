package v1

import "github.com/gogf/gf/v2/frame/g"

type AISessionListReq struct {
	g.Meta   `path:"/sessions" method:"get" tags:"医生AI查看" summary:"获取AI会话列表"`
	Page     *int   `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int   `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	PetID    *int64 `json:"pet_id" p:"pet_id" v:"min:1#pet_id不合法" dc:"宠物ID筛选"`
}

type AISessionListItem struct {
	ID        int64  `json:"id" dc:"会话ID"`
	SessionNo string `json:"session_no" dc:"会话编号"`
	PetID     int64  `json:"pet_id" dc:"宠物ID"`
	PetName   string `json:"pet_name" dc:"宠物名称"`
	ModelName string `json:"model_name" dc:"模型名称"`
	Status    int    `json:"status" dc:"状态"`
	CreatedAt string `json:"created_at" dc:"创建时间"`
}

type AISessionListRes struct {
	List       []AISessionListItem `json:"list" dc:"会话列表"`
	Pagination Pagination          `json:"pagination" dc:"分页信息"`
}
