package api

import (
	"context"
	"example/util/env"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

var conn *grpc.ClientConn

var (
	grpcHost = env.Var("EXAMPLE_GRPC_HOST").Default("localhost").Desc("gRPC host").String()
	grpcPort = env.Var("EXAMPLE_GRPC_PORT").Default("8080").Desc("gRPC port").String()
	httpHost = env.Var("EXAMPLE_HTTP_ADDR").Default("localhost").Desc("HTTP address").String()
	httpPort = env.Var("EXAMPLE_HTTP_PORT").Default("3000").Desc("HTTP port").String()
)

func GrpcHost() string {
	return grpcHost
}

func GrpcPort() string {
	return grpcPort
}

func GrpcAddress() string {
	return grpcHost + ":" + grpcPort
}

func HttpHost() string {
	return httpHost
}

func HttpPort() string {
	return httpPort
}

func HttpAddress() string {
	return httpHost + ":" + httpPort
}

func Connect(address string) (*grpc.ClientConn, error) {
	log.Infoln("connecting to grpc server...")
	var err error
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted, codes.Unavailable),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err = grpc.DialContext(ctx, address,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
