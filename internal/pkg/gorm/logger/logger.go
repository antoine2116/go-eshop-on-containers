package customLogger

import (
	"context"
	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	Logger   *zap.Logger
	LogLevel gormLogger.LogLevel
}

func NewGormLogger(logger *zap.Logger) *GormLogger {
	return &GormLogger{
		Logger:   logger,
		LogLevel: gormLogger.Info,
	}
}

func (l *GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *GormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		l.Logger.Sugar().Info(msg, data)
	}
}

func (l *GormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		l.Logger.Sugar().Info(msg, data)
	}
}

func (l *GormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		l.Logger.Sugar().Info(msg, data)
	}
}

func (l *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case l.LogLevel >= gormLogger.Info:
		sql, rows := fc()
		l.Logger.Sugar().Info("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
	case l.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		l.Logger.Sugar().Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
	case err != nil && l.LogLevel >= gormLogger.Error:
		sql, rows := fc()
		l.Logger.Sugar().Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
	}
}
