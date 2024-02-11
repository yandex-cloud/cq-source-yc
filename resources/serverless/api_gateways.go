package serverless

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
)

func APIGatewayGateways() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_apigateway_gateways",
		Description: `https://cloud.yandex.ru/docs/api-gateway/apigateway/api-ref/grpc/apigateway_service#ApiGateway1`,
		Multiplex:   client.FolderMultiplex,
		Resolver:    fetchAPIGateways,
		Transform:   client.TransformWithStruct(&apigateway.ApiGateway{}, client.PrimaryKeyIdTransformer),
		Relations:   schema.Tables{APIGatewayOpenapiSpecs()},
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchAPIGateways(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.Serverless().APIGateway().ApiGateway().ApiGatewayIterator(ctx, &apigateway.ListApiGatewayRequest{FolderId: folderId})
	for it.Next() {
		res <- it.Value()
	}

	return it.Error()
}

func APIGatewayOpenapiSpecs() *schema.Table {
	return &schema.Table{
		Name:        "yc_serverless_apigateway_openapi_specs",
		Description: `https://cloud.yandex.ru/docs/api-gateway/apigateway/api-ref/grpc/apigateway_service#GetOpenapiSpecResponse`,
		Resolver:    fetchAPIGatewayOpenapiSpecs,
		Transform:   client.TransformWithStruct(&apigateway.GetOpenapiSpecResponse{}, transformers.WithSkipFields("ApiGatewayId")),
		Columns:     schema.ColumnList{client.ParentIdColumn},
	}
}

func fetchAPIGatewayOpenapiSpecs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	apigw, ok := parent.Item.(*apigateway.ApiGateway)
	if !ok {
		return fmt.Errorf("parent is not type of *apigateway.ApiGateway: %+v", apigw)
	}

	resp, err := c.SDK.Serverless().APIGateway().ApiGateway().GetOpenapiSpec(ctx, &apigateway.GetOpenapiSpecRequest{ApiGatewayId: apigw.Id})
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
