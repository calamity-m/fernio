package middleware

import (
	"log/slog"
	"time"

	"github.com/calamity-m/fernio/lib/logging"
	"github.com/gin-gonic/gin"
)

var (
	RequestIdHeader string = "X-Request-Id"
)

func LoggingMiddleware(logger *logging.Logger) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		ctx.Next()

		/*
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path
		*/

		// Get params
		attrs := []slog.Attr{
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
			slog.String("query", ctx.Request.URL.RawQuery),
			slog.Int("status", ctx.Writer.Status()),
			slog.String("client-ip", ctx.ClientIP()),
			slog.String("user-agent", ctx.Request.UserAgent()),
		}

		// attrs = append(attrs, slog.String("request-id", ctx.Request.Header.Get("request-id")))

		end := time.Now()
		attrs = append(attrs, slog.Duration("latency", end.Sub(start)))

		ctx.Set(logging.RequestIdKey, "muh-val")

		logger.LogAttrs(ctx, slog.LevelDebug, "Processed Request", attrs...)
	}

}
