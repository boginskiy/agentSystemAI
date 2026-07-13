package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	appl, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = appl.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

}
