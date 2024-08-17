package logger

import (
	"github.com/sirupsen/logrus"
)

// var Logger      *zap.Logger
// var Logger *zap.SugaredLogger

var Logger *logrus.Logger

func init() {
	Logger = InitLogrusLogger()

	// Logger = InitZapLogger()
	// Logger = ZapLogger.Sugar()
}
