package common

import (
	"context"
	"strings"

	v1 "PetCare/api/common/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) HospitalOptions(ctx context.Context, req *v1.HospitalOptionReq) (res *v1.HospitalOptionRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	items, err := service.Common.HospitalOptions(ctx)
	if err != nil {
		return nil, err
	}

	list := make(v1.HospitalOptionRes, 0, len(items))
	for _, item := range items {
		list = append(list, v1.HospitalOptionItem{
			ID:           item.ID,
			HospitalName: item.HospitalName,
		})
	}
	return &list, nil
}

func (c *Controller) HospitalDoctorOptions(ctx context.Context, req *v1.HospitalDoctorOptionReq) (res *v1.HospitalDoctorOptionRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	items, err := service.Common.HospitalDoctorOptions(ctx, service.CommonHospitalDoctorOptionsInput{
		HospitalID: req.HospitalID,
	})
	if err != nil {
		return nil, err
	}

	list := make(v1.HospitalDoctorOptionRes, 0, len(items))
	for _, item := range items {
		list = append(list, v1.HospitalDoctorOptionItem{
			ID:         item.ID,
			DoctorName: item.DoctorName,
			Title:      item.Title,
		})
	}
	return &list, nil
}

func (c *Controller) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "上传成功")

	output, err := service.Common.Upload(ctx, service.CommonUploadInput{
		File:    req.File,
		BizType: req.BizType,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UploadRes{
		FileURL:  buildCommonAbsoluteURL(ctx, output.FileURL),
		FileName: output.FileName,
	}, nil
}

func buildCommonAbsoluteURL(ctx context.Context, fileURL string) string {
	if strings.TrimSpace(fileURL) == "" {
		return ""
	}
	if strings.HasPrefix(fileURL, "http://") || strings.HasPrefix(fileURL, "https://") {
		return fileURL
	}
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return fileURL
	}
	return strings.TrimRight(r.GetSchema()+"://"+r.Host, "/") + "/" + strings.TrimLeft(fileURL, "/")
}
