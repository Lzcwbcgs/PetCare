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

type openAICompatibleProvider struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

type openAIChatStreamRequest struct {
	Model    string          `json:"model"`
	Messages []AIChatMessage `json:"messages"`
	Stream   bool            `json:"stream"`
}

type openAIChatStreamChunk struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

type openAIEmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

type openAIEmbeddingResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}

func newOpenAICompatibleProvider(ctx context.Context) AIProvider {
	baseURL := strings.TrimRight(aiConfigString(ctx, "ai.providers.api.baseUrl", "https://api.deepseek.com/v1"), "/")
	if baseURL == "" {
		baseURL = "https://api.deepseek.com/v1"
	}
	return &openAICompatibleProvider{
		baseURL: baseURL,
		apiKey:  aiConfigString(ctx, "ai.providers.api.apiKey", ""),
		client: &http.Client{
			Timeout: time.Duration(aiRequestTimeoutSeconds(ctx)) * time.Second,
		},
	}
}

func (p *openAICompatibleProvider) Name() string {
	return "openai-compatible"
}

func (p *openAICompatibleProvider) ChatStream(ctx context.Context, req AIChatRequest, onChunk func(AIChatChunk) error) (AIChatResult, error) {
	payload, err := json.Marshal(openAIChatStreamRequest{
		Model:    req.Model,
		Messages: req.Messages,
		Stream:   true,
	})
	if err != nil {
		return AIChatResult{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/chat/completions", bytes.NewReader(payload))
	if err != nil {
		return AIChatResult{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if strings.TrimSpace(p.apiKey) != "" {
		httpReq.Header.Set("Authorization", "Bearer "+strings.TrimSpace(p.apiKey))
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return AIChatResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return AIChatResult{}, fmt.Errorf("openai-compatible stream failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var result AIChatResult
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ":") {
			continue
		}
		if !strings.HasPrefix(line, "data:") {
			continue
		}

		payloadLine := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if payloadLine == "[DONE]" {
			break
		}

		var chunk openAIChatStreamChunk
		if err = json.Unmarshal([]byte(payloadLine), &chunk); err != nil {
			continue
		}
		if len(chunk.Choices) == 0 {
			continue
		}

		content := chunk.Choices[0].Delta.Content
		if content != "" {
			if err = onChunk(AIChatChunk{Content: content}); err != nil {
				return AIChatResult{}, err
			}
		}
		if chunk.Choices[0].FinishReason != "" {
			result.FinishReason = chunk.Choices[0].FinishReason
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

func (p *openAICompatibleProvider) Embed(ctx context.Context, texts []string, model string) ([][]float64, error) {
	payload, err := json.Marshal(openAIEmbeddingRequest{
		Model: model,
		Input: texts,
	})
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/embeddings", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if strings.TrimSpace(p.apiKey) != "" {
		httpReq.Header.Set("Authorization", "Bearer "+strings.TrimSpace(p.apiKey))
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("openai-compatible embedding failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var embeddingResp openAIEmbeddingResponse
	if err = json.NewDecoder(resp.Body).Decode(&embeddingResp); err != nil {
		return nil, err
	}

	vectors := make([][]float64, 0, len(embeddingResp.Data))
	for _, item := range embeddingResp.Data {
		vectors = append(vectors, item.Embedding)
	}
	return vectors, nil
}

func (p *openAICompatibleProvider) HealthCheck(ctx context.Context) error {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, p.baseURL+"/models", nil)
	if err != nil {
		return err
	}
	if strings.TrimSpace(p.apiKey) != "" {
		httpReq.Header.Set("Authorization", "Bearer "+strings.TrimSpace(p.apiKey))
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("openai-compatible health check failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}
