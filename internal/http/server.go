package http

import (
	"context"
	"example/api"
	example "example/generated"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
)

type Server interface {
	Serve() error
}

type server struct {
}

func (*server) Serve() error {
	defer glog.Flush()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := example.RegisterAppHandlerFromEndpoint(ctx, mux, api.GrpcAddress(), opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(api.HttpAddress(), mux)
}

func NewServer() Server {
	return &server{}
}
