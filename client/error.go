package client

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IgnoreErrorHandler(err error) bool {
	if grpcStatus, ok := status.FromError(err); ok && grpcStatus != nil && grpcStatus.Code() == codes.PermissionDenied {
		return true
	}

	return false
}
