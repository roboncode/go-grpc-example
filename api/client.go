package api

import (
	"context"
	"example/util/env"
	"example/util/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

var conn *grpc.ClientConn

var (
	grpcAddr = env.Var("EXAMPLE_GRPC_ADDR").Default(":8080").Desc("gRPC address").String()
	httpAddr = env.Var("EXAMPLE_HTTP_ADDR").Default(":3000").Desc("HTTP address").String()
)

func GrpcAddress() string {
	return grpcAddr
}

func HttpAddress() string {
	return httpAddr
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
