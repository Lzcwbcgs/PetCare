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

func (s aiService) SendMessageStream(ctx context.Context, in AISendMessageInput, emit func(event AIStreamEvent)) error {
	messageContent := strings.TrimSpace(in.MessageContent)
	if messageContent == "" {
		return consts.NewBadRequestError("message_content is required")
	}
	if in.MessageType <= 0 {
		in.MessageType = aiMessageTypeText
	}

	session, err := aiLoadAccessibleSession(ctx, in.SessionID, in.SenderUserID, in.SenderRole)
	if err != nil {
		return err
	}

	senderType, err := aiSenderTypeFromRole(in.SenderRole)
	if err != nil {
		return err
	}

	now := time.Now()
	userMessageResult, err := dao.AiMessage.Ctx(ctx).Data(do.AiMessage{
		SessionId:      session.ID,
		SenderType:     senderType,
		SenderId:       in.SenderUserID,
		MessageContent: messageContent,
		MessageType:    in.MessageType,
		CreatedAt:      now,
	}).Insert()
	if err != nil {
		return consts.WrapInternalError(err, "save user message failed")
	}
	userMessageID, err := userMessageResult.LastInsertId()
	if err != nil {
		return consts.WrapInternalError(err, "read user message id failed")
	}

	_, _ = dao.AiSession.Ctx(ctx).
		Where(dao.AiSession.Columns().Id, session.ID).
		Data(do.AiSession{LastMessageAt: now, UpdatedAt: now}).
		Update()

	emit(AIStreamEvent{
		Type: "start",
		Payload: map[string]any{
			"session_id": session.ID,
			"message_id": userMessageID,
		},
	})

	history, err := aiLoadRecentSessionMessages(ctx, session.ID, aiHistoryLimit(ctx))
	if err != nil {
		emit(AIStreamEvent{Type: "error", Payload: map[string]any{"message": "load conversation history failed"}})
		return nil
	}

	providerType := aiNormalizeProviderType(session.ModelType)
	if providerType == "" {
		providerType = aiDefaultProvider(ctx)
	}

	provider, err := aiProviderByType(ctx, providerType)
	if err != nil {
		emit(AIStreamEvent{Type: "error", Payload: map[string]any{"message": err.Error()}})
		return nil
	}

	modelName := strings.TrimSpace(session.ModelName)
	if modelName == "" {
		modelName = aiDefaultChatModel(ctx, providerType)
	}
	if modelName == "" {
		emit(AIStreamEvent{Type: "error", Payload: map[string]any{"message": "chat model is not configured"}})
		return nil
	}

	chatMessages := make([]AIChatMessage, 0, len(history)+1)
	chatMessages = append(chatMessages, AIChatMessage{Role: "system", Content: aiSystemPrompt(ctx)})
	var retrievedChunks []string
	if session.RagEnabled == 1 {
		retrievedChunks, _ = aiRetrieveKnowledgeChunks(ctx, messageContent)
	}
	if len(retrievedChunks) > 0 {
		chatMessages = append(chatMessages, AIChatMessage{
			Role:    "system",
			Content: aiBuildKnowledgePrompt(retrievedChunks),
		})
	}
	chatMessages = append(chatMessages, history...)

	var responseBuilder strings.Builder
	chatResult, chatErr := provider.ChatStream(ctx, AIChatRequest{
		Model:    modelName,
		Messages: chatMessages,
	}, func(chunk AIChatChunk) error {
		if chunk.Content == "" {
			return nil
		}
		responseBuilder.WriteString(chunk.Content)
		emit(AIStreamEvent{
			Type: "chunk",
			Payload: map[string]any{
				"content": chunk.Content,
			},
		})
		return nil
	})
	if chatErr != nil {
		emit(AIStreamEvent{Type: "error", Payload: map[string]any{"message": "model invocation failed"}})
		return nil
	}

	aiContent := strings.TrimSpace(responseBuilder.String())
	if aiContent == "" {
		emit(AIStreamEvent{Type: "error", Payload: map[string]any{"message": "model returned empty content"}})
		return nil
	}

	aiMessageResult, err := dao.AiMessage.Ctx(ctx).Data(do.AiMessage{
		SessionId:        session.ID,
		SenderType:       aiSenderTypeAI,
		MessageContent:   aiContent,
		MessageType:      aiMessageTypeText,
		ProviderType:     providerType,
		ProviderName:     provider.Name(),
		PromptTokens:     chatResult.PromptTokens,
		CompletionTokens: chatResult.CompletionTokens,
		FinishReason:     chatResult.FinishReason,
		CreatedAt:        time.Now(),
	}).Insert()
	if err != nil {
		emit(AIStreamEvent{Type: "error", Payload: map[string]any{"message": "save ai message failed"}})
		return nil
	}
	aiMessageID, err := aiMessageResult.LastInsertId()
	if err != nil {
		emit(AIStreamEvent{Type: "error", Payload: map[string]any{"message": "read ai message id failed"}})
		return nil
	}

	analysisResult := aiBuildAnalysisResult(aiContent)
	_, _ = dao.AiAnalysisRecord.Ctx(ctx).Data(do.AiAnalysisRecord{
		PetId:            session.PetID,
		SessionId:        session.ID,
		AnalysisType:     aiAnalysisTypeRisk,
		InputSource:      aiInputSourceChat,
		SummaryTitle:     "Core analysis",
		AnalysisResult:   analysisResult,
		LlmBasedResult:   aiContent,
		RiskLevel:        aiInferRiskLevel(aiContent),
		ReviewedByDoctor: 0,
		CreatedAt:        time.Now(),
	}).Insert()

	_, _ = dao.AiSession.Ctx(ctx).
		Where(dao.AiSession.Columns().Id, session.ID).
		Data(do.AiSession{
			ProviderName:   provider.Name(),
			SessionSummary: aiBuildSessionSummary(aiContent),
			RetrievalCount: len(retrievedChunks),
			LastMessageAt:  time.Now(),
			UpdatedAt:      time.Now(),
		}).
		Update()

	finishReason := strings.TrimSpace(chatResult.FinishReason)
	if finishReason == "" {
		finishReason = "stop"
	}
	emit(AIStreamEvent{
		Type: "done",
		Payload: map[string]any{
			"message_id":    aiMessageID,
			"finish_reason": finishReason,
		},
	})
	RecordOperationLogByRole(
		ctx,
		in.SenderRole,
		in.SenderUserID,
		"ai_session",
		"send_message",
		fmt.Sprintf("send ai message session_id=%d", session.ID),
	)
	return nil
}

func (s aiService) ListMessages(ctx context.Context, in AIMessageListInput) (*AIMessageListOutput, error) {
	_, err := aiLoadAccessibleSession(ctx, in.SessionID, in.RequesterUserID, in.RequesterRole)
	if err != nil {
		return nil, err
	}

	page := in.Page
	size := in.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	if size > 100 {
		size = 100
	}

	model := dao.AiMessage.Ctx(ctx).Where(dao.AiMessage.Columns().SessionId, in.SessionID)
	total, err := model.Clone().Count()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query ai message total failed")
	}

	records, err := model.Page(page, size).OrderAsc(dao.AiMessage.Columns().Id).All()
	if err != nil {
		return nil, consts.WrapInternalError(err, "query ai messages failed")
	}

	items := make([]AIMessageItem, 0, len(records))
	for _, record := range records {
		var senderID *int64
		if record[dao.AiMessage.Columns().SenderId].Val() != nil {
			id := record[dao.AiMessage.Columns().SenderId].Int64()
			senderID = &id
		}

		items = append(items, AIMessageItem{
			ID:             record[dao.AiMessage.Columns().Id].Int64(),
			SenderType:     record[dao.AiMessage.Columns().SenderType].Int(),
			SenderID:       senderID,
			MessageContent: record[dao.AiMessage.Columns().MessageContent].String(),
			MessageType:    record[dao.AiMessage.Columns().MessageType].Int(),
			ProviderType:   record[aiMessageColumnProviderType].String(),
			ProviderName:   record[aiMessageColumnProviderName].String(),
			FinishReason:   record[aiMessageColumnFinishReason].String(),
			CreatedAt:      formatGTime(record[dao.AiMessage.Columns().CreatedAt].GTime()),
		})
	}

	RecordOperationLogByRole(
		ctx,
		in.RequesterRole,
		in.RequesterUserID,
		"ai_session",
		"list_messages",
		fmt.Sprintf("list ai messages session_id=%d", in.SessionID),
	)

	return &AIMessageListOutput{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}
