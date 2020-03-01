package main

import (
	"superdecimal/gmicro/pkg/utils"
	"superdecimal/gmicro/services/calc-api/calculator"
	"superdecimal/gmicro/services/calc-api/config"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck

	conf, err := config.Read()
	if err != nil {
		logger.Fatal("Failed to read config", zap.Error(err))
	}

	logger.Info("Starting calc-api...", zap.Int("port", conf.Port))

	calcServer := calculator.NewServer()

	go func() {
		if err := calcServer.Listen(conf.Port); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	utils.Wait(calcServer, logger)
}
