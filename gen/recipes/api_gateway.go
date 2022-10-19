package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
)

func ApiGateway() []*Resource {
	return []*Resource{
		{
			Service:      "api_gateway",
			SubService:   "api_gateways",
			Struct:       new(apigateway.ApiGateway),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
