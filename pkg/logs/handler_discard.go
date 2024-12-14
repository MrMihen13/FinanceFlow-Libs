package logs

import (
	"context"
	"log/slog"
)

type discardHandler struct{}

func NewDiscardHandler() slog.Handler { return &discardHandler{} }

func (d *discardHandler) Enabled(context.Context, slog.Level) bool { return false }

func (d *discardHandler) Handle(context.Context, slog.Record) error { return nil }

func (d *discardHandler) WithAttrs([]slog.Attr) slog.Handler { return d }

func (d *discardHandler) WithGroup(string) slog.Handler { return d }
