package admin

import (
	"context"

	v1 "PetCare/api/admin/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type KnowledgeController struct{}

func NewKnowledge() *KnowledgeController {
	return &KnowledgeController{}
}

func (c *KnowledgeController) Upload(ctx context.Context, req *v1.KnowledgeUploadReq) (res *v1.KnowledgeUploadRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "上传成功，已进入向量化队列")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.AdminKnowledge.Upload(ctx, service.AdminKnowledgeUploadInput{
		OperatorID: claims.UserID,
		File:       req.File,
		Category:   req.Category,
		Title:      req.Title,
	})
	if err != nil {
		return nil, err
	}

	return &v1.KnowledgeUploadRes{
		KnowledgeID: output.KnowledgeID,
		FileName:    output.FileName,
		Status:      output.Status,
	}, nil
}

func (c *KnowledgeController) Status(ctx context.Context, req *v1.KnowledgeStatusReq) (res *v1.KnowledgeStatusRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := authClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	output, err := service.AdminKnowledge.Status(ctx, service.AdminKnowledgeStatusInput{
		OperatorID:  claims.UserID,
		KnowledgeID: req.KnowledgeID,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.KnowledgeStatusItem, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, v1.KnowledgeStatusItem{
			KnowledgeID:    item.KnowledgeID,
			FileName:       item.FileName,
			Status:         item.Status,
			Progress:       item.Progress,
			ChunkTotal:     item.ChunkTotal,
			EmbeddedChunks: item.EmbeddedChunks,
			VectorCount:    item.VectorCount,
			ErrorMessage:   item.ErrorMessage,
			UpdatedAt:      item.UpdatedAt,
		})
	}
	return &v1.KnowledgeStatusRes{List: items}, nil
}
