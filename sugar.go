package log

import (
	"go.uber.org/zap"
)

func SDPanic(args ...interface{}) {
	if isSugared() {
		sugar.DPanic(args...)
	}
}

func SDPanicf(template string, args ...interface{}) {
	if isSugared() {
		sugar.DPanicf(template, args...)
	}
}

func SDPanicln(args ...interface{}) {
	if isSugared() {
		sugar.DPanicln(args...)
	}
}

func SDPanicw(msg string, keysAndValues ...interface{}) {
	if isSugared() {
		sugar.DPanicw(msg, keysAndValues...)
	}
}

func SDebug(args ...interface{}) {
	if isSugared() {
		sugar.Debug(args...)
	}
}

func SDebugf(template string, args ...interface{}) {
	if isSugared() {
		sugar.Debugf(template, args...)
	}
}

func SDebugln(args ...interface{}) {
	if isSugared() {
		sugar.Debugln(args...)
	}
}

func SDebugw(msg string, keysAndValues ...interface{}) {
	if isSugared() {
		sugar.Debugw(msg, keysAndValues...)
	}
}

func Desugar() *zap.Logger {
	if isSugared() {
		logger = sugar.Desugar()
		return logger
	}

	return nil
}

func SError(args ...interface{}) {
	if isSugared() {
		sugar.Error(args...)
	}
}

func SErrorf(template string, args ...interface{}) {
	if isSugared() {
		sugar.Errorf(template, args...)
	}
}

func SErrorln(args ...interface{}) {
	if isSugared() {
		sugar.Errorln(args...)
	}
}

func SErrorw(msg string, keysAndValues ...interface{}) {
	if isSugared() {
		sugar.Errorw(msg, keysAndValues...)
	}
}

func SFatal(args ...interface{}) {
	if isSugared() {
		sugar.Fatal(args...)
	}
}

func SFatalf(template string, args ...interface{}) {
	if isSugared() {
		sugar.Fatalf(template, args...)
	}
}

func SFatalln(args ...interface{}) {
	if isSugared() {
		sugar.Fatalln(args...)
	}
}

func SFatalw(msg string, keysAndValues ...interface{}) {
	if isSugared() {
		sugar.Fatalw(msg, keysAndValues...)
	}
}

func SInfo(args ...interface{}) {
	if isSugared() {
		sugar.Info(args...)
	}
}

func SInfof(template string, args ...interface{}) {
	if isSugared() {
		sugar.Infof(template, args...)
	}
}

func SInfoln(args ...interface{}) {
	if isSugared() {
		sugar.Infoln(args...)
	}
}

func SInfow(msg string, keysAndValues ...interface{}) {
	if isSugared() {
		sugar.Infow(msg, keysAndValues...)
	}
}

func SNamed(s string) *zap.SugaredLogger {
	if isSugared() {
		sugar = sugar.Named(s)
		return sugar
	}

	return nil
}

func SPanic(args ...interface{}) {
	if isSugared() {
		sugar.Panic(args...)
	}
}

func SPanicf(template string, args ...interface{}) {
	if isSugared() {
		sugar.Panicf(template, args...)
	}
}

func SPanicln(args ...interface{}) {
	if isSugared() {
		sugar.Panicln(args...)
	}
}

func SPanicw(msg string, keysAndValues ...interface{}) {
	if isSugared() {
		sugar.Panicw(msg, keysAndValues...)
	}
}

func SSync() error {
	if isSugared() {
		return sugar.Sync()
	}

	return nil
}

func SWarn(args ...interface{}) {
	if isSugared() {
		sugar.Warn(args...)
	}
}

func SWarnf(template string, args ...interface{}) {
	if isSugared() {
		sugar.Warnf(template, args...)
	}
}

func SWarnln(args ...interface{}) {
	if isSugared() {
		sugar.Warnln(args...)
	}
}

func SWarnw(msg string, keysAndValues ...interface{}) {
	if isSugared() {
		sugar.Warnw(msg, keysAndValues...)
	}
}

func SWith(args ...interface{}) *zap.SugaredLogger {
	if isSugared() {
		sugar = sugar.With(args...)
		return sugar
	}

	return nil
}

func SWithOptions(opts ...zap.Option) *zap.SugaredLogger {
	if isSugared() {
		sugar = sugar.WithOptions(opts...)
		return sugar
	}

	return nil
}

func isSugared() bool {
	if !isSetup() {
		return false
	}

	if sugar != nil {
		Info("creating sugared log")
		sugar = logger.Sugar()
	}

	return true
}
