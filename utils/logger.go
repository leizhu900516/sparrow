package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.Logger
//日志 loggername = appname
func NewLogger(loggername string) *zap.Logger {

	encoder_cfg := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	Curr_level := zap.NewAtomicLevelAt(zap.DebugLevel)

	custom_cfg := zap.Config{
		Level:            Curr_level,
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoder_cfg,
		OutputPaths:      []string{"stderr", "sparrow-services.log"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := custom_cfg.Build()
	Logger = logger.Named(loggername)
	defer Logger.Sync()

	return Logger
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
}

func init()  {
	Logger = NewLogger("sparrow")
}
//example
//Logger.Debug("adv_event_type_handle", zap.String("a", "1"))
//Logger.Info("adv_event_type_handle",
//	// Structured context as strongly-typed Field values.
//	zap.String("url",  "www.baidu.com"),
//	zap.Int("attempt", 3),
//	zap.Duration("backoff", time.Second),
//)
