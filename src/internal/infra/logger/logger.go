package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger interface
type Logger interface {
	Info(msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
}

// ZapLogger implementa Logger usando zap
type ZapLogger struct {
	logger *zap.SugaredLogger
}

// NewLogger cria e retorna uma instância da interface Logger
func NewLogger() Logger {
	// l := zap.Must(zap.NewProduction())
	// return &ZapLogger{logger: l}

	encoder := zap.NewProductionEncoderConfig()
	zapConfig := zap.NewProductionConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig = encoder
	zapConfig.Development = false
	zapConfig.Encoding = "json"

	l, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{logger: l.Sugar()}

}

func (z *ZapLogger) Info(msg string, fields ...interface{}) {
	args := append([]interface{}{msg}, fields...)
	z.logger.Info(args...)
}

func (z *ZapLogger) Debug(msg string, fields ...interface{}) {
	args := append([]interface{}{msg}, fields...)
	z.logger.Debug(args...)
}

func (z *ZapLogger) Error(msg string, fields ...interface{}) {
	args := append([]interface{}{msg}, fields...)
	z.logger.Error(args...)
}
