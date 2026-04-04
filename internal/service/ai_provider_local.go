package service

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type localOllamaProvider struct {
	baseURL string
	client  *http.Client
}

type ollamaChatRequest struct {
	Model    string          `json:"model"`
	Messages []AIChatMessage `json:"messages"`
	Stream   bool            `json:"stream"`
}

type ollamaChatChunk struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
	Done       bool   `json:"done"`
	DoneReason string `json:"done_reason"`
}

type ollamaEmbedRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type ollamaEmbedResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
}

type ollamaEmbeddingsFallbackRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ollamaEmbeddingsFallbackResponse struct {
	Embedding []float64 `json:"embedding"`
}

func newLocalOllamaProvider(ctx context.Context) AIProvider {
	baseURL := strings.TrimRight(aiConfigString(ctx, "ai.providers.local.baseUrl", "http://127.0.0.1:11434"), "/")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:11434"
	}
	return &localOllamaProvider{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: time.Duration(aiRequestTimeoutSeconds(ctx)) * time.Second,
		},
	}
}

func (p *localOllamaProvider) Name() string {
	return "ollama"
}

func (p *localOllamaProvider) ChatStream(ctx context.Context, req AIChatRequest, onChunk func(AIChatChunk) error) (AIChatResult, error) {
	payload, err := json.Marshal(ollamaChatRequest{
		Model:    req.Model,
		Messages: req.Messages,
		Stream:   true,
	})
	if err != nil {
		return AIChatResult{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/api/chat", bytes.NewReader(payload))
	if err != nil {
		return AIChatResult{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return AIChatResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return AIChatResult{}, fmt.Errorf("ollama stream failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var result AIChatResult
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var chunk ollamaChatChunk
		if err = json.Unmarshal([]byte(line), &chunk); err != nil {
			continue
		}
		if chunk.Message.Content != "" {
			if err = onChunk(AIChatChunk{Content: chunk.Message.Content}); err != nil {
				return AIChatResult{}, err
			}
		}
		if chunk.Done {
			result.FinishReason = strings.TrimSpace(chunk.DoneReason)
			if result.FinishReason == "" {
				result.FinishReason = "stop"
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return AIChatResult{}, err
	}
	if result.FinishReason == "" {
		result.FinishReason = "stop"
	}
	return result, nil
}

func (p *localOllamaProvider) Embed(ctx context.Context, texts []string, model string) ([][]float64, error) {
	if len(texts) == 0 {
		return [][]float64{}, nil
	}

	vectors, err := p.embedBatch(ctx, texts, model)
	if err == nil {
		return vectors, nil
	}

	fallbackVectors, fallbackErr := p.embedFallback(ctx, texts, model)
	if fallbackErr != nil {
		return nil, fmt.Errorf("ollama embedding failed: %v; fallback failed: %v", err, fallbackErr)
	}
	return fallbackVectors, nil
}

func (p *localOllamaProvider) embedBatch(ctx context.Context, texts []string, model string) ([][]float64, error) {
	payload, err := json.Marshal(ollamaEmbedRequest{
		Model: model,
		Input: texts,
	})
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/api/embed", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return nil, fmt.Errorf("ollama /api/embed failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var embeddingResp ollamaEmbedResponse
	if err = json.NewDecoder(resp.Body).Decode(&embeddingResp); err != nil {
		return nil, err
	}
	return embeddingResp.Embeddings, nil
}

func (p *localOllamaProvider) embedFallback(ctx context.Context, texts []string, model string) ([][]float64, error) {
	vectors := make([][]float64, 0, len(texts))
	for _, text := range texts {
		payload, err := json.Marshal(ollamaEmbeddingsFallbackRequest{
			Model:  model,
			Prompt: text,
		})
		if err != nil {
			return nil, err
		}

		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/api/embeddings", bytes.NewReader(payload))
		if err != nil {
			return nil, err
		}
		httpReq.Header.Set("Content-Type", "application/json")

		resp, err := p.client.Do(httpReq)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
			resp.Body.Close()
			return nil, fmt.Errorf("ollama /api/embeddings failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
		}

		var embeddingResp ollamaEmbeddingsFallbackResponse
		if err = json.NewDecoder(resp.Body).Decode(&embeddingResp); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		vectors = append(vectors, embeddingResp.Embedding)
	}
	return vectors, nil
}

func (p *localOllamaProvider) HealthCheck(ctx context.Context) error {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, p.baseURL+"/api/tags", nil)
	if err != nil {
		return err
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("ollama health check failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}
