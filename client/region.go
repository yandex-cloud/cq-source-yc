package client

import (
	"strings"
)

type Region string

const (
	RegionRU Region = "ru-central1"
	RegionKZ Region = "kz1"
)

var RegionEndpoints = map[Region]string{
	RegionRU: "api.cloud.yandex.net:443",
	RegionKZ: "api.yandexcloud.kz:443",
}

var regionByHost = func() map[string]Region {
	result := make(map[string]Region, len(RegionEndpoints))
	for region, endpoint := range RegionEndpoints {
		result[endpointHost(endpoint)] = region
	}
	return result
}()

func RegionFromEndpoint(endpoint string) (Region, bool) {
	region, ok := regionByHost[endpointHost(endpoint)]
	return region, ok
}

// endpointHost strips the scheme, the port and the path off an endpoint,
// leaving the bare lowercase host.
func endpointHost(endpoint string) string {
	host := endpoint
	if _, after, ok := strings.Cut(host, "://"); ok {
		host = after
	}
	host, _, _ = strings.Cut(host, "/")
	if i := strings.LastIndex(host, ":"); i >= 0 {
		host = host[:i]
	}
	return strings.ToLower(host)
}
