package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// A client whose installation does not offer the service must be turned away
// before the multiplexer reaches for the resource hierarchy, which is why this
// passes a client that has none.
func TestMultiplexSkipsUnavailableService(t *testing.T) {
	multiplexers := map[string]func(Service) schema.Multiplexer{
		"organization": OrganizationMultiplex,
		"cloud":        CloudMultiplex,
		"folder":       FolderMultiplex,
		"global":       GlobalMultiplex,
	}

	for name, multiplexer := range multiplexers {
		t.Run(name, func(t *testing.T) {
			c := &Client{
				Region:   RegionKZ,
				Logger:   zerolog.Nop(),
				services: newServiceSet([]string{"compute"}),
			}
			assert.Empty(t, multiplexer(ServiceBaremetal)(c))
		})
	}
}

func TestGlobalMultiplexKeepsAvailableService(t *testing.T) {
	c := &Client{
		Region:   RegionKZ,
		Logger:   zerolog.Nop(),
		services: newServiceSet([]string{"compute"}),
	}
	assert.Equal(t, []schema.ClientMeta{c}, GlobalMultiplex(ServiceCompute)(c))
}

func TestServiceAvailable(t *testing.T) {
	discovered := &Client{
		Region:   RegionKZ,
		services: newServiceSet([]string{"compute", "vpc"}),
	}
	assert.True(t, discovered.serviceAvailable(ServiceCompute))
	assert.False(t, discovered.serviceAvailable(ServiceBaremetal))
	assert.False(t, discovered.serviceAvailable("no-such-service"))

	t.Run("undiscovered_client_filters_nothing", func(t *testing.T) {
		c := &Client{Region: RegionKZ}
		assert.True(t, c.serviceAvailable(ServiceBaremetal))
		assert.True(t, c.serviceAvailable("no-such-service"))
	})

	t.Run("datalens_follows_the_region", func(t *testing.T) {
		// DataLens is an HTTP API, so discovery never lists it and the region
		// has to answer for it
		ru := &Client{Region: RegionRU, services: newServiceSet([]string{"compute"})}
		kz := &Client{Region: RegionKZ, services: newServiceSet([]string{"compute"})}
		unknown := &Client{services: newServiceSet([]string{"compute"})}

		assert.True(t, ru.serviceAvailable(ServiceDataLens))
		assert.False(t, kz.serviceAvailable(ServiceDataLens))
		assert.True(t, unknown.serviceAvailable(ServiceDataLens))
	})
}
