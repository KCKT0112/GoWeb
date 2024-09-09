package utils

import (
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig holds configuration for the logger
type LoggerConfig struct {
	Level string
}

// Global logger
var Logger *zap.Logger

// InitializeLogger initializes the zap logger based on LoggerConfig
func InitializeLogger(config LoggerConfig) {
	logDir := "log"

	// Ensure the log directory exists
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	// Create a lumberjack logger for log rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logDir + "/" + time.Now().Format("2006-01-02") + ".log",
		MaxSize:    10,   // megabytes
		MaxBackups: 7,    // number of backups
		MaxAge:     28,   // days
		Compress:   true, // compress log files
	}

	// Define the log encoding (JSON in this case)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "msg"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	// Set the log level based on the configuration
	var level zapcore.Level
	switch strings.ToLower(config.Level) {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	// Define the log writer and encoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberjackLogger),
		level,
	)

	// Create the logger
	Logger = zap.New(core)
}
