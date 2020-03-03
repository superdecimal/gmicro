package main

import (
	"fmt"
	"net"

	"superdecimal/gmicro/services/calc-api/calculator"
	"superdecimal/gmicro/services/calc-api/config"
	"superdecimal/gmicro/services/calc-api/health"

	gmrpc "superdecimal/gmicro/pkg/proto"
	hrpc "superdecimal/gmicro/pkg/proto/health"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck

	conf, err := config.Read()
	if err != nil {
		logger.Fatal("Failed to read config", zap.Error(err))
	}

	logger.Info("Starting calc-api...", zap.Int("port", conf.Port))

	// Start new grpc server
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	go func() {
		// Init server implementations
		srv := calculator.NewServer(logger)
		hsrv := health.NewServer()

		// Register implementations on our grpc serve
		gmrpc.RegisterCalculatorAPIServer(grpcServer, srv)
		hrpc.RegisterHealthServer(grpcServer, hsrv)

		// setup listener
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", conf.Port))
		if err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}

		// serve
		err = grpcServer.Serve(lis)
		if err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	Wait(grpcServer, logger)
}
