package log

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func Setup() error {
	if logger != nil {
		logger.Warn("attempting to reinitialize logger")
		return nil
	}

	tmp, err := zap.NewProduction()
	if err != nil {
		log.Printf("unable to initialize zap logger: %v", err)
		return err
	}

	logger = tmp
	return nil
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	if isSetup() {
		return logger.Check(lvl, msg)
	}

	return nil
}

func Core() zapcore.Core {
	if isSetup() {
		return logger.Core()
	}

	return zapcore.NewNopCore()
}

func DPanic(msg string, fields ...zap.Field) {
	if isSetup() {
		logger.DPanic(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if isSetup() {
		logger.Debug(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if isSetup() {
		logger.Error(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if isSetup() {
		logger.Fatal(msg, fields...)
	}
}

func Info(msg string, fields ...zap.Field) {
	if isSetup() {
		logger.Info(msg, fields...)
	}
}

func Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	if isSetup() {
		logger.Log(lvl, msg, fields...)
	}
}

func Named(s string) *zap.Logger {
	if isSetup() {
		logger = logger.Named(s)
		return logger
	}

	return nil
}

func Panic(msg string, fields ...zap.Field) {
	if isSetup() {
		logger.Panic(msg, fields...)
	}
}

func Sugar() *zap.SugaredLogger {
	if isSetup() {
		sugar = logger.Sugar()
		return sugar
	}

	return nil
}

func Sync() error {
	if isSetup() {
		return logger.Sync()
	}

	return nil
}

func Warn(msg string, fields ...zap.Field) {
	if isSetup() {
		logger.Warn(msg, fields...)
	}
}

func With(fields ...zap.Field) *zap.Logger {
	if isSetup() {
		logger = logger.With(fields...)
		return logger
	}

	return nil
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	if isSetup() {
		logger = logger.WithOptions(opts...)
		return logger
	}

	return nil
}

func isSetup() bool {
	if logger != nil {
		return true
	}

	log.Print("logger is not setup")
	return false
}
