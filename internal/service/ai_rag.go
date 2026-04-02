package service

import (
	"context"
	"strings"
)

func aiRetrieveKnowledgeChunks(ctx context.Context, query string) ([]string, error) {
	question := strings.TrimSpace(query)
	if question == "" {
		return nil, nil
	}
	if !aiConfigBool(ctx, "rag.enabled", true) {
		return nil, nil
	}
	if aiVectorStoreType(ctx) != "qdrant" {
		return nil, nil
	}

	embeddingProviderType := aiEmbeddingProvider(ctx)
	provider, err := aiProviderByType(ctx, embeddingProviderType)
	if err != nil {
		return nil, err
	}

	embeddingModel := aiDefaultEmbeddingModel(ctx, embeddingProviderType)
	vectors, err := provider.Embed(ctx, []string{question}, embeddingModel)
	if err != nil {
		return nil, err
	}
	if len(vectors) == 0 {
		return nil, nil
	}

	qdrant := newQdrantClient(ctx)
	hits, err := qdrant.search(ctx, vectors[0], aiRagTopK(ctx))
	if err != nil {
		return nil, err
	}

	chunks := make([]string, 0, len(hits))
	for _, hit := range hits {
		if hit.Payload == nil {
			continue
		}
		content, ok := hit.Payload["content"].(string)
		if !ok || strings.TrimSpace(content) == "" {
			continue
		}
		chunks = append(chunks, strings.TrimSpace(content))
	}
	return chunks, nil
}

func aiBuildKnowledgePrompt(chunks []string) string {
	if len(chunks) == 0 {
		return ""
	}
	builder := strings.Builder{}
	builder.WriteString("Reference knowledge (use when relevant):\n")
	for i, chunk := range chunks {
		builder.WriteString("- ")
		builder.WriteString(chunk)
		if i != len(chunks)-1 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
