// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// PetVaccinationRecord is the golang structure for table pet_vaccination_record.
type PetVaccinationRecord struct {
	Id              int64     `json:"id"              orm:"id"               description:"疫苗记录ID"`  // 疫苗记录ID
	PetId           int64     `json:"petId"           orm:"pet_id"           description:"宠物ID"`    // 宠物ID
	VaccineName     string    `json:"vaccineName"     orm:"vaccine_name"     description:"疫苗名称"`    // 疫苗名称
	VaccinationDate time.Time `json:"vaccinationDate" orm:"vaccination_date" description:"接种日期"`    // 接种日期
	NextDueDate     time.Time `json:"nextDueDate"     orm:"next_due_date"    description:"下次应接种日期"` // 下次应接种日期
	HospitalName    string    `json:"hospitalName"    orm:"hospital_name"    description:"接种机构"`    // 接种机构
	Remark          string    `json:"remark"          orm:"remark"           description:"备注"`      // 备注
	CreatedAt       time.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`    // 创建时间
	UpdatedAt       time.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`    // 更新时间
}
