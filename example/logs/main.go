package main

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/MrMihen13/FinanceFlow-Libs/pkg/logs"
)

func runLogs(ctx context.Context, logger *logs.Logger) {
	l := logger.With("with", "something")

	l.Info(ctx, "This is info log")
	l.Warn(ctx, "This is warning log")
	l.Error(ctx, "This is error log")
	l.Debug(ctx, "This is debug log")
	l.Trace(ctx, "This is trace log")

	l.Info(ctx, "Info log with String", logs.String("key", "value"))
	l.Info(ctx, "Info log with Int64", logs.Int64("key", 12344))
	l.Info(ctx, "Info log with Int", logs.Int("key", 12352))
	l.Info(ctx, "Info log with Uint64", logs.Uint64("key", 241234))
	l.Info(ctx, "Info log with Float64", logs.Float64("key", 3.14159))
	l.Info(ctx, "Info log with Bool", logs.Bool("key", true))
	l.Info(ctx, "Info log with Duration", logs.Duration("key", time.Second))
	l.Info(ctx, "Info log with Time", logs.Time("key", time.Now()))
	l.Info(ctx, "Info log with Err", logs.Err(errors.New("error")))
}

func main() {
	ctx := context.Background()

	runLogs(ctx, logs.NewLogger(&logs.Options{HandlerOptions: &slog.HandlerOptions{Level: slog.LevelDebug}}))
	runLogs(ctx, logs.NewPrettyLogger(&logs.Options{HandlerOptions: &slog.HandlerOptions{Level: slog.LevelDebug}}))
	runLogs(ctx, logs.NewDiscardLogger())
}
