package logs

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(level slog.Level) {
	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	Logger = slog.New(handler)
}

func GetLogger() *slog.Logger {
	return Logger
}
