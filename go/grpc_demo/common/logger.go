package common

import (
	"log/slog"
	"os"
)

func GetLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
}
