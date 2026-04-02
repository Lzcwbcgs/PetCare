package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

type (
	DoctorNotificationListInput struct {
		DoctorID int64
		Page     int
		Size     int
	}

	DoctorNotificationItem struct {
		ID               int64
		NotificationType int
		Title            string
		Content          string
		AppointmentID    int64
		SendTime         string
		Status           int
	}

	DoctorNotificationListOutput struct {
		Items []DoctorNotificationItem
		Total int
		Page  int
		Size  int
	}

	DoctorNotificationReadInput struct {
		DoctorID       int64
		NotificationID int64
	}
)

type IDoctorNotification interface {
	List(ctx context.Context, in DoctorNotificationListInput) (*DoctorNotificationListOutput, error)
	Read(ctx context.Context, in DoctorNotificationReadInput) error
}

var DoctorNotification IDoctorNotification = doctorNotificationService{}

type doctorNotificationService struct{}

func (s doctorNotificationService) List(ctx context.Context, in DoctorNotificationListInput) (*DoctorNotificationListOutput, error) {
	var (
		page = in.Page
		size = in.Size
		cols = dao.Notification.Columns()
	)

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	model := dao.Notification.Ctx(ctx).Where(cols.DoctorId, in.DoctorID)

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生通知列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生通知列表失败")
	}

	items := make([]DoctorNotificationItem, 0, len(records))
	for _, record := range records {
		sendTime := ""
		if !record[cols.SendTime].IsNil() {
			sendTime = record[cols.SendTime].GTime().Format(appointmentTimeLayout)
		}

		items = append(items, DoctorNotificationItem{
			ID:               record[cols.Id].Int64(),
			NotificationType: record[cols.NotificationType].Int(),
			Title:            record[cols.Title].String(),
			Content:          record[cols.Content].String(),
			AppointmentID:    record[cols.AppointmentId].Int64(),
			SendTime:         sendTime,
			Status:           record[cols.Status].Int(),
		})
	}

	return &DoctorNotificationListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s doctorNotificationService) Read(ctx context.Context, in DoctorNotificationReadInput) error {
	record, err := dao.Notification.Ctx(ctx).
		Where(dao.Notification.Columns().Id, in.NotificationID).
		Where(dao.Notification.Columns().DoctorId, in.DoctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询医生通知失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("通知不存在")
	}

	if record[dao.Notification.Columns().Status].Int() == 3 {
		return nil
	}

	result, err := dao.Notification.Ctx(ctx).
		Where(dao.Notification.Columns().Id, in.NotificationID).
		Where(dao.Notification.Columns().DoctorId, in.DoctorID).
		Data(do.Notification{
			Status:    3,
			UpdatedAt: time.Now(),
		}).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "标记通知已读失败")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("通知不存在")
	}
	return nil
}
