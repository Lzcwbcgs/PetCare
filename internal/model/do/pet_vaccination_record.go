// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PetVaccinationRecord is the golang structure of table pet_vaccination_record for DAO operations like Where/Data.
type PetVaccinationRecord struct {
	g.Meta          `orm:"table:pet_vaccination_record, do:true"`
	Id              any // 疫苗记录ID
	PetId           any // 宠物ID
	VaccineName     any // 疫苗名称
	VaccinationDate any // 接种日期
	NextDueDate     any // 下次应接种日期
	HospitalName    any // 接种机构
	Remark          any // 备注
	CreatedAt       any // 创建时间
	UpdatedAt       any // 更新时间
}
