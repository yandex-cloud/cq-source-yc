package yc

import (
	"context"

	ycsdkv2 "github.com/yandex-cloud/go-sdk/v2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Conn resolves a raw gRPC connection for a fully-qualified service method
// (e.g. "yandex.cloud.ai.files.v1.FileService.List") via the v2 SDK's endpoint
// resolver. Used for services that have no high-level Go SDK: the returned
// connection feeds the generated <pkg>.New<Svc>ServiceClient(conn) constructors.
func Conn(ctx context.Context, sdk *ycsdkv2.SDK, method protoreflect.FullName) (grpc.ClientConnInterface, error) {
	return sdk.GetConnection(ctx, method)
}
