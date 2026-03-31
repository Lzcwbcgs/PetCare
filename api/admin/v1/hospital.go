package v1

import "github.com/gogf/gf/v2/frame/g"

type Pagination struct {
	Page     int `json:"page" dc:"当前页码"`
	PageSize int `json:"page_size" dc:"每页数量"`
	Total    int `json:"total" dc:"总数"`
}

type HospitalCreateReq struct {
	g.Meta       `path:"/hospitals" method:"post" tags:"医院管理" summary:"新增医院"`
	HospitalName string  `json:"hospital_name" v:"required|length:1,100#请填写医院名称|医院名称长度需在1到100位之间" dc:"医院名称"`
	Address      string  `json:"address" v:"required|length:1,255#请填写医院地址|医院地址长度需在1到255位之间" dc:"医院地址"`
	Phone        string  `json:"phone" v:"required|length:1,32#请填写医院联系电话|医院联系电话长度需在1到32位之间" dc:"医院联系电话"`
	Description  *string `json:"description" v:"max-length:65535#医院简介长度不能超过65535" dc:"医院简介"`
	Status       int     `json:"status" v:"required|in:0,1#请填写医院状态|医院状态不合法" dc:"状态：1启用，0停用"`
}

type HospitalCreateRes struct {
	HospitalID int64 `json:"hospital_id" dc:"新建医院ID"`
}

type HospitalListReq struct {
	g.Meta   `path:"/hospitals" method:"get" tags:"医院管理" summary:"获取医院列表"`
	Page     *int    `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int    `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	Status   *int    `json:"status" p:"status" v:"in:0,1#医院状态不合法" dc:"状态：1启用，0停用"`
	Keyword  *string `json:"keyword" p:"keyword" v:"max-length:100#关键字长度不能超过100" dc:"医院名称关键字"`
}

type HospitalListItem struct {
	ID           int64  `json:"id" dc:"医院ID"`
	HospitalName string `json:"hospital_name" dc:"医院名称"`
	Address      string `json:"address" dc:"医院地址"`
	Phone        string `json:"phone" dc:"医院联系电话"`
	Status       int    `json:"status" dc:"状态"`
	CreatedAt    string `json:"created_at" dc:"创建时间"`
}

type HospitalListRes struct {
	List       []HospitalListItem `json:"list" dc:"医院列表"`
	Pagination Pagination         `json:"pagination" dc:"分页信息"`
}

type HospitalDetailReq struct {
	g.Meta     `path:"/hospitals/{hospital_id}" method:"get" tags:"医院管理" summary:"获取医院详情"`
	HospitalID int64 `json:"hospital_id" p:"hospital_id" v:"required|min:1#医院ID不能为空|医院ID不合法" dc:"医院ID"`
}

type HospitalDetailRes struct {
	ID           int64  `json:"id" dc:"医院ID"`
	HospitalName string `json:"hospital_name" dc:"医院名称"`
	Address      string `json:"address" dc:"医院地址"`
	Phone        string `json:"phone" dc:"医院联系电话"`
	Description  string `json:"description" dc:"医院简介"`
	Status       int    `json:"status" dc:"状态"`
	CreatedAt    string `json:"created_at" dc:"创建时间"`
	UpdatedAt    string `json:"updated_at" dc:"更新时间"`
}

type HospitalUpdateReq struct {
	g.Meta       `path:"/hospitals/{hospital_id}" method:"put" tags:"医院管理" summary:"修改医院"`
	HospitalID   int64   `json:"hospital_id" p:"hospital_id" v:"required|min:1#医院ID不能为空|医院ID不合法" dc:"医院ID"`
	HospitalName *string `json:"hospital_name" v:"length:1,100#医院名称长度需在1到100位之间" dc:"医院名称"`
	Address      *string `json:"address" v:"length:1,255#医院地址长度需在1到255位之间" dc:"医院地址"`
	Phone        *string `json:"phone" v:"length:1,32#医院联系电话长度需在1到32位之间" dc:"医院联系电话"`
	Description  *string `json:"description" v:"max-length:65535#医院简介长度不能超过65535" dc:"医院简介"`
	Status       *int    `json:"status" v:"in:0,1#医院状态不合法" dc:"状态：1启用，0停用"`
}

type HospitalUpdateRes struct{}

type HospitalDeleteReq struct {
	g.Meta     `path:"/hospitals/{hospital_id}" method:"delete" tags:"医院管理" summary:"删除医院"`
	HospitalID int64 `json:"hospital_id" p:"hospital_id" v:"required|min:1#医院ID不能为空|医院ID不合法" dc:"医院ID"`
}

type HospitalDeleteRes struct{}
