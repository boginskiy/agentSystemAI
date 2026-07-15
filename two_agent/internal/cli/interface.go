package cli

import (
	"context"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/tools"
)

type Printer interface {
	Print(mess string) (int, error)
}

type Formater interface {
	LineMess(mess string) string
	LineMessWithErr(mess string, err error) string
	EnterCommand(mess string) string
	WelcomeMess() string
}

type Scanner interface {
	Scan() (string, error)
}

type Commander interface {
	Run(ctx context.Context) error
}

type Chatterer interface {
	AddTool(tool tools.Toolmaker)
	Processing(command string) error
}
