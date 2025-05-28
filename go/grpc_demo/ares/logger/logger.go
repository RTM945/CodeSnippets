package logger

import (
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

var baseLogger *zap.Logger

func init() {
	// 1. 构造一个自定义 Config
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // 默认 DEBUG
		Development:      true,                                 // 开发模式：Caller 信息更丰富
		Encoding:         "console",                            // 纯文本控制台输出
		OutputPaths:      []string{"stdout"},                   // 输出到标准输出
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,                   // 彩色等级词，如 DEBUG/INFO
			EncodeTime:     zapcore.TimeEncoderOfLayout("2025-05-21 12:17:33"), // 可读时间格式
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // file:line
		},
	}

	var err error
	baseLogger, err = cfg.Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}

	// 可选：替换全局 Logger，这样 zap.L()/zap.S() 也能用
	// zap.ReplaceGlobals(baseLogger)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc
		_ = zap.L().Sync()
	}()
}

// GetLogger 带tag的子logger
func GetLogger(tag string) *zap.SugaredLogger {
	return baseLogger.With(zap.String("tag", tag)).Sugar()
}
