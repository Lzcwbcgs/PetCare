package service

import (
	"context"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	operatorTypeAdmin  = 1
	operatorTypeDoctor = 2
	operatorTypeUser   = 3
)

type OperationLogRecordInput struct {
	OperatorType    int
	OperatorID      int64
	OperationModule string
	OperationType   string
	OperationDesc   string
}

func RecordOperationLog(ctx context.Context, in OperationLogRecordInput) {
	module := strings.TrimSpace(in.OperationModule)
	opType := strings.TrimSpace(in.OperationType)
	if in.OperatorType <= 0 || in.OperatorID <= 0 || module == "" || opType == "" {
		return
	}

	var ip string
	if r := g.RequestFromCtx(ctx); r != nil {
		ip = r.GetClientIp()
	}

	_, _ = dao.OperationLog.Ctx(ctx).Data(do.OperationLog{
		OperatorType:    in.OperatorType,
		OperatorId:      in.OperatorID,
		OperationModule: module,
		OperationType:   opType,
		OperationDesc:   strings.TrimSpace(in.OperationDesc),
		IpAddress:       ip,
		CreatedAt:       time.Now(),
	}).Insert()
}

func RecordOperationLogByRole(ctx context.Context, role string, userID int64, module, opType, desc string) {
	operatorType, ok := roleToOperatorType(role)
	if !ok {
		return
	}
	RecordOperationLog(ctx, OperationLogRecordInput{
		OperatorType:    operatorType,
		OperatorID:      userID,
		OperationModule: module,
		OperationType:   opType,
		OperationDesc:   desc,
	})
}

func roleToOperatorType(role string) (int, bool) {
	switch NormalizeRole(role) {
	case consts.RoleAdmin:
		return operatorTypeAdmin, true
	case consts.RoleDoctor:
		return operatorTypeDoctor, true
	case consts.RoleUser:
		return operatorTypeUser, true
	default:
		return 0, false
	}
}
