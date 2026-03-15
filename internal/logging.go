package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func setupLogger() (*os.File, error) {
	logPath := "app.log"

	// Open the file for appending, create if it doesn't exist
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String(a.Key, a.Value.Time().Format("2006-01-02T15:04:05.000Z07:00"))
			}
			// Optional: Shorten the file path to just the filename
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}
			return a
		},
	}

	// Create the logger
	logger := slog.New(slog.NewTextHandler(f, opts))
	slog.SetDefault(logger)

	return f, nil
}
