package service

import (
	"context"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

const dateLayout = "2006-01-02"

type (
	VaccinationCreateInput struct {
		UserID          int64
		PetID           int64
		VaccineName     string
		VaccinationDate *string
		NextDueDate     *string
		HospitalName    *string
		Remark          *string
	}
	VaccinationListInput struct {
		UserID int64
		PetID  int64
		Page   int
		Size   int
	}
	VaccinationCreateOutput struct {
		ID int64
	}
	VaccinationItem struct {
		ID              int64
		PetID           int64
		VaccineName     string
		VaccinationDate string
		NextDueDate     string
		HospitalName    string
		Remark          string
		CreatedAt       string
		UpdatedAt       string
	}
	VaccinationListOutput struct {
		Items []VaccinationItem
		Total int
		Page  int
		Size  int
	}
)

type IPetVaccination interface {
	Create(ctx context.Context, in VaccinationCreateInput) (*VaccinationCreateOutput, error)
	List(ctx context.Context, in VaccinationListInput) (*VaccinationListOutput, error)
}

var PetVaccination IPetVaccination = petVaccinationService{}

type petVaccinationService struct{}

func (s petVaccinationService) Create(ctx context.Context, in VaccinationCreateInput) (*VaccinationCreateOutput, error) {
	if err := checkPetOwner(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	data := do.PetVaccinationRecord{
		PetId:       in.PetID,
		VaccineName: in.VaccineName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if in.HospitalName != nil {
		data.HospitalName = *in.HospitalName
	}
	if in.Remark != nil {
		data.Remark = *in.Remark
	}
	if in.VaccinationDate != nil && strings.TrimSpace(*in.VaccinationDate) != "" {
		t, err := time.ParseInLocation(dateLayout, *in.VaccinationDate, time.Local)
		if err != nil {
			return nil, consts.NewBadRequestError("接种日期格式不正确")
		}
		data.VaccinationDate = t
	}
	if in.NextDueDate != nil && strings.TrimSpace(*in.NextDueDate) != "" {
		t, err := time.ParseInLocation(dateLayout, *in.NextDueDate, time.Local)
		if err != nil {
			return nil, consts.NewBadRequestError("下次接种日期格式不正确")
		}
		data.NextDueDate = t
	}

	result, err := dao.PetVaccinationRecord.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增疫苗记录失败")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "新增疫苗记录失败")
	}
	return &VaccinationCreateOutput{ID: id}, nil
}

func (s petVaccinationService) List(ctx context.Context, in VaccinationListInput) (*VaccinationListOutput, error) {
	if err := checkPetOwner(ctx, in.PetID, in.UserID); err != nil {
		return nil, err
	}

	page, size := in.Page, in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	if size > 100 {
		size = 100
	}

	cols := dao.PetVaccinationRecord.Columns()
	model := dao.PetVaccinationRecord.Ctx(ctx).Where(cols.PetId, in.PetID)

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询疫苗记录失败")
	}
	records, err := model.Page(page, size).OrderDesc(cols.Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "查询疫苗记录失败")
	}

	items := make([]VaccinationItem, 0, len(records))
	for _, r := range records {
		vacDate, nextDate := "", ""
		if !r[cols.VaccinationDate].IsNil() {
			vacDate = r[cols.VaccinationDate].GTime().Format(dateLayout)
		}
		if !r[cols.NextDueDate].IsNil() {
			nextDate = r[cols.NextDueDate].GTime().Format(dateLayout)
		}
		items = append(items, VaccinationItem{
			ID:              r[cols.Id].Int64(),
			PetID:           r[cols.PetId].Int64(),
			VaccineName:     r[cols.VaccineName].String(),
			VaccinationDate: vacDate,
			NextDueDate:     nextDate,
			HospitalName:    r[cols.HospitalName].String(),
			Remark:          r[cols.Remark].String(),
			CreatedAt:       r[cols.CreatedAt].GTime().Format(petTimeLayout),
			UpdatedAt:       r[cols.UpdatedAt].GTime().Format(petTimeLayout),
		})
	}
	return &VaccinationListOutput{Items: items, Total: total, Page: page, Size: size}, nil
}

func checkPetOwner(ctx context.Context, petID, userID int64) error {
	cols := dao.Pet.Columns()
	count, err := dao.Pet.Ctx(ctx).Where(cols.Id, petID).Where(cols.UserId, userID).Count()
	if err != nil {
		return consts.WrapInternalError(err, "查询宠物信息失败")
	}
	if count == 0 {
		return consts.NewNotFoundError("宠物不存在")
	}
	return nil
}
