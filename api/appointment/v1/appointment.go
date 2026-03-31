package v1

import "github.com/gogf/gf/v2/frame/g"

type Pagination struct {
	Page     int `json:"page" dc:"当前页码"`
	PageSize int `json:"page_size" dc:"每页数量"`
	Total    int `json:"total" dc:"总数"`
}

type AppointmentCreateReq struct {
	g.Meta             `path:"/" method:"post" tags:"预约管理" summary:"创建预约"`
	PetID              int64  `json:"pet_id" v:"required|min:1#请填写宠物ID|宠物ID不合法" dc:"宠物ID"`
	HospitalID         int64  `json:"hospital_id" v:"required|min:1#请填写医院ID|医院ID不合法" dc:"医院ID"`
	DoctorID           int64  `json:"doctor_id" v:"required|min:1#请填写医生ID|医生ID不合法" dc:"医生ID"`
	AppointmentType    int    `json:"appointment_type" v:"required|in:1,2#请填写预约类型|预约类型不合法" dc:"预约类型：1体检预约，2看病预约"`
	SymptomDescription string `json:"symptom_description" v:"max-length:65535#症状描述长度不能超过65535" dc:"症状描述"`
	AppointmentTime    string `json:"appointment_time" v:"required|datetime#请填写预约时间|预约时间格式不正确" dc:"预约时间 (yyyy-MM-dd HH:mm:ss)"`
}

type AppointmentCreateRes struct {
	AppointmentID int64  `json:"appointment_id" dc:"预约ID"`
	AppointmentNo string `json:"appointment_no" dc:"预约单号"`
	Status        int    `json:"status" dc:"状态"`
}

type AppointmentListReq struct {
	g.Meta          `path:"/" method:"get" tags:"预约管理" summary:"获取预约列表"`
	Page            *int `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize        *int `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	Status          *int `json:"status" p:"status" v:"in:1,2,3,4#预约状态不合法" dc:"状态：1待就诊，2已完成，3已取消，4已过期"`
	AppointmentType *int `json:"appointment_type" p:"appointment_type" v:"in:1,2#预约类型不合法" dc:"预约类型：1体检预约，2看病预约"`
}

type AppointmentListItem struct {
	ID              int64  `json:"id" dc:"预约ID"`
	AppointmentNo   string `json:"appointment_no" dc:"预约单号"`
	PetID           int64  `json:"pet_id" dc:"宠物ID"`
	PetName         string `json:"pet_name" dc:"宠物名称"`
	HospitalID      int64  `json:"hospital_id" dc:"医院ID"`
	HospitalName    string `json:"hospital_name" dc:"医院名称"`
	DoctorID        int64  `json:"doctor_id" dc:"医生ID"`
	DoctorName      string `json:"doctor_name" dc:"医生姓名"`
	AppointmentType int    `json:"appointment_type" dc:"预约类型"`
	AppointmentTime string `json:"appointment_time" dc:"预约时间"`
	Status          int    `json:"status" dc:"状态"`
}

type AppointmentListRes struct {
	List       []AppointmentListItem `json:"list" dc:"预约列表"`
	Pagination Pagination            `json:"pagination" dc:"分页信息"`
}

type AppointmentDetailReq struct {
	g.Meta        `path:"/{appointment_id}" method:"get" tags:"预约管理" summary:"获取预约详情"`
	AppointmentID int64 `json:"appointment_id" p:"appointment_id" v:"required|min:1#预约ID不能为空|预约ID不合法" dc:"预约ID"`
}

type AppointmentDetailRes struct {
	ID                 int64  `json:"id" dc:"预约ID"`
	AppointmentNo      string `json:"appointment_no" dc:"预约单号"`
	UserID             int64  `json:"user_id" dc:"用户ID"`
	UserNickname       string `json:"user_nickname" dc:"用户昵称"`
	PetID              int64  `json:"pet_id" dc:"宠物ID"`
	PetName            string `json:"pet_name" dc:"宠物名称"`
	HospitalID         int64  `json:"hospital_id" dc:"医院ID"`
	HospitalName       string `json:"hospital_name" dc:"医院名称"`
	DoctorID           int64  `json:"doctor_id" dc:"医生ID"`
	DoctorName         string `json:"doctor_name" dc:"医生姓名"`
	AppointmentType    int    `json:"appointment_type" dc:"预约类型"`
	SymptomDescription string `json:"symptom_description" dc:"症状描述"`
	AppointmentTime    string `json:"appointment_time" dc:"预约时间"`
	ReminderTime       string `json:"reminder_time" dc:"提醒时间"`
	Status             int    `json:"status" dc:"状态"`
	Source             int    `json:"source" dc:"来源"`
	CreatedAt          string `json:"created_at" dc:"创建时间"`
}

type AppointmentCancelReq struct {
	g.Meta        `path:"/{appointment_id}/cancel" method:"put" tags:"预约管理" summary:"取消预约"`
	AppointmentID int64  `json:"appointment_id" p:"appointment_id" v:"required|min:1#预约ID不能为空|预约ID不合法" dc:"预约ID"`
	CancelReason  string `json:"cancel_reason" v:"required|max-length:255#请填写取消原因|取消原因长度不能超过255" dc:"取消原因"`
}

type AppointmentCancelRes struct{}
