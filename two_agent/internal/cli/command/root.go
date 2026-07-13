package command

import (
	"context"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli/output"
)

type Root struct {
	Printer  output.Printer
	Formater output.Formater
}

func NewRoot(printer output.Printer, formater output.Formater) *Root {
	return &Root{
		Printer:  printer,
		Formater: formater,
	}
}

func (r *Root) Run(ctx context.Context) error {
	r.Printer.Print(r.Formater.WelcomeMess())
}

// Токен
// Дальнейшие шаги?

// А если я задаю задачу, как мне понять когда использовать tools
//  а когда LLM

// Например, создай базу данных в докер контейнере, и наполни ее такими данными, приложить
// csv filr ?
