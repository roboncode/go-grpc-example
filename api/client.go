package api

import (
	"aaa/pkg"
	"aaa/tools/env"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

var (
	GrpcAddr = env.Var("AAA_GRPC_ADDR").Default(":8080").Desc("gRPC address").String()
	//HttpAddr = env.Var("AAA_HTTP_ADDR").Default(":3000").Desc("HTTP address").String()
)

func Connect() (pkg.AppClient, error) {
	var err error
	conn, err = grpc.Dial(GrpcAddr, grpc.WithInsecure())
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
