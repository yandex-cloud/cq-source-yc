package client

import (
	"context"

	"github.com/yandex-cloud/cq-source-yc/client/yc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GRPCConn resolves a raw gRPC connection for a fully-qualified service method
// (e.g. "yandex.cloud.ai.files.v1.FileService.List") via the v2 SDK's endpoint
// resolver. Use it for services that have no high-level Go SDK (no released
// github.com/yandex-cloud/go-sdk/services/* module): the returned connection
// feeds the generated <pkg>.New<Svc>ServiceClient(conn) constructors directly.
func (c *Client) GRPCConn(ctx context.Context, method protoreflect.FullName) (grpc.ClientConnInterface, error) {
	return yc.Conn(ctx, c.SDKv2, method)
}
