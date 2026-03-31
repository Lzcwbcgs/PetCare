package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type MedicalReportUploadReq struct {
	g.Meta          `path:"/medical-records/{medical_record_id}/reports" method:"post" mime:"multipart/form-data" tags:"医疗报告模块" summary:"上传医疗报告"`
	MedicalRecordID int64             `json:"medical_record_id" p:"medical_record_id" v:"required|min:1#病历ID不能为空|病历ID不合法" dc:"病历ID"`
	ReportTitle     string            `json:"report_title" v:"required|max-length:255#请填写报告标题|报告标题长度不能超过255" dc:"报告标题"`
	ReportType      string            `json:"report_type" v:"max-length:100#报告类型长度不能超过100" dc:"报告类型"`
	ReportContent   string            `json:"report_content" v:"max-length:65535#报告内容长度不能超过65535" dc:"报告文字内容"`
	File            *ghttp.UploadFile `json:"file" type:"file" dc:"报告附件"`
}

type MedicalReportUploadRes struct {
	ReportID int64  `json:"report_id" dc:"报告ID"`
	FileURL  string `json:"file_url" dc:"报告文件地址"`
}

type MedicalReportListReq struct {
	g.Meta          `path:"/medical-records/{medical_record_id}/reports" method:"get" tags:"医疗报告模块" summary:"获取报告列表"`
	MedicalRecordID int64 `json:"medical_record_id" p:"medical_record_id" v:"required|min:1#病历ID不能为空|病历ID不合法" dc:"病历ID"`
}

type MedicalReportListItem struct {
	ID            int64  `json:"id" dc:"报告ID"`
	ReportTitle   string `json:"report_title" dc:"报告标题"`
	ReportType    string `json:"report_type" dc:"报告类型"`
	FileURL       string `json:"file_url" dc:"报告文件地址"`
	ReportContent string `json:"report_content" dc:"报告文字内容"`
	UploadedAt    string `json:"uploaded_at" dc:"上传时间"`
}

type MedicalReportListRes struct {
	List []MedicalReportListItem `json:"list" dc:"报告列表"`
}

type MedicalReportDeleteReq struct {
	g.Meta   `path:"/reports/{report_id}" method:"delete" tags:"医疗报告模块" summary:"删除医疗报告"`
	ReportID int64 `json:"report_id" p:"report_id" v:"required|min:1#报告ID不能为空|报告ID不合法" dc:"报告ID"`
}

type MedicalReportDeleteRes struct{}
