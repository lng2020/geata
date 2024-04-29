package logger

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewGinLoggerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		params := map[string]string{}
		for _, p := range c.Params {
			params[p.Key] = p.Value
		}

		c.Next()

		status := c.Writer.Status()
		method := c.Request.Method
		host := c.Request.Host
		route := c.FullPath()
		end := time.Now()
		latency := end.Sub(start)
		ip := c.ClientIP()

		attributes := []slog.Attr{
			slog.Duration("latency", latency),
			slog.String("method", method),
			slog.String("host", host),
			slog.String("path", path),
			slog.String("query", query),
			slog.Any("params", params),
			slog.String("route", route),
			slog.String("ip", ip),
			slog.Int("status", status),
		}

		level := slog.LevelInfo
		msg := "Incoming request"
		if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
			level = slog.LevelWarn
			msg = c.Errors.String()
		} else if status >= http.StatusInternalServerError {
			level = slog.LevelError
			msg = c.Errors.String()
		}

		logger.LogAttrs(c.Request.Context(), level, msg, attributes...)
	}
}
