package logger

import (
	"context"
	"os"

	"day-36/domain"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(env string, logFilePath string) (*ZapLogger, error) {
	var encoderConfig zapcore.EncoderConfig

	if env == "production" {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleSyncer := zapcore.AddSync(os.Stdout)
	var core zapcore.Core

	if logFilePath != "" {
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			fileSyncer := zapcore.AddSync(file)
			jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
			fileCore := zapcore.NewCore(jsonEncoder, fileSyncer, zap.InfoLevel)
			consoleCore := zapcore.NewCore(consoleEncoder, consoleSyncer, zap.DebugLevel)
			core = zapcore.NewTee(consoleCore, fileCore)
		} else {
			core = zapcore.NewCore(consoleEncoder, consoleSyncer, zap.DebugLevel)
		}
	} else {
		core = zapcore.NewCore(consoleEncoder, consoleSyncer, zap.DebugLevel)
	}

	zapLog := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &ZapLogger{logger: zapLog}, nil
}

func (z *ZapLogger) extractFields(ctx context.Context, keysAndValues ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(keysAndValues)/2+1)
	if reqID, ok := ctx.Value(domain.RequestIDKey).(string); ok && reqID != "" {
		fields = append(fields, zap.String("request_id", reqID))
	}
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key, ok := keysAndValues[i].(string)
			if !ok {
				key = "invalid_key"
			}
			fields = append(fields, zap.Any(key, keysAndValues[i+1]))
		}
	}
	return fields
}

func (z *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Info(msg, z.extractFields(ctx, keysAndValues...)...)
}

func (z *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Warn(msg, z.extractFields(ctx, keysAndValues...)...)
}

func (z *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Error(msg, z.extractFields(ctx, keysAndValues...)...)
}

func (z *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Debug(msg, z.extractFields(ctx, keysAndValues...)...)
}

func (z *ZapLogger) Sync() error {
	return z.logger.Sync()
}
