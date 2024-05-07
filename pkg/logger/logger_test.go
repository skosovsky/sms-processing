package logger_test

import (
	"context"
	"log/slog"
	"testing"

	log "sms-processing/pkg/logger"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := log.NewLogger()

	if logger == nil {
		t.Errorf("logger should not be nil")
	}

	if log.Default() != logger {
		t.Errorf("logger should be default logger")
	}

	if slog.Default() != log.Default() {
		t.Errorf("logger should be default logger")
	}

	if logger != log.Default() {
		t.Errorf("logger should be default logger")
	}

	logger = log.NewLogger(log.WithIsJSON(true))
	if slog.Default() != logger {
		t.Errorf("logger should be default logger")
	}

	_ = log.NewLogger(log.WithIsJSON(false))
	if slog.Default() == logger {
		t.Errorf("logger should NOT be default logger")
	}

	logger = log.NewLogger(log.WithLevel("info"))
	if logger.Handler().Enabled(ctx, log.LevelDebug) {
		t.Errorf("logger should NOT be enabled for debug level")
	}

	logger = log.NewLogger(log.WithLevel("warn"))
	if logger.Handler().Enabled(ctx, log.LevelDebug) {
		t.Errorf("logger should NOT be enabled for debug level")
	}

	logger = log.NewLogger(log.WithLevel("debug"))
	enabled := []log.Level{log.LevelDebug, log.LevelInfo, log.LevelWarn, log.LevelError, log.LevelFatal}
	for _, level := range enabled {
		if !logger.Handler().Enabled(ctx, level) {
			t.Errorf("logger should be enabled for all levels")
		}
	}

	logger = log.NewLogger(log.WithLevel("abcdef"))
	if !logger.Handler().Enabled(ctx, log.LevelInfo) {
		t.Errorf("logger should be enabled for info level")
	}

	if logger.Handler().Enabled(ctx, log.LevelDebug) {
		t.Errorf("logger should NOT be enabled for info level")
	}

	logger = log.NewLogger(log.WithAddSource(true))
	if slog.Default() != logger {
		t.Errorf("logger should be default logger")
	}

	logger = log.NewLogger(log.WithSetDefault(false))
	if slog.Default() == logger {
		t.Errorf("logger should NOT be default logger")
	}

	if logger == log.Default() {
		t.Errorf("logger should NOT be default logger")
	}
}
