package logger

import (
	"ares/pkg/logger/interfaces"
	zapwarpper "ares/pkg/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is the default logger
var Log = initLogger()

func initLogger() interfaces.Logger {
	return zapwarpper.New(zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // 默认 DEBUG
		Development:      true,                                 // 开发模式：Caller 信息更丰富
		Encoding:         "console",                            // 纯文本控制台输出
		OutputPaths:      []string{"stdout"},                   // 输出到标准输出
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,                   // 彩色等级词，如 DEBUG/INFO
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 可读时间格式
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // file:line
		},
	})

}

// SetLogger rewrites the default logger
func SetLogger(l interfaces.Logger) {
	if l != nil {
		Log = l
	}
}

// GetLogger 带tag的子logger
func GetLogger(tag string) interfaces.Logger {
	return Log.WithField("tag", tag)
}
