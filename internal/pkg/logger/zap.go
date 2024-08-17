package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitZapLogger() *zap.Logger {
	encoder := getEncoder()

	// First, define our level-handling logic.

	// level: debug,info,warning
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	// level: error, dpanic, panic, fatal
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	infoLevelWriterSyncer := getWriterSyncer("info")
	errorLevelWriterSyncer := getWriterSyncer("error")

	infoMultiWriteSyncer := zapcore.NewMultiWriteSyncer(infoLevelWriterSyncer, os.Stdout)
	errorMultiWriteSyncer := zapcore.NewMultiWriteSyncer(errorLevelWriterSyncer, os.Stdout)

	core := zapcore.NewCore(encoder, infoMultiWriteSyncer, infoLevel)
	errorCore := zapcore.NewCore(encoder, errorMultiWriteSyncer, errorLevel)

	coreArr := []zapcore.Core{core, errorCore}

	// export
	zapLogger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) // zap.AddCaller() will add line number and file name
	return zapLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer(level string) zapcore.WriteSyncer {
	lumberWriteSyncer := &lumberjack.Logger{
		Filename: fmt.Sprintf("logs/%s.log", level),
		MaxSize:  10, // megabytes
		Compress: true,
	}
	// file, _ := os.Create("logs/app.log")
	return zapcore.Lock(zapcore.AddSync(lumberWriteSyncer))
}
