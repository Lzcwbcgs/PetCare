// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MedicalRecord is the golang structure of table medical_record for DAO operations like Where/Data.
type MedicalRecord struct {
	g.Meta               `orm:"table:medical_record, do:true"`
	Id                   any // 病历主键ID
	AppointmentId        any // 关联预约ID
	PetId                any // 宠物ID
	UserId               any // 用户ID
	DoctorId             any // 接诊医生ID
	ChiefComplaint       any // 主诉
	PresentHistory       any // 现病史
	PhysicalExamination  any // 体格检查结果
	PreliminaryDiagnosis any // 初步诊断
	TreatmentPlan        any // 治疗方案
	Prescription         any // 处方建议
	DoctorAdvice         any // 医嘱
	VisitTime            any // 就诊时间
	Status               any // 状态：1已创建，2已完成，3已归档
	CreatedAt            any // 创建时间
	UpdatedAt            any // 更新时间
}
