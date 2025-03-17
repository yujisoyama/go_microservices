package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcException(code codes.Code, customMsg string, err error) error {
	return status.Errorf(codes.Internal, customMsg, err)
}