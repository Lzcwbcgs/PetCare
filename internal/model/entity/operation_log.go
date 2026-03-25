// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// OperationLog is the golang structure for table operation_log.
type OperationLog struct {
	Id              int64     `json:"id"              orm:"id"               description:"日志ID"`               // 日志ID
	OperatorType    int       `json:"operatorType"    orm:"operator_type"    description:"操作人类型：1管理员，2医生，3用户"` // 操作人类型：1管理员，2医生，3用户
	OperatorId      int64     `json:"operatorId"      orm:"operator_id"      description:"操作人ID"`              // 操作人ID
	OperationModule string    `json:"operationModule" orm:"operation_module" description:"操作模块"`               // 操作模块
	OperationType   string    `json:"operationType"   orm:"operation_type"   description:"操作类型"`               // 操作类型
	OperationDesc   string    `json:"operationDesc"   orm:"operation_desc"   description:"操作描述"`               // 操作描述
	IpAddress       string    `json:"ipAddress"       orm:"ip_address"       description:"IP地址"`               // IP地址
	CreatedAt       time.Time `json:"createdAt"       orm:"created_at"       description:"操作时间"`               // 操作时间
}
