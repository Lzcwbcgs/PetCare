// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Pet is the golang structure for table pet.
type Pet struct {
	Id         int64     `json:"id"         orm:"id"         description:"宠物主键ID"`              // 宠物主键ID
	UserId     int64     `json:"userId"     orm:"user_id"    description:"所属用户ID"`              // 所属用户ID
	PetName    string    `json:"petName"    orm:"pet_name"   description:"宠物名字"`                // 宠物名字
	PetType    string    `json:"petType"    orm:"pet_type"   description:"宠物类型，当前默认猫"`          // 宠物类型，当前默认猫
	AvatarUrl  string    `json:"avatarUrl"  orm:"avatar_url" description:"宠物头像URL"`             // 宠物头像URL
	Gender     int       `json:"gender"     orm:"gender"     description:"性别：1公，2母，0未知"`        // 性别：1公，2母，0未知
	Age        int       `json:"age"        orm:"age"        description:"年龄"`                  // 年龄
	AgeUnit    string    `json:"ageUnit"    orm:"age_unit"   description:"年龄单位：month/月，year/岁"` // 年龄单位：month/月，year/岁
	Breed      string    `json:"breed"      orm:"breed"      description:"品种"`                  // 品种
	Weight     float64   `json:"weight"     orm:"weight"     description:"体重（kg）"`              // 体重（kg）
	Sterilized int       `json:"sterilized" orm:"sterilized" description:"是否绝育：1是，0否"`          // 是否绝育：1是，0否
	Remark     string    `json:"remark"     orm:"remark"     description:"备注"`                  // 备注
	Status     int       `json:"status"     orm:"status"     description:"状态：1正常，0停用/删除"`       // 状态：1正常，0停用/删除
	CreatedAt  time.Time `json:"createdAt"  orm:"created_at" description:"创建时间"`                // 创建时间
	UpdatedAt  time.Time `json:"updatedAt"  orm:"updated_at" description:"更新时间"`                // 更新时间
}
