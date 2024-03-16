package log

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
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
	L  *zap.Logger
	S  *zap.SugaredLogger
	W  *WrappedLogger
	SL *slog.Logger
)

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	L = logger
	S = logger.Sugar()
	W = &WrappedLogger{S}
	SL = slog.New(zapslog.NewHandler(logger.Core(), nil))
}
