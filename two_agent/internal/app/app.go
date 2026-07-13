package app

import (
	"context"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/config"
	"github.com/boginskiy/agentSystemAI/two_agent/pkg/logger"
)

type App struct {
	Cfg  config.Config
	Logg logger.Logger
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

	return nil
}

func (a *App) initModules(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	return nil
}

func (a *App) initLogger(ctx context.Context) error {
	return nil
}
