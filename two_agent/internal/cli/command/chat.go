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

	fmt.Println(splitcommand)

	return nil
}

func (c *Chat) AddTool(tool tools.Toolmaker) {
	c.Tools[tool.CallCommand()] = tool
}

func (c *Chat) split(line string) []string {
	return strings.Split(line, ":")
}
