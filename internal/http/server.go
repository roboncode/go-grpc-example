package http

import (
	"context"
	"errors"
	"example/api"
	example "example/generated"
	"example/util/check"
	"fmt"
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
	if check.PortAvailable(api.HttpPort()) == false {
		return errors.New(fmt.Sprintf("HTTP address %s is in use", api.HttpAddress()))
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := example.RegisterAppServiceHandlerFromEndpoint(ctx, mux, api.GrpcAddress(), opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(api.HttpAddress(), mux)
}

func NewServer() Server {
	return &server{}
}
