package middleware

import (
	"log/slog"
	"time"

	"github.com/calamity-m/fernio/lib/logging"
	"github.com/gin-gonic/gin"
)

func Logger(logger *logging.Logger) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()

		// Set ctx for logged vars
		ctx.Set("environment", logger.Environment)
		ctx.Set("system", logger.System)

		// Process Request
		ctx.Next()

		// Get params
		attrs := []slog.Attr{
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
			slog.String("query", ctx.Request.URL.RawQuery),
			slog.Int("status", ctx.Writer.Status()),
			slog.String("client-ip", ctx.ClientIP()),
			slog.String("user-agent", ctx.Request.UserAgent()),
		}

		end := time.Now()
		attrs = append(attrs, slog.Duration("latency", end.Sub(start)))

		logger.LogAttrs(ctx, slog.LevelDebug, "Processed Request", attrs...)
	}

}
