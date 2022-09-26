package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/apigateway/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/containers/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/functions/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/triggers/v1"
)

func Serverless() []*Resource {
	return []*Resource{
		{
			Service:      "serverless",
			SubService:   "api_gateways",
			Struct:       new(apigateway.ApiGateway),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "serverless",
			SubService:   "functions",
			Struct:       new(functions.Function),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "serverless",
			SubService:   "containers",
			Struct:       new(containers.Container),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
		{
			Service:      "serverless",
			SubService:   "triggers",
			Struct:       new(triggers.Trigger),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
