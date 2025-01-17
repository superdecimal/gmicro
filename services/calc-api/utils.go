package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

type Stopper interface {
	GracefulStop()
}

func Wait(service Stopper, logger *zap.Logger) {
	var gracefulStop = make(chan os.Signal)

	// Get Notified for incoming signals
	signal.Notify(gracefulStop, syscall.SIGTERM) //nolint:staticcheck
	signal.Notify(gracefulStop, syscall.SIGINT)  //nolint:staticcheck

	// Wait for signal
	sig := <-gracefulStop

	logger.Info("Terminating...", zap.String("signal", sig.String()))

	service.GracefulStop()
}
