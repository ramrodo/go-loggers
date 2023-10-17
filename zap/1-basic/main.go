package main

import (
	"errors"

	"github.com/ramrodo/go-loggers/logger"
	"go.uber.org/zap"
)

func main() {
	l := logger.Get()
	defer l.Sync()

	// logging messages using a consoleEncoder
	l.Debug("This is a Debug message")

	l.Info("This is a Info message")

	l.Warn("This is a Warn message")

	l.Error("This is a Error message")
	l.Error("This is a Error message with an error object included", zap.Error(errors.New("new Error")))

	l.Fatal("This is a Fatal error (calls os.Exit(1), even if logging at FatalLevel is disabled)")
	l.Info("This message should not be logged because the previous call to Fatal called os.Exit(1)")
}
