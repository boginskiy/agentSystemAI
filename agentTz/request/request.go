package request

import "agentTz/models"

// ChatRequest — тело запроса к /api/chat.
type ChatRequest struct {
	Model    string               `json:"model"`
	Messages []models.ChatMessage `json:"messages"`
	Tools    []models.Tool        `json:"tools,omitempty"`
	Stream   bool                 `json:"stream"`
	Think    bool                 `json:"think"`
}
