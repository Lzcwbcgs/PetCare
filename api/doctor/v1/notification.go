package v1

import "github.com/gogf/gf/v2/frame/g"

type NotificationListReq struct {
	g.Meta   `path:"/notifications" method:"get" tags:"医生通知模块" summary:"获取医生通知列表"`
	Page     *int `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
}

type NotificationListItem struct {
	ID               int64  `json:"id" dc:"通知ID"`
	NotificationType int    `json:"notification_type" dc:"通知类型"`
	Title            string `json:"title" dc:"通知标题"`
	Content          string `json:"content" dc:"通知内容"`
	AppointmentID    int64  `json:"appointment_id" dc:"关联预约ID"`
	SendTime         string `json:"send_time" dc:"发送时间"`
	Status           int    `json:"status" dc:"状态"`
}

type NotificationListRes struct {
	List       []NotificationListItem `json:"list" dc:"通知列表"`
	Pagination Pagination             `json:"pagination" dc:"分页信息"`
}

type NotificationReadReq struct {
	g.Meta         `path:"/notifications/{notification_id}/read" method:"put" tags:"医生通知模块" summary:"标记医生通知已读"`
	NotificationID int64 `json:"notification_id" p:"notification_id" v:"required|min:1#通知ID不能为空|通知ID不合法" dc:"通知ID"`
}

type NotificationReadRes struct{}
