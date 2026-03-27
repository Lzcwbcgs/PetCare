package v1

import "github.com/gogf/gf/v2/frame/g"

type AdminProfileReq struct {
	g.Meta `path:"/profile" method:"get" tags:"管理员资料" summary:"获取管理员个人信息"`
}

type AdminProfileRes struct {
	ID        int64  `json:"id" dc:"管理员ID"`
	Username  string `json:"username" dc:"用户名"`
	RealName  string `json:"real_name" dc:"姓名"`
	Phone     string `json:"phone" dc:"联系电话"`
	Status    int    `json:"status" dc:"状态"`
	CreatedAt string `json:"created_at" dc:"创建时间"`
}
