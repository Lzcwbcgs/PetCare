// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Pet is the golang structure of table pet for DAO operations like Where/Data.
type Pet struct {
	g.Meta     `orm:"table:pet, do:true"`
	Id         any // 宠物主键ID
	UserId     any // 所属用户ID
	PetName    any // 宠物名字
	PetType    any // 宠物类型，当前默认猫
	AvatarUrl  any // 宠物头像URL
	Gender     any // 性别：1公，2母，0未知
	Age        any // 年龄
	AgeUnit    any // 年龄单位：month/月，year/岁
	Breed      any // 品种
	Weight     any // 体重（kg）
	Sterilized any // 是否绝育：1是，0否
	Remark     any // 备注
	Status     any // 状态：1正常，0停用/删除
	CreatedAt  any // 创建时间
	UpdatedAt  any // 更新时间
}
