package v1

import "github.com/gogf/gf/v2/frame/g"

type DoctorCreateReq struct {
	g.Meta     `path:"/doctors" method:"post" tags:"医生管理" summary:"新增医生"`
	HospitalID int64   `json:"hospital_id" v:"required|min:1#请填写所属医院|所属医院不合法" dc:"所属医院ID"`
	Username   string  `json:"username" v:"required|length:3,32#请填写医生登录账号|医生登录账号长度需在3到32位之间" dc:"医生登录账号"`
	Password   string  `json:"password" v:"required|length:6,64#请填写登录密码|登录密码长度需在6到64位之间" dc:"登录密码"`
	DoctorName string  `json:"doctor_name" v:"required|length:1,50#请填写医生姓名|医生姓名长度需在1到50位之间" dc:"医生姓名"`
	Gender     int     `json:"gender" v:"required|in:0,1,2#请填写医生性别|医生性别不合法" dc:"性别：1男，2女，0未知"`
	Phone      string  `json:"phone" v:"required|phone#请填写手机号|手机号格式不正确" dc:"手机号"`
	Email      *string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Title      string  `json:"title" v:"required|length:1,50#请填写职称|职称长度需在1到50位之间" dc:"职称"`
	Specialty  string  `json:"specialty" v:"required|length:1,255#请填写擅长领域|擅长领域长度需在1到255位之间" dc:"擅长领域"`
	AvatarURL  *string `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255" dc:"头像URL"`
	Intro      *string `json:"intro" v:"max-length:65535#医生简介长度不能超过65535" dc:"医生简介"`
	Status     int     `json:"status" v:"required|in:0,1#请填写医生状态|医生状态不合法" dc:"状态：1在职可接诊，0停用"`
}

type DoctorCreateRes struct {
	DoctorID int64 `json:"doctor_id" dc:"新建医生ID"`
}

type DoctorListReq struct {
	g.Meta     `path:"/doctors" method:"get" tags:"医生管理" summary:"获取医生列表"`
	Page       *int    `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize   *int    `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	HospitalID *int64  `json:"hospital_id" p:"hospital_id" v:"min:1#所属医院不合法" dc:"所属医院ID"`
	Status     *int    `json:"status" p:"status" v:"in:0,1#医生状态不合法" dc:"状态：1在职可接诊，0停用"`
	Keyword    *string `json:"keyword" p:"keyword" v:"max-length:50#关键字长度不能超过50" dc:"医生姓名关键字"`
}

type DoctorListItem struct {
	ID           int64  `json:"id" dc:"医生ID"`
	HospitalID   int64  `json:"hospital_id" dc:"所属医院ID"`
	HospitalName string `json:"hospital_name" dc:"所属医院名称"`
	Username     string `json:"username" dc:"医生登录账号"`
	DoctorName   string `json:"doctor_name" dc:"医生姓名"`
	Title        string `json:"title" dc:"职称"`
	Specialty    string `json:"specialty" dc:"擅长领域"`
	Phone        string `json:"phone" dc:"手机号"`
	Status       int    `json:"status" dc:"状态"`
	CreatedAt    string `json:"created_at" dc:"创建时间"`
}

type DoctorListRes struct {
	List       []DoctorListItem `json:"list" dc:"医生列表"`
	Pagination Pagination       `json:"pagination" dc:"分页信息"`
}

type DoctorDetailReq struct {
	g.Meta   `path:"/doctors/{doctor_id}" method:"get" tags:"医生管理" summary:"获取医生详情"`
	DoctorID int64 `json:"doctor_id" p:"doctor_id" v:"required|min:1#医生ID不能为空|医生ID不合法" dc:"医生ID"`
}

type DoctorDetailRes struct {
	ID         int64  `json:"id" dc:"医生ID"`
	HospitalID int64  `json:"hospital_id" dc:"所属医院ID"`
	Username   string `json:"username" dc:"医生登录账号"`
	DoctorName string `json:"doctor_name" dc:"医生姓名"`
	Gender     int    `json:"gender" dc:"性别"`
	Phone      string `json:"phone" dc:"手机号"`
	Email      string `json:"email" dc:"邮箱"`
	Title      string `json:"title" dc:"职称"`
	Specialty  string `json:"specialty" dc:"擅长领域"`
	AvatarURL  string `json:"avatar_url" dc:"头像URL"`
	Intro      string `json:"intro" dc:"医生简介"`
	Status     int    `json:"status" dc:"状态"`
	CreatedAt  string `json:"created_at" dc:"创建时间"`
	UpdatedAt  string `json:"updated_at" dc:"更新时间"`
}

type DoctorUpdateReq struct {
	g.Meta     `path:"/doctors/{doctor_id}" method:"put" tags:"医生管理" summary:"修改医生信息"`
	DoctorID   int64   `json:"doctor_id" p:"doctor_id" v:"required|min:1#医生ID不能为空|医生ID不合法" dc:"医生ID"`
	HospitalID *int64  `json:"hospital_id" v:"min:1#所属医院不合法" dc:"所属医院ID"`
	DoctorName *string `json:"doctor_name" v:"length:1,50#医生姓名长度需在1到50位之间" dc:"医生姓名"`
	Gender     *int    `json:"gender" v:"in:0,1,2#医生性别不合法" dc:"性别：1男，2女，0未知"`
	Phone      *string `json:"phone" v:"phone#手机号格式不正确" dc:"手机号"`
	Email      *string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Title      *string `json:"title" v:"length:1,50#职称长度需在1到50位之间" dc:"职称"`
	Specialty  *string `json:"specialty" v:"length:1,255#擅长领域长度需在1到255位之间" dc:"擅长领域"`
	AvatarURL  *string `json:"avatar_url" v:"max-length:255#头像地址长度不能超过255" dc:"头像URL"`
	Intro      *string `json:"intro" v:"max-length:65535#医生简介长度不能超过65535" dc:"医生简介"`
	Status     *int    `json:"status" v:"in:0,1#医生状态不合法" dc:"状态：1在职可接诊，0停用"`
}

type DoctorUpdateRes struct{}

type DoctorDeleteReq struct {
	g.Meta   `path:"/doctors/{doctor_id}" method:"delete" tags:"医生管理" summary:"删除医生"`
	DoctorID int64 `json:"doctor_id" p:"doctor_id" v:"required|min:1#医生ID不能为空|医生ID不合法" dc:"医生ID"`
}

type DoctorDeleteRes struct{}
