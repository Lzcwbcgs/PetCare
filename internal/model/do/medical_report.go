// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MedicalReport is the golang structure of table medical_report for DAO operations like Where/Data.
type MedicalReport struct {
	g.Meta          `orm:"table:medical_report, do:true"`
	Id              any // 报告主键ID
	MedicalRecordId any // 病历ID
	DoctorId        any // 上传医生ID
	ReportTitle     any // 报告标题
	ReportType      any // 报告类型，如检查报告、检验报告、处置报告
	FileUrl         any // 报告文件路径
	ReportContent   any // 报告文字内容
	UploadedAt      any // 上传时间
}
