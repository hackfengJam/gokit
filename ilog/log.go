package ilog

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// Logger logger
type Logger struct {
	l *zap.Logger
}

var defaultLogger struct {
	l    *Logger
	once sync.Once
}

func getDefaultLogger() *Logger {
	defaultLogger.once.Do(func() {
		defaultLogger.l = NewLogger()
	})
	return defaultLogger.l
}

// NewLogger init logger
func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{
		l: logger,
	}
}

// Sync log sync
func Sync() {
	_ = getDefaultLogger().l.Sync()
}

// Info default alias for Logger.Info.
func Info(msg string, fields ...zap.Field) {
	getDefaultLogger().l.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

// Infof default alias for Logger.Infof.
func Infof(format string, args ...interface{}) {
	getDefaultLogger().l.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, args))
}

// Debug default alias for Logger.Debug.
func Debug(msg string, fields ...zap.Field) {
	getDefaultLogger().l.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

// Debugf default alias for Logger.Debugf.
func Debugf(format string, args ...interface{}) {
	getDefaultLogger().l.WithOptions(zap.AddCallerSkip(1)).Debug(fmt.Sprintf(format, args))
}

// Error default alias for Logger.Error.
func Error(msg string, fields ...zap.Field) {
	getDefaultLogger().l.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

// Errorf default alias for Logger.Errorf.
func Errorf(format string, args ...interface{}) {
	getDefaultLogger().l.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprintf(format, args))
}
