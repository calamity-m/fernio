package logging

import (
	"log/slog"
	"os"
)

type Logger struct {
	config Config
	*slog.Logger
}

func New(opts ...func(*Logger)) *Logger {
	l := &Logger{}

	for _, fn := range opts {
		fn(l)
	}

	// Initialize will set defaults based on empty config
	l.initalize()
	return l
}

func WithConfig(c Config) func(*Logger) {
	return func(l *Logger) {
		l.config = c
	}
}

func (l *Logger) initalize() {
	if l.config.Structured {
		l.Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l.config.Level, AddSource: l.config.AddSource}))
	} else {
		l.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: l.config.Level, AddSource: l.config.AddSource}))
	}
}
