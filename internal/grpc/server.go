package grpc

import (
	"example/api"
	"example/generated"
	"example/internal/grpc/interceptors"
	"example/util/log"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	Serve(server example.AppServiceServer) error
	Instance() *grpc.Server
}

type server struct {
	gs *grpc.Server
}

func (s *server) Instance() *grpc.Server {
	if s.gs == nil {
		s.gs = grpc.NewServer(
			grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor, interceptors.ValidateInterceptor),
		)
		grpc_prometheus.Register(s.gs)
	}
	return s.gs
}

func (s *server) Serve(server example.AppServiceServer) error {
	lis, err := net.Listen("tcp", api.GrpcAddress())
	if err != nil {
		return err
	}

	example.RegisterAppServiceServer(s.Instance(), server)

	log.Infof("Listening to gRPC on %s\n", api.GrpcAddress())
	return s.Instance().Serve(lis)
}

func NewServer() Server {
	return &server{}
}
