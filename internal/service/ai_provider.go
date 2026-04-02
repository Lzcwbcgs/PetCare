package service

import (
	"context"
	"strings"

	"PetCare/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	aiProviderTypeAPI   = "api"
	aiProviderTypeLocal = "local"
)

type (
	AIChatRequest struct {
		Model    string
		Messages []AIChatMessage
	}

	AIChatMessage struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}

	AIChatChunk struct {
		Content string
	}

	AIChatResult struct {
		FinishReason     string
		PromptTokens     int
		CompletionTokens int
	}

	AIProvider interface {
		Name() string
		ChatStream(ctx context.Context, req AIChatRequest, onChunk func(AIChatChunk) error) (AIChatResult, error)
		Embed(ctx context.Context, texts []string, model string) ([][]float64, error)
		HealthCheck(ctx context.Context) error
	}
)

func aiProviderByType(ctx context.Context, providerType string) (AIProvider, error) {
	switch aiNormalizeProviderType(providerType) {
	case aiProviderTypeAPI:
		return newOpenAICompatibleProvider(ctx), nil
	case aiProviderTypeLocal:
		return newLocalOllamaProvider(ctx), nil
	default:
		return nil, consts.NewBadRequestError("unsupported model_type")
	}
}

func aiNormalizeProviderType(providerType string) string {
	switch strings.ToLower(strings.TrimSpace(providerType)) {
	case aiProviderTypeAPI:
		return aiProviderTypeAPI
	case aiProviderTypeLocal:
		return aiProviderTypeLocal
	default:
		return ""
	}
}

func aiDefaultProvider(ctx context.Context) string {
	providerType := aiNormalizeProviderType(aiConfigString(ctx, "ai.defaultProvider", aiProviderTypeAPI))
	if providerType == "" {
		return aiProviderTypeAPI
	}
	return providerType
}

func aiEmbeddingProvider(ctx context.Context) string {
	providerType := aiNormalizeProviderType(aiConfigString(ctx, "ai.embeddingProvider", ""))
	if providerType == "" {
		return aiDefaultProvider(ctx)
	}
	return providerType
}

func aiDefaultChatModel(ctx context.Context, providerType string) string {
	switch aiNormalizeProviderType(providerType) {
	case aiProviderTypeLocal:
		return aiConfigString(ctx, "ai.providers.local.chatModel", "qwen2.5:7b")
	default:
		return aiConfigString(ctx, "ai.providers.api.chatModel", "gpt-4o-mini")
	}
}

func aiDefaultEmbeddingModel(ctx context.Context, providerType string) string {
	switch aiNormalizeProviderType(providerType) {
	case aiProviderTypeLocal:
		return aiConfigString(ctx, "ai.providers.local.embeddingModel", "bge-m3")
	default:
		return aiConfigString(ctx, "ai.providers.api.embeddingModel", "text-embedding-3-small")
	}
}

func aiHistoryLimit(ctx context.Context) int {
	limit := aiConfigInt(ctx, "ai.historyLimit", 20)
	if limit <= 0 {
		return 20
	}
	if limit > 100 {
		return 100
	}
	return limit
}

func aiRequestTimeoutSeconds(ctx context.Context) int {
	seconds := aiConfigInt(ctx, "ai.requestTimeoutSeconds", 120)
	if seconds <= 0 {
		return 120
	}
	if seconds > 600 {
		return 600
	}
	return seconds
}

func aiSystemPrompt(ctx context.Context) string {
	prompt := strings.TrimSpace(aiConfigString(ctx, "ai.systemPrompt", ""))
	if prompt == "" {
		return "You are a veterinary consultation assistant. Give concise, safe, and practical advice. If there are clear warning signs, recommend seeking in-person veterinary care promptly."
	}
	return prompt
}

func aiRagEnabledDefault(ctx context.Context) int {
	if aiConfigBool(ctx, "rag.enabled", false) {
		return 1
	}
	return 0
}

func aiRagTopK(ctx context.Context) int {
	topK := aiConfigInt(ctx, "rag.topK", 4)
	if topK <= 0 {
		return 4
	}
	if topK > 20 {
		return 20
	}
	return topK
}

func aiChunkSize(ctx context.Context) int {
	size := aiConfigInt(ctx, "rag.chunkSize", 700)
	if size < 100 {
		return 100
	}
	if size > 2000 {
		return 2000
	}
	return size
}

func aiChunkOverlap(ctx context.Context) int {
	overlap := aiConfigInt(ctx, "rag.chunkOverlap", 120)
	if overlap < 0 {
		return 0
	}
	if overlap > 500 {
		return 500
	}
	return overlap
}

func aiVectorStoreType(ctx context.Context) string {
	return strings.ToLower(strings.TrimSpace(aiConfigString(ctx, "rag.vectorStore.type", "qdrant")))
}

func aiQdrantBaseURL(ctx context.Context) string {
	baseURL := strings.TrimSpace(aiConfigString(ctx, "rag.vectorStore.qdrant.baseUrl", "http://127.0.0.1:6333"))
	return strings.TrimRight(baseURL, "/")
}

func aiQdrantCollection(ctx context.Context) string {
	name := strings.TrimSpace(aiConfigString(ctx, "rag.vectorStore.qdrant.collection", "petcare_knowledge"))
	if name == "" {
		return "petcare_knowledge"
	}
	return name
}

func aiConfigString(ctx context.Context, pattern string, def string) string {
	value, err := g.Cfg().GetWithEnv(ctx, pattern, def)
	if err != nil || value == nil {
		return def
	}
	return value.String()
}

func aiConfigInt(ctx context.Context, pattern string, def int) int {
	value, err := g.Cfg().GetWithEnv(ctx, pattern, def)
	if err != nil || value == nil {
		return def
	}
	return value.Int()
}

func aiConfigBool(ctx context.Context, pattern string, def bool) bool {
	value, err := g.Cfg().GetWithEnv(ctx, pattern, def)
	if err != nil || value == nil {
		return def
	}
	return value.Bool()
}
