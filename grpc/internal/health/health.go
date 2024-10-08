package health

import (
	"context"

	"github.com/AmitSuresh/grpc/proto/health"
	"go.uber.org/zap"
)

type healthHandler struct {
	l *zap.Logger
}

func (hh *healthHandler) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (hh *healthHandler) Watch(req *health.HealthCheckRequest, s health.Health_WatchServer) error {
	return nil
}

type HealthHandler interface {
	Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error)
	Watch(req *health.HealthCheckRequest, s health.Health_WatchServer) error
}

func NewHealthHandler(l *zap.Logger) health.HealthServer {
	return &healthHandler{
		l: l,
	}
}
