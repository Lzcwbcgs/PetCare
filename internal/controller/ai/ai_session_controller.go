package ai

import (
	"context"

	v1 "PetCare/api/ai/v1"
	"PetCare/internal/consts"
	"PetCare/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *Controller) CreateSession(ctx context.Context, req *v1.SessionCreateReq) (res *v1.SessionCreateRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "Session created successfully")

	claims, err := requireAIRoles(ctx, consts.RoleUser)
	if err != nil {
		return nil, err
	}

	output, err := service.AI.CreateSession(ctx, service.AISessionCreateInput{
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
		PetID:           req.PetID,
		HospitalID:      req.HospitalID,
		DoctorID:        req.DoctorID,
		ModelType:       req.ModelType,
		ModelName:       req.ModelName,
	})
	if err != nil {
		return nil, err
	}

	return &v1.SessionCreateRes{
		SessionID: output.SessionID,
		SessionNo: output.SessionNo,
		Status:    output.Status,
	}, nil
}

func (c *Controller) Detail(ctx context.Context, req *v1.SessionDetailReq) (res *v1.SessionDetailRes, err error) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.CtxKeyResponseMessage, "success")

	claims, err := requireAIRoles(ctx, consts.RoleUser, consts.RoleDoctor)
	if err != nil {
		return nil, err
	}

	output, err := service.AI.Detail(ctx, service.AISessionDetailInput{
		SessionID:       req.SessionID,
		RequesterUserID: claims.UserID,
		RequesterRole:   claims.Role,
	})
	if err != nil {
		return nil, err
	}

	return &v1.SessionDetailRes{
		ID:             output.ID,
		SessionNo:      output.SessionNo,
		PetID:          output.PetID,
		SourceType:     output.SourceType,
		ModelType:      output.ModelType,
		ModelName:      output.ModelName,
		SessionSummary: output.SessionSummary,
		Status:         output.Status,
		CreatedAt:      output.CreatedAt,
		UpdatedAt:      output.UpdatedAt,
	}, nil
}
