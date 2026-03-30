package pet

import (
	"context"
	"strconv"
	"strings"

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

	claims, err := requirePetRoles(ctx, consts.RoleUser)
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
			ID:      item.ID,
			PetName: item.PetName,
			PetType: item.PetType,
			Gender:  item.Gender,
			Age:     item.Age,
			AgeUnit: item.AgeUnit,
			Breed:   item.Breed,
			Weight:  formatPetWeight(item.Weight),
			Status:  item.Status,
		})
	}

	return &v1.PetListRes{
		List: items,
		Pagination: v1.Pagination{
			Page:     output.Page,
			PageSize: output.Size,
			Total:    output.Total,
		},
	}, nil
}

func (c *Controller) Detail(ctx context.Context, req *v1.PetDetailReq) (res *v1.PetDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requirePetRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	item, err := service.Pet.Detail(ctx, service.PetDetailInput{
		UserID: claims.UserID,
		PetID:  req.PetID,
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
		Weight:     formatPetWeight(item.Weight),
		Sterilized: item.Sterilized,
		Remark:     item.Remark,
		Status:     item.Status,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}, nil
}

func (c *Controller) Create(ctx context.Context, req *v1.PetCreateReq) (res *v1.PetCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "新增成功")

	claims, err := requirePetRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	weight, err := parsePetWeight(req.Weight)
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
		Weight:     weight,
		Sterilized: req.Sterilized,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &v1.PetCreateRes{PetID: output.ID}, nil
}

func (c *Controller) Update(ctx context.Context, req *v1.PetUpdateReq) (res *v1.PetUpdateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "修改成功")

	claims, err := requirePetRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	weight, err := parsePetWeight(req.Weight)
	if err != nil {
		return nil, err
	}

	err = service.Pet.Update(ctx, service.PetUpdateInput{
		UserID:     claims.UserID,
		PetID:      req.PetID,
		PetName:    req.PetName,
		PetType:    req.PetType,
		AvatarURL:  req.AvatarURL,
		Gender:     req.Gender,
		Age:        req.Age,
		AgeUnit:    req.AgeUnit,
		Breed:      req.Breed,
		Weight:     weight,
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

	claims, err := requirePetRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	err = service.Pet.Delete(ctx, service.PetDeleteInput{
		UserID: claims.UserID,
		PetID:  req.PetID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PetDeleteRes{}, nil
}

func parsePetWeight(weight *string) (*float64, error) {
	if weight == nil {
		return nil, nil
	}

	value := strings.TrimSpace(*weight)
	if value == "" {
		return nil, consts.NewBadRequestError("体重格式不正确")
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, consts.NewBadRequestError("体重格式不正确")
	}
	if parsed < 0 || parsed > 999.99 {
		return nil, consts.NewBadRequestError("体重不能小于0且不能超过999.99kg")
	}

	return &parsed, nil
}

func formatPetWeight(weight float64) string {
	return strconv.FormatFloat(weight, 'f', 2, 64)
}
