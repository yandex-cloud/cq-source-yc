package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
	ycsdkv2 "github.com/yandex-cloud/go-sdk/v2"
	"github.com/yandex-cloud/go-sdk/v2/pkg/endpoints"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestWithV2Routes(t *testing.T) {
	base := (&ycsdkv2.SDK{}).WithEndpoints(endpoints.NewPrefixEndpointsResolver(endpoints.PrefixToEndpoint{
		"yandex.cloud.ai.assistants.v1": endpoints.NewEndpointParams("assistant.api.cloud.yandex.net:443"),
	}))
	// as discovery lists them in RU, where go-sdk/v2 routes none of the three
	sdk := withV2Routes(base, []*endpoint.ApiEndpoint{
		{Id: "fomo-dataset", Address: "fomo-dataset.api.cloud.yandex.net:443"},
		{Id: "fomo-tuning", Address: "fomo-tuning.api.cloud.yandex.net:443"},
		{Id: "ai-foundation-models", Address: "llm.api.cloud.yandex.net:443"},
		{Id: "ai-assistants", Address: "assistant.api.cloud.yandex.net:443"},
	})

	for method, addr := range map[protoreflect.FullName]string{
		"yandex.cloud.ai.dataset.v1.DatasetService.List":                "fomo-dataset.api.cloud.yandex.net:443",
		"yandex.cloud.ai.tuning.v1.TuningService.List":                  "fomo-tuning.api.cloud.yandex.net:443",
		"yandex.cloud.ai.batch_inference.v1.BatchInferenceService.List": "llm.api.cloud.yandex.net:443",
		"yandex.cloud.ai.assistants.v1.AssistantService.List":           "assistant.api.cloud.yandex.net:443",
	} {
		ep, err := sdk.GetEndpoint(method)
		require.NoError(t, err, method)
		assert.Equal(t, addr, ep.Addr, method)
	}
}

func TestWithV2RoutesServiceNotOffered(t *testing.T) {
	sdk := withV2Routes((&ycsdkv2.SDK{}).WithEndpoints(endpoints.NewPrefixEndpointsResolver(endpoints.PrefixToEndpoint{})), nil)

	_, err := sdk.GetEndpoint("yandex.cloud.ai.dataset.v1.DatasetService.List")
	require.Error(t, err)
}

func TestV2RoutesAreGatedServices(t *testing.T) {
	for prefix, service := range v2Routes {
		assert.Contains(t, serviceIDs, service, "%s routes to a service with no pinned endpoint id", prefix)
	}
}
