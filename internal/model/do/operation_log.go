// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OperationLog is the golang structure of table operation_log for DAO operations like Where/Data.
type OperationLog struct {
	g.Meta          `orm:"table:operation_log, do:true"`
	Id              any // 日志ID
	OperatorType    any // 操作人类型：1管理员，2医生，3用户
	OperatorId      any // 操作人ID
	OperationModule any // 操作模块
	OperationType   any // 操作类型
	OperationDesc   any // 操作描述
	IpAddress       any // IP地址
	CreatedAt       any // 操作时间
}
