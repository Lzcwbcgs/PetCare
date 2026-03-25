// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MedicalRecordDao is the data access object for the table medical_record.
type MedicalRecordDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  MedicalRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// MedicalRecordColumns defines and stores column names for the table medical_record.
type MedicalRecordColumns struct {
	Id                   string // 病历主键ID
	AppointmentId        string // 关联预约ID
	PetId                string // 宠物ID
	UserId               string // 用户ID
	DoctorId             string // 接诊医生ID
	ChiefComplaint       string // 主诉
	PresentHistory       string // 现病史
	PhysicalExamination  string // 体格检查结果
	PreliminaryDiagnosis string // 初步诊断
	TreatmentPlan        string // 治疗方案
	Prescription         string // 处方建议
	DoctorAdvice         string // 医嘱
	VisitTime            string // 就诊时间
	Status               string // 状态：1已创建，2已完成，3已归档
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
}

// medicalRecordColumns holds the columns for the table medical_record.
var medicalRecordColumns = MedicalRecordColumns{
	Id:                   "id",
	AppointmentId:        "appointment_id",
	PetId:                "pet_id",
	UserId:               "user_id",
	DoctorId:             "doctor_id",
	ChiefComplaint:       "chief_complaint",
	PresentHistory:       "present_history",
	PhysicalExamination:  "physical_examination",
	PreliminaryDiagnosis: "preliminary_diagnosis",
	TreatmentPlan:        "treatment_plan",
	Prescription:         "prescription",
	DoctorAdvice:         "doctor_advice",
	VisitTime:            "visit_time",
	Status:               "status",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewMedicalRecordDao creates and returns a new DAO object for table data access.
func NewMedicalRecordDao(handlers ...gdb.ModelHandler) *MedicalRecordDao {
	return &MedicalRecordDao{
		group:    "default",
		table:    "medical_record",
		columns:  medicalRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MedicalRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MedicalRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MedicalRecordDao) Columns() MedicalRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MedicalRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MedicalRecordDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MedicalRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
