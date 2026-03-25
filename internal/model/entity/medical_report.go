// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// MedicalReport is the golang structure for table medical_report.
type MedicalReport struct {
	Id              int64     `json:"id"              orm:"id"                description:"报告主键ID"`               // 报告主键ID
	MedicalRecordId int64     `json:"medicalRecordId" orm:"medical_record_id" description:"病历ID"`                 // 病历ID
	DoctorId        int64     `json:"doctorId"        orm:"doctor_id"         description:"上传医生ID"`               // 上传医生ID
	ReportTitle     string    `json:"reportTitle"     orm:"report_title"      description:"报告标题"`                 // 报告标题
	ReportType      string    `json:"reportType"      orm:"report_type"       description:"报告类型，如检查报告、检验报告、处置报告"` // 报告类型，如检查报告、检验报告、处置报告
	FileUrl         string    `json:"fileUrl"         orm:"file_url"          description:"报告文件路径"`               // 报告文件路径
	ReportContent   string    `json:"reportContent"   orm:"report_content"    description:"报告文字内容"`               // 报告文字内容
	UploadedAt      time.Time `json:"uploadedAt"      orm:"uploaded_at"       description:"上传时间"`                 // 上传时间
}
