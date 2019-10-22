package log

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	MODE_DEV        = "dev"
	MODE_BETA       = "beta"
	MODE_PROD       = "prod"
	LOG_ROTATE_SIZE = 100
	LOG_FILENAME    = "sso.log"
)

type LogSettings struct {
	EnableConsole bool
	ConsoleLevel  string
	EnableFile    bool
	FileLevel     string
	FileLocation  string
}

var logger *zap.SugaredLogger
var (
	Debug, Warn, Info, Error, DPanic, Panic, Fatal func(template string, args ...interface{})
	//Debugf, Warnf, Infof, Errorf, DPainc, Panicf, Fatalf func(template string, args ...interface{})
)

func ConfigZapLog(s *LogSettings) {
	var allCore []zapcore.Core

	if s.EnableConsole {
		level := zap.DebugLevel
		if s.ConsoleLevel == "INFO" {
			level = zap.InfoLevel
		} else if s.ConsoleLevel == "WARN" {
			level = zap.WarnLevel
		} else if s.ConsoleLevel == "ERROR" {
			level = zap.ErrorLevel
		}
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05 "))
		}
		allCore = append(allCore, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			level),
		)
	}

	if s.EnableFile {
		level := zap.DebugLevel
		if s.FileLevel == "INFO" {
			level = zap.InfoLevel
		} else if s.FileLevel == "WARN" {
			level = zap.WarnLevel
		} else if s.FileLevel == "ERROR" {
			level = zap.ErrorLevel
		}

		hook := lumberjack.Logger{
			Filename:   s.FileLocation,
			MaxSize:    LOG_ROTATE_SIZE, // megabytes
			MaxBackups: 30,
			MaxAge:     1,    //days
			Compress:   true, // disabled by default
		}
		w := zapcore.AddSync(&hook)

		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05 "))
		}
		allCore = append(allCore, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			w,
			level),
		)
	}

	core := zapcore.NewTee(allCore...)
	logger = zap.New(core).WithOptions(zap.AddCaller()).Sugar()

	Debug = logger.Debugf
	Warn = logger.Warnf
	Info = logger.Infof
	Error = logger.Errorf
	DPanic = logger.DPanicf
	Panic = logger.Panicf
	Fatal = logger.Fatalf
}

func Sync() {
	logger.Sync()
}
