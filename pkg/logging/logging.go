package logging

import (
	"fmt"
	"log/slog"
)

type Logger struct {
	System          string
	Environment     string
	RequestIdHeader string
	Config          LogConfig
	*slog.Logger
}

func WithConfig(c LogConfig) func(*Logger) {
	return func(l *Logger) {
		l.Config = c
	}
}

func WithSystem(s string) func(*Logger) {
	return func(l *Logger) {
		l.System = s
	}
}

func WithEnvironment(e string) func(*Logger) {
	return func(l *Logger) {
		l.Environment = e
	}
}

func WithRequestIdHeader(h string) func(*Logger) {
	return func(l *Logger) {
		l.RequestIdHeader = h
	}
}

func New(opts ...func(*Logger)) (*Logger, error) {
	l := &Logger{}

	for _, fn := range opts {
		fn(l)
	}

	if l.System == "" || l.Environment == "" || l.RequestIdHeader == "" {
		return nil, fmt.Errorf("cannot have empty System, Environment or RequestIdHeader - %s, %s, %s", l.System, l.Environment, l.RequestIdHeader)
	}

	// Initialize will set defaults based on empty config
	handler := NewServerHandler().WithSystem(l.System).WithEnvironment(l.Environment).WithRequestIdHeader(l.RequestIdHeader)

	if l.Config.Structured {
		handler = handler.WithJsonBase(l.Config.Level, l.Config.AddSource)
	} else {
		handler = handler.WithTextBase(l.Config.Level, l.Config.AddSource)
	}

	l.Logger = slog.New(handler)

	return l, nil
}
