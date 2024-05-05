package gbootstrap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !cfg.Mode.Develop {
		logMode = zapcore.InfoLevel
	}

	var cores []zapcore.Core

	logOutput := cfg.Log.Output

	// 控制台日志输出配置
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// 控制台日志输出
	if logOutput == "console" || logOutput == "both" {
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logMode)
		cores = append(cores, consoleCore)
	}

	// 文件日志输出
	if logOutput == "file" || logOutput == "both" {
		fileCore := zapcore.NewCore(getEncoder(), getWriteSyncer(), logMode)
		cores = append(cores, fileCore)
	}

	// 如果没有配置或配置错误，默认输出到控制台
	if len(cores) == 0 {
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logMode)
		cores = append(cores, consoleCore)
	}

	core := zapcore.NewTee(cores...)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	separator := string(filepath.Separator)
	rootDir, _ := os.Getwd()
	logFilePath := rootDir + separator + "log" + separator + time.Now().Format(time.DateOnly) + ".txt"
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    cfg.Log.MaxSize,    //日志文件最大尺寸(M)
		MaxBackups: cfg.Log.MaxBackups, //保留旧文件的最大个数
		MaxAge:     cfg.Log.MaxAge,     //保留旧文件的最大大天数
		Compress:   false,              //是否压缩
	}
	return zapcore.AddSync(lumberjackSyncer)
}
