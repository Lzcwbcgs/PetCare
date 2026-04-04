package v1

import "github.com/gogf/gf/v2/frame/g"

type AISessionMessageListReq struct {
	g.Meta    `path:"/sessions/{session_id}/messages" method:"get" tags:"医生AI查看" summary:"获取AI会话消息记录"`
	SessionID int64 `json:"session_id" p:"session_id" v:"required|min:1#会话ID不能为空|会话ID不合法" dc:"会话ID"`
	Page      *int  `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize  *int  `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
}

type AISessionMessageItem struct {
	ID             int64  `json:"id" dc:"消息ID"`
	SenderType     int    `json:"sender_type" dc:"发送者类型"`
	SenderID       *int64 `json:"sender_id" dc:"发送者ID"`
	MessageContent string `json:"message_content" dc:"消息内容"`
	CreatedAt      string `json:"created_at" dc:"创建时间"`
}

type AISessionMessageListRes struct {
	List []AISessionMessageItem `json:"list" dc:"消息列表"`
}
