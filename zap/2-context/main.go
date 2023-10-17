package main

import (
	"context"

	"github.com/ramrodo/go-loggers/logger"
)

func main() {
	zapLogger := logger.Get()
	defer zapLogger.Sync()

	ctx := context.Background()

	// Wrap the context given with a new context.Context
	// containing the zap logger as a key-value pair
	ctx = logger.WithCtx(ctx, zapLogger)

	// Get the logger from the context
	l := logger.FromCtx(ctx)

	l.Debug("This is a Debug message")

	l.Info("This is a Info message")

	l.Warn("This is a Warn message")
}
