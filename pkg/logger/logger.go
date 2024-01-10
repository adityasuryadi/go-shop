package logger

import "go.uber.org/zap"

type Logger struct {
	log *zap.SugaredLogger
}

func NewLogger() *Logger {
	log := InitLogger()
	return &Logger{
		log: log,
	}
}

func InitLogger() *zap.SugaredLogger {
	logger := zap.Must(zap.NewProduction()).WithOptions(zap.AddCallerSkip(1))

	defer logger.Sync()
	return logger.Sugar()
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	l.log.Errorf(message, args...)
}

func (l *Logger) Panicf(message string, args ...interface{}) {
	l.log.Panicf(message, args...)
}
