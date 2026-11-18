package client

import (
	"context"

	"github.com/yandex-cloud/cq-source-yc/client/yc"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
	ycsdkv2 "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/pkg/endpoints"
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

// v2Routes are the proto packages go-sdk/v2 cannot route by itself: it keys
// them on endpoint ids discovery does not list (ai-fomo-dataset,
// ai-fomo-tuning instead of fomo-dataset, fomo-tuning) or has no entry for at
// all (batch_inference), so every call fails with "endpoint not found for
// method". They are routed to the endpoint of the service their tables are
// gated on, so that gating and routing agree.
var v2Routes = map[protoreflect.FullName]Service{
	"yandex.cloud.ai.dataset.v1":         ServiceAIDatasets,
	"yandex.cloud.ai.tuning.v1":          ServiceAITuning,
	"yandex.cloud.ai.batch_inference.v1": ServiceAIBatchInference,
}

// withV2Routes returns sdk routing v2Routes to the discovered endpoints of
// their services. A service the installation does not offer is left alone:
// its tables are skipped anyway.
func withV2Routes(sdk *ycsdkv2.SDK, discovered []*endpoint.ApiEndpoint) *ycsdkv2.SDK {
	addresses := make(map[Service]string, len(discovered))
	for _, ep := range discovered {
		addresses[Service(ep.GetId())] = ep.GetAddress()
	}
	for prefix, service := range v2Routes {
		if addr, ok := addresses[service]; ok {
			sdk = sdk.WithEndpoint(prefix, endpoints.NewEndpointParams(addr))
		}
	}
	return sdk
}
