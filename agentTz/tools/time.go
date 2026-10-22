package tools

import (
	"agentTz/models"
	"fmt"
	"time"
)

func TimeTool() models.Tool {
	description := fmt.Sprintln(
		"Returns the current real-world time for a given timezone. ",
		"Use this whenever the user asks about the current time, date, or day ",
		"in any city or region. Never guess the time from your own knowledge.",
	)

	return models.Tool{
		Type: "function",
		Function: models.ToolSchema{
			Name:        "get_current_time",
			Description: description,
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"timezone": map[string]interface{}{
						"type":        "string",
						"description": "IANA timezone name, e.g. 'Europe/Moscow' or 'America/New_York'",
					},
				},
				"required": []string{"timezone"},
			},
		},
	}
}

func GetCurrentTime(timezone string) string {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Sprintf(`{"error": "unknown timezone: %s"}`, timezone)
	}

	now := time.Now().In(loc)
	return fmt.Sprintf(
		`{"timezone": %q, "time": %q, "unix": %d}`,
		timezone,
		now.Format("2006-01-02 15:04:05 MST"),
		now.Unix(),
	)
}
