package api

import (
	"context"
	"example/generated"
	"example/tools/env"
	"example/tools/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

var conn *grpc.ClientConn

var (
	GrpcAddr = env.Var("EXAMPLE_GRPC_ADDR").Default(":8080").Desc("gRPC address").String()
	//HttpAddr = env.Var("EXAMPLE_HTTP_ADDR").Default(":3000").Desc("HTTP address").String()
)

func Connect() (pkg.AppClient, error) {
	log.Infoln("connecting to grpc server...")
	var err error
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted, codes.Unavailable),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err = grpc.DialContext(ctx, GrpcAddr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)

	if err != nil {
		return nil, err
	}
	client := pkg.NewAppClient(conn)
	return client, nil
}

func Disconnect() error {
	if conn != nil {
		if err := conn.Close(); err != nil {
			return err
		}
	}
	return nil
}
