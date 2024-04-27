package logging

import (
	"context"
	"log/slog"
	"os"
)

type ServerHandler struct {
	System          string
	Environment     string
	RequestIdHeader string
	slog.Handler
}

func retrieveCtxString(ctx context.Context, key any) string {
	if val, ok := ctx.Value(key).(string); ok {
		return val
	}

	return "unknown"
}

func (h *ServerHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(slog.String(h.RequestIdHeader, retrieveCtxString(ctx, h.RequestIdHeader)))
	r.AddAttrs(slog.String("environment", retrieveCtxString(ctx, "environment")))
	r.AddAttrs(slog.String("system", retrieveCtxString(ctx, "system")))

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

func (h *ServerHandler) WithRequestIdHeader(k string) *ServerHandler {
	h.RequestIdHeader = k
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
