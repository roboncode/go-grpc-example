package grpc

import (
	"errors"
	"example/api"
	"example/internal/grpc/interceptors"
	"example/util/check"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"net"
)

type Options struct {
	ServiceRegistration func(s *grpc.Server)
}

type Server interface {
	Serve() error
	Instance() *grpc.Server
}

type server struct {
	gs      *grpc.Server
	options *Options
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

func (s *server) Serve() error {
	if check.PortAvailable(api.GrpcPort()) == false {
		return errors.New(fmt.Sprintf("GRPC address %s is in use", api.GrpcAddress()))
	}

	lis, err := net.Listen("tcp", api.GrpcAddress())
	if err != nil {
		return err
	}

	grpcServer := s.Instance()
	s.options.ServiceRegistration(grpcServer)

	return s.Instance().Serve(lis)
}

func NewServer(options *Options) Server {
	return &server{
		options: options,
	}
}
