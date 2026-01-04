package logger

import (
	"github.com/racibaz/go-arch/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (*ZapLogger, error) {

	cfg := config.Get()
	config := zap.NewDevelopmentConfig()

	if cfg.IsProduction() {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	sugar := logger.Sugar()
	return &ZapLogger{logger: sugar}, nil
}

func (l *ZapLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}

func (l *ZapLogger) Info(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}

func (l *ZapLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}

func (l *ZapLogger) Error(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}

func (l *ZapLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args...)
}
