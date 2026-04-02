package v1

import "github.com/gogf/gf/v2/frame/g"

type OperationLogListReq struct {
	g.Meta          `path:"/operation-logs" method:"get" tags:"操作日志模块" summary:"获取操作日志列表"`
	Page            *int    `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize        *int    `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	OperatorType    *int    `json:"operator_type" p:"operator_type" v:"in:1,2,3#操作人类型不合法" dc:"操作人类型：1管理员，2医生，3用户"`
	OperationModule *string `json:"operation_module" p:"operation_module" v:"max-length:50#操作模块长度不能超过50" dc:"操作模块"`
}

type OperationLogListItem struct {
	ID              int64  `json:"id" dc:"日志ID"`
	OperatorType    int    `json:"operator_type" dc:"操作人类型"`
	OperatorID      int64  `json:"operator_id" dc:"操作人ID"`
	OperationModule string `json:"operation_module" dc:"操作模块"`
	OperationType   string `json:"operation_type" dc:"操作类型"`
	OperationDesc   string `json:"operation_desc" dc:"操作描述"`
	IPAddress       string `json:"ip_address" dc:"IP地址"`
	CreatedAt       string `json:"created_at" dc:"创建时间"`
}

type OperationLogListRes struct {
	List       []OperationLogListItem `json:"list" dc:"日志列表"`
	Pagination Pagination             `json:"pagination" dc:"分页信息"`
}

type OperationLogDetailReq struct {
	g.Meta `path:"/operation-logs/{log_id}" method:"get" tags:"操作日志模块" summary:"获取操作日志详情"`
	LogID  int64 `json:"log_id" p:"log_id" v:"required|min:1#日志ID不能为空|日志ID不合法" dc:"日志ID"`
}

type OperationLogDetailRes struct {
	ID              int64  `json:"id" dc:"日志ID"`
	OperatorType    int    `json:"operator_type" dc:"操作人类型"`
	OperatorID      int64  `json:"operator_id" dc:"操作人ID"`
	OperationModule string `json:"operation_module" dc:"操作模块"`
	OperationType   string `json:"operation_type" dc:"操作类型"`
	OperationDesc   string `json:"operation_desc" dc:"操作描述"`
	IPAddress       string `json:"ip_address" dc:"IP地址"`
	CreatedAt       string `json:"created_at" dc:"创建时间"`
}
