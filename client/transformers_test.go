package client_test

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/resources/compute"
	"github.com/yandex-cloud/cq-source-yc/resources/kubernetes"
	computepb "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	k8s "github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var update = flag.Bool("update", false, "update golden files")

func TestTypeTransformer(t *testing.T) {
	tests := []struct {
		name     string
		typ      reflect.Type
		wantType arrow.DataType
	}{
		{
			name:     "[]*k8s.Location → JSON",
			typ:      reflect.TypeOf([]*k8s.Location{}),
			wantType: cqtypes.ExtensionTypes.JSON,
		},
		{
			name:     "*k8s.Master → JSON",
			typ:      reflect.TypeOf((*k8s.Master)(nil)),
			wantType: cqtypes.ExtensionTypes.JSON,
		},
		{
			name:     "*timestamppb.Timestamp → Timestamp_us",
			typ:      reflect.TypeOf((*timestamppb.Timestamp)(nil)),
			wantType: arrow.FixedWidthTypes.Timestamp_us,
		},
		{
			name:     "[]*timestamppb.Timestamp → ListOf(Timestamp_us)",
			typ:      reflect.TypeOf([]*timestamppb.Timestamp{}),
			wantType: arrow.ListOf(arrow.FixedWidthTypes.Timestamp_us),
		},
		{
			name:     "k8s.Cluster_Status (enum) → String",
			typ:      reflect.TypeOf(k8s.Cluster_RUNNING),
			wantType: arrow.BinaryTypes.String,
		},
		{
			name:     "*wrapperspb.Int64Value → Int64",
			typ:      reflect.TypeOf((*wrapperspb.Int64Value)(nil)),
			wantType: arrow.PrimitiveTypes.Int64,
		},
		{
			name:     "[]byte → nil",
			typ:      reflect.TypeOf([]byte{}),
			wantType: nil,
		},
		{
			name:     "[]k8s.Cluster_Status (enum slice) → ListOf(String)",
			typ:      reflect.TypeOf([]k8s.Cluster_Status{}),
			wantType: arrow.ListOf(arrow.BinaryTypes.String),
		},
		{
			name:     "[]string → nil (SDK handles)",
			typ:      reflect.TypeOf([]string{}),
			wantType: nil,
		},
		{
			name:     "map[string]string → nil",
			typ:      reflect.TypeOf(map[string]string{}),
			wantType: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := reflect.StructField{
				Name: "TestField",
				Type: tt.typ,
			}
			got, err := client.TypeTransformer(field)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantType, got)
		})
	}
}

func TestResolverTransformer(t *testing.T) {
	tests := []struct {
		name    string
		typ     reflect.Type
		wantNil bool
	}{
		{
			name:    "[]*k8s.Location (proto message slice)",
			typ:     reflect.TypeOf([]*k8s.Location{}),
			wantNil: false,
		},
		{
			name:    "*k8s.Master (single proto message)",
			typ:     reflect.TypeOf((*k8s.Master)(nil)),
			wantNil: false,
		},
		{
			name:    "*timestamppb.Timestamp",
			typ:     reflect.TypeOf((*timestamppb.Timestamp)(nil)),
			wantNil: false,
		},
		{
			name:    "k8s.Cluster_Status (enum)",
			typ:     reflect.TypeOf(k8s.Cluster_RUNNING),
			wantNil: false,
		},
		{
			name:    "*wrapperspb.Int64Value (wrapper)",
			typ:     reflect.TypeOf((*wrapperspb.Int64Value)(nil)),
			wantNil: false,
		},
		{
			name:    "[]byte → nil",
			typ:     reflect.TypeOf([]byte{}),
			wantNil: true,
		},
		{
			name:    "[]string → nil",
			typ:     reflect.TypeOf([]string{}),
			wantNil: true,
		},
		{
			name:    "map[string]string → nil",
			typ:     reflect.TypeOf(map[string]string{}),
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := reflect.StructField{
				Name: "TestField",
				Type: tt.typ,
			}
			got := client.ResolverTransformer(field, "TestField")
			if tt.wantNil {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestResolverTransformer_DottedPath(t *testing.T) {
	// Dotted paths arise from WithUnwrapStructFields (e.g. "CloudStatus.Id").
	// ResolverTransformer must not panic and should return nil for plain types.
	field := reflect.StructField{
		Name: "TestField",
		Type: reflect.TypeOf(""),
	}
	assert.NotPanics(t, func() {
		got := client.ResolverTransformer(field, "Parent.Child")
		assert.Nil(t, got, "plain string field with dotted path should return nil resolver")
	})
}

// snapshotTransform runs the full CQ transform pipeline on a proto message
// and compares the resulting JSON against a golden file in testdata/.
// Run with -update to regenerate golden files.
func snapshotTransform(t *testing.T, table *schema.Table, msg proto.Message) {
	t.Helper()

	table.Columns = schema.ColumnList{} // drop extra columns (cloud_id etc.)
	require.NoError(t, table.Transform(table))

	resource := schema.NewResourceData(table, nil, msg)
	for _, col := range table.Columns {
		require.NoError(t, col.Resolver(t.Context(), nil, resource, col))
	}

	arr := array.RecordToStructArray(resource.GetValues().ToArrowRecord(resource.Table.ToArrowSchema()))
	got, err := json.MarshalIndent(arr.GetOneForMarshal(0), "", "  ")
	t.Log(string(got))
	require.NoError(t, err)

	golden := filepath.Join("testdata", table.Name+".json")
	if *update {
		require.NoError(t, os.MkdirAll("testdata", 0o755))
		require.NoError(t, os.WriteFile(golden, append(got, '\n'), 0o644))
		t.Logf("updated %s", golden)
		return
	}

	want, err := os.ReadFile(golden)
	require.NoError(t, err, "golden file missing; run with -update")
	assert.JSONEq(t, string(want), string(got))
}

func TestSnapshot_k8s_Cluster(t *testing.T) {
	cluster := &k8s.Cluster{
		Id:          "test-cluster-id",
		FolderId:    "test-folder-id",
		CreatedAt:   &timestamppb.Timestamp{Seconds: 1577836800}, // 2020-01-01T00:00:00Z
		Name:        "test-cluster",
		Description: "test cluster description",
		Labels:      map[string]string{"env": "test", "managed": "true"},
		Status:      k8s.Cluster_RUNNING,
		Health:      k8s.Cluster_HEALTHY,
		NetworkId:   "test-network-id",
		Master: &k8s.Master{
			MasterType: &k8s.Master_RegionalMaster{
				RegionalMaster: &k8s.RegionalMaster{
					RegionId:          "ru-central1",
					InternalV4Address: "10.0.0.1",
				},
			},
			Locations:       []*k8s.Location{{ZoneId: "ru-central1-b", SubnetId: "test-subnet-id"}},
			EtcdClusterSize: 3,
			Version:         "1.33",
			Endpoints:       &k8s.MasterEndpoints{InternalV4Endpoint: "https://10.0.0.1"},
			MasterAuth:      &k8s.MasterAuth{ClusterCaCertificate: "dGVzdC1jYS1jZXJ0"},
			VersionInfo:     &k8s.VersionInfo{CurrentVersion: "1.33"},
			MaintenancePolicy: &k8s.MasterMaintenancePolicy{
				MaintenanceWindow: &k8s.MaintenanceWindow{
					Policy: &k8s.MaintenanceWindow_Anytime{Anytime: &k8s.AnytimeMaintenanceWindow{}},
				},
			},
			SecurityGroupIds: []string{"test-sg-1", "test-sg-2", "test-sg-3"},
			MasterLogging: &k8s.MasterLogging{
				Enabled:                  true,
				Destination:              &k8s.MasterLogging_LogGroupId{LogGroupId: "test-log-group-id"},
				AuditEnabled:             true,
				ClusterAutoscalerEnabled: true,
				KubeApiserverEnabled:     true,
				EventsEnabled:            true,
			},
			Resources: &k8s.MasterResources{Cores: 2, CoreFraction: 100, Memory: 8589934592},
			ScalePolicy: &k8s.MasterScalePolicy{
				ScaleType: &k8s.MasterScalePolicy_AutoScale_{
					AutoScale: &k8s.MasterScalePolicy_AutoScale{MinResourcePresetId: "s-c2-m8"},
				},
			},
		},
		IpAllocationPolicy: &k8s.IPAllocationPolicy{
			ClusterIpv4CidrBlock: "10.64.0.0/16",
			NodeIpv4CidrMaskSize: 24,
			ServiceIpv4CidrBlock: "10.65.0.0/16",
			ClusterIpv6CidrBlock: "fc00::64:0/112",
			ServiceIpv6CidrBlock: "fc00::65:0/112",
		},
		ServiceAccountId:     "test-sa-id",
		NodeServiceAccountId: "test-node-sa-id",
		ReleaseChannel:       k8s.ReleaseChannel_STABLE,
		KmsProvider:          &k8s.KMSProvider{KeyId: "test-kms-key-id"},
		NetworkImplementation: &k8s.Cluster_Cilium{
			Cilium: &k8s.Cilium{RoutingMode: k8s.Cilium_TUNNEL},
		},
	}

	snapshotTransform(t, kubernetes.Clusters(), cluster)
}

func TestSnapshot_compute_Instance(t *testing.T) {
	instance := &computepb.Instance{
		Id:          "test-instance-id",
		FolderId:    "test-folder-id",
		CreatedAt:   &timestamppb.Timestamp{Seconds: 1577836800}, // 2020-01-01T00:00:00Z
		Name:        "test-instance",
		Description: "test instance description",
		Labels:      map[string]string{"env": "test", "team": "platform"},
		ZoneId:      "ru-central1-a",
		PlatformId:  "standard-v3",
		Resources:   &computepb.Resources{Memory: 8589934592, Cores: 4, CoreFraction: 100},
		Status:      computepb.Instance_RUNNING,
		Metadata:    map[string]string{"user-data": "#cloud-config"},
		MetadataOptions: &computepb.MetadataOptions{
			GceHttpEndpoint:   computepb.MetadataOption_ENABLED,
			AwsV1HttpEndpoint: computepb.MetadataOption_ENABLED,
		},
		BootDisk: &computepb.AttachedDisk{
			Mode:       computepb.AttachedDisk_READ_WRITE,
			DeviceName: "boot",
			DiskId:     "test-boot-disk-id",
			AutoDelete: true,
		},
		SecondaryDisks: []*computepb.AttachedDisk{
			{Mode: computepb.AttachedDisk_READ_WRITE, DeviceName: "data", DiskId: "test-data-disk-id", AutoDelete: false},
			{Mode: computepb.AttachedDisk_READ_ONLY, DeviceName: "logs", DiskId: "test-logs-disk-id", AutoDelete: true},
		},
		NetworkInterfaces: []*computepb.NetworkInterface{
			{
				Index:      "0",
				MacAddress: "00:00:5e:00:53:01",
				SubnetId:   "test-subnet-id",
				PrimaryV4Address: &computepb.PrimaryAddress{
					Address: "10.0.0.10",
					OneToOneNat: &computepb.OneToOneNat{
						Address:   "198.51.100.1",
						IpVersion: computepb.IpVersion_IPV4,
					},
				},
				SecurityGroupIds: []string{"test-sg-1", "test-sg-2"},
			},
		},
		Fqdn:              "test-instance.ru-central1.internal",
		SchedulingPolicy:  &computepb.SchedulingPolicy{Preemptible: false},
		ServiceAccountId:  "test-sa-id",
		NetworkSettings:   &computepb.NetworkSettings{Type: computepb.NetworkSettings_STANDARD},
		PlacementPolicy:   &computepb.PlacementPolicy{},
		MaintenancePolicy: computepb.MaintenancePolicy_RESTART,
		HardwareGeneration: &computepb.HardwareGeneration{
			Features: &computepb.HardwareGeneration_LegacyFeatures{
				LegacyFeatures: &computepb.LegacyHardwareFeatures{
					PciTopology: computepb.PCITopology_PCI_TOPOLOGY_V1,
				},
			},
		},
	}

	snapshotTransform(t, compute.Instances(), instance)
}
