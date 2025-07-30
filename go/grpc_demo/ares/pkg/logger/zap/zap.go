package zap

import (
	"ares/logger/interfaces"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
)

type zapImpl struct {
	*zap.SugaredLogger
}

func (z zapImpl) WithFields(fields map[string]interface{}) interfaces.Logger {
	return zapImpl{z.SugaredLogger.With(fields)}
}

func (z zapImpl) WithField(key string, value interface{}) interfaces.Logger {
	return zapImpl{z.SugaredLogger.With(key, value)}
}

func (z zapImpl) WithError(err error) interfaces.Logger {
	return zapImpl{z.SugaredLogger.With("error", err)}
}

func (z zapImpl) GetInternalLogger() any {
	return z.SugaredLogger
}

func New(cfg zap.Config) interfaces.Logger {
	logger := zap.Must(cfg.Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)))
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		_ = logger.Sync()
	}()
	return &zapImpl{SugaredLogger: logger.Sugar()}
}
