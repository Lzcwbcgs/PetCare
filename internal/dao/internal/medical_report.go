// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MedicalReportDao is the data access object for the table medical_report.
type MedicalReportDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  MedicalReportColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// MedicalReportColumns defines and stores column names for the table medical_report.
type MedicalReportColumns struct {
	Id              string // 报告主键ID
	MedicalRecordId string // 病历ID
	DoctorId        string // 上传医生ID
	ReportTitle     string // 报告标题
	ReportType      string // 报告类型，如检查报告、检验报告、处置报告
	FileUrl         string // 报告文件路径
	ReportContent   string // 报告文字内容
	UploadedAt      string // 上传时间
}

// medicalReportColumns holds the columns for the table medical_report.
var medicalReportColumns = MedicalReportColumns{
	Id:              "id",
	MedicalRecordId: "medical_record_id",
	DoctorId:        "doctor_id",
	ReportTitle:     "report_title",
	ReportType:      "report_type",
	FileUrl:         "file_url",
	ReportContent:   "report_content",
	UploadedAt:      "uploaded_at",
}

// NewMedicalReportDao creates and returns a new DAO object for table data access.
func NewMedicalReportDao(handlers ...gdb.ModelHandler) *MedicalReportDao {
	return &MedicalReportDao{
		group:    "default",
		table:    "medical_report",
		columns:  medicalReportColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MedicalReportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MedicalReportDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MedicalReportDao) Columns() MedicalReportColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MedicalReportDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MedicalReportDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MedicalReportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
