package v1

import "github.com/gogf/gf/v2/frame/g"

type NotificationCreateReq struct {
	g.Meta           `path:"/notifications" method:"post" tags:"系统通知管理" summary:"发布系统通知"`
	ReceiverType     string  `json:"receiver_type" v:"required|in:user,doctor#请填写接收人类型|接收人类型不合法" dc:"接收人类型：user/doctor"`
	ReceiverIDs      []int64 `json:"receiver_ids" v:"required#请填写接收人ID列表" dc:"接收人ID列表"`
	NotificationType int     `json:"notification_type" v:"required|in:2#请填写通知类型|通知类型不合法" dc:"通知类型：2系统通知"`
	Title            string  `json:"title" v:"required|length:1,255#请填写通知标题|通知标题长度需在1到255位之间" dc:"通知标题"`
	Content          string  `json:"content" v:"required|max-length:65535#请填写通知内容|通知内容长度不能超过65535" dc:"通知内容"`
	SendTime         string  `json:"send_time" v:"required|datetime#请填写发送时间|发送时间格式不正确" dc:"发送时间 (yyyy-MM-dd HH:mm:ss)"`
}

type NotificationCreateRes struct {
	Count int `json:"count" dc:"发布成功数量"`
}

type NotificationListReq struct {
	g.Meta           `path:"/notifications" method:"get" tags:"系统通知管理" summary:"获取系统通知列表"`
	Page             *int `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize         *int `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	NotificationType *int `json:"notification_type" p:"notification_type" v:"in:2#通知类型不合法" dc:"通知类型：2系统通知"`
}

type NotificationListItem struct {
	ID               int64  `json:"id" dc:"通知ID"`
	NotificationType int    `json:"notification_type" dc:"通知类型"`
	Title            string `json:"title" dc:"通知标题"`
	Content          string `json:"content" dc:"通知内容"`
	SendTime         string `json:"send_time" dc:"发送时间"`
	Status           int    `json:"status" dc:"状态"`
	CreatedAt        string `json:"created_at" dc:"创建时间"`
}

type NotificationListRes struct {
	List       []NotificationListItem `json:"list" dc:"通知列表"`
	Pagination Pagination             `json:"pagination" dc:"分页信息"`
}

type NotificationDeleteReq struct {
	g.Meta         `path:"/notifications/{notification_id}" method:"delete" tags:"系统通知管理" summary:"撤回系统通知"`
	NotificationID int64 `json:"notification_id" p:"notification_id" v:"required|min:1#通知ID不能为空|通知ID不合法" dc:"通知ID"`
}

type NotificationDeleteRes struct{}
