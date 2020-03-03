package calculator

import (
	"context"
	"io"
	gmrpc "superdecimal/gmicro/pkg/proto"

	"go.uber.org/zap"
)

type server struct {
	logger *zap.Logger
}

type Server interface {
	gmrpc.CalculatorAPIServer
}

func NewServer(logger *zap.Logger) Server {
	return &server{logger: logger}
}

func (srv *server) Add(
	ctx context.Context,
	req *gmrpc.AddRequest,
) (
	*gmrpc.AddResponse,
	error,
) {
	logger := srv.logger

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
	logger := srv.logger

	logger.Info("Sum method called")

	total := int32(0)

	for {
		// Start receiving items from the stream
		num, err := stream.Recv()
		logger.Info("Received num", zap.Int32("num", num.GetNum()))

		// if EOF we are done
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
