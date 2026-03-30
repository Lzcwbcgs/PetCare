package pet

import (
	"context"

	v1 "PetCare/api/pet/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *Controller) AllergyCreate(ctx context.Context, req *v1.AllergyCreateReq) (res *v1.AllergyCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	claims, err := requirePetRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	output, err := service.PetAllergy.Create(ctx, service.AllergyCreateInput{
		UserID:             claims.UserID,
		PetID:              req.PetID,
		Allergen:           req.Allergen,
		SymptomDescription: req.SymptomDescription,
		SeverityLevel:      req.SeverityLevel,
		Remark:             req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AllergyCreateRes{ID: output.ID}, nil
}

func (c *Controller) AllergyList(ctx context.Context, req *v1.AllergyListReq) (res *v1.AllergyListRes, err error) {
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

	output, err := service.PetAllergy.List(ctx, service.AllergyListInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		PetID:           req.PetID,
		Page:            page,
		Size:            size,
		SeverityLevel:   req.SeverityLevel,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.AllergyItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.AllergyItem{
			ID:                 item.ID,
			Allergen:           item.Allergen,
			SymptomDescription: item.SymptomDescription,
			SeverityLevel:      item.SeverityLevel,
			Remark:             item.Remark,
		})
	}
	return &v1.AllergyListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}
