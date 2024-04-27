package logging

import (
	"context"
	"log/slog"
	"os"
)

var RequestIdKey string = "request-id"

type ServerHandler struct {
	System       string
	Environment  string
	RequestIdKey string
	slog.Handler
}

func (h *ServerHandler) Handle(ctx context.Context, r slog.Record) error {
	if val, ok := ctx.Value(RequestIdKey).(string); ok {
		r.AddAttrs(slog.String(RequestIdKey, val))
	} else {
		r.AddAttrs(slog.String(RequestIdKey, "unknown"))
	}
	return h.Handler.Handle(ctx, r)
}

func (h *ServerHandler) WithSystem(s string) *ServerHandler {
	h.System = s
	return h
}

func (h *ServerHandler) WithEnvironment(e string) *ServerHandler {
	h.Environment = e
	return h
}

func (h *ServerHandler) WithRequestIdKey(k string) *ServerHandler {
	h.RequestIdKey = k
	return h
}

func (h *ServerHandler) WithJsonBase(l slog.Level, addSrc bool) *ServerHandler {
	h.Handler = slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     l,
			AddSource: addSrc,
		},
	)
	return h
}

func (h *ServerHandler) WithTextBase(l slog.Level, addSrc bool) *ServerHandler {
	h.Handler = slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     l,
			AddSource: addSrc,
		},
	)
	return h
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{}
}
