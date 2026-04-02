package admin

import (
	"context"

	v1 "PetCare/api/admin/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type DoctorController struct{}

func NewDoctor() *DoctorController {
	return &DoctorController{}
}

func (c *DoctorController) Create(ctx context.Context, req *v1.DoctorCreateReq) (res *v1.DoctorCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	output, err := service.AdminDoctor.Create(ctx, service.AdminDoctorCreateInput{
		HospitalID: req.HospitalID,
		Username:   req.Username,
		Password:   req.Password,
		DoctorName: req.DoctorName,
		Gender:     req.Gender,
		Phone:      req.Phone,
		Email:      req.Email,
		Title:      req.Title,
		Specialty:  req.Specialty,
		AvatarURL:  req.AvatarURL,
		Intro:      req.Intro,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DoctorCreateRes{DoctorID: output.ID}, nil
}

func (c *DoctorController) List(ctx context.Context, req *v1.DoctorListReq) (res *v1.DoctorListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	var (
		page int
		size int
	)
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		size = *req.PageSize
	}

	output, err := service.AdminDoctor.List(ctx, service.AdminDoctorListInput{
		Page:       page,
		Size:       size,
		HospitalID: req.HospitalID,
		Status:     req.Status,
		Keyword:    req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.DoctorListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.DoctorListItem{
			ID:           item.ID,
			HospitalID:   item.HospitalID,
			HospitalName: item.HospitalName,
			Username:     item.Username,
			DoctorName:   item.DoctorName,
			Title:        item.Title,
			Specialty:    item.Specialty,
			Phone:        item.Phone,
			Status:       item.Status,
			CreatedAt:    item.CreatedAt,
		})
	}

	return &v1.DoctorListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *DoctorController) Detail(ctx context.Context, req *v1.DoctorDetailReq) (res *v1.DoctorDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	item, err := service.AdminDoctor.Detail(ctx, service.AdminDoctorDetailInput{
		DoctorID: req.DoctorID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.DoctorDetailRes{
		ID:         item.ID,
		HospitalID: item.HospitalID,
		Username:   item.Username,
		DoctorName: item.DoctorName,
		Gender:     item.Gender,
		Phone:      item.Phone,
		Email:      item.Email,
		Title:      item.Title,
		Specialty:  item.Specialty,
		AvatarURL:  item.AvatarURL,
		Intro:      item.Intro,
		Status:     item.Status,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}, nil
}

func (c *DoctorController) Update(ctx context.Context, req *v1.DoctorUpdateReq) (res *v1.DoctorUpdateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "修改成功")

	err = service.AdminDoctor.Update(ctx, service.AdminDoctorUpdateInput{
		DoctorID:   req.DoctorID,
		HospitalID: req.HospitalID,
		DoctorName: req.DoctorName,
		Gender:     req.Gender,
		Phone:      req.Phone,
		Email:      req.Email,
		Title:      req.Title,
		Specialty:  req.Specialty,
		AvatarURL:  req.AvatarURL,
		Intro:      req.Intro,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DoctorUpdateRes{}, nil
}

func (c *DoctorController) Delete(ctx context.Context, req *v1.DoctorDeleteReq) (res *v1.DoctorDeleteRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "删除成功")

	err = service.AdminDoctor.Delete(ctx, service.AdminDoctorDeleteInput{
		DoctorID: req.DoctorID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DoctorDeleteRes{}, nil
}
