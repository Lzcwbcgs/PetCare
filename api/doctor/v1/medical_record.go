package v1

import "github.com/gogf/gf/v2/frame/g"

type MedicalRecordCreateReq struct {
	g.Meta               `path:"/medical-records" method:"post" tags:"病历管理" summary:"创建病历记录"`
	AppointmentID        int64  `json:"appointment_id" v:"required|min:1#请填写预约ID|预约ID不合法" dc:"关联预约ID"`
	PetID                int64  `json:"pet_id" v:"required|min:1#请填写宠物ID|宠物ID不合法" dc:"宠物ID"`
	UserID               int64  `json:"user_id" v:"required|min:1#请填写用户ID|用户ID不合法" dc:"用户ID"`
	ChiefComplaint       string `json:"chief_complaint" v:"required|max-length:65535#请填写主诉|主诉长度不能超过65535" dc:"主诉"`
	PresentHistory       string `json:"present_history" v:"required|max-length:65535#请填写现病史|现病史长度不能超过65535" dc:"现病史"`
	PhysicalExamination  string `json:"physical_examination" v:"required|max-length:65535#请填写体格检查结果|体格检查结果长度不能超过65535" dc:"体格检查结果"`
	PreliminaryDiagnosis string `json:"preliminary_diagnosis" v:"required|max-length:65535#请填写初步诊断|初步诊断长度不能超过65535" dc:"初步诊断"`
	TreatmentPlan        string `json:"treatment_plan" v:"required|max-length:65535#请填写治疗方案|治疗方案长度不能超过65535" dc:"治疗方案"`
	Prescription         string `json:"prescription" v:"required|max-length:65535#请填写处方建议|处方建议长度不能超过65535" dc:"处方建议"`
	DoctorAdvice         string `json:"doctor_advice" v:"required|max-length:65535#请填写医嘱|医嘱长度不能超过65535" dc:"医嘱"`
	VisitTime            string `json:"visit_time" v:"required|datetime#请填写就诊时间|就诊时间格式不正确" dc:"就诊时间 (yyyy-MM-dd HH:mm:ss)"`
	Status               int    `json:"status" v:"required|in:1,2,3#请填写病历状态|病历状态不合法" dc:"状态：1已创建，2已完成，3已归档"`
}

type MedicalRecordCreateRes struct {
	MedicalRecordID int64 `json:"medical_record_id" dc:"病历ID"`
}

type MedicalRecordUpdateReq struct {
	g.Meta               `path:"/medical-records/{medical_record_id}" method:"put" tags:"病历管理" summary:"更新病历记录"`
	MedicalRecordID      int64   `json:"medical_record_id" p:"medical_record_id" v:"required|min:1#病历ID不能为空|病历ID不合法" dc:"病历ID"`
	ChiefComplaint       *string `json:"chief_complaint" v:"max-length:65535#主诉长度不能超过65535" dc:"主诉"`
	PresentHistory       *string `json:"present_history" v:"max-length:65535#现病史长度不能超过65535" dc:"现病史"`
	PhysicalExamination  *string `json:"physical_examination" v:"max-length:65535#体格检查结果长度不能超过65535" dc:"体格检查结果"`
	PreliminaryDiagnosis *string `json:"preliminary_diagnosis" v:"max-length:65535#初步诊断长度不能超过65535" dc:"初步诊断"`
	TreatmentPlan        *string `json:"treatment_plan" v:"max-length:65535#治疗方案长度不能超过65535" dc:"治疗方案"`
	Prescription         *string `json:"prescription" v:"max-length:65535#处方建议长度不能超过65535" dc:"处方建议"`
	DoctorAdvice         *string `json:"doctor_advice" v:"max-length:65535#医嘱长度不能超过65535" dc:"医嘱"`
	VisitTime            *string `json:"visit_time" v:"datetime#就诊时间格式不正确" dc:"就诊时间 (yyyy-MM-dd HH:mm:ss)"`
	Status               *int    `json:"status" v:"in:1,2,3#病历状态不合法" dc:"状态：1已创建，2已完成，3已归档"`
}

type MedicalRecordUpdateRes struct{}

type MedicalRecordDetailReq struct {
	g.Meta          `path:"/medical-records/{medical_record_id}" method:"get" tags:"病历管理" summary:"获取病历详情"`
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
