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
