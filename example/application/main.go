package main

import (
	"context"
	"log/slog"

	"github.com/MrMihen13/FinanceFlow-Libs/pkg/application"
	"github.com/MrMihen13/FinanceFlow-Libs/pkg/logs"
)

func main() {
	ctx := context.Background()

	logger := logs.NewLogger(&logs.Options{HandlerOptions: &slog.HandlerOptions{Level: slog.LevelDebug}})

	app, err := application.NewAppWithContext(ctx, logger)
	if err != nil {
		panic(err)
	}

	app.Run()
}
