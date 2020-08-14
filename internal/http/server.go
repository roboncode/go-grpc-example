package http

import (
	"context"
	"example/api"
	example "example/generated"
	"example/util/log"
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
	//defer glog.Flush()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := example.RegisterAppServiceHandlerFromEndpoint(ctx, mux, api.GrpcAddress(), opts)
	if err != nil {
		return err
	}

	log.Infof("Listening to HTTP on %s\n", api.HttpAddress())

	return http.ListenAndServe(api.HttpAddress(), mux)
}

func NewServer() Server {
	return &server{}
}
