package command

import (
	"fmt"
	"strings"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/tools"
)

type Chat struct {
	Tools map[string]tools.Toolmaker
}

func NewChat() *Chat {
	return &Chat{
		Tools: make(map[string]tools.Toolmaker, 10),
	}
}

func (c *Chat) Processing(command string) error {
	splitcommand := c.split(command)

	if len(splitcommand) == 1 && (splitcommand[0] == "" || splitcommand[0] == " ") {
		return fmt.Errorf("empty input")
	}

	commandInput := splitcommand[0]

	// Проверка, есть ли введеная команда в списке инструментов.
	if tool, ok := c.Tools[commandInput]; ok {
		tool.Do(splitcommand)

		return nil
	}

	// Команды нет в списке инструментов. Отправляем запрос LLM.
	fmt.Println("Call LLM")

	return nil
}

func (c *Chat) AddTool(tool tools.Toolmaker) {
	c.Tools[tool.CallCommand()] = tool
}

func (c *Chat) split(line string) []string {
	return strings.Split(line, ":")
}
