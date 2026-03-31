package v1

import "github.com/gogf/gf/v2/frame/g"

type Pagination struct {
	Page     int `json:"page" dc:"当前页码"`
	PageSize int `json:"page_size" dc:"每页数量"`
	Total    int `json:"total" dc:"总数"`
}

type ListReq struct {
	g.Meta   `path:"/" method:"get" tags:"通知模块" summary:"获取通知列表"`
	Page     *int `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	Status   *int `json:"status" p:"status" v:"in:0,1,2,3#通知状态不合法" dc:"状态"`
}

type ListItem struct {
	ID               int64  `json:"id" dc:"通知ID"`
	NotificationType int    `json:"notification_type" dc:"通知类型"`
	Title            string `json:"title" dc:"通知标题"`
	Content          string `json:"content" dc:"通知内容"`
	AppointmentID    int64  `json:"appointment_id" dc:"关联预约ID"`
	SendTime         string `json:"send_time" dc:"发送时间"`
	Status           int    `json:"status" dc:"状态"`
}

type ListRes struct {
	List       []ListItem `json:"list" dc:"通知列表"`
	Pagination Pagination `json:"pagination" dc:"分页信息"`
}

type DetailReq struct {
	g.Meta         `path:"/{notification_id}" method:"get" tags:"通知模块" summary:"获取通知详情"`
	NotificationID int64 `json:"notification_id" p:"notification_id" v:"required|min:1#通知ID不能为空|通知ID不合法" dc:"通知ID"`
}

type DetailRes struct {
	ID               int64  `json:"id" dc:"通知ID"`
	NotificationType int    `json:"notification_type" dc:"通知类型"`
	Title            string `json:"title" dc:"通知标题"`
	Content          string `json:"content" dc:"通知内容"`
	AppointmentID    int64  `json:"appointment_id" dc:"关联预约ID"`
	SendTime         string `json:"send_time" dc:"发送时间"`
	Status           int    `json:"status" dc:"状态"`
}

type ReadReq struct {
	g.Meta         `path:"/{notification_id}/read" method:"put" tags:"通知模块" summary:"标记通知已读"`
	NotificationID int64 `json:"notification_id" p:"notification_id" v:"required|min:1#通知ID不能为空|通知ID不合法" dc:"通知ID"`
}

type ReadRes struct{}
