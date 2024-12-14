package logs

import (
	"log/slog"
)

const (
	LevelTrace slog.Level = -8
	LevelFatal slog.Level = 12
	LevelPanic slog.Level = 13
)

type Options struct {
	*slog.HandlerOptions
}
