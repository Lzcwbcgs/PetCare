// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// MedicalRecord is the golang structure for table medical_record.
type MedicalRecord struct {
	Id                   int64     `json:"id"                   orm:"id"                    description:"病历主键ID"`            // 病历主键ID
	AppointmentId        int64     `json:"appointmentId"        orm:"appointment_id"        description:"关联预约ID"`            // 关联预约ID
	PetId                int64     `json:"petId"                orm:"pet_id"                description:"宠物ID"`              // 宠物ID
	UserId               int64     `json:"userId"               orm:"user_id"               description:"用户ID"`              // 用户ID
	DoctorId             int64     `json:"doctorId"             orm:"doctor_id"             description:"接诊医生ID"`            // 接诊医生ID
	ChiefComplaint       string    `json:"chiefComplaint"       orm:"chief_complaint"       description:"主诉"`                // 主诉
	PresentHistory       string    `json:"presentHistory"       orm:"present_history"       description:"现病史"`               // 现病史
	PhysicalExamination  string    `json:"physicalExamination"  orm:"physical_examination"  description:"体格检查结果"`            // 体格检查结果
	PreliminaryDiagnosis string    `json:"preliminaryDiagnosis" orm:"preliminary_diagnosis" description:"初步诊断"`              // 初步诊断
	TreatmentPlan        string    `json:"treatmentPlan"        orm:"treatment_plan"        description:"治疗方案"`              // 治疗方案
	Prescription         string    `json:"prescription"         orm:"prescription"          description:"处方建议"`              // 处方建议
	DoctorAdvice         string    `json:"doctorAdvice"         orm:"doctor_advice"         description:"医嘱"`                // 医嘱
	VisitTime            time.Time `json:"visitTime"            orm:"visit_time"            description:"就诊时间"`              // 就诊时间
	Status               int       `json:"status"               orm:"status"                description:"状态：1已创建，2已完成，3已归档"` // 状态：1已创建，2已完成，3已归档
	CreatedAt            time.Time `json:"createdAt"            orm:"created_at"            description:"创建时间"`              // 创建时间
	UpdatedAt            time.Time `json:"updatedAt"            orm:"updated_at"            description:"更新时间"`              // 更新时间
}
