package client_test

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/kubernetes"
	k8s "github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
)

func TestResolveOneofField(t *testing.T) {
	table := kubernetes.Clusters()
	table.Columns = schema.ColumnList{} // drop cloud_id: requires a live client
	require.NoError(t, table.Transform(table))

	col := func(name string) schema.Column {
		c := table.Columns.Get(name)
		require.NotNil(t, c, "column %q not found in table", name)
		return *c
	}

	tests := []struct {
		name     string
		cluster  *k8s.Cluster
		column   schema.Column
		wantJSON string
	}{
		{
			name: "network_implementation/cilium",
			cluster: &k8s.Cluster{
				NetworkImplementation: &k8s.Cluster_Cilium{
					Cilium: &k8s.Cilium{RoutingMode: k8s.Cilium_TUNNEL},
				},
			},
			column:   col("network_implementation"),
			wantJSON: `{"cilium":{"routing_mode":"TUNNEL"}}`,
		},
		{
			name:     "network_implementation/nil",
			cluster:  &k8s.Cluster{},
			column:   col("network_implementation"),
			wantJSON: "",
		},
		{
			name: "internet_gateway/set",
			cluster: &k8s.Cluster{
				InternetGateway: &k8s.Cluster_GatewayIpv4Address{
					GatewayIpv4Address: "1.2.3.4",
				},
			},
			column:   col("internet_gateway"),
			wantJSON: `{"gateway_ipv4_address":"1.2.3.4"}`,
		},
		{
			name:     "internet_gateway/nil",
			cluster:  &k8s.Cluster{},
			column:   col("internet_gateway"),
			wantJSON: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := schema.NewResourceData(table, nil, tt.cluster)
			assert.NoError(t, tt.column.Resolver(context.Background(), nil, resource, tt.column))

			got := resource.Get(tt.column.Name)
			if tt.wantJSON == "" {
				assert.False(t, got.IsValid())
			} else {
				assert.JSONEq(t, tt.wantJSON, string(got.(*scalar.JSON).Value))
			}
		})
	}
}

// TestOneofPathsAreFlat verifies the assumption documented in ResolveOneofField:
// paths for oneof fields are always a single dot-free segment.
// This holds because the CQ SDK struct transformer only recurses into anonymous/embedded
// structs — proto message sub-fields are serialised as JSON blobs and never unwrapped,
// and oneof fields (interface types) are never structs, so the transformer never
// descends into them.
func TestOneofPathsAreFlat(t *testing.T) {
	var oneofPaths []string
	spy := transformers.WithResolverTransformer(func(field reflect.StructField, path string) schema.ColumnResolver {
		if _, ok := field.Tag.Lookup("protobuf_oneof"); ok {
			oneofPaths = append(oneofPaths, path)
		}
		return nil
	})

	table := &schema.Table{
		Name:      "test",
		Transform: client.TransformWithStruct(&k8s.Cluster{}, spy),
	}
	require.NoError(t, table.Transform(table))

	require.NotEmpty(t, oneofPaths, "expected to find oneof fields in k8s.Cluster")
	for _, path := range oneofPaths {
		assert.False(t, strings.Contains(path, "."), "oneof path must be flat, got: %q", path)
	}
}
