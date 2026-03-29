package v1

import "github.com/gogf/gf/v2/frame/g"

type PetListReq struct {
	g.Meta  `path:"/" method:"get" tags:"宠物档案" summary:"获取宠物列表"`
	Page    *int    `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	Size    *int    `json:"size" v:"min:1|max:100#每页数量必须在1-100之间" dc:"每页数量"`
	PetName *string `json:"pet_name" v:"max-length:50#宠物名字长度不能超过50位" dc:"宠物名字(模糊匹配)"`
	PetType *string `json:"pet_type" v:"max-length:20#宠物类型长度不能超过20位" dc:"宠物类型"`
	Status  *int    `json:"status" v:"in:0,1#状态不合法" dc:"状态：1正常，0停用/删除"`
}

type PetListItem struct {
	ID         int64   `json:"id" dc:"宠物主键ID"`
	UserID     int64   `json:"user_id" dc:"所属用户ID"`
	PetName    string  `json:"pet_name" dc:"宠物名字"`
	PetType    string  `json:"pet_type" dc:"宠物类型"`
	AvatarURL  string  `json:"avatar_url" dc:"宠物头像URL"`
	Gender     int     `json:"gender" dc:"性别：1公，2母，0未知"`
	Age        int     `json:"age" dc:"年龄"`
	AgeUnit    string  `json:"age_unit" dc:"年龄单位：month/月，year/岁"`
	Breed      string  `json:"breed" dc:"品种"`
	Weight     float64 `json:"weight" dc:"体重（kg）"`
	Sterilized int     `json:"sterilized" dc:"是否绝育：1是，0否"`
	Remark     string  `json:"remark" dc:"备注"`
	Status     int     `json:"status" dc:"状态：1正常，0停用/删除"`
	CreatedAt  string  `json:"created_at" dc:"创建时间"`
	UpdatedAt  string  `json:"updated_at" dc:"更新时间"`
}

type PetListRes struct {
	List  []PetListItem `json:"list" dc:"宠物列表"`
	Total int           `json:"total" dc:"总数"`
	Page  int           `json:"page" dc:"当前页"`
	Size  int           `json:"size" dc:"每页数量"`
}

type PetDetailReq struct {
	g.Meta `path:"/{id}" method:"get" tags:"宠物档案" summary:"获取宠物详情"`
	ID     int64 `json:"id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物主键ID"`
}

type PetDetailRes struct {
	ID         int64   `json:"id" dc:"宠物主键ID"`
	UserID     int64   `json:"user_id" dc:"所属用户ID"`
	PetName    string  `json:"pet_name" dc:"宠物名字"`
	PetType    string  `json:"pet_type" dc:"宠物类型"`
	AvatarURL  string  `json:"avatar_url" dc:"宠物头像URL"`
	Gender     int     `json:"gender" dc:"性别：1公，2母，0未知"`
	Age        int     `json:"age" dc:"年龄"`
	AgeUnit    string  `json:"age_unit" dc:"年龄单位：month/月，year/岁"`
	Breed      string  `json:"breed" dc:"品种"`
	Weight     float64 `json:"weight" dc:"体重（kg）"`
	Sterilized int     `json:"sterilized" dc:"是否绝育：1是，0否"`
	Remark     string  `json:"remark" dc:"备注"`
	Status     int     `json:"status" dc:"状态：1正常，0停用/删除"`
	CreatedAt  string  `json:"created_at" dc:"创建时间"`
	UpdatedAt  string  `json:"updated_at" dc:"更新时间"`
}

type PetCreateReq struct {
	g.Meta     `path:"/" method:"post" tags:"宠物档案" summary:"新增宠物档案"`
	PetName    string   `json:"pet_name" v:"required|length:1,50#请输入宠物名字|宠物名字长度需在1到50位之间" dc:"宠物名字"`
	PetType    *string  `json:"pet_type" v:"max-length:20#宠物类型长度不能超过20位" dc:"宠物类型"`
	AvatarURL  *string  `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"宠物头像URL"`
	Gender     *int     `json:"gender" v:"in:0,1,2#性别不合法" dc:"性别：1公，2母，0未知"`
	Age        *int     `json:"age" v:"min:0#年龄不能小于0" dc:"年龄"`
	AgeUnit    *string  `json:"age_unit" v:"in:month,year#年龄单位不合法" dc:"年龄单位：month/月，year/岁"`
	Breed      *string  `json:"breed" v:"max-length:50#品种长度不能超过50位" dc:"品种"`
	Weight     *float64 `json:"weight" v:"min:0|max:999.99#体重不能小于0|体重不能超过999.99kg" dc:"体重（kg）"`
	Sterilized *int     `json:"sterilized" v:"in:0,1#是否绝育不合法" dc:"是否绝育：1是，0否"`
	Remark     *string  `json:"remark" v:"max-length:65535#备注长度不能超过65535位" dc:"备注"`
}

type PetCreateRes struct {
	ID int64 `json:"id" dc:"新建宠物ID"`
}

type PetUpdateReq struct {
	g.Meta     `path:"/{id}" method:"put" tags:"宠物档案" summary:"更新宠物档案"`
	ID         int64    `json:"id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物主键ID"`
	PetName    *string  `json:"pet_name" v:"length:1,50#宠物名字长度需在1到50位之间" dc:"宠物名字"`
	PetType    *string  `json:"pet_type" v:"max-length:20#宠物类型长度不能超过20位" dc:"宠物类型"`
	AvatarURL  *string  `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"宠物头像URL"`
	Gender     *int     `json:"gender" v:"in:0,1,2#性别不合法" dc:"性别：1公，2母，0未知"`
	Age        *int     `json:"age" v:"min:0#年龄不能小于0" dc:"年龄"`
	AgeUnit    *string  `json:"age_unit" v:"in:month,year#年龄单位不合法" dc:"年龄单位：month/月，year/岁"`
	Breed      *string  `json:"breed" v:"max-length:50#品种长度不能超过50位" dc:"品种"`
	Weight     *float64 `json:"weight" v:"min:0|max:999.99#体重不能小于0|体重不能超过999.99kg" dc:"体重（kg）"`
	Sterilized *int     `json:"sterilized" v:"in:0,1#是否绝育不合法" dc:"是否绝育：1是，0否"`
	Remark     *string  `json:"remark" v:"max-length:65535#备注长度不能超过65535位" dc:"备注"`
}

type PetUpdateRes struct{}

type PetDeleteReq struct {
	g.Meta `path:"/{id}" method:"delete" tags:"宠物档案" summary:"删除宠物档案"`
	ID     int64 `json:"id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物主键ID"`
}

type PetDeleteRes struct{}
