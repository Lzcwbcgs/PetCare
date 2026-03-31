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
	"github.com/gogf/gf/v2/os/gtime"
)

const appointmentTimeLayout = "2006-01-02 15:04:05"

type (
	AppointmentCreateInput struct {
		UserID             int64
		PetID              int64
		HospitalID         int64
		DoctorID           int64
		AppointmentType    int
		SymptomDescription string
		AppointmentTime    string
	}

	AppointmentCreateOutput struct {
		ID            int64
		AppointmentNo string
		Status        int
	}

	AppointmentListInput struct {
		UserID          int64
		Page            int
		Size            int
		Status          *int
		AppointmentType *int
	}

	AppointmentDetailInput struct {
		RequesterUserID int64
		RequesterRole   string
		AppointmentID   int64
	}

	AppointmentCancelInput struct {
		UserID        int64
		AppointmentID int64
		CancelReason  string
	}

	AppointmentItem struct {
		ID                 int64
		AppointmentNo      string
		UserID             int64
		UserNickname       string
		PetID              int64
		PetName            string
		HospitalID         int64
		HospitalName       string
		DoctorID           int64
		DoctorName         string
		AppointmentType    int
		SymptomDescription string
		AppointmentTime    string
		ReminderTime       string
		Status             int
		Source             int
		CreatedAt          string
		UpdatedAt          string
	}

	AppointmentListOutput struct {
		Items []AppointmentItem
		Total int
		Page  int
		Size  int
	}
)

type IAppointment interface {
	Create(ctx context.Context, in AppointmentCreateInput) (*AppointmentCreateOutput, error)
	List(ctx context.Context, in AppointmentListInput) (*AppointmentListOutput, error)
	Detail(ctx context.Context, in AppointmentDetailInput) (*AppointmentItem, error)
	Cancel(ctx context.Context, in AppointmentCancelInput) error
}

var Appointment IAppointment = appointmentService{}

type appointmentService struct{}

func (s appointmentService) Create(ctx context.Context, in AppointmentCreateInput) (*AppointmentCreateOutput, error) {
	if err := ensureUserOwnedPet(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}
	if err := ensureActiveHospital(ctx, in.HospitalID); err != nil {
		return nil, err
	}
	if err := ensureActiveDoctor(ctx, in.DoctorID, in.HospitalID); err != nil {
		return nil, err
	}
	if in.AppointmentType == 2 && strings.TrimSpace(in.SymptomDescription) == "" {
		return nil, consts.NewBadRequestError("看病预约请填写症状描述")
	}

	appointmentTime, err := time.ParseInLocation(appointmentTimeLayout, in.AppointmentTime, time.Local)
	if err != nil {
		return nil, consts.NewBadRequestError("预约时间格式不正确")
	}

	now := time.Now()
	reminderTime := appointmentTime.Add(-time.Hour)

	var output AppointmentCreateOutput
	err = dao.Appointment.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		appointmentNo, err := generateAppointmentNo(ctx, tx, now)
		if err != nil {
			return err
		}

		result, err := tx.Model(dao.Appointment.Table()).Data(do.Appointment{
			AppointmentNo:      appointmentNo,
			UserId:             in.UserID,
			PetId:              in.PetID,
			HospitalId:         in.HospitalID,
			DoctorId:           in.DoctorID,
			AppointmentType:    in.AppointmentType,
			SymptomDescription: strings.TrimSpace(in.SymptomDescription),
			AppointmentTime:    appointmentTime,
			ReminderTime:       reminderTime,
			Status:             1,
			Source:             1,
			CreatedAt:          now,
			UpdatedAt:          now,
		}).Insert()
		if err != nil {
			return consts.WrapInternalError(err, "创建预约失败")
		}

		lastID, err := result.LastInsertId()
		if err != nil {
			return consts.WrapInternalError(err, "获取预约ID失败")
		}

		output = AppointmentCreateOutput{
			ID:            lastID,
			AppointmentNo: appointmentNo,
			Status:        1,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (s appointmentService) List(ctx context.Context, in AppointmentListInput) (*AppointmentListOutput, error) {
	var (
		page = in.Page
		size = in.Size
		cols = dao.Appointment.Columns()
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

	model := dao.Appointment.Ctx(ctx).Where(cols.UserId, in.UserID)
	if in.Status != nil {
		model = model.Where(cols.Status, *in.Status)
	}
	if in.AppointmentType != nil {
		model = model.Where(cols.AppointmentType, *in.AppointmentType)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询预约列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询预约列表失败")
	}

	items, err := enrichAppointmentItems(ctx, records)
	if err != nil {
		return nil, err
	}

	return &AppointmentListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s appointmentService) Detail(ctx context.Context, in AppointmentDetailInput) (*AppointmentItem, error) {
	record, err := dao.Appointment.Ctx(ctx).
		Where(dao.Appointment.Columns().Id, in.AppointmentID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询预约详情失败")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("预约不存在")
	}

	if err = ensureAppointmentReadable(record, in.RequesterUserID, in.RequesterRole); err != nil {
		return nil, err
	}

	items, err := enrichAppointmentItems(ctx, gdb.Result{record})
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, consts.NewNotFoundError("预约不存在")
	}
	return &items[0], nil
}

func (s appointmentService) Cancel(ctx context.Context, in AppointmentCancelInput) error {
	record, err := dao.Appointment.Ctx(ctx).
		Where(dao.Appointment.Columns().Id, in.AppointmentID).
		Where(dao.Appointment.Columns().UserId, in.UserID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询预约失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("预约不存在")
	}

	status := record[dao.Appointment.Columns().Status].Int()
	switch status {
	case 2:
		return consts.NewConflictError("预约已完成，无法取消")
	case 3:
		return consts.NewConflictError("预约已取消")
	case 4:
		return consts.NewConflictError("预约已过期，无法取消")
	}

	appointmentTime := record[dao.Appointment.Columns().AppointmentTime].GTime()
	if appointmentTime != nil && appointmentTime.Before(gtime.New(time.Now())) {
		return consts.NewConflictError("预约已过期，无法取消")
	}

	result, err := dao.Appointment.Ctx(ctx).
		Where(dao.Appointment.Columns().Id, in.AppointmentID).
		Where(dao.Appointment.Columns().UserId, in.UserID).
		Data(do.Appointment{
			Status:    3,
			UpdatedAt: time.Now(),
		}).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "取消预约失败")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return consts.WrapInternalError(err, "获取更新结果失败")
	}
	if rowsAffected == 0 {
		return consts.NewNotFoundError("预约不存在")
	}
	return nil
}

func generateAppointmentNo(ctx context.Context, tx gdb.TX, now time.Time) (string, error) {
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	count, err := tx.Model(dao.Appointment.Table()).
		WhereGTE(dao.Appointment.Columns().CreatedAt, startOfDay).
		WhereLT(dao.Appointment.Columns().CreatedAt, endOfDay).
		Count()
	if err != nil {
		return "", consts.WrapInternalError(err, "生成预约单号失败")
	}

	return fmt.Sprintf("APT%s%04d", now.Format("20060102"), count+1), nil
}

func ensureUserOwnedPet(ctx context.Context, petID, userID int64) error {
	count, err := dao.Pet.Ctx(ctx).
		Where(dao.Pet.Columns().Id, petID).
		Where(dao.Pet.Columns().UserId, userID).
		Count()
	if err != nil {
		return consts.WrapInternalError(err, "查询宠物信息失败")
	}
	if count == 0 {
		return consts.NewNotFoundError("宠物不存在")
	}
	return nil
}

func ensureActiveDoctor(ctx context.Context, doctorID, hospitalID int64) error {
	record, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, doctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询医生信息失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("医生不存在")
	}
	if record[dao.Doctor.Columns().Status].Int() != 1 {
		return consts.NewConflictError("医生已停用")
	}
	if record[dao.Doctor.Columns().HospitalId].Int64() != hospitalID {
		return consts.NewConflictError("医生不属于当前医院")
	}
	return nil
}

func ensureAppointmentReadable(record gdb.Record, requesterUserID int64, requesterRole string) error {
	switch NormalizeRole(requesterRole) {
	case consts.RoleUser:
		if record[dao.Appointment.Columns().UserId].Int64() != requesterUserID {
			return consts.NewForbiddenError("")
		}
	case consts.RoleDoctor:
		if record[dao.Appointment.Columns().DoctorId].Int64() != requesterUserID {
			return consts.NewForbiddenError("")
		}
	case consts.RoleAdmin:
		return nil
	default:
		return consts.NewForbiddenError("")
	}
	return nil
}

func enrichAppointmentItems(ctx context.Context, records gdb.Result) ([]AppointmentItem, error) {
	petNameMap, err := loadPetNameMap(ctx, collectAppointmentPetIDs(records))
	if err != nil {
		return nil, err
	}
	hospitalNameMap, err := loadHospitalNameMap(ctx, collectAppointmentHospitalIDs(records))
	if err != nil {
		return nil, err
	}
	doctorNameMap, err := loadDoctorNameMap(ctx, collectAppointmentDoctorIDs(records))
	if err != nil {
		return nil, err
	}
	userNicknameMap, err := loadUserNicknameMap(ctx, collectAppointmentUserIDs(records))
	if err != nil {
		return nil, err
	}

	items := make([]AppointmentItem, 0, len(records))
	for _, record := range records {
		item := appointmentItemFromRecord(record)
		item.PetName = petNameMap[item.PetID]
		item.HospitalName = hospitalNameMap[item.HospitalID]
		item.DoctorName = doctorNameMap[item.DoctorID]
		item.UserNickname = userNicknameMap[item.UserID]
		items = append(items, item)
	}
	return items, nil
}

func collectAppointmentPetIDs(records gdb.Result) []int64 {
	return collectUniqueInt64(records, func(record gdb.Record) int64 {
		return record[dao.Appointment.Columns().PetId].Int64()
	})
}

func collectAppointmentHospitalIDs(records gdb.Result) []int64 {
	return collectUniqueInt64(records, func(record gdb.Record) int64 {
		return record[dao.Appointment.Columns().HospitalId].Int64()
	})
}

func collectAppointmentDoctorIDs(records gdb.Result) []int64 {
	return collectUniqueInt64(records, func(record gdb.Record) int64 {
		return record[dao.Appointment.Columns().DoctorId].Int64()
	})
}

func collectAppointmentUserIDs(records gdb.Result) []int64 {
	return collectUniqueInt64(records, func(record gdb.Record) int64 {
		return record[dao.Appointment.Columns().UserId].Int64()
	})
}

func collectUniqueInt64(records gdb.Result, extractor func(record gdb.Record) int64) []int64 {
	idMap := make(map[int64]struct{})
	ids := make([]int64, 0)
	for _, record := range records {
		id := extractor(record)
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

func loadPetNameMap(ctx context.Context, petIDs []int64) (map[int64]string, error) {
	nameMap := make(map[int64]string, len(petIDs))
	if len(petIDs) == 0 {
		return nameMap, nil
	}

	records, err := dao.Pet.Ctx(ctx).
		WhereIn(dao.Pet.Columns().Id, petIDs).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物信息失败")
	}
	for _, record := range records {
		nameMap[record[dao.Pet.Columns().Id].Int64()] = record[dao.Pet.Columns().PetName].String()
	}
	return nameMap, nil
}

func loadDoctorNameMap(ctx context.Context, doctorIDs []int64) (map[int64]string, error) {
	nameMap := make(map[int64]string, len(doctorIDs))
	if len(doctorIDs) == 0 {
		return nameMap, nil
	}

	records, err := dao.Doctor.Ctx(ctx).
		WhereIn(dao.Doctor.Columns().Id, doctorIDs).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生信息失败")
	}
	for _, record := range records {
		nameMap[record[dao.Doctor.Columns().Id].Int64()] = record[dao.Doctor.Columns().DoctorName].String()
	}
	return nameMap, nil
}

func loadUserNicknameMap(ctx context.Context, userIDs []int64) (map[int64]string, error) {
	nameMap := make(map[int64]string, len(userIDs))
	if len(userIDs) == 0 {
		return nameMap, nil
	}

	records, err := dao.User.Ctx(ctx).
		WhereIn(dao.User.Columns().Id, userIDs).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询用户信息失败")
	}
	for _, record := range records {
		nameMap[record[dao.User.Columns().Id].Int64()] = record[dao.User.Columns().Nickname].String()
	}
	return nameMap, nil
}

func appointmentItemFromRecord(record gdb.Record) AppointmentItem {
	cols := dao.Appointment.Columns()
	return AppointmentItem{
		ID:                 record[cols.Id].Int64(),
		AppointmentNo:      record[cols.AppointmentNo].String(),
		UserID:             record[cols.UserId].Int64(),
		PetID:              record[cols.PetId].Int64(),
		HospitalID:         record[cols.HospitalId].Int64(),
		DoctorID:           record[cols.DoctorId].Int64(),
		AppointmentType:    record[cols.AppointmentType].Int(),
		SymptomDescription: record[cols.SymptomDescription].String(),
		AppointmentTime:    record[cols.AppointmentTime].GTime().Format(appointmentTimeLayout),
		ReminderTime:       record[cols.ReminderTime].GTime().Format(appointmentTimeLayout),
		Status:             record[cols.Status].Int(),
		Source:             record[cols.Source].Int(),
		CreatedAt:          record[cols.CreatedAt].GTime().Format(appointmentTimeLayout),
		UpdatedAt:          record[cols.UpdatedAt].GTime().Format(appointmentTimeLayout),
	}
}
