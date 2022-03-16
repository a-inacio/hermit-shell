package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() {
	config := zap.NewDevelopmentConfig()
	color := os.Getenv("LOG_COLOR_ENABLE")
	if color == "true" {
		config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	} else {
		config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()

	zap.ReplaceGlobals(logger)
}

func GetLogger() *zap.SugaredLogger {
	return zap.S()
}
