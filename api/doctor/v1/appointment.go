package v1

import "github.com/gogf/gf/v2/frame/g"

type Pagination struct {
	Page     int `json:"page" dc:"当前页码"`
	PageSize int `json:"page_size" dc:"每页数量"`
	Total    int `json:"total" dc:"总数"`
}

type AppointmentListReq struct {
	g.Meta   `path:"/" method:"get" tags:"医生预约管理" summary:"获取医生预约列表"`
	Page     *int `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"页码"`
	PageSize *int `json:"page_size" p:"page_size" v:"min:1|max:100#每页数量必须在1到100之间" dc:"每页数量"`
	Status   *int `json:"status" p:"status" v:"in:1,2,3,4#预约状态不合法" dc:"状态：1待就诊，2已完成，3已取消，4已过期"`
}

type AppointmentListItem struct {
	ID              int64  `json:"id" dc:"预约ID"`
	AppointmentNo   string `json:"appointment_no" dc:"预约单号"`
	PetID           int64  `json:"pet_id" dc:"宠物ID"`
	PetName         string `json:"pet_name" dc:"宠物名称"`
	UserID          int64  `json:"user_id" dc:"用户ID"`
	UserNickname    string `json:"user_nickname" dc:"用户昵称"`
	AppointmentType int    `json:"appointment_type" dc:"预约类型"`
	AppointmentTime string `json:"appointment_time" dc:"预约时间"`
	Status          int    `json:"status" dc:"状态"`
}

type AppointmentListRes struct {
	List       []AppointmentListItem `json:"list" dc:"预约列表"`
	Pagination Pagination            `json:"pagination" dc:"分页信息"`
}

type AppointmentDetailReq struct {
	g.Meta        `path:"/{appointment_id}/detail" method:"get" tags:"医生预约管理" summary:"获取预约接诊详情"`
	AppointmentID int64 `json:"appointment_id" p:"appointment_id" v:"required|min:1#预约ID不能为空|预约ID不合法" dc:"预约ID"`
}

type AppointmentDetailAppointment struct {
	ID                 int64  `json:"id" dc:"预约ID"`
	AppointmentNo      string `json:"appointment_no" dc:"预约单号"`
	AppointmentType    int    `json:"appointment_type" dc:"预约类型"`
	SymptomDescription string `json:"symptom_description" dc:"症状描述"`
	AppointmentTime    string `json:"appointment_time" dc:"预约时间"`
	Status             int    `json:"status" dc:"状态"`
}

type AppointmentDetailPet struct {
	ID         int64  `json:"id" dc:"宠物ID"`
	PetName    string `json:"pet_name" dc:"宠物名称"`
	PetType    string `json:"pet_type" dc:"宠物类型"`
	Gender     int    `json:"gender" dc:"性别"`
	Age        int    `json:"age" dc:"年龄"`
	AgeUnit    string `json:"age_unit" dc:"年龄单位"`
	Breed      string `json:"breed" dc:"品种"`
	Weight     string `json:"weight" dc:"体重"`
	Sterilized int    `json:"sterilized" dc:"是否绝育"`
	Remark     string `json:"remark" dc:"备注"`
}

type AppointmentDetailMedicalHistory struct {
	ID          int64  `json:"id" dc:"病史记录ID"`
	HistoryType string `json:"history_type" dc:"病史类型"`
	Description string `json:"description" dc:"病史描述"`
	DiagnosedAt string `json:"diagnosed_at" dc:"确诊时间"`
	IsCurrent   int    `json:"is_current" dc:"是否当前存在"`
}

type AppointmentDetailVaccination struct {
	ID              int64  `json:"id" dc:"疫苗记录ID"`
	VaccineName     string `json:"vaccine_name" dc:"疫苗名称"`
	VaccinationDate string `json:"vaccination_date" dc:"接种日期"`
	NextDueDate     string `json:"next_due_date" dc:"下次应接种日期"`
}

type AppointmentDetailAllergy struct {
	ID                 int64  `json:"id" dc:"过敏记录ID"`
	Allergen           string `json:"allergen" dc:"过敏源"`
	SymptomDescription string `json:"symptom_description" dc:"症状描述"`
	SeverityLevel      int    `json:"severity_level" dc:"严重程度"`
}

type AppointmentDetailRes struct {
	Appointment      AppointmentDetailAppointment      `json:"appointment" dc:"预约信息"`
	Pet              AppointmentDetailPet              `json:"pet" dc:"宠物信息"`
	MedicalHistories []AppointmentDetailMedicalHistory `json:"medical_histories" dc:"病史记录"`
	Vaccinations     []AppointmentDetailVaccination    `json:"vaccinations" dc:"疫苗记录"`
	Allergies        []AppointmentDetailAllergy        `json:"allergies" dc:"过敏记录"`
}

type AppointmentUpdateStatusReq struct {
	g.Meta        `path:"/{appointment_id}/status" method:"put" tags:"医生预约管理" summary:"更新预约状态"`
	AppointmentID int64 `json:"appointment_id" p:"appointment_id" v:"required|min:1#预约ID不能为空|预约ID不合法" dc:"预约ID"`
	Status        int   `json:"status" v:"required|in:2,4#请填写预约状态|预约状态不合法" dc:"预约状态：2已完成，4已过期"`
}

type AppointmentUpdateStatusRes struct{}
