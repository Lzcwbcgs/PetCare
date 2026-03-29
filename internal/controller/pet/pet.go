package pet

import (
	"context"

	v1 "PetCare/api/pet/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) List(ctx context.Context, req *v1.PetListReq) (res *v1.PetListRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
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
	if req.Size != nil {
		size = *req.Size
	}

	output, err := service.Pet.List(ctx, service.PetListInput{
		UserID:  claims.UserID,
		Page:    page,
		Size:    size,
		PetName: req.PetName,
		PetType: req.PetType,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.PetListItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.PetListItem{
			ID:         item.ID,
			UserID:     item.UserID,
			PetName:    item.PetName,
			PetType:    item.PetType,
			AvatarURL:  item.AvatarURL,
			Gender:     item.Gender,
			Age:        item.Age,
			AgeUnit:    item.AgeUnit,
			Breed:      item.Breed,
			Weight:     item.Weight,
			Sterilized: item.Sterilized,
			Remark:     item.Remark,
			Status:     item.Status,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}

	return &v1.PetListRes{
		List:  items,
		Total: output.Total,
		Page:  output.Page,
		Size:  output.Size,
	}, nil
}

func (c *Controller) Detail(ctx context.Context, req *v1.PetDetailReq) (res *v1.PetDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	item, err := service.Pet.Detail(ctx, service.PetDetailInput{
		UserID: claims.UserID,
		PetID:  req.ID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.PetDetailRes{
		ID:         item.ID,
		UserID:     item.UserID,
		PetName:    item.PetName,
		PetType:    item.PetType,
		AvatarURL:  item.AvatarURL,
		Gender:     item.Gender,
		Age:        item.Age,
		AgeUnit:    item.AgeUnit,
		Breed:      item.Breed,
		Weight:     item.Weight,
		Sterilized: item.Sterilized,
		Remark:     item.Remark,
		Status:     item.Status,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}, nil
}

func (c *Controller) Create(ctx context.Context, req *v1.PetCreateReq) (res *v1.PetCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "创建成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.Pet.Create(ctx, service.PetCreateInput{
		UserID:     claims.UserID,
		PetName:    req.PetName,
		PetType:    req.PetType,
		AvatarURL:  req.AvatarURL,
		Gender:     req.Gender,
		Age:        req.Age,
		AgeUnit:    req.AgeUnit,
		Breed:      req.Breed,
		Weight:     req.Weight,
		Sterilized: req.Sterilized,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &v1.PetCreateRes{ID: output.ID}, nil
}

func (c *Controller) Update(ctx context.Context, req *v1.PetUpdateReq) (res *v1.PetUpdateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "更新成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Pet.Update(ctx, service.PetUpdateInput{
		UserID:     claims.UserID,
		PetID:      req.ID,
		PetName:    req.PetName,
		PetType:    req.PetType,
		AvatarURL:  req.AvatarURL,
		Gender:     req.Gender,
		Age:        req.Age,
		AgeUnit:    req.AgeUnit,
		Breed:      req.Breed,
		Weight:     req.Weight,
		Sterilized: req.Sterilized,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PetUpdateRes{}, nil
}

func (c *Controller) Delete(ctx context.Context, req *v1.PetDeleteReq) (res *v1.PetDeleteRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "删除成功")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Pet.Delete(ctx, service.PetDeleteInput{
		UserID: claims.UserID,
		PetID:  req.ID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PetDeleteRes{}, nil
}
