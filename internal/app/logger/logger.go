package logger

import (
	"log/slog"
	"os"
)

var opt = &slog.HandlerOptions{
	Level: slog.LevelInfo,
}

var levelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func InitLogger(logLevel string) {
	level, ok := levelMap[logLevel]
	if !ok {
		level = slog.LevelInfo
	}
	opt.Level = level
	handler := slog.NewTextHandler(os.Stdout, opt)
	slog.SetDefault(slog.New(handler))
}

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}

func StringAttr(key, value string) slog.Attr {
	return slog.String(key, value)
}
