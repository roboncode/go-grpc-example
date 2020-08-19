package healthcheck

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Server interface {
	Serve(grpcServer *grpc.Server) error
}

type server struct {
	serviceName  string
	healthServer *health.Server
}

func (p *server) Serve(grpcServer *grpc.Server) error {
	healthpb.RegisterHealthServer(grpcServer, p.healthServer)
	logrus.Infoln("health check enabled")
	return nil
}

func (p *server) SetStatus(servingStatus healthpb.HealthCheckResponse_ServingStatus) {
	p.healthServer.SetServingStatus(p.serviceName, servingStatus)
}

func NewServer(serviceName string) Server {
	healthServer := health.NewServer()
	healthServer.SetServingStatus(serviceName, healthpb.HealthCheckResponse_SERVING)
	return &server{
		healthServer: healthServer,
	}
}
