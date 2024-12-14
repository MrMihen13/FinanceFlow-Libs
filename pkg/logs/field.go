package logs

import (
	"log/slog"
	"time"
)

func Any(key string, value any) slog.Attr                { return slog.Any(key, value) }
func String(key, value string) slog.Attr                 { return slog.String(key, value) }
func Int64(key string, value int64) slog.Attr            { return slog.Int64(key, value) }
func Int(key string, value int) slog.Attr                { return slog.Int(key, value) }
func Uint64(key string, value uint64) slog.Attr          { return slog.Uint64(key, value) }
func Float64(key string, value float64) slog.Attr        { return slog.Float64(key, value) }
func Bool(key string, value bool) slog.Attr              { return slog.Bool(key, value) }
func Time(key string, value time.Time) slog.Attr         { return slog.Time(key, value) }
func Duration(key string, value time.Duration) slog.Attr { return slog.Duration(key, value) }
func Err(val error) slog.Attr                            { return Any("error", val.Error()) }
