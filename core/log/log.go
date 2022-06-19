package log

import (
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Init global logger, make sure only one instance in server
func Init(log *lumberjack.Logger) *zap.Logger {
	once.Do(func() {
		writer := getLogWriter(log)
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
		logger = zap.New(core, zap.AddCaller())
	})

	return logger
}

// getEncoder adjust log message
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter set up the way log output
func getLogWriter(log *lumberjack.Logger) zapcore.WriteSyncer {
	return zapcore.AddSync(log)
}
