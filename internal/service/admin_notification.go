package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
)

const (
	adminNotificationReceiverUser   = "user"
	adminNotificationReceiverDoctor = "doctor"
	adminNotificationTypeSystem     = 2
)

type (
	AdminNotificationCreateInput struct {
		ReceiverType     string
		ReceiverIDs      []int64
		NotificationType int
		Title            string
		Content          string
		SendTime         string
	}

	AdminNotificationCreateOutput struct {
		Count int
	}

	AdminNotificationListInput struct {
		Page             int
		Size             int
		NotificationType *int
	}

	AdminNotificationDeleteInput struct {
		NotificationID int64
	}

	AdminNotificationItem struct {
		ID               int64
		NotificationType int
		Title            string
		Content          string
		SendTime         string
		Status           int
		CreatedAt        string
	}

	AdminNotificationListOutput struct {
		Items []AdminNotificationItem
		Total int
		Page  int
		Size  int
	}
)

type IAdminNotification interface {
	Create(ctx context.Context, in AdminNotificationCreateInput) (*AdminNotificationCreateOutput, error)
	List(ctx context.Context, in AdminNotificationListInput) (*AdminNotificationListOutput, error)
	Delete(ctx context.Context, in AdminNotificationDeleteInput) error
}

var AdminNotification IAdminNotification = adminNotificationService{}

type adminNotificationService struct{}

func (s adminNotificationService) Create(ctx context.Context, in AdminNotificationCreateInput) (*AdminNotificationCreateOutput, error) {
	var receiverType = normalizeAdminNotificationReceiverType(in.ReceiverType)

	receiverIDs := uniquePositiveInt64s(in.ReceiverIDs)
	if len(receiverIDs) == 0 {
		return nil, consts.NewBadRequestError("请填写接收人ID列表")
	}
	if err := ensureNotificationReceiversExist(ctx, receiverType, receiverIDs); err != nil {
		return nil, err
	}

	sendTime, err := time.ParseInLocation(appointmentTimeLayout, in.SendTime, time.Local)
	if err != nil {
		return nil, consts.NewBadRequestError("发送时间格式不正确")
	}

	now := time.Now()
	err = dao.Notification.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, receiverID := range receiverIDs {
			data := do.Notification{
				UserId:           int64(0),
				DoctorId:         int64(0),
				AppointmentId:    int64(0),
				NotificationType: in.NotificationType,
				Title:            in.Title,
				Content:          in.Content,
				SendTime:         sendTime,
				Status:           1,
				CreatedAt:        now,
				UpdatedAt:        now,
			}

			switch receiverType {
			case adminNotificationReceiverUser:
				data.UserId = receiverID
			case adminNotificationReceiverDoctor:
				data.DoctorId = receiverID
			default:
				return consts.NewBadRequestError("接收人类型不合法")
			}

			if _, err = tx.Model(dao.Notification.Table()).Data(data).Insert(); err != nil {
				return consts.WrapInternalError(err, "发布系统通知失败")
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &AdminNotificationCreateOutput{Count: len(receiverIDs)}, nil
}

func (s adminNotificationService) List(ctx context.Context, in AdminNotificationListInput) (*AdminNotificationListOutput, error) {
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

	model := dao.Notification.Ctx(ctx).Where(cols.NotificationType, adminNotificationTypeSystem)
	if in.NotificationType != nil {
		model = model.Where(cols.NotificationType, *in.NotificationType)
	}

	records, err := model.OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询系统通知列表失败")
	}

	items := make([]AdminNotificationItem, 0, len(records))
	seen := make(map[string]int, len(records))
	for _, record := range records {
		key := adminNotificationBatchKey(record)
		if index, ok := seen[key]; ok {
			if items[index].Status != 2 && record[cols.Status].Int() == 2 {
				items[index].Status = 2
			}
			continue
		}

		seen[key] = len(items)
		items = append(items, adminNotificationItemFromRecord(record))
	}

	total := len(items)
	start := (page - 1) * size
	if start > total {
		start = total
	}
	end := start + size
	if end > total {
		end = total
	}
	items = items[start:end]

	return &AdminNotificationListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s adminNotificationService) Delete(ctx context.Context, in AdminNotificationDeleteInput) error {
	cols := dao.Notification.Columns()

	record, err := dao.Notification.Ctx(ctx).
		Where(cols.Id, in.NotificationID).
		Where(cols.NotificationType, adminNotificationTypeSystem).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询系统通知失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("系统通知不存在")
	}

	model := dao.Notification.Ctx(ctx).
		Where(cols.NotificationType, adminNotificationTypeSystem).
		Where(cols.Title, record[cols.Title].String()).
		Where(cols.Content, record[cols.Content].String()).
		Where(cols.SendTime, record[cols.SendTime].Time()).
		Where(cols.CreatedAt, record[cols.CreatedAt].Time())

	switch adminNotificationReceiverTypeFromRecord(record) {
	case adminNotificationReceiverUser:
		model = model.Where(cols.DoctorId, 0).WhereGT(cols.UserId, 0)
	case adminNotificationReceiverDoctor:
		model = model.Where(cols.UserId, 0).WhereGT(cols.DoctorId, 0)
	}

	result, err := model.Delete()
	if err != nil {
		return consts.WrapInternalError(err, "撤回系统通知失败")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取删除结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("系统通知不存在")
	}
	return nil
}

func adminNotificationItemFromRecord(record gdb.Record) AdminNotificationItem {
	cols := dao.Notification.Columns()

	sendTime := ""
	if !record[cols.SendTime].IsNil() {
		sendTime = record[cols.SendTime].GTime().Format(appointmentTimeLayout)
	}

	createdAt := ""
	if !record[cols.CreatedAt].IsNil() {
		createdAt = record[cols.CreatedAt].GTime().Format(appointmentTimeLayout)
	}

	status := record[cols.Status].Int()
	if status == 3 {
		status = 1
	}

	return AdminNotificationItem{
		ID:               record[cols.Id].Int64(),
		NotificationType: record[cols.NotificationType].Int(),
		Title:            record[cols.Title].String(),
		Content:          record[cols.Content].String(),
		SendTime:         sendTime,
		Status:           status,
		CreatedAt:        createdAt,
	}
}

func adminNotificationBatchKey(record gdb.Record) string {
	cols := dao.Notification.Columns()
	return fmt.Sprintf(
		"%s|%d|%s|%s|%s|%s",
		adminNotificationReceiverTypeFromRecord(record),
		record[cols.NotificationType].Int(),
		record[cols.Title].String(),
		record[cols.Content].String(),
		record[cols.SendTime].Time().Format(time.RFC3339Nano),
		record[cols.CreatedAt].Time().Format(time.RFC3339Nano),
	)
}

func adminNotificationReceiverTypeFromRecord(record gdb.Record) string {
	cols := dao.Notification.Columns()
	if record[cols.UserId].Int64() > 0 {
		return adminNotificationReceiverUser
	}
	return adminNotificationReceiverDoctor
}

func ensureNotificationReceiversExist(ctx context.Context, receiverType string, receiverIDs []int64) error {
	switch receiverType {
	case adminNotificationReceiverUser:
		count, err := dao.User.Ctx(ctx).
			WhereIn(dao.User.Columns().Id, receiverIDs).
			Count()
		if err != nil {
			return consts.WrapInternalError(err, "校验接收用户失败")
		}
		if count != len(receiverIDs) {
			return consts.NewNotFoundError("接收用户不存在")
		}
		return nil

	case adminNotificationReceiverDoctor:
		count, err := dao.Doctor.Ctx(ctx).
			WhereIn(dao.Doctor.Columns().Id, receiverIDs).
			Count()
		if err != nil {
			return consts.WrapInternalError(err, "校验接收医生失败")
		}
		if count != len(receiverIDs) {
			return consts.NewNotFoundError("接收医生不存在")
		}
		return nil

	default:
		return consts.NewBadRequestError("接收人类型不合法")
	}
}

func uniquePositiveInt64s(ids []int64) []int64 {
	if len(ids) == 0 {
		return nil
	}

	result := make([]int64, 0, len(ids))
	seen := make(map[int64]struct{}, len(ids))
	for _, id := range ids {
		if id <= 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		result = append(result, id)
	}
	return result
}

func normalizeAdminNotificationReceiverType(receiverType string) string {
	return strings.ToLower(strings.TrimSpace(receiverType))
}
