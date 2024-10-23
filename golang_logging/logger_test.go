package golang_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()
	// logger.SetLevel(logrus.TraceLevel)
	logger.Trace("this a trace")
	logger.Debug("this a debug")
	logger.Info("this a Info")
	logger.Warn("this is a Warn")
	logger.Error("this is a Error")

	// logger.Panic("this is a Panic")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Watchamachalit")
	logger.Warn("Watchamachalit")
	logger.Error("Watchamachalit")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	DisableTimestamp: true,
	// 	DisableQuote:     true,
	// })
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{ //FieldMap is used to rename the field names witin JSON formatter
			logrus.FieldKeyTime: "@timestamp",
			// logrus.FieldKeyMsg:  "@message",
			logrus.FieldKeyLevel: "nil",
		},
	})

	logger.Info("Watchamachalit")
	logger.Warn("Watchamachalit")
	logger.Error("Watchamachalit")
}

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"Username": "Rifqi",
		"IsActive": "True",
	}).Info("Yes")
}
