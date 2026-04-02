package pet

import (
	"context"

	v1 "PetCare/api/pet/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *Controller) VaccinationCreate(ctx context.Context, req *v1.VaccinationCreateReq) (res *v1.VaccinationCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	claims, err := requirePetRoles(ctx, consts.RoleUser)
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

	claims, err := requirePetRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

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

	output, err := service.PetVaccination.List(ctx, service.VaccinationListInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		PetID:           req.PetID,
		Page:            page,
		Size:            size,
		VaccineName:     req.VaccineName,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.VaccinationItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.VaccinationItem{
			ID:              item.ID,
			VaccineName:     item.VaccineName,
			VaccinationDate: item.VaccinationDate,
			NextDueDate:     item.NextDueDate,
			HospitalName:    item.HospitalName,
		})
	}
	return &v1.VaccinationListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}
