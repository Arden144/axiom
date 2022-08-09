package log

import (
	"log"

	"go.uber.org/zap"
)

type WrappedLogger struct {
	*zap.SugaredLogger
}

func (l *WrappedLogger) Trace(args ...interface{}) {
	l.Debug(args)
}

func (l *WrappedLogger) Tracef(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

var (
	L *zap.Logger
	S *zap.SugaredLogger
	W *WrappedLogger
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	L = logger
	S = logger.Sugar()
	W = &WrappedLogger{S}
}