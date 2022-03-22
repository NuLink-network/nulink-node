package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var logger = logrus.New()

func InitLogger(level logrus.Level) {
	logger.SetLevel(level)
	logger.SetReportCaller(true)

	logger.SetFormatter(&logrus.TextFormatter{
		QuoteEmptyFields: true,
	})

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

}
