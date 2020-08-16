package grpc

import (
	"errors"
	"example/api"
	"example/generated"
	"example/internal/grpc/interceptors"
	"example/util/check"
	"fmt"
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
	if check.PortAvailable(api.GrpcPort()) == false {
		return errors.New(fmt.Sprintf("GRPC address %s is in use", api.GrpcAddress()))
	}

	lis, err := net.Listen("tcp", api.GrpcAddress())
	if err != nil {
		return err
	}

	example.RegisterAppServiceServer(s.Instance(), server)
	return s.Instance().Serve(lis)
}

func NewServer() Server {
	return &server{}
}
