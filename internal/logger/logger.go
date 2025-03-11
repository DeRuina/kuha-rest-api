package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

// Initialize the global zap logger
func Init() {
	rawLogger, _ := zap.NewProduction()
	Logger = rawLogger.Sugar()
}

// Ensures logs are flushed before the application exits
func Cleanup() {
	_ = Logger.Sync()
}
