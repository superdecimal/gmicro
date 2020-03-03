package calculator

import (
	"context"
	"fmt"
	"io"
	"net"
	gmrpc "superdecimal/gmicro/pkg/proto"
	hrpc "superdecimal/gmicro/pkg/proto/health"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	grpcServer *grpc.Server
}

type Server interface {
	gmrpc.CalculatorAPIServer

	Listen(port int) error
	Stop()
}

func NewServer() Server {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	srv := &server{
		grpcServer: grpcServer,
	}
	hsrv := &health{}

	gmrpc.RegisterCalculatorAPIServer(grpcServer, srv)
	hrpc.RegisterHealthServer(grpcServer, hsrv)

	return srv
}

func (srv *server) Stop() {
	srv.grpcServer.GracefulStop()
}

func (srv *server) Listen(
	port int,
) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}

	err = srv.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func (srv *server) Add(
	ctx context.Context,
	req *gmrpc.AddRequest,
) (
	*gmrpc.AddResponse,
	error,
) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck

	// get the inputs
	a := req.GetA()
	b := req.GetB()

	logger.Info(
		"Add method called",
		zap.Int32("a", a),
		zap.Int32("b", b),
	)

	// process
	result := a + b

	logger.Info(
		"Add method finished",
		zap.Int32("result", result),
	)

	return &gmrpc.AddResponse{
		Result: result,
	}, nil
}

func (srv *server) Sum(
	stream gmrpc.CalculatorAPI_SumServer,
) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck

	logger.Info("Sum method called")

	total := int32(0)

	for {
		num, err := stream.Recv()
		logger.Info("Received num", zap.Int32("num", num.GetNum()))

		if err == io.EOF {
			if scerr := stream.SendAndClose(
				&gmrpc.SumResponse{
					Result: total,
				}); scerr != nil {
				logger.Error("Failed to send result", zap.Error(err))
			}

			break
		}

		if err != nil {
			logger.Error("Failed to receive", zap.Error(err))
			return err
		}

		total += num.GetNum()
	}

	logger.Info(
		"Sum method finished",
		zap.Int32("result", total),
	)

	return nil
}
