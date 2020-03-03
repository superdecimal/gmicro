package health

import (
	"context"
	hrpc "superdecimal/gmicro/pkg/proto/health"
)

type Health struct {
}

func NewServer() *Health {
	return &Health{}
}

func (hsrv *Health) Check(
	ctx context.Context,
	req *hrpc.HealthCheckRequest,
) (
	*hrpc.HealthCheckResponse,
	error,
) {
	return &hrpc.HealthCheckResponse{
		Status: hrpc.HealthCheckResponse_SERVING,
	}, nil
}

func (hsrv *Health) Watch(
	req *hrpc.HealthCheckRequest,
	stream hrpc.Health_WatchServer,
) error {
	for {
		if err := stream.Send(
			&hrpc.HealthCheckResponse{
				Status: hrpc.HealthCheckResponse_SERVING,
			}); err != nil {
			return err
		}
	}
}
