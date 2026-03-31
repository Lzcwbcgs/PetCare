package v1

import "github.com/gogf/gf/v2/frame/g"

type Pagination struct {
	Page     int `json:"page" dc:"当前页码"`
	PageSize int `json:"page_size" dc:"每页数量"`
	Total    int `json:"total" dc:"总数"`
}

type MedicalRecordListReq struct {
	g.Meta   `path:"/" method:"get" tags:"病历与报告查看" summary:"获取我的宠物病历列表"`
	Page     *int   `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int   `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	PetID    *int64 `json:"pet_id" p:"pet_id" v:"min:1#宠物ID不合法" dc:"宠物ID"`
}

type MedicalRecordListItem struct {
	ID                   int64  `json:"id" dc:"病历ID"`
	AppointmentID        int64  `json:"appointment_id" dc:"预约ID"`
	PetID                int64  `json:"pet_id" dc:"宠物ID"`
	DoctorID             int64  `json:"doctor_id" dc:"医生ID"`
	DoctorName           string `json:"doctor_name" dc:"医生姓名"`
	PreliminaryDiagnosis string `json:"preliminary_diagnosis" dc:"初步诊断"`
	VisitTime            string `json:"visit_time" dc:"就诊时间"`
	Status               int    `json:"status" dc:"状态"`
}

type MedicalRecordListRes struct {
	List       []MedicalRecordListItem `json:"list" dc:"病历列表"`
	Pagination Pagination              `json:"pagination" dc:"分页信息"`
}

type MedicalRecordDetailReq struct {
	g.Meta          `path:"/{medical_record_id}" method:"get" tags:"病历与报告查看" summary:"获取病历详情"`
	MedicalRecordID int64 `json:"medical_record_id" p:"medical_record_id" v:"required|min:1#病历ID不能为空|病历ID不合法" dc:"病历ID"`
}

type MedicalRecordDetailRes struct {
	ID                   int64  `json:"id" dc:"病历ID"`
	AppointmentID        int64  `json:"appointment_id" dc:"预约ID"`
	PetID                int64  `json:"pet_id" dc:"宠物ID"`
	UserID               int64  `json:"user_id" dc:"用户ID"`
	DoctorID             int64  `json:"doctor_id" dc:"医生ID"`
	ChiefComplaint       string `json:"chief_complaint" dc:"主诉"`
	PresentHistory       string `json:"present_history" dc:"现病史"`
	PhysicalExamination  string `json:"physical_examination" dc:"体格检查结果"`
	PreliminaryDiagnosis string `json:"preliminary_diagnosis" dc:"初步诊断"`
	TreatmentPlan        string `json:"treatment_plan" dc:"治疗方案"`
	Prescription         string `json:"prescription" dc:"处方建议"`
	DoctorAdvice         string `json:"doctor_advice" dc:"医嘱"`
	VisitTime            string `json:"visit_time" dc:"就诊时间"`
	Status               int    `json:"status" dc:"状态"`
}

type MedicalRecordReportListReq struct {
	g.Meta          `path:"/{medical_record_id}/reports" method:"get" tags:"病历与报告查看" summary:"获取病历报告列表"`
	MedicalRecordID int64 `json:"medical_record_id" p:"medical_record_id" v:"required|min:1#病历ID不能为空|病历ID不合法" dc:"病历ID"`
}

type MedicalRecordReportListItem struct {
	ID            int64  `json:"id" dc:"报告ID"`
	ReportTitle   string `json:"report_title" dc:"报告标题"`
	ReportType    string `json:"report_type" dc:"报告类型"`
	FileURL       string `json:"file_url" dc:"报告文件地址"`
	ReportContent string `json:"report_content" dc:"报告文字内容"`
	UploadedAt    string `json:"uploaded_at" dc:"上传时间"`
}

type MedicalRecordReportListRes struct {
	List []MedicalRecordReportListItem `json:"list" dc:"报告列表"`
}
