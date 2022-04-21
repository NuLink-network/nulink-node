package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Create a new instance of the logger. You can have any number of instances.
var logger = logrus.New()

func InitLogger(level logrus.Level, path string) {
	logger.SetLevel(level)
	logger.SetReportCaller(true)

	logger.SetFormatter(&logrus.TextFormatter{
		QuoteEmptyFields: true,
	})

	file, err := os.OpenFile("node.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Warn("Failed to log to file, using default stderr")
	}

	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "node.log",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	})

}

func Logger() *logrus.Logger {
	return logger
}
