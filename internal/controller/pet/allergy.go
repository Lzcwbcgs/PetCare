package pet

import (
	"context"

	v1 "PetCare/api/pet/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *Controller) AllergyCreate(ctx context.Context, req *v1.AllergyCreateReq) (res *v1.AllergyCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "创建成功")

	claims, err := authClaimsFromCtx(ctx)
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

	output, err := service.PetAllergy.List(ctx, service.AllergyListInput{
		UserID:        claims.UserID,
		PetID:         req.PetID,
		Page:          page,
		Size:          size,
		SeverityLevel: req.SeverityLevel,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.AllergyItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.AllergyItem{
			ID:                 item.ID,
			PetID:              item.PetID,
			Allergen:           item.Allergen,
			SymptomDescription: item.SymptomDescription,
			SeverityLevel:      item.SeverityLevel,
			Remark:             item.Remark,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		})
	}
	return &v1.AllergyListRes{
		List:  items,
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
	}, nil
}
