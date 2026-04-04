package service

import (
	"context"
	"strings"

	"PetCare/internal/consts"
	"PetCare/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	AdminOperationLogListInput struct {
		Page            int
		Size            int
		OperatorType    *int
		OperationModule *string
	}

	AdminOperationLogItem struct {
		ID              int64
		OperatorType    int
		OperatorID      int64
		OperationModule string
		OperationType   string
		OperationDesc   string
		IPAddress       string
		CreatedAt       string
	}

	AdminOperationLogListOutput struct {
		Items []AdminOperationLogItem
		Total int
		Page  int
		Size  int
	}

	AdminOperationLogDetailInput struct {
		LogID int64
	}
)

type IAdminOperationLog interface {
	List(ctx context.Context, in AdminOperationLogListInput) (*AdminOperationLogListOutput, error)
	Detail(ctx context.Context, in AdminOperationLogDetailInput) (*AdminOperationLogItem, error)
}

var AdminOperationLog IAdminOperationLog = adminOperationLogService{}

type adminOperationLogService struct{}

func (s adminOperationLogService) List(ctx context.Context, in AdminOperationLogListInput) (*AdminOperationLogListOutput, error) {
	page := in.Page
	size := in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	model := dao.OperationLog.Ctx(ctx)
	if in.OperatorType != nil {
		model = model.Where(dao.OperationLog.Columns().OperatorType, *in.OperatorType)
	}
	if in.OperationModule != nil && strings.TrimSpace(*in.OperationModule) != "" {
		model = model.Where(dao.OperationLog.Columns().OperationModule, strings.TrimSpace(*in.OperationModule))
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query operation logs failed")
	}

	records, err := model.Page(page, size).OrderDesc(dao.OperationLog.Columns().Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query operation logs failed")
	}

	items := make([]AdminOperationLogItem, 0, len(records))
	for _, record := range records {
		items = append(items, adminOperationLogFromRecord(record))
	}

	return &AdminOperationLogListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s adminOperationLogService) Detail(ctx context.Context, in AdminOperationLogDetailInput) (*AdminOperationLogItem, error) {
	record, err := dao.OperationLog.Ctx(ctx).Where(dao.OperationLog.Columns().Id, in.LogID).One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query operation log detail failed")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("operation log not found")
	}

	item := adminOperationLogFromRecord(record)
	return &item, nil
}

func adminOperationLogFromRecord(record gdb.Record) AdminOperationLogItem {
	return AdminOperationLogItem{
		ID:              record[dao.OperationLog.Columns().Id].Int64(),
		OperatorType:    record[dao.OperationLog.Columns().OperatorType].Int(),
		OperatorID:      record[dao.OperationLog.Columns().OperatorId].Int64(),
		OperationModule: record[dao.OperationLog.Columns().OperationModule].String(),
		OperationType:   record[dao.OperationLog.Columns().OperationType].String(),
		OperationDesc:   record[dao.OperationLog.Columns().OperationDesc].String(),
		IPAddress:       record[dao.OperationLog.Columns().IpAddress].String(),
		CreatedAt:       formatGTime(record[dao.OperationLog.Columns().CreatedAt].GTime()),
	}
}
