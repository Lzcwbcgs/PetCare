package service

import (
	"context"

	"PetCare/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	AISessionCreateInput struct {
		RequesterUserID int64
		RequesterRole   string
		PetID           int64
		HospitalID      *int64
		DoctorID        *int64
		ModelType       string
		ModelName       string
		RagEnabled      *int
	}

	AISessionCreateOutput struct {
		SessionID int64
		SessionNo string
		Status    int
	}

	AISendMessageInput struct {
		SessionID      int64
		SenderUserID   int64
		SenderRole     string
		MessageContent string
		MessageType    int
	}

	AIStreamEvent struct {
		Type    string
		Payload any
	}

	AISessionDetailInput struct {
		SessionID       int64
		RequesterUserID int64
		RequesterRole   string
	}

	AISessionDetailOutput struct {
		ID             int64
		SessionNo      string
		PetID          int64
		SourceType     int
		ModelType      string
		ModelName      string
		ProviderName   string
		SessionSummary string
		RagEnabled     int
		Status         int
		LastMessageAt  string
		CreatedAt      string
		UpdatedAt      string
	}

	AIDoctorSessionListInput struct {
		DoctorID int64
		Page     int
		Size     int
		PetID    *int64
	}

	AIDoctorSessionItem struct {
		ID        int64
		SessionNo string
		PetID     int64
		PetName   string
		ModelName string
		Status    int
		CreatedAt string
	}

	AIDoctorSessionListOutput struct {
		Items []AIDoctorSessionItem
		Total int
		Page  int
		Size  int
	}

	AIMessageListInput struct {
		SessionID       int64
		RequesterUserID int64
		RequesterRole   string
		Page            int
		Size            int
	}

	AIMessageItem struct {
		ID             int64
		SenderType     int
		SenderID       *int64
		MessageContent string
		MessageType    int
		ProviderType   string
		ProviderName   string
		FinishReason   string
		CreatedAt      string
	}

	AIMessageListOutput struct {
		Items []AIMessageItem
		Total int
		Page  int
		Size  int
	}

	AIAnalysisListInput struct {
		SessionID       int64
		RequesterUserID int64
		RequesterRole   string
	}

	AIAnalysisItem struct {
		ID               int64
		AnalysisType     int
		InputSource      int
		SummaryTitle     string
		AnalysisResult   string
		RuleBasedResult  string
		LlmBasedResult   string
		RiskLevel        *int
		ReviewedByDoctor int
		CreatedAt        string
	}

	AIAnalysisListOutput struct {
		Items []AIAnalysisItem
	}
)

type IAI interface {
	CreateSession(ctx context.Context, in AISessionCreateInput) (*AISessionCreateOutput, error)
	SendMessageStream(ctx context.Context, in AISendMessageInput, emit func(event AIStreamEvent)) error
	Detail(ctx context.Context, in AISessionDetailInput) (*AISessionDetailOutput, error)
	ListDoctorSessions(ctx context.Context, in AIDoctorSessionListInput) (*AIDoctorSessionListOutput, error)
	ListMessages(ctx context.Context, in AIMessageListInput) (*AIMessageListOutput, error)
	ListAnalysisRecords(ctx context.Context, in AIAnalysisListInput) (*AIAnalysisListOutput, error)
}

var AI IAI = aiService{}

type aiService struct{}

type aiSessionData struct {
	ID             int64
	SessionNo      string
	UserID         int64
	PetID          int64
	DoctorID       int64
	SourceType     int
	ModelType      string
	ModelName      string
	ProviderName   string
	SessionSummary string
	RagEnabled     int
	Status         int
	LastMessageAt  *gtime.Time
	CreatedAt      *gtime.Time
	UpdatedAt      *gtime.Time
}

func aiSessionDataFromRecord(record gdb.Record) aiSessionData {
	cols := dao.AiSession.Columns()
	return aiSessionData{
		ID:             record[cols.Id].Int64(),
		SessionNo:      record[cols.SessionNo].String(),
		UserID:         record[cols.UserId].Int64(),
		PetID:          record[cols.PetId].Int64(),
		DoctorID:       record[cols.DoctorId].Int64(),
		SourceType:     record[cols.SourceType].Int(),
		ModelType:      record[cols.ModelType].String(),
		ModelName:      record[cols.ModelName].String(),
		ProviderName:   record[aiSessionColumnProviderName].String(),
		SessionSummary: record[cols.SessionSummary].String(),
		RagEnabled:     record[aiSessionColumnRagEnabled].Int(),
		Status:         record[cols.Status].Int(),
		LastMessageAt:  record[aiSessionColumnLastMessageAt].GTime(),
		CreatedAt:      record[cols.CreatedAt].GTime(),
		UpdatedAt:      record[cols.UpdatedAt].GTime(),
	}
}
