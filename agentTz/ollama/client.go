package ollama

import (
	"agentTz/request"
	"agentTz/response"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultBaseURL = "http://localhost:11434"

// Client — обёртка над HTTP-клиентом для Ollama.
type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	return &Client{
		baseURL: baseURL,
		// Таймаут большой: первый запрос может занять 30+ секунд на загрузку модели.
		http: &http.Client{Timeout: 120 * time.Second},
	}
}

// Chat отправляет запрос к Ollama и возвращает ответ.
func (c *Client) Chat(ctx context.Context, req request.ChatRequest) (*response.ChatResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// Request
	httpReq, err := http.NewRequestWithContext(
		// /api/chat
		// /v1/chat/completions
		ctx, http.MethodPost, c.baseURL+"/v1/chat/completions", bytes.NewReader(body),
	)

	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// Response
	res, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama returned %d: %s", res.StatusCode, string(respBody))
	}

	var chatRes response.ChatResponse

	if err := json.Unmarshal(respBody, &chatRes); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &chatRes, nil
}
