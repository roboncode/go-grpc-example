package api

import (
	aaa "aaa/generated"
	"aaa/tools/env"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

var (
	GrpcAddr = env.Var("PERSON_GRPC_ADDR").Default(":8080").Desc("gRPC address").String()
	//HttpAddr = env.String("PERSON_HTTP_ADDR", ":3000", "HTTP address")
)

func Connect() (aaa.AppClient, error) {
	var err error
	conn, err = grpc.Dial(GrpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := aaa.NewAppClient(conn)
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
