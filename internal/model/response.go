package model

// Response is the unified HTTP response structure for the project.
type Response struct {
	Code    int    `json:"code"    dc:"业务状态码"`
	Message string `json:"message" dc:"响应消息"`
	Data    any    `json:"data"    dc:"响应数据"`
}

