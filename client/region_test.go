package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yandex-cloud/cq-source-yc/client"
)

func TestRegionFromEndpoint(t *testing.T) {
	tests := []struct {
		endpoint string
		want     client.Region
		wantOk   bool
	}{
		{endpoint: "api.cloud.yandex.net:443", want: client.RegionRU, wantOk: true},
		{endpoint: "api.yandexcloud.kz:443", want: client.RegionKZ, wantOk: true},
		{endpoint: "api.cloud.yandex.net", want: client.RegionRU, wantOk: true},
		{endpoint: "API.YandexCloud.KZ:443", want: client.RegionKZ, wantOk: true},
		{endpoint: "https://api.yandexcloud.kz:443/", want: client.RegionKZ, wantOk: true},
		{endpoint: "api.cloud-preprod.yandex.net:443"},
		{endpoint: "localhost:7777"},
		{endpoint: ""},
	}

	for _, tt := range tests {
		t.Run(tt.endpoint, func(t *testing.T) {
			region, ok := client.RegionFromEndpoint(tt.endpoint)
			assert.Equal(t, tt.wantOk, ok)
			assert.Equal(t, tt.want, region)
		})
	}
}

// Every region we claim to know must be reachable from its own endpoint,
// otherwise the availability data can never be applied to it.
func TestRegionEndpointsRoundTrip(t *testing.T) {
	for region, endpoint := range client.RegionEndpoints {
		got, ok := client.RegionFromEndpoint(endpoint)
		assert.True(t, ok, "endpoint %q of region %q is not recognised", endpoint, region)
		assert.Equal(t, region, got)
	}
}
