package command

import (
	"context"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli"
)

type Root struct {
	Printer   cli.Printer
	Formater  cli.Formater
	Scanner   cli.Scanner
	Chatterer cli.Chatterer
}

func NewRoot(printer cli.Printer, formater cli.Formater, scanner cli.Scanner, chatterer cli.Chatterer) *Root {
	return &Root{
		Printer:   printer,
		Formater:  formater,
		Scanner:   scanner,
		Chatterer: chatterer,
	}
}

func (r *Root) Run(ctx context.Context) error {
	r.Printer.Print(r.Formater.WelcomeMess())

	for {
		r.Printer.Print(r.Formater.EnterCommand("Enter the command"))

		select {
		case <-ctx.Done():
			r.Printer.Print(r.Formater.LineMess("The session is over"))
			return nil

		default:
			// Считываем данные от пользователя и подаем их в обработку
			command, err := r.Scanner.Scan()
			if err != nil {
				return err
			}

			if command == "/exit" || command == "/q" {
				r.Printer.Print(r.Formater.LineMess("The session is over"))
				return nil
			}

			err = r.Chatterer.Processing(command)
			if err != nil {
				r.Printer.Print(r.Formater.LineMessWithErr("Repeat the command again", err))
			}
		}
	}
}

// А если я задаю задачу, как мне понять когда использовать tools
//  а когда LLM

// Например, создай базу данных в докер контейнере, и наполни ее такими данными, приложить
// csv filr ?

// 1
// Создать контейнер с постегрес

// /create container: <описание параметров>

// Проверка, есть ли созданный ранее файл с настройками контейнера, если есть то запускаем его!
// Проверка, есть ли запущенный контейнер ? Запущенные контейнере, после тесты доожны быть уничтожены.
// Если нет, то модель должна сгенерировать новый Dockerfile c настройками Postgres
