package models

// Tool — описание инструмента для модели.
type Tool struct {
	Type     string     `json:"type"` // всегда "function"
	Function ToolSchema `json:"function"`
}

type ToolSchema struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}
