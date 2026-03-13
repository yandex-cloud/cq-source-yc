package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/kubernetes"
	k8s "github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
					GatewayIpv4Address: "198.51.100.1",
				},
			},
			column:   col("internet_gateway"),
			wantJSON: `{"gateway_ipv4_address":"198.51.100.1"}`,
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

func TestResolveOneofField_DottedPath(t *testing.T) {
	// Simulate WithUnwrapStructFields("Master") on Cluster — the SDK will
	// produce dotted paths like "Master.MasterType" for nested oneof fields.
	cluster := &k8s.Cluster{
		Master: &k8s.Master{
			MasterType: &k8s.Master_RegionalMaster{
				RegionalMaster: &k8s.RegionalMaster{
					RegionId:          "ru-central1",
					InternalV4Address: "10.0.0.1",
				},
			},
		},
	}

	table := &schema.Table{
		Name: "test_nested_oneof",
		Transform: client.TransformWithStruct(cluster,
			transformers.WithUnwrapStructFields("Master"),
		),
	}
	table.Columns = schema.ColumnList{}
	require.NoError(t, table.Transform(table))

	col := table.Columns.Get("master_master_type")
	require.NotNil(t, col, "column master_master_type not found; columns: %v", table.Columns.Names())

	resource := schema.NewResourceData(table, nil, cluster)
	require.NoError(t, col.Resolver(t.Context(), nil, resource, *col))

	got := resource.Get("master_master_type")
	require.True(t, got.IsValid())
	assert.JSONEq(t,
		`{"regional_master":{"region_id":"ru-central1","internal_v4_address":"10.0.0.1"}}`,
		string(got.(*scalar.JSON).Value),
	)

	// nil parent: Master is nil → should produce null column, not panic.
	clusterNilMaster := &k8s.Cluster{}
	resource2 := schema.NewResourceData(table, nil, clusterNilMaster)
	require.NoError(t, col.Resolver(t.Context(), nil, resource2, *col))
	assert.False(t, resource2.Get("master_master_type").IsValid())
}

// resolveSlice is a helper that sets up a table from the given item,
// transforms it, resolves the named column, and returns both the column and the resolved scalar.
func resolveSlice(t *testing.T, item any, columnName string) (schema.Column, scalar.Scalar) {
	t.Helper()
	table := &schema.Table{Name: "test", Transform: client.TransformWithStruct(item)}
	require.NoError(t, table.Transform(table))
	col := table.Columns.Get(columnName)
	require.NotNil(t, col)
	resource := schema.NewResourceData(table, nil, item)
	require.NoError(t, col.Resolver(t.Context(), nil, resource, *col))
	return *col, resource.Get(columnName)
}

func TestResolveProtoSlice(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		master := &k8s.Master{
			Locations: []*k8s.Location{
				{ZoneId: "ru-central1-a", SubnetId: "subnet-1"},
				{ZoneId: "ru-central1-b", SubnetId: "subnet-2"},
			},
		}
		_, got := resolveSlice(t, master, "locations")
		require.True(t, got.IsValid())
		assert.JSONEq(t,
			`[{"zone_id":"ru-central1-a","subnet_id":"subnet-1"},{"zone_id":"ru-central1-b","subnet_id":"subnet-2"}]`,
			string(got.(*scalar.JSON).Value),
		)
	})

	t.Run("nil_slice", func(t *testing.T) {
		master := &k8s.Master{}
		_, got := resolveSlice(t, master, "locations")
		assert.False(t, got.IsValid(), "nil slice should produce null column")
	})

	t.Run("empty_slice", func(t *testing.T) {
		type testStruct struct {
			Locations []*k8s.Location
		}
		item := &testStruct{Locations: []*k8s.Location{}}
		_, got := resolveSlice(t, item, "locations")
		assert.False(t, got.IsValid(), "empty slice should produce null column")
	})

	t.Run("timestamp", func(t *testing.T) {
		type testStruct struct {
			Timestamps []*timestamppb.Timestamp
		}
		ts1 := &timestamppb.Timestamp{Seconds: 1000000000}
		ts2 := &timestamppb.Timestamp{Seconds: 2000000000}
		item := &testStruct{Timestamps: []*timestamppb.Timestamp{ts1, ts2}}

		_, got := resolveSlice(t, item, "timestamps")
		require.True(t, got.IsValid())

		list := got.(*scalar.List)
		require.Len(t, list.Value, 2)
		assert.Equal(t, time.Unix(1000000000, 0).UTC(), list.Value[0].Get())
		assert.Equal(t, time.Unix(2000000000, 0).UTC(), list.Value[1].Get())
	})

	t.Run("enum", func(t *testing.T) {
		type testStruct struct {
			Statuses []k8s.Cluster_Status
		}
		item := &testStruct{
			Statuses: []k8s.Cluster_Status{k8s.Cluster_RUNNING, k8s.Cluster_STOPPED},
		}
		_, got := resolveSlice(t, item, "statuses")
		require.True(t, got.IsValid())

		list := got.(*scalar.List)
		require.Len(t, list.Value, 2)
		assert.Equal(t, "RUNNING", list.Value[0].String())
		assert.Equal(t, "STOPPED", list.Value[1].String())
	})

	t.Run("wrapper", func(t *testing.T) {
		type testStruct struct {
			Values []*wrapperspb.StringValue
		}
		item := &testStruct{
			Values: []*wrapperspb.StringValue{
				wrapperspb.String("hello"),
				wrapperspb.String("world"),
			},
		}
		_, got := resolveSlice(t, item, "values")
		require.True(t, got.IsValid())

		list := got.(*scalar.List)
		require.Len(t, list.Value, 2)
		assert.Equal(t, "hello", list.Value[0].String())
		assert.Equal(t, "world", list.Value[1].String())
	})
}
