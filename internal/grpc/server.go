package grpc

import (
	"example/api"
	"example/generated"
	"example/internal/grpc/interceptors"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	Serve(server example.AppServer) error
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

func (s *server) Serve(server example.AppServer) error {
	lis, err := net.Listen("tcp", api.GrpcAddress())
	if err != nil {
		return err
	}

	example.RegisterAppServer(s.Instance(), server)

	return s.Instance().Serve(lis)
}

func NewServer() Server {
	return &server{}
}
