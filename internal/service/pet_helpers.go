package service

import (
	"context"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
)

const (
	dateFormat     = "2006-01-02"
	dateTimeFormat = "2006-01-02 15:04:05"
)

func parseDate(value string) (time.Time, error) {
	return time.ParseInLocation(dateFormat, value, time.Local)
}

func formatDate(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.Format(dateFormat)
}

func formatDateTime(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.Format(dateTimeFormat)
}

func ensurePetAccessible(ctx context.Context, petID int64, userID int64) error {
	cols := dao.Pet.Columns()
	model := dao.Pet.Ctx(ctx).
		Where(cols.Id, petID).
		Where(cols.Status, 1)
	if userID > 0 {
		model = model.Where(cols.UserId, userID)
	}
	record, err := model.One()
	if err != nil {
		return consts.WrapInternalError(err, "查询宠物失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("宠物不存在")
	}
	return nil
}
