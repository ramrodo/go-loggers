package main

import (
	"github.com/ramrodo/go-loggers/logger"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(logger.Get())
}

func main() {
	zap.L().Info("Hello from Zap!")
}
