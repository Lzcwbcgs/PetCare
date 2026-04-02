package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"
	"PetCare/internal/model/do"
)

func (s aiService) CreateSession(ctx context.Context, in AISessionCreateInput) (*AISessionCreateOutput, error) {
	if !RoleAllowed(in.RequesterRole, consts.RoleUser) {
		return nil, consts.NewForbiddenError("")
	}
	if err := aiEnsureUserOwnedPet(ctx, in.PetID, in.RequesterUserID); err != nil {
		return nil, err
	}
	if in.HospitalID != nil {
		if err := aiEnsureActiveHospital(ctx, *in.HospitalID); err != nil {
			return nil, err
		}
	}
	if in.DoctorID != nil {
		if err := aiEnsureActiveDoctor(ctx, *in.DoctorID, in.HospitalID); err != nil {
			return nil, err
		}
	}

	providerType := aiNormalizeProviderType(in.ModelType)
	if providerType == "" {
		providerType = aiDefaultProvider(ctx)
	}
	modelName := strings.TrimSpace(in.ModelName)
	if modelName == "" {
		modelName = aiDefaultChatModel(ctx, providerType)
	}
	if modelName == "" {
		return nil, consts.NewBadRequestError("model_name is required")
	}

	ragEnabled := aiRagEnabledDefault(ctx)
	if in.RagEnabled != nil {
		ragEnabled = *in.RagEnabled
	}

	now := time.Now()
	sessionNo, err := generateAISessionNo(ctx, now)
	if err != nil {
		return nil, err
	}

	var (
		hospitalID any
		doctorID   any
	)
	if in.HospitalID != nil {
		hospitalID = *in.HospitalID
	}
	if in.DoctorID != nil {
		doctorID = *in.DoctorID
	}

	insertResult, err := dao.AiSession.Ctx(ctx).Data(do.AiSession{
		SessionNo:      sessionNo,
		UserId:         in.RequesterUserID,
		PetId:          in.PetID,
		HospitalId:     hospitalID,
		DoctorId:       doctorID,
		SourceType:     1,
		ModelType:      providerType,
		ModelName:      modelName,
		ProviderName:   providerType,
		RagEnabled:     ragEnabled,
		RetrievalCount: 0,
		SyncToAdmin:    1,
		Status:         aiSessionStatusActive,
		CreatedAt:      now,
		UpdatedAt:      now,
	}).Insert()
	if err != nil {
		return nil, consts.WrapInternalError(err, "create AI session failed")
	}

	sessionID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, consts.WrapInternalError(err, "read session id failed")
	}

	RecordOperationLogByRole(
		ctx,
		in.RequesterRole,
		in.RequesterUserID,
		"ai_session",
		"create",
		fmt.Sprintf("create ai session id=%d", sessionID),
	)

	return &AISessionCreateOutput{
		SessionID: sessionID,
		SessionNo: sessionNo,
		Status:    aiSessionStatusActive,
	}, nil
}

func (s aiService) Detail(ctx context.Context, in AISessionDetailInput) (*AISessionDetailOutput, error) {
	session, err := aiLoadAccessibleSession(ctx, in.SessionID, in.RequesterUserID, in.RequesterRole)
	if err != nil {
		return nil, err
	}
	RecordOperationLogByRole(
		ctx,
		in.RequesterRole,
		in.RequesterUserID,
		"ai_session",
		"detail",
		fmt.Sprintf("query ai session detail id=%d", in.SessionID),
	)

	return &AISessionDetailOutput{
		ID:             session.ID,
		SessionNo:      session.SessionNo,
		PetID:          session.PetID,
		SourceType:     session.SourceType,
		ModelType:      session.ModelType,
		ModelName:      session.ModelName,
		ProviderName:   session.ProviderName,
		SessionSummary: session.SessionSummary,
		RagEnabled:     session.RagEnabled,
		Status:         session.Status,
		LastMessageAt:  formatGTime(session.LastMessageAt),
		CreatedAt:      formatGTime(session.CreatedAt),
		UpdatedAt:      formatGTime(session.UpdatedAt),
	}, nil
}

func (s aiService) ListDoctorSessions(ctx context.Context, in AIDoctorSessionListInput) (*AIDoctorSessionListOutput, error) {
	page := in.Page
	size := in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	model := dao.AiSession.Ctx(ctx).Where(dao.AiSession.Columns().DoctorId, in.DoctorID)
	if in.PetID != nil && *in.PetID > 0 {
		model = model.Where(dao.AiSession.Columns().PetId, *in.PetID)
	}

	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query ai sessions failed")
	}

	records, err := model.Page(page, size).OrderDesc(dao.AiSession.Columns().Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query ai sessions failed")
	}

	petIDs := make([]int64, 0, len(records))
	for _, record := range records {
		petID := record[dao.AiSession.Columns().PetId].Int64()
		if petID > 0 {
			petIDs = append(petIDs, petID)
		}
	}
	petNameMap, err := loadPetNameMap(ctx, petIDs)
	if err != nil {
		return nil, err
	}

	items := make([]AIDoctorSessionItem, 0, len(records))
	for _, record := range records {
		petID := record[dao.AiSession.Columns().PetId].Int64()
		items = append(items, AIDoctorSessionItem{
			ID:        record[dao.AiSession.Columns().Id].Int64(),
			SessionNo: record[dao.AiSession.Columns().SessionNo].String(),
			PetID:     petID,
			PetName:   petNameMap[petID],
			ModelName: record[dao.AiSession.Columns().ModelName].String(),
			Status:    record[dao.AiSession.Columns().Status].Int(),
			CreatedAt: formatGTime(record[dao.AiSession.Columns().CreatedAt].GTime()),
		})
	}

	RecordOperationLog(ctx, OperationLogRecordInput{
		OperatorType:    operatorTypeDoctor,
		OperatorID:      in.DoctorID,
		OperationModule: "doctor_ai_view",
		OperationType:   "list_sessions",
		OperationDesc:   "doctor query ai sessions",
	})

	return &AIDoctorSessionListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}
