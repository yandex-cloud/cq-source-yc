package serverless

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/pkg/errors"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/mcpgateway/v1"
	mcpgatewaysdk "github.com/yandex-cloud/go-sdk/services/serverless/mcpgateway/v1"
)

func McpGateways() *schema.Table {
	return &schema.Table{
		Name:                "yc_serverless_mcpgateway_gateways",
		Description:         `https://aistudio.yandex.ru/docs/ru/ai-studio/mcp-gateway/api-ref/grpc/McpGateway/get.html#yandex.cloud.serverless.mcpgateway.v1.McpGateway`,
		Multiplex:           client.FolderMultiplex,
		Resolver:            fetchMcpGateways,
		PreResourceResolver: getMcpGateway,
		Transform:           client.TransformWithStruct(&mcpgateway.McpGateway{}, client.PrimaryKeyIdTransformer),
	}
}

func fetchMcpGateways(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	it := mcpgatewaysdk.NewMcpGatewayClient(c.SDKv2).Iterator(ctx, &mcpgateway.ListMcpGatewayRequest{FolderId: c.FolderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}

func getMcpGateway(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)

	item, err := mcpgatewaysdk.NewMcpGatewayClient(c.SDKv2).Get(ctx, &mcpgateway.GetMcpGatewayRequest{
		McpGatewayId: resource.Item.(*mcpgateway.McpGatewayPreview).Id,
	})

	if err != nil {
		return errors.WithStack(err)
	}

	resource.SetItem(item)
	return nil
}
