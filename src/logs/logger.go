package logs

import (
	"fmt"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(level slog.Level) {
	fmt.Println(level)
	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	Logger = slog.New(handler)
}

func GetLogger() *slog.Logger {
	return Logger
}
