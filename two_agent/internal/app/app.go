package app

import (
	"context"
	"os"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli"
	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli/command"
	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli/input"
	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli/output"
	"github.com/boginskiy/agentSystemAI/two_agent/internal/config"
	"github.com/boginskiy/agentSystemAI/two_agent/internal/tools/container"
	"github.com/boginskiy/agentSystemAI/two_agent/pkg/logger"
)

type App struct {
	Cfg       config.Config
	Logg      logger.Logger
	Commander cli.Commander
	Chatterer *command.Chat
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initModules(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.Commander.Run(ctx)
}

func (a *App) initModules(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initChat,
		a.initCommander,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initCommander(ctx context.Context) error {
	formater := output.NewFormat()
	printer := output.NewOutput()
	scanner := input.NewScanner(os.Stdin)

	a.Commander = command.NewRoot(printer, formater, scanner, a.Chatterer)
	return nil
}

func (a *App) initChat(ctx context.Context) error {
	chatterer := command.NewChat()

	// Tools
	createrContainer := container.NewContainer("/create container")

	chatterer.AddTool(createrContainer)

	a.Chatterer = chatterer
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	a.Cfg = config.NewConf()
	return nil
}

func (a *App) initLogger(ctx context.Context) error {
	return nil
}
