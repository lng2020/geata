package logger

import (
	"fmt"
	"log/slog"

	xormlog "xorm.io/xorm/log"
)

type XORMLogger struct {
	logger *slog.Logger
}

func NewXORMLogger() *XORMLogger {
	return &XORMLogger{
		logger: slog.Default(),
	}
}

func (l *XORMLogger) Debug(v ...interface{}) {
	l.logger.Debug(fmt.Sprint(v...))
}

func (l *XORMLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, v...))
}

func (l *XORMLogger) Error(v ...interface{}) {
	l.logger.Error(fmt.Sprint(v...))
}

func (l *XORMLogger) Errorf(format string, v ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, v...))
}

func (l *XORMLogger) Info(v ...interface{}) {
	l.logger.Info(fmt.Sprint(v...))
}

func (l *XORMLogger) Infof(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, v...))
}

func (l *XORMLogger) Warn(v ...interface{}) {
	l.logger.Warn(fmt.Sprint(v...))
}

func (l *XORMLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, v...))
}

func (l *XORMLogger) Level() xormlog.LogLevel {
	return xormlog.LogLevel(opt.Level.Level())
}

func (l *XORMLogger) SetLevel(level xormlog.LogLevel) {
	// We do not need to implement this method
}

func (l *XORMLogger) ShowSQL(show ...bool) {
	// we do not need to implement this method
}

func (l *XORMLogger) IsShowSQL() bool {
	return false
}
