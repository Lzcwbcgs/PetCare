package pet

import (
	"context"

	v1 "PetCare/api/pet/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *Controller) VaccinationCreate(ctx context.Context, req *v1.VaccinationCreateReq) (res *v1.VaccinationCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "创建成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.PetVaccination.Create(ctx, service.VaccinationCreateInput{
		UserID:          claims.UserID,
		PetID:           req.PetID,
		VaccineName:     req.VaccineName,
		VaccinationDate: req.VaccinationDate,
		NextDueDate:     req.NextDueDate,
		HospitalName:    req.HospitalName,
		Remark:          req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.VaccinationCreateRes{ID: output.ID}, nil
}

func (c *Controller) VaccinationList(ctx context.Context, req *v1.VaccinationListReq) (res *v1.VaccinationListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	var page, size int
	if req.Page != nil {
		page = *req.Page
	}
	if req.Size != nil {
		size = *req.Size
	}

	output, err := service.PetVaccination.List(ctx, service.VaccinationListInput{
		UserID: claims.UserID,
		PetID:  req.PetID,
		Page:   page,
		Size:   size,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.VaccinationItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.VaccinationItem{
			ID:              item.ID,
			PetID:           item.PetID,
			VaccineName:     item.VaccineName,
			VaccinationDate: item.VaccinationDate,
			NextDueDate:     item.NextDueDate,
			HospitalName:    item.HospitalName,
			Remark:          item.Remark,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}
	return &v1.VaccinationListRes{
		List:  items,
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
	}, nil
}
