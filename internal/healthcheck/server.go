package healthcheck

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Server interface {
	Serve(grpcServer *grpc.Server) error
}

type server struct {
}

func (*server) Serve(grpcServer *grpc.Server) error {
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	//glog.Infoln("Listening to HealthCheck on " + api.GrpcAddress())

	return nil
}

func NewServer() Server {
	return &server{}
}
