package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"PetCare/internal/consts"
	"PetCare/internal/dao"

	"github.com/gogf/gf/v2/os/gtime"
)

func generateAISessionNo(ctx context.Context, now time.Time) (string, error) {
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	count, err := dao.AiSession.Ctx(ctx).
		WhereGTE(dao.AiSession.Columns().CreatedAt, startOfDay).
		WhereLT(dao.AiSession.Columns().CreatedAt, endOfDay).
		Count()
	if err != nil {
		return "", consts.WrapInternalError(err, "generate AI session no failed")
	}

	return fmt.Sprintf("AIS%s%04d", now.Format("20060102"), count+1), nil
}

func aiEnsureUserOwnedPet(ctx context.Context, petID int64, userID int64) error {
	count, err := dao.Pet.Ctx(ctx).
		Where(dao.Pet.Columns().Id, petID).
		Where(dao.Pet.Columns().UserId, userID).
		Count()
	if err != nil {
		return consts.WrapInternalError(err, "query pet failed")
	}
	if count == 0 {
		return consts.NewNotFoundError("pet not found")
	}
	return nil
}

func aiEnsureActiveHospital(ctx context.Context, hospitalID int64) error {
	record, err := dao.Hospital.Ctx(ctx).
		Where(dao.Hospital.Columns().Id, hospitalID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "query hospital failed")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("hospital not found")
	}
	if record[dao.Hospital.Columns().Status].Int() != 1 {
		return consts.NewConflictError("hospital is disabled")
	}
	return nil
}

func aiEnsureActiveDoctor(ctx context.Context, doctorID int64, hospitalID *int64) error {
	record, err := dao.Doctor.Ctx(ctx).
		Where(dao.Doctor.Columns().Id, doctorID).
		One()
	if err != nil {
		return consts.WrapInternalError(err, "query doctor failed")
	}
	if record.IsEmpty() {
		return consts.NewNotFoundError("doctor not found")
	}
	if record[dao.Doctor.Columns().Status].Int() != 1 {
		return consts.NewConflictError("doctor is disabled")
	}
	if hospitalID != nil && record[dao.Doctor.Columns().HospitalId].Int64() != *hospitalID {
		return consts.NewConflictError("doctor does not belong to hospital")
	}
	return nil
}

func aiSenderTypeFromRole(role string) (int, error) {
	switch NormalizeRole(role) {
	case consts.RoleUser:
		return aiSenderTypeUser, nil
	case consts.RoleDoctor:
		return aiSenderTypeDoctor, nil
	default:
		return 0, consts.NewForbiddenError("")
	}
}

func aiLoadAccessibleSession(ctx context.Context, sessionID int64, requesterID int64, requesterRole string) (*aiSessionData, error) {
	record, err := dao.AiSession.Ctx(ctx).
		Where(dao.AiSession.Columns().Id, sessionID).
		One()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query ai session failed")
	}
	if record.IsEmpty() {
		return nil, consts.NewNotFoundError("session not found")
	}

	session := aiSessionDataFromRecord(record)
	switch NormalizeRole(requesterRole) {
	case consts.RoleUser:
		if session.UserID != requesterID {
			return nil, consts.NewForbiddenError("")
		}
	case consts.RoleDoctor:
		if session.DoctorID <= 0 || session.DoctorID != requesterID {
			return nil, consts.NewForbiddenError("")
		}
	default:
		return nil, consts.NewForbiddenError("")
	}
	return &session, nil
}

func aiLoadRecentSessionMessages(ctx context.Context, sessionID int64, limit int) ([]AIChatMessage, error) {
	records, err := dao.AiMessage.Ctx(ctx).
		Where(dao.AiMessage.Columns().SessionId, sessionID).
		OrderDesc(dao.AiMessage.Columns().Id).
		Limit(limit).
		All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query ai message history failed")
	}

	messages := make([]AIChatMessage, 0, len(records))
	for i := len(records) - 1; i >= 0; i-- {
		record := records[i]
		role := "user"
		if record[dao.AiMessage.Columns().SenderType].Int() == aiSenderTypeAI {
			role = "assistant"
		}
		messages = append(messages, AIChatMessage{
			Role:    role,
			Content: record[dao.AiMessage.Columns().MessageContent].String(),
		})
	}
	return messages, nil
}

func aiBuildSessionSummary(content string) string {
	summary := strings.TrimSpace(content)
	if summary == "" {
		return ""
	}
	if len(summary) > 120 {
		return summary[:120]
	}
	return summary
}

func aiBuildAnalysisResult(content string) string {
	result := strings.TrimSpace(content)
	if result == "" {
		return "No analysis generated"
	}
	return result
}

func aiInferRiskLevel(content string) int {
	lower := strings.ToLower(content)
	if strings.Contains(lower, "emergency") || strings.Contains(lower, "urgent") || strings.Contains(content, "尽快就医") || strings.Contains(content, "立即就医") {
		return 3
	}
	if strings.Contains(lower, "monitor") || strings.Contains(lower, "observe") || strings.Contains(content, "建议观察") || strings.Contains(content, "可能") {
		return 2
	}
	return 1
}

func formatGTime(value *gtime.Time) string {
	if value == nil || value.IsZero() {
		return ""
	}
	return value.Format(aiTimeLayout)
}
