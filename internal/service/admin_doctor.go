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

const adminDoctorTimeLayout = "2006-01-02 15:04:05"

type (
	AdminDoctorCreateInput struct {
		HospitalID int64
		Username   string
		Password   string
		DoctorName string
		Gender     int
		Phone      string
		Email      *string
		Title      string
		Specialty  string
		AvatarURL  *string
		Intro      *string
		Status     int
	}

	AdminDoctorCreateOutput struct {
		ID int64
	}

	AdminDoctorListInput struct {
		Page       int
		Size       int
		HospitalID *int64
		Status     *int
		Keyword    *string
	}

	AdminDoctorDetailInput struct {
		DoctorID int64
	}

	AdminDoctorUpdateInput struct {
		DoctorID   int64
		HospitalID *int64
		DoctorName *string
		Gender     *int
		Phone      *string
		Email      *string
		Title      *string
		Specialty  *string
		AvatarURL  *string
		Intro      *string
		Status     *int
	}

	AdminDoctorDeleteInput struct {
		DoctorID int64
	}

	AdminDoctorItem struct {
		ID           int64
		HospitalID   int64
		HospitalName string
		Username     string
		DoctorName   string
		Gender       int
		Phone        string
		Email        string
		Title        string
		Specialty    string
		AvatarURL    string
		Intro        string
		Status       int
		CreatedAt    string
		UpdatedAt    string
	}

	AdminDoctorListOutput struct {
		Items []AdminDoctorItem
		Total int
		Page  int
		Size  int
	}
)

type IAdminDoctor interface {
	Create(ctx context.Context, in AdminDoctorCreateInput) (*AdminDoctorCreateOutput, error)
	List(ctx context.Context, in AdminDoctorListInput) (*AdminDoctorListOutput, error)
	Detail(ctx context.Context, in AdminDoctorDetailInput) (*AdminDoctorItem, error)
	Update(ctx context.Context, in AdminDoctorUpdateInput) error
	Delete(ctx context.Context, in AdminDoctorDeleteInput) error
}

var AdminDoctor IAdminDoctor = adminDoctorService{}

type adminDoctorService struct{}

func (s adminDoctorService) Create(ctx context.Context, in AdminDoctorCreateInput) (*AdminDoctorCreateOutput, error) {
	if err := ensureActiveHospital(ctx, in.HospitalID); err != nil {
		return nil, err
	}

	now := time.Now()
	data := do.Doctor{
		HospitalId:   in.HospitalID,
		Username:     in.Username,
		PasswordHash: hashPassword(in.Password),
		DoctorName:   in.DoctorName,
		Gender:       in.Gender,
		Phone:        in.Phone,
		Title:        in.Title,
		Specialty:    in.Specialty,
		Status:       in.Status,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if in.Email != nil {
		data.Email = *in.Email
	}
	if in.AvatarURL != nil {
		data.AvatarUrl = *in.AvatarURL
	}
	if in.Intro != nil {
		data.Intro = *in.Intro
	}

	result, err := dao.Doctor.Ctx(ctx).Data(data).Insert()
	if err != nil {
		if isDuplicateErr(err) {
			return nil, consts.NewConflictError("医生登录账号已存在")
		}
		return nil, consts.WrapInternalError(err, "新增医生失败")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "获取医生ID失败")
	}
	return &AdminDoctorCreateOutput{ID: id}, nil
}

func (s adminDoctorService) List(ctx context.Context, in AdminDoctorListInput) (*AdminDoctorListOutput, error) {
	var (
		page = in.Page
		size = in.Size
		cols = dao.Doctor.Columns()
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

	model := dao.Doctor.Ctx(ctx)
	if in.HospitalID != nil {
		model = model.Where(cols.HospitalId, *in.HospitalID)
	}
	if in.Status != nil {
		model = model.Where(cols.Status, *in.Status)
	}
	if in.Keyword != nil && strings.TrimSpace(*in.Keyword) != "" {
		model = model.WhereLike(cols.DoctorName, "%"+strings.TrimSpace(*in.Keyword)+"%")
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生列表失败")
	}

	hospitalNameMap, err := loadHospitalNameMap(ctx, collectDoctorHospitalIDs(records))
	if err != nil {
		return nil, err
	}

	items := make([]AdminDoctorItem, 0, len(records))
	for _, record := range records {
		item := adminDoctorItemFromRecord(record)
		item.HospitalName = hospitalNameMap[item.HospitalID]
		items = append(items, item)
	}

	return &AdminDoctorListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s adminDoctorService) Detail(ctx context.Context, in AdminDoctorDetailInput) (*AdminDoctorItem, error) {
	record, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, in.DoctorID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生详情失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("医生不存在")
	}

	item := adminDoctorItemFromRecord(record)
	return &item, nil
}

func (s adminDoctorService) Update(ctx context.Context, in AdminDoctorUpdateInput) error {
	if !hasAdminDoctorUpdates(in) {
		return consts.NewBadRequestError("至少提供一个更新字段")
	}
	if in.HospitalID != nil {
		if err := ensureActiveHospital(ctx, *in.HospitalID); err != nil {
			return err
		}
	}

	data := do.Doctor{
		UpdatedAt: time.Now(),
	}
	if in.HospitalID != nil {
		data.HospitalId = *in.HospitalID
	}
	if in.DoctorName != nil {
		data.DoctorName = *in.DoctorName
	}
	if in.Gender != nil {
		data.Gender = *in.Gender
	}
	if in.Phone != nil {
		data.Phone = *in.Phone
	}
	if in.Email != nil {
		data.Email = *in.Email
	}
	if in.Title != nil {
		data.Title = *in.Title
	}
	if in.Specialty != nil {
		data.Specialty = *in.Specialty
	}
	if in.AvatarURL != nil {
		data.AvatarUrl = *in.AvatarURL
	}
	if in.Intro != nil {
		data.Intro = *in.Intro
	}
	if in.Status != nil {
		data.Status = *in.Status
	}

	result, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, in.DoctorID).
		Data(data).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "修改医生信息失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("医生不存在")
	}
	return nil
}

func (s adminDoctorService) Delete(ctx context.Context, in AdminDoctorDeleteInput) error {
	result, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, in.DoctorID).
		Data(do.Doctor{
			Status:    0,
			UpdatedAt: time.Now(),
		}).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "删除医生失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取删除结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("医生不存在")
	}
	return nil
}

func hasAdminDoctorUpdates(in AdminDoctorUpdateInput) bool {
	return in.HospitalID != nil ||
		in.DoctorName != nil ||
		in.Gender != nil ||
		in.Phone != nil ||
		in.Email != nil ||
		in.Title != nil ||
		in.Specialty != nil ||
		in.AvatarURL != nil ||
		in.Intro != nil ||
		in.Status != nil
}

func ensureActiveHospital(ctx context.Context, hospitalID int64) error {
	record, err := dao.Hospital.Ctx(ctx).
		Where(dao.Hospital.Columns().Id, hospitalID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询医院信息失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("医院不存在")
	}
	if record[dao.Hospital.Columns().Status].Int() != 1 {
		return consts.NewConflictError("医院已停用")
	}
	return nil
}

func collectDoctorHospitalIDs(records gdb.Result) []int64 {
	idMap := make(map[int64]struct{})
	ids := make([]int64, 0)
	for _, record := range records {
		id := record[dao.Doctor.Columns().HospitalId].Int64()
		if id <= 0 {
			continue
		}
		if _, ok := idMap[id]; ok {
			continue
		}
		idMap[id] = struct{}{}
		ids = append(ids, id)
	}
	return ids
}

func loadHospitalNameMap(ctx context.Context, hospitalIDs []int64) (map[int64]string, error) {
	nameMap := make(map[int64]string, len(hospitalIDs))
	if len(hospitalIDs) == 0 {
		return nameMap, nil
	}

	records, err := dao.Hospital.Ctx(ctx).
		WhereIn(dao.Hospital.Columns().Id, hospitalIDs).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医院信息失败")
	}

	for _, record := range records {
		nameMap[record[dao.Hospital.Columns().Id].Int64()] = record[dao.Hospital.Columns().HospitalName].String()
	}
	return nameMap, nil
}

func adminDoctorItemFromRecord(record gdb.Record) AdminDoctorItem {
	cols := dao.Doctor.Columns()
	return AdminDoctorItem{
		ID:         record[cols.Id].Int64(),
		HospitalID: record[cols.HospitalId].Int64(),
		Username:   record[cols.Username].String(),
		DoctorName: record[cols.DoctorName].String(),
		Gender:     record[cols.Gender].Int(),
		Phone:      record[cols.Phone].String(),
		Email:      record[cols.Email].String(),
		Title:      record[cols.Title].String(),
		Specialty:  record[cols.Specialty].String(),
		AvatarURL:  record[cols.AvatarUrl].String(),
		Intro:      record[cols.Intro].String(),
		Status:     record[cols.Status].Int(),
		CreatedAt:  record[cols.CreatedAt].Time().Format(adminDoctorTimeLayout),
		UpdatedAt:  record[cols.UpdatedAt].Time().Format(adminDoctorTimeLayout),
	}
}
