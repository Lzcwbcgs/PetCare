package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	NotificationListInput struct {
		RequesterUserID int64
		RequesterRole   string
		Page            int
		Size            int
		Status          *int
	}

	NotificationItem struct {
		ID               int64
		NotificationType int
		Title            string
		Content          string
		AppointmentID    int64
		SendTime         string
		Status           int
	}

	NotificationListOutput struct {
		Items []NotificationItem
		Total int
		Page  int
		Size  int
	}

	NotificationDetailInput struct {
		RequesterUserID int64
		RequesterRole   string
		NotificationID  int64
	}

	NotificationReadInput struct {
		RequesterUserID int64
		RequesterRole   string
		NotificationID  int64
	}
)

type INotification interface {
	List(ctx context.Context, in NotificationListInput) (*NotificationListOutput, error)
	Detail(ctx context.Context, in NotificationDetailInput) (*NotificationItem, error)
	Read(ctx context.Context, in NotificationReadInput) error
}

var Notification INotification = notificationService{}

type notificationService struct{}

func (s notificationService) List(ctx context.Context, in NotificationListInput) (*NotificationListOutput, error) {
	var (
		page  = in.Page
		size  = in.Size
		cols  = dao.Notification.Columns()
		model = s.notificationModel(ctx, in.RequesterUserID, in.RequesterRole)
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
	if model == nil {
		return nil, consts.NewForbiddenError("")
	}
	if in.Status != nil {
		model = model.Where(cols.Status, *in.Status)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询通知列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询通知列表失败")
	}

	items := make([]NotificationItem, 0, len(records))
	for _, record := range records {
		items = append(items, notificationItemFromRecord(record))
	}

	return &NotificationListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s notificationService) Detail(ctx context.Context, in NotificationDetailInput) (*NotificationItem, error) {
	record, err := s.readableNotification(ctx, in.RequesterUserID, in.RequesterRole, in.NotificationID)
	if err != nil {
		return nil, err
	}

	item := notificationItemFromRecord(record)
	return &item, nil
}

func (s notificationService) Read(ctx context.Context, in NotificationReadInput) error {
	record, err := s.readableNotification(ctx, in.RequesterUserID, in.RequesterRole, in.NotificationID)
	if err != nil {
		return err
	}
	if record[dao.Notification.Columns().Status].Int() == 3 {
		return nil
	}

	model := s.notificationModel(ctx, in.RequesterUserID, in.RequesterRole)
	if model == nil {
		return consts.NewForbiddenError("")
	}

	result, err := model.
		Where(dao.Notification.Columns().Id, in.NotificationID).
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

func (s notificationService) readableNotification(ctx context.Context, requesterUserID int64, requesterRole string, notificationID int64) (gdb.Record, error) {
	model := s.notificationModel(ctx, requesterUserID, requesterRole)
	if model == nil {
		return nil, consts.NewForbiddenError("")
	}

	record, err := model.
		Where(dao.Notification.Columns().Id, notificationID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询通知详情失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("通知不存在")
	}
	return record, nil
}

func (s notificationService) notificationModel(ctx context.Context, requesterUserID int64, requesterRole string) *gdb.Model {
	switch NormalizeRole(requesterRole) {
	case consts.RoleUser:
		return dao.Notification.Ctx(ctx).Where(dao.Notification.Columns().UserId, requesterUserID)
	case consts.RoleDoctor:
		return dao.Notification.Ctx(ctx).Where(dao.Notification.Columns().DoctorId, requesterUserID)
	default:
		return nil
	}
}

func notificationItemFromRecord(record gdb.Record) NotificationItem {
	cols := dao.Notification.Columns()

	sendTime := ""
	if !record[cols.SendTime].IsNil() {
		sendTime = record[cols.SendTime].GTime().Format(appointmentTimeLayout)
	}

	return NotificationItem{
		ID:               record[cols.Id].Int64(),
		NotificationType: record[cols.NotificationType].Int(),
		Title:            record[cols.Title].String(),
		Content:          record[cols.Content].String(),
		AppointmentID:    record[cols.AppointmentId].Int64(),
		SendTime:         sendTime,
		Status:           record[cols.Status].Int(),
	}
}
