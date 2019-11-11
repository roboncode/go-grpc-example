package api

import (
	aaa "aaa/generated"
	"aaa/tools/env"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

var (
	GrpcAddr = env.String("PERSON_GRPC_ADDR", ":8080", "gRPC address")
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
