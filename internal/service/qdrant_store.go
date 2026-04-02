package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type qdrantPoint struct {
	ID      any
	Vector  []float64
	Payload map[string]any
}

type qdrantSearchHit struct {
	ID      any
	Score   float64
	Payload map[string]any
}

type qdrantClient struct {
	baseURL    string
	collection string
	httpClient *http.Client
}

func newQdrantClient(ctx context.Context) *qdrantClient {
	return &qdrantClient{
		baseURL:    aiQdrantBaseURL(ctx),
		collection: aiQdrantCollection(ctx),
		httpClient: &http.Client{Timeout: time.Duration(aiRequestTimeoutSeconds(ctx)) * time.Second},
	}
}

func (c *qdrantClient) ensureCollection(ctx context.Context, vectorSize int) error {
	if vectorSize <= 0 {
		return fmt.Errorf("invalid vector size")
	}

	collectionURL := c.baseURL + "/collections/" + url.PathEscape(c.collection)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, collectionURL, nil)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
	resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	if resp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("qdrant query collection failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(respBody)))
	}

	createBody, _ := json.Marshal(map[string]any{
		"vectors": map[string]any{
			"size":     vectorSize,
			"distance": "Cosine",
		},
	})
	createReq, err := http.NewRequestWithContext(ctx, http.MethodPut, collectionURL, bytes.NewReader(createBody))
	if err != nil {
		return err
	}
	createReq.Header.Set("Content-Type", "application/json")

	createResp, err := c.httpClient.Do(createReq)
	if err != nil {
		return err
	}
	defer createResp.Body.Close()
	if createResp.StatusCode < 200 || createResp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(createResp.Body, 2048))
		return fmt.Errorf("qdrant create collection failed: status=%d body=%s", createResp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}

func (c *qdrantClient) upsertPoints(ctx context.Context, points []qdrantPoint) error {
	if len(points) == 0 {
		return nil
	}
	requestBody, err := json.Marshal(map[string]any{"points": points})
	if err != nil {
		return err
	}

	upsertURL := c.baseURL + "/collections/" + url.PathEscape(c.collection) + "/points?wait=true"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, upsertURL, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return fmt.Errorf("qdrant upsert failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}

func (c *qdrantClient) search(ctx context.Context, vector []float64, limit int) ([]qdrantSearchHit, error) {
	if len(vector) == 0 {
		return nil, nil
	}
	if limit <= 0 {
		limit = 4
	}

	requestBody, err := json.Marshal(map[string]any{
		"vector":       vector,
		"limit":        limit,
		"with_payload": true,
	})
	if err != nil {
		return nil, err
	}

	searchURL := c.baseURL + "/collections/" + url.PathEscape(c.collection) + "/points/search"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, searchURL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return nil, fmt.Errorf("qdrant search failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var response struct {
		Result []struct {
			ID      any            `json:"id"`
			Score   float64        `json:"score"`
			Payload map[string]any `json:"payload"`
		} `json:"result"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	hits := make([]qdrantSearchHit, 0, len(response.Result))
	for _, item := range response.Result {
		hits = append(hits, qdrantSearchHit{
			ID:      item.ID,
			Score:   item.Score,
			Payload: item.Payload,
		})
	}
	return hits, nil
}
