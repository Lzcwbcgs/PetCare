package appointment

import (
	"context"

	v1 "PetCare/api/appointment/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Create(ctx context.Context, req *v1.AppointmentCreateReq) (res *v1.AppointmentCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "预约成功")

	claims, err := requireAppointmentRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	output, err := service.Appointment.Create(ctx, service.AppointmentCreateInput{
		UserID:             claims.UserID,
		PetID:              req.PetID,
		HospitalID:         req.HospitalID,
		DoctorID:           req.DoctorID,
		AppointmentType:    req.AppointmentType,
		SymptomDescription: req.SymptomDescription,
		AppointmentTime:    req.AppointmentTime,
	})
	if err != nil {
		return nil, err
	}

	return &v1.AppointmentCreateRes{
		AppointmentID: output.ID,
		AppointmentNo: output.AppointmentNo,
		Status:        output.Status,
	}, nil
}

func (c *Controller) List(ctx context.Context, req *v1.AppointmentListReq) (res *v1.AppointmentListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireAppointmentRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	var (
		page int
		size int
	)
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		size = *req.PageSize
	}

	output, err := service.Appointment.List(ctx, service.AppointmentListInput{
		UserID:          claims.UserID,
		Page:            page,
		Size:            size,
		Status:          req.Status,
		AppointmentType: req.AppointmentType,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.AppointmentListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.AppointmentListItem{
			ID:              item.ID,
			AppointmentNo:   item.AppointmentNo,
			PetID:           item.PetID,
			PetName:         item.PetName,
			HospitalID:      item.HospitalID,
			HospitalName:    item.HospitalName,
			DoctorID:        item.DoctorID,
			DoctorName:      item.DoctorName,
			AppointmentType: item.AppointmentType,
			AppointmentTime: item.AppointmentTime,
			Status:          item.Status,
		})
	}

	return &v1.AppointmentListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *Controller) Detail(ctx context.Context, req *v1.AppointmentDetailReq) (res *v1.AppointmentDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireAppointmentRoles(ctx, consts.RoleUser, consts.RoleDoctor, consts.RoleAdmin)
	if err != nil {
		return nil, err
	}

	item, err := service.Appointment.Detail(ctx, service.AppointmentDetailInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		AppointmentID:   req.AppointmentID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.AppointmentDetailRes{
		ID:                 item.ID,
		AppointmentNo:      item.AppointmentNo,
		UserID:             item.UserID,
		UserNickname:       item.UserNickname,
		PetID:              item.PetID,
		PetName:            item.PetName,
		HospitalID:         item.HospitalID,
		HospitalName:       item.HospitalName,
		DoctorID:           item.DoctorID,
		DoctorName:         item.DoctorName,
		AppointmentType:    item.AppointmentType,
		SymptomDescription: item.SymptomDescription,
		AppointmentTime:    item.AppointmentTime,
		ReminderTime:       item.ReminderTime,
		Status:             item.Status,
		Source:             item.Source,
		CreatedAt:          item.CreatedAt,
	}, nil
}

func (c *Controller) Cancel(ctx context.Context, req *v1.AppointmentCancelReq) (res *v1.AppointmentCancelRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "取消成功")

	claims, err := requireAppointmentRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	err = service.Appointment.Cancel(ctx, service.AppointmentCancelInput{
		UserID:        claims.UserID,
		AppointmentID: req.AppointmentID,
		CancelReason:  req.CancelReason,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AppointmentCancelRes{}, nil
}
