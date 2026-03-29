package pet

import (
	"context"

	v1 "PetCare/api/pet/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *Controller) MedicalHistoryCreate(ctx context.Context, req *v1.MedicalHistoryCreateReq) (res *v1.MedicalHistoryCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "创建成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.PetMedicalHistory.Create(ctx, service.MedicalHistoryCreateInput{
		UserID:      claims.UserID,
		PetID:       req.PetID,
		HistoryType: req.HistoryType,
		Description: req.Description,
		DiagnosedAt: req.DiagnosedAt,
		IsCurrent:   req.IsCurrent,
	})
	if err != nil {
		return nil, err
	}
	return &v1.MedicalHistoryCreateRes{ID: output.ID}, nil
}

func (c *Controller) MedicalHistoryList(ctx context.Context, req *v1.MedicalHistoryListReq) (res *v1.MedicalHistoryListRes, err error) {
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

	output, err := service.PetMedicalHistory.List(ctx, service.MedicalHistoryListInput{
		UserID:    claims.UserID,
		PetID:     req.PetID,
		Page:      page,
		Size:      size,
		IsCurrent: req.IsCurrent,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.MedicalHistoryItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.MedicalHistoryItem{
			ID:          item.ID,
			PetID:       item.PetID,
			HistoryType: item.HistoryType,
			Description: item.Description,
			DiagnosedAt: item.DiagnosedAt,
			IsCurrent:   item.IsCurrent,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return &v1.MedicalHistoryListRes{
		List:  items,
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
	}, nil
}
