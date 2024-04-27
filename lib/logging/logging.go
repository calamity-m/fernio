package logging

import (
	"log/slog"
)

type Logger struct {
	Config LogConfig
	*slog.Logger
}

func WithConfig(c LogConfig) func(*Logger) {
	return func(l *Logger) {
		l.Config = c
	}
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

func (l *Logger) initalize() {

	handler := NewServerHandler().WithSystem("").WithEnvironment("").WithRequestIdKey("")

	if l.Config.Structured {
		handler = handler.WithJsonBase(l.Config.Level, l.Config.AddSource)
	} else {
		handler = handler.WithTextBase(l.Config.Level, l.Config.AddSource)
	}

	l.Logger = slog.New(handler)
}
