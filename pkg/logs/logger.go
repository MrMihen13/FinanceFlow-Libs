package logs

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(opts *Options) *Logger {
	return &Logger{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, opts.HandlerOptions)),
	}
}

func NewPrettyLogger(opts *Options) *Logger {
	return &Logger{
		Logger: slog.New(NewPrettyHandler(os.Stdout, opts)),
	}
}

func NewDiscardLogger() *Logger {
	return &Logger{
		Logger: slog.New(NewDiscardHandler()),
	}
}

func (l *Logger) logAndCapture(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.Log(ctx, level, msg, args...)
}

func (l *Logger) clone() *Logger {
	clonedLogger := &Logger{
		Logger: l.Logger,
	}
	return clonedLogger
}

func (l *Logger) Info(ctx context.Context, msg string, args ...any) {
	l.logAndCapture(ctx, slog.LevelInfo, msg, args...)
}

func (l *Logger) Warn(ctx context.Context, msg string, args ...any) {
	l.logAndCapture(ctx, slog.LevelWarn, msg, args...)
}

func (l *Logger) Error(ctx context.Context, msg string, args ...any) {
	l.logAndCapture(ctx, slog.LevelError, msg, args...)
}

func (l *Logger) Debug(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, slog.LevelDebug, msg, args...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, args ...any) {
	l.logAndCapture(ctx, slog.LevelError, msg, args...)
	os.Exit(1)
}

func (l *Logger) Trace(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelTrace, msg, args...)
}

func (l *Logger) Panic(ctx context.Context, msg string, args ...any) {
	l.logAndCapture(ctx, slog.LevelError, msg, args...)
	panic(msg)
}
func (l *Logger) With(attrs ...any) *Logger {
	c := l.clone()
	return &Logger{Logger: c.Logger.With(attrs...)}
}
