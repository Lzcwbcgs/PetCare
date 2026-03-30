package v1

import "github.com/gogf/gf/v2/frame/g"

type Pagination struct {
	Page     int `json:"page" dc:"当前页码"`
	PageSize int `json:"page_size" dc:"每页数量"`
	Total    int `json:"total" dc:"总数"`
}

type PetListReq struct {
	g.Meta   `path:"/" method:"get" tags:"宠物档案" summary:"获取宠物列表"`
	Page     *int    `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int    `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1-100之间" dc:"每页数量"`
	PetName  *string `json:"pet_name" p:"pet_name" v:"max-length:50#宠物名字长度不能超过50位" dc:"宠物名字(模糊匹配)"`
	PetType  *string `json:"pet_type" p:"pet_type" v:"max-length:20#宠物类型长度不能超过20位" dc:"宠物类型"`
	Status   *int    `json:"status" p:"status" v:"in:0,1#状态不合法" dc:"状态：1正常，0停用/删除"`
}

type PetListItem struct {
	ID      int64  `json:"id" dc:"宠物主键ID"`
	PetName string `json:"pet_name" dc:"宠物名字"`
	PetType string `json:"pet_type" dc:"宠物类型"`
	Gender  int    `json:"gender" dc:"性别：1公，2母，0未知"`
	Age     int    `json:"age" dc:"年龄"`
	AgeUnit string `json:"age_unit" dc:"年龄单位：month/月，year/岁"`
	Breed   string `json:"breed" dc:"品种"`
	Weight  string `json:"weight" dc:"体重（kg）"`
	Status  int    `json:"status" dc:"状态：1正常，0停用/删除"`
}

type PetListRes struct {
	List       []PetListItem `json:"list" dc:"宠物列表"`
	Pagination Pagination    `json:"pagination" dc:"分页信息"`
}

type PetDetailReq struct {
	g.Meta `path:"/{pet_id}" method:"get" tags:"宠物档案" summary:"获取宠物详情"`
	PetID  int64 `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物主键ID"`
}

type PetDetailRes struct {
	ID         int64  `json:"id" dc:"宠物主键ID"`
	UserID     int64  `json:"user_id" dc:"所属用户ID"`
	PetName    string `json:"pet_name" dc:"宠物名字"`
	PetType    string `json:"pet_type" dc:"宠物类型"`
	AvatarURL  string `json:"avatar_url" dc:"宠物头像URL"`
	Gender     int    `json:"gender" dc:"性别：1公，2母，0未知"`
	Age        int    `json:"age" dc:"年龄"`
	AgeUnit    string `json:"age_unit" dc:"年龄单位：month/月，year/岁"`
	Breed      string `json:"breed" dc:"品种"`
	Weight     string `json:"weight" dc:"体重（kg）"`
	Sterilized int    `json:"sterilized" dc:"是否绝育：1是，0否"`
	Remark     string `json:"remark" dc:"备注"`
	Status     int    `json:"status" dc:"状态：1正常，0停用/删除"`
	CreatedAt  string `json:"created_at" dc:"创建时间"`
	UpdatedAt  string `json:"updated_at" dc:"更新时间"`
}

type PetCreateReq struct {
	g.Meta     `path:"/" method:"post" tags:"宠物档案" summary:"新增宠物档案"`
	PetName    string  `json:"pet_name" v:"required|length:1,50#请输入宠物名字|宠物名字长度需在1到50位之间" dc:"宠物名字"`
	PetType    *string `json:"pet_type" v:"max-length:20#宠物类型长度不能超过20位" dc:"宠物类型"`
	AvatarURL  *string `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"宠物头像URL"`
	Gender     *int    `json:"gender" v:"in:0,1,2#性别不合法" dc:"性别：1公，2母，0未知"`
	Age        *int    `json:"age" v:"min:0#年龄不能小于0" dc:"年龄"`
	AgeUnit    *string `json:"age_unit" v:"in:month,year#年龄单位不合法" dc:"年龄单位：month/月，year/岁"`
	Breed      *string `json:"breed" v:"max-length:50#品种长度不能超过50位" dc:"品种"`
	Weight     *string `json:"weight" v:"max-length:16#体重格式不合法" dc:"体重（kg，字符串格式，最多2位小数）"`
	Sterilized *int    `json:"sterilized" v:"in:0,1#是否绝育不合法" dc:"是否绝育：1是，0否"`
	Remark     *string `json:"remark" v:"max-length:65535#备注长度不能超过65535位" dc:"备注"`
}

type PetCreateRes struct {
	PetID int64 `json:"pet_id" dc:"新建宠物ID"`
}

type PetUpdateReq struct {
	g.Meta     `path:"/{pet_id}" method:"put" tags:"宠物档案" summary:"修改宠物档案"`
	PetID      int64   `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物主键ID"`
	PetName    *string `json:"pet_name" v:"length:1,50#宠物名字长度需在1到50位之间" dc:"宠物名字"`
	PetType    *string `json:"pet_type" v:"max-length:20#宠物类型长度不能超过20位" dc:"宠物类型"`
	AvatarURL  *string `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255位" dc:"宠物头像URL"`
	Gender     *int    `json:"gender" v:"in:0,1,2#性别不合法" dc:"性别：1公，2母，0未知"`
	Age        *int    `json:"age" v:"min:0#年龄不能小于0" dc:"年龄"`
	AgeUnit    *string `json:"age_unit" v:"in:month,year#年龄单位不合法" dc:"年龄单位：month/月，year/岁"`
	Breed      *string `json:"breed" v:"max-length:50#品种长度不能超过50位" dc:"品种"`
	Weight     *string `json:"weight" v:"max-length:16#体重格式不合法" dc:"体重（kg，字符串格式，最多2位小数）"`
	Sterilized *int    `json:"sterilized" v:"in:0,1#是否绝育不合法" dc:"是否绝育：1是，0否"`
	Remark     *string `json:"remark" v:"max-length:65535#备注长度不能超过65535位" dc:"备注"`
}

type PetUpdateRes struct{}

type PetDeleteReq struct {
	g.Meta `path:"/{pet_id}" method:"delete" tags:"宠物档案" summary:"删除宠物档案"`
	PetID  int64 `json:"pet_id" p:"pet_id" v:"required|min:1#宠物ID不能为空|宠物ID不合法" dc:"宠物主键ID"`
}

type PetDeleteRes struct{}
