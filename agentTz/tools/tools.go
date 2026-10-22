package tools

import (
	"encoding/json"
	"fmt"
)

// ExecuteTool выполняет инструмент по имени и возвращает результат как строку.
func ExecuteTool(name string, args string) string {
	switch name {
	case "get_current_time":
		var a struct {
			Timezone string `json:"timezone"`
		}

		if err := json.Unmarshal(json.RawMessage(args), &a); err != nil {
			return fmt.Sprintf(`{"error": "bad arguments: %s"}`, err)
		}
		return GetCurrentTime(a.Timezone)

	default:
		return fmt.Sprintf(`{"error": "unknown tool: %s"}`, name)
	}
}
