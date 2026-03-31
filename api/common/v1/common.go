package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HospitalOptionReq struct {
	g.Meta `path:"/hospitals/options" method:"get" tags:"公共辅助接口" summary:"获取医院下拉数据"`
}

type HospitalOptionItem struct {
	ID           int64  `json:"id" dc:"医院ID"`
	HospitalName string `json:"hospital_name" dc:"医院名称"`
}

type HospitalOptionRes []HospitalOptionItem

type HospitalDoctorOptionReq struct {
	g.Meta     `path:"/hospitals/{hospital_id}/doctors/options" method:"get" tags:"公共辅助接口" summary:"获取医院医生下拉数据"`
	HospitalID int64 `json:"hospital_id" p:"hospital_id" v:"required|min:1#医院ID不能为空|医院ID不合法" dc:"医院ID"`
}

type HospitalDoctorOptionItem struct {
	ID         int64  `json:"id" dc:"医生ID"`
	DoctorName string `json:"doctor_name" dc:"医生姓名"`
	Title      string `json:"title" dc:"职称"`
}

type HospitalDoctorOptionRes []HospitalDoctorOptionItem

type UploadReq struct {
	g.Meta  `path:"/upload" method:"post" mime:"multipart/form-data" tags:"公共辅助接口" summary:"通用文件上传"`
	File    *ghttp.UploadFile `json:"file" type:"file" v:"required#请上传文件" dc:"上传文件"`
	BizType *string           `json:"biz_type" v:"max-length:50#业务类型长度不能超过50" dc:"业务类型，如 avatar、report"`
}

type UploadRes struct {
	FileURL  string `json:"file_url" dc:"文件地址"`
	FileName string `json:"file_name" dc:"文件名"`
}
