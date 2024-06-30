package vbuy

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Encoding string

const (
	JSON    Encoding = "json"
	Console Encoding = "console"
)

var OutputPaths = []string{"stderr"}

func NewLogger(level string, env string) *zap.Logger {
	var config zap.Config
	if env == Production {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		config = zap.Config{
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Encoding:          string(JSON),
			EncoderConfig:     encoderConfig,
		}
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		config = zap.Config{
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: false,
			Encoding:          string(Console),
			EncoderConfig:     encoderConfig,
		}
	}
	config.Level = zap.NewAtomicLevelAt(ToZapLevel(level))
	config.OutputPaths = OutputPaths
	config.ErrorOutputPaths = OutputPaths
	return zap.Must(config.Build())
}

func NewLoggerWithFields(l *zap.Logger, fields ...zap.Field) *zap.Logger {
	return l.With(fields...)
}

func ToZapLevel(level string) zapcore.Level {
	switch level {
	case "info":
		return zap.InfoLevel
	default:
		return zap.DebugLevel
	}
}
