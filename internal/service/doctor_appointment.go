package service

import (
	"context"
	"strconv"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

const doctorAppointmentDateLayout = "2006-01-02"

type (
	DoctorAppointmentListInput struct {
		DoctorID int64
		Page     int
		Size     int
		Status   *int
	}

	DoctorAppointmentItem struct {
		ID              int64
		AppointmentNo   string
		PetID           int64
		PetName         string
		UserID          int64
		UserNickname    string
		AppointmentType int
		AppointmentTime string
		Status          int
	}

	DoctorAppointmentListOutput struct {
		Items []DoctorAppointmentItem
		Total int
		Page  int
		Size  int
	}

	DoctorAppointmentDetailInput struct {
		DoctorID      int64
		AppointmentID int64
	}

	DoctorAppointmentUpdateStatusInput struct {
		DoctorID      int64
		AppointmentID int64
		Status        int
	}

	DoctorAppointmentSummary struct {
		ID                 int64
		AppointmentNo      string
		AppointmentType    int
		SymptomDescription string
		AppointmentTime    string
		Status             int
	}

	DoctorAppointmentPet struct {
		ID         int64
		PetName    string
		PetType    string
		Gender     int
		Age        int
		AgeUnit    string
		Breed      string
		Weight     string
		Sterilized int
		Remark     string
	}

	DoctorAppointmentMedicalHistory struct {
		ID          int64
		HistoryType string
		Description string
		DiagnosedAt string
		IsCurrent   int
	}

	DoctorAppointmentVaccination struct {
		ID              int64
		VaccineName     string
		VaccinationDate string
		NextDueDate     string
	}

	DoctorAppointmentAllergy struct {
		ID                 int64
		Allergen           string
		SymptomDescription string
		SeverityLevel      int
	}

	DoctorAppointmentDetailOutput struct {
		Appointment      DoctorAppointmentSummary
		Pet              DoctorAppointmentPet
		MedicalHistories []DoctorAppointmentMedicalHistory
		Vaccinations     []DoctorAppointmentVaccination
		Allergies        []DoctorAppointmentAllergy
	}
)

type IDoctorAppointment interface {
	List(ctx context.Context, in DoctorAppointmentListInput) (*DoctorAppointmentListOutput, error)
	Detail(ctx context.Context, in DoctorAppointmentDetailInput) (*DoctorAppointmentDetailOutput, error)
	UpdateStatus(ctx context.Context, in DoctorAppointmentUpdateStatusInput) error
}

var DoctorAppointment IDoctorAppointment = doctorAppointmentService{}

type doctorAppointmentService struct{}

func (s doctorAppointmentService) List(ctx context.Context, in DoctorAppointmentListInput) (*DoctorAppointmentListOutput, error) {
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

	model := dao.Appointment.Ctx(ctx).Where(cols.DoctorId, in.DoctorID)
	if in.Status != nil {
		model = model.Where(cols.Status, *in.Status)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生预约列表失败")
	}

	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询医生预约列表失败")
	}

	petNameMap, err := loadPetNameMap(ctx, collectAppointmentPetIDs(records))
	if err != nil {
		return nil, err
	}
	userNicknameMap, err := loadUserNicknameMap(ctx, collectAppointmentUserIDs(records))
	if err != nil {
		return nil, err
	}

	items := make([]DoctorAppointmentItem, 0, len(records))
	for _, record := range records {
		items = append(items, DoctorAppointmentItem{
			ID:              record[cols.Id].Int64(),
			AppointmentNo:   record[cols.AppointmentNo].String(),
			PetID:           record[cols.PetId].Int64(),
			PetName:         petNameMap[record[cols.PetId].Int64()],
			UserID:          record[cols.UserId].Int64(),
			UserNickname:    userNicknameMap[record[cols.UserId].Int64()],
			AppointmentType: record[cols.AppointmentType].Int(),
			AppointmentTime: record[cols.AppointmentTime].GTime().Format(appointmentTimeLayout),
			Status:          record[cols.Status].Int(),
		})
	}

	return &DoctorAppointmentListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (s doctorAppointmentService) Detail(ctx context.Context, in DoctorAppointmentDetailInput) (*DoctorAppointmentDetailOutput, error) {
	appointmentRecord, err := dao.Appointment.Ctx(ctx).
		Where(dao.Appointment.Columns().Id, in.AppointmentID).
		Where(dao.Appointment.Columns().DoctorId, in.DoctorID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询预约接诊详情失败")
	}
	if appointmentRecord.IsEmpty() {
		return nil, consts.NewNotFoundError("预约不存在")
	}

	petRecord, err := dao.Pet.Ctx(ctx).
		Where(dao.Pet.Columns().Id, appointmentRecord[dao.Appointment.Columns().PetId].Int64()).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询宠物信息失败")
	}
	if petRecord.IsEmpty() {
		return nil, consts.NewNotFoundError("宠物不存在")
	}

	medicalHistories, err := loadDoctorAppointmentMedicalHistories(ctx, petRecord[dao.Pet.Columns().Id].Int64())
	if err != nil {
		return nil, err
	}
	vaccinations, err := loadDoctorAppointmentVaccinations(ctx, petRecord[dao.Pet.Columns().Id].Int64())
	if err != nil {
		return nil, err
	}
	allergies, err := loadDoctorAppointmentAllergies(ctx, petRecord[dao.Pet.Columns().Id].Int64())
	if err != nil {
		return nil, err
	}

	return &DoctorAppointmentDetailOutput{
		Appointment: DoctorAppointmentSummary{
			ID:                 appointmentRecord[dao.Appointment.Columns().Id].Int64(),
			AppointmentNo:      appointmentRecord[dao.Appointment.Columns().AppointmentNo].String(),
			AppointmentType:    appointmentRecord[dao.Appointment.Columns().AppointmentType].Int(),
			SymptomDescription: appointmentRecord[dao.Appointment.Columns().SymptomDescription].String(),
			AppointmentTime:    appointmentRecord[dao.Appointment.Columns().AppointmentTime].GTime().Format(appointmentTimeLayout),
			Status:             appointmentRecord[dao.Appointment.Columns().Status].Int(),
		},
		Pet: DoctorAppointmentPet{
			ID:         petRecord[dao.Pet.Columns().Id].Int64(),
			PetName:    petRecord[dao.Pet.Columns().PetName].String(),
			PetType:    petRecord[dao.Pet.Columns().PetType].String(),
			Gender:     petRecord[dao.Pet.Columns().Gender].Int(),
			Age:        petRecord[dao.Pet.Columns().Age].Int(),
			AgeUnit:    petRecord[dao.Pet.Columns().AgeUnit].String(),
			Breed:      petRecord[dao.Pet.Columns().Breed].String(),
			Weight:     formatDoctorAppointmentWeight(petRecord[dao.Pet.Columns().Weight].Float64()),
			Sterilized: petRecord[dao.Pet.Columns().Sterilized].Int(),
			Remark:     petRecord[dao.Pet.Columns().Remark].String(),
		},
		MedicalHistories: medicalHistories,
		Vaccinations:     vaccinations,
		Allergies:        allergies,
	}, nil
}

func (s doctorAppointmentService) UpdateStatus(ctx context.Context, in DoctorAppointmentUpdateStatusInput) error {
	record, err := dao.Appointment.Ctx(ctx).
		Where(dao.Appointment.Columns().Id, in.AppointmentID).
		Where(dao.Appointment.Columns().DoctorId, in.DoctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "查询预约失败")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("预约不存在")
	}

	currentStatus := record[dao.Appointment.Columns().Status].Int()
	if currentStatus == in.Status {
		return nil
	}
	if currentStatus == 3 {
		return consts.NewConflictError("预约已取消，无法更新状态")
	}
	if currentStatus == 2 {
		return consts.NewConflictError("预约已完成，无法更新状态")
	}
	if currentStatus == 4 {
		return consts.NewConflictError("预约已过期，无法更新状态")
	}

	result, err := dao.Appointment.Ctx(ctx).
		Where(dao.Appointment.Columns().Id, in.AppointmentID).
		Where(dao.Appointment.Columns().DoctorId, in.DoctorID).
		Data(do.Appointment{
			Status:    in.Status,
			UpdatedAt: time.Now(),
		}).
		Update()
	if err != nil {
		return consts.WrapInternalError(err, "更新预约状态失败")
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

func loadDoctorAppointmentMedicalHistories(ctx context.Context, petID int64) ([]DoctorAppointmentMedicalHistory, error) {
	records, err := dao.PetMedicalHistory.Ctx(ctx).
		Where(dao.PetMedicalHistory.Columns().PetId, petID).
		OrderDesc(dao.PetMedicalHistory.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询病史记录失败")
	}

	items := make([]DoctorAppointmentMedicalHistory, 0, len(records))
	for _, record := range records {
		diagnosedAt := ""
		if !record[dao.PetMedicalHistory.Columns().DiagnosedAt].IsNil() {
			diagnosedAt = record[dao.PetMedicalHistory.Columns().DiagnosedAt].GTime().Format(appointmentTimeLayout)
		}
		items = append(items, DoctorAppointmentMedicalHistory{
			ID:          record[dao.PetMedicalHistory.Columns().Id].Int64(),
			HistoryType: record[dao.PetMedicalHistory.Columns().HistoryType].String(),
			Description: record[dao.PetMedicalHistory.Columns().Description].String(),
			DiagnosedAt: diagnosedAt,
			IsCurrent:   record[dao.PetMedicalHistory.Columns().IsCurrent].Int(),
		})
	}
	return items, nil
}

func loadDoctorAppointmentVaccinations(ctx context.Context, petID int64) ([]DoctorAppointmentVaccination, error) {
	records, err := dao.PetVaccinationRecord.Ctx(ctx).
		Where(dao.PetVaccinationRecord.Columns().PetId, petID).
		OrderDesc(dao.PetVaccinationRecord.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询疫苗记录失败")
	}

	items := make([]DoctorAppointmentVaccination, 0, len(records))
	for _, record := range records {
		vaccinationDate := ""
		nextDueDate := ""
		if !record[dao.PetVaccinationRecord.Columns().VaccinationDate].IsNil() {
			vaccinationDate = record[dao.PetVaccinationRecord.Columns().VaccinationDate].GTime().Format(doctorAppointmentDateLayout)
		}
		if !record[dao.PetVaccinationRecord.Columns().NextDueDate].IsNil() {
			nextDueDate = record[dao.PetVaccinationRecord.Columns().NextDueDate].GTime().Format(doctorAppointmentDateLayout)
		}
		items = append(items, DoctorAppointmentVaccination{
			ID:              record[dao.PetVaccinationRecord.Columns().Id].Int64(),
			VaccineName:     record[dao.PetVaccinationRecord.Columns().VaccineName].String(),
			VaccinationDate: vaccinationDate,
			NextDueDate:     nextDueDate,
		})
	}
	return items, nil
}

func loadDoctorAppointmentAllergies(ctx context.Context, petID int64) ([]DoctorAppointmentAllergy, error) {
	records, err := dao.PetAllergyRecord.Ctx(ctx).
		Where(dao.PetAllergyRecord.Columns().PetId, petID).
		OrderDesc(dao.PetAllergyRecord.Columns().Id).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询过敏记录失败")
	}

	items := make([]DoctorAppointmentAllergy, 0, len(records))
	for _, record := range records {
		items = append(items, DoctorAppointmentAllergy{
			ID:                 record[dao.PetAllergyRecord.Columns().Id].Int64(),
			Allergen:           record[dao.PetAllergyRecord.Columns().Allergen].String(),
			SymptomDescription: record[dao.PetAllergyRecord.Columns().SymptomDescription].String(),
			SeverityLevel:      record[dao.PetAllergyRecord.Columns().SeverityLevel].Int(),
		})
	}
	return items, nil
}

func formatDoctorAppointmentWeight(weight float64) string {
	return strconv.FormatFloat(weight, 'f', 2, 64)
}
