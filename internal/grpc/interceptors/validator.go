package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	Validate() error
}

// ValidateInterceptor executes the grpc message validator, if it exists.
func ValidateInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	validator, ok := req.(Validator)
	if ok {
		err := validator.Validate()
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	return handler(ctx, req)
}
