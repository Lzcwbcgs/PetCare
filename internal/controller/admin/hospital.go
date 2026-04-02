package admin

import (
	"context"

	v1 "PetCare/api/admin/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type HospitalController struct{}

func NewHospital() *HospitalController {
	return &HospitalController{}
}

func (c *HospitalController) Create(ctx context.Context, req *v1.HospitalCreateReq) (res *v1.HospitalCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	output, err := service.AdminHospital.Create(ctx, service.AdminHospitalCreateInput{
		HospitalName: req.HospitalName,
		Address:      req.Address,
		Phone:        req.Phone,
		Description:  req.Description,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.HospitalCreateRes{HospitalID: output.ID}, nil
}

func (c *HospitalController) List(ctx context.Context, req *v1.HospitalListReq) (res *v1.HospitalListRes, err error) {
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

	output, err := service.AdminHospital.List(ctx, service.AdminHospitalListInput{
		Page:    page,
		Size:    size,
		Status:  req.Status,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.HospitalListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.HospitalListItem{
			ID:           item.ID,
			HospitalName: item.HospitalName,
			Address:      item.Address,
			Phone:        item.Phone,
			Status:       item.Status,
			CreatedAt:    item.CreatedAt,
		})
	}

	return &v1.HospitalListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *HospitalController) Detail(ctx context.Context, req *v1.HospitalDetailReq) (res *v1.HospitalDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	item, err := service.AdminHospital.Detail(ctx, service.AdminHospitalDetailInput{
		HospitalID: req.HospitalID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.HospitalDetailRes{
		ID:           item.ID,
		HospitalName: item.HospitalName,
		Address:      item.Address,
		Phone:        item.Phone,
		Description:  item.Description,
		Status:       item.Status,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}, nil
}

func (c *HospitalController) Update(ctx context.Context, req *v1.HospitalUpdateReq) (res *v1.HospitalUpdateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "修改成功")

	err = service.AdminHospital.Update(ctx, service.AdminHospitalUpdateInput{
		HospitalID:   req.HospitalID,
		HospitalName: req.HospitalName,
		Address:      req.Address,
		Phone:        req.Phone,
		Description:  req.Description,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.HospitalUpdateRes{}, nil
}

func (c *HospitalController) Delete(ctx context.Context, req *v1.HospitalDeleteReq) (res *v1.HospitalDeleteRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "删除成功")

	err = service.AdminHospital.Delete(ctx, service.AdminHospitalDeleteInput{
		HospitalID: req.HospitalID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.HospitalDeleteRes{}, nil
}
