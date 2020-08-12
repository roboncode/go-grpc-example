package grpc

import (
	"example/api"
	"example/generated"
	"example/internal/grpc/interceptors"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"net"
)

type Serverer interface {
	Serve(server example.AppServer) error
	Server() *grpc.Server
}

type Server struct {
	gs *grpc.Server
}

func (s *Server) Server() *grpc.Server {
	if s.gs == nil {
		s.gs = grpc.NewServer(
			grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor, interceptors.ValidateInterceptor),
		)
		grpc_prometheus.Register(s.gs)
	}
	return s.gs
}

func (s *Server) Serve(server example.AppServer) error {
	lis, err := net.Listen("tcp", api.Address())
	if err != nil {
		return err
	}

	example.RegisterAppServer(s.Server(), server)

	return s.Server().Serve(lis)
}

func NewServer() Serverer {
	return &Server{}
}
