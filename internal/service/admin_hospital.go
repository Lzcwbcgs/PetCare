package service

import (
	"context"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
)

const adminHospitalTimeLayout = "2006-01-02 15:04:05"

type (
	AdminHospitalCreateInput struct {
		HospitalName string
		Address      string
		Phone        string
		Description  *string
		Status       int
	}

	AdminHospitalCreateOutput struct {
		ID int64
	}

	AdminHospitalListInput struct {
		Page    int
		Size    int
		Status  *int
		Keyword *string
	}

	AdminHospitalDetailInput struct {
		HospitalID int64
	}

	AdminHospitalUpdateInput struct {
		HospitalID   int64
		HospitalName *string
		Address      *string
		Phone        *string
		Description  *string
		Status       *int
	}

	AdminHospitalDeleteInput struct {
		HospitalID int64
	}

	AdminHospitalItem struct {
		ID           int64
		HospitalName string
		Address      string
		Phone        string
		Description  string
		Status       int
		CreatedAt    string
		UpdatedAt    string
	}

	AdminHospitalListOutput struct {
		Items []AdminHospitalItem
		Total int
		Page  int
		Size  int
	}
)

type IAdminHospital interface {
	Create(ctx context.Context, in AdminHospitalCreateInput) (*AdminHospitalCreateOutput, error)
	List(ctx context.Context, in AdminHospitalListInput) (*AdminHospitalListOutput, error)
	Detail(ctx context.Context, in AdminHospitalDetailInput) (*AdminHospitalItem, error)
	Update(ctx context.Context, in AdminHospitalUpdateInput) error
	Delete(ctx context.Context, in AdminHospitalDeleteInput) error
}

var AdminHospital IAdminHospital = adminHospitalService{}

type adminHospitalService struct{}

func (s adminHospitalService) Create(ctx context.Context, in AdminHospitalCreateInput) (*AdminHospitalCreateOutput, error) {
	now := time.Now()
	data := do.Hospital{
		HospitalName: in.HospitalName,
		Address:      in.Address,
		Phone:        in.Phone,
		Status:       in.Status,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if in.Description != nil {
		data.Description = *in.Description
	}

	result, err := dao.Hospital.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增医院失败")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取医院ID失败")
	}
	return &AdminHospitalCreateOutput{ID: id}, nil
}

func (s adminHospitalService) List(ctx context.Context, in AdminHospitalListInput) (*AdminHospitalListOutput, error) {
	var (
		page = in.Page
		size = in.Size
		cols = dao.Hospital.Columns()
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

	model := dao.Hospital.Ctx(ctx)
	if in.Status != nil {
		model = model.Where(cols.Status, *in.Status)
	}
	if in.Keyword != nil && strings.TrimSpace(*in.Keyword) != "" {
		model = model.WhereLike(cols.HospitalName, "%"+strings.TrimSpace(*in.Keyword)+"%")
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医院列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医院列表失败")
	}

	items := make([]AdminHospitalItem, 0, len(records))
	for _, record := range records {
		items = append(items, adminHospitalItemFromRecord(record))
	}

	return &AdminHospitalListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s adminHospitalService) Detail(ctx context.Context, in AdminHospitalDetailInput) (*AdminHospitalItem, error) {
	record, err := dao.Hospital.Ctx(ctx).
		Where(dao.Hospital.Columns().Id, in.HospitalID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医院详情失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("医院不存在")
	}

	item := adminHospitalItemFromRecord(record)
	return &item, nil
}

func (s adminHospitalService) Update(ctx context.Context, in AdminHospitalUpdateInput) error {
	if !hasAdminHospitalUpdates(in) {
		return consts.NewBadRequestError("至少提供一个更新字段")
	}

	data := do.Hospital{
		UpdatedAt: time.Now(),
	}
	if in.HospitalName != nil {
		data.HospitalName = *in.HospitalName
	}
	if in.Address != nil {
		data.Address = *in.Address
	}
	if in.Phone != nil {
		data.Phone = *in.Phone
	}
	if in.Description != nil {
		data.Description = *in.Description
	}
	if in.Status != nil {
		data.Status = *in.Status
	}

	result, err := dao.Hospital.Ctx(ctx).
		Where(dao.Hospital.Columns().Id, in.HospitalID).
		Data(data).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "修改医院失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("医院不存在")
	}
	return nil
}

func (s adminHospitalService) Delete(ctx context.Context, in AdminHospitalDeleteInput) error {
	var (
		hospitalCols = dao.Hospital.Columns()
		doctorCols   = dao.Doctor.Columns()
		now          = time.Now()
	)

	err := dao.Hospital.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		hospitalResult, err := tx.Model(dao.Hospital.Table()).
			Where(hospitalCols.Id, in.HospitalID).
			Data(do.Hospital{
				Status:    0,
				UpdatedAt: now,
			}).
			Update()
		if err != nil {
			return consts.WrapInternalError(err, "删除医院失败")
		}

		rowsAffected, err := hospitalResult.RowsAffected()
		if err != nil {
			return consts.WrapInternalError(err, "获取删除结果失败")
		}
		if rowsAffected == 0 {
			return consts.NewNotFoundError("医院不存在")
		}

		_, err = tx.Model(dao.Doctor.Table()).
			Where(doctorCols.HospitalId, in.HospitalID).
			Data(do.Doctor{
				Status:    0,
				UpdatedAt: now,
			}).
			Update()
		if err != nil {
			return consts.WrapInternalError(err, "同步停用医院医生失败")
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func hasAdminHospitalUpdates(in AdminHospitalUpdateInput) bool {
	return in.HospitalName != nil ||
		in.Address != nil ||
		in.Phone != nil ||
		in.Description != nil ||
		in.Status != nil
}

func adminHospitalItemFromRecord(record gdb.Record) AdminHospitalItem {
	cols := dao.Hospital.Columns()
	return AdminHospitalItem{
		ID:           record[cols.Id].Int64(),
		HospitalName: record[cols.HospitalName].String(),
		Address:      record[cols.Address].String(),
		Phone:        record[cols.Phone].String(),
		Description:  record[cols.Description].String(),
		Status:       record[cols.Status].Int(),
		CreatedAt:    record[cols.CreatedAt].Time().Format(adminHospitalTimeLayout),
		UpdatedAt:    record[cols.UpdatedAt].Time().Format(adminHospitalTimeLayout),
	}
}
