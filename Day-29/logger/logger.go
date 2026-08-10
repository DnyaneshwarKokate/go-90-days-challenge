package logger

import (
	"context"
	"log/slog"
	"os"

	"day-29/domain"

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

	var core zapcore.Core

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	if env == "production" {
		consoleEncoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	consoleSyncer := zapcore.AddSync(os.Stdout)

	var syncers []zapcore.WriteSyncer
	syncers = append(syncers, consoleSyncer)

	if logFilePath != "" {
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
			fileSyncer := zapcore.AddSync(file)
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
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Info(msg, fields...)
}

func (z *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Warn(msg, fields...)
}

func (z *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Error(msg, fields...)
}

func (z *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fields := z.extractFields(ctx, keysAndValues...)
	z.logger.Debug(msg, fields...)
}

func (z *ZapLogger) Sync() error {
	return z.logger.Sync()
}

func (z *ZapLogger) GetRawZapLogger() *zap.Logger {
	return z.logger
}

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(env string) *SlogLogger {
	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return &SlogLogger{logger: slog.New(handler)}
}

func (s *SlogLogger) extractAttrs(ctx context.Context, keysAndValues ...interface{}) []interface{} {
	attrs := make([]interface{}, 0, len(keysAndValues)+2)

	if reqID, ok := ctx.Value(domain.RequestIDKey).(string); ok && reqID != "" {
		attrs = append(attrs, "request_id", reqID)
	}

	attrs = append(attrs, keysAndValues...)
	return attrs
}

func (s *SlogLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.InfoContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}

func (s *SlogLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.WarnContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}

func (s *SlogLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.ErrorContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}

func (s *SlogLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.logger.DebugContext(ctx, msg, s.extractAttrs(ctx, keysAndValues...)...)
}
