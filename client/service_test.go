package client

import (
	"context"
	"crypto/tls"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// serviceIDs pins the endpoint id every Service constant resolves to. Most of
// them come from go-sdk constants, so a dependency bump that renames one would
// otherwise quietly change which tables we sync.
var serviceIDs = map[Service]string{
	ServiceAIAssistants:          "ai-assistants",
	ServiceAIBatchInference:      "ai-foundation-models",
	ServiceAIDatasets:            "fomo-dataset",
	ServiceAIFiles:               "ai-files",
	ServiceAITuning:              "fomo-tuning",
	ServiceALB:                   "alb",
	ServiceAuditTrails:           "audittrails",
	ServiceBaremetal:             "baremetal",
	ServiceCDN:                   "cdn",
	ServiceCIC:                   "cic",
	ServiceCloudRegistry:         "cloud-registry",
	ServiceCompute:               "compute",
	ServiceContainerRegistry:     "container-registry",
	ServiceDataLens:              "datalens",
	ServiceDataSphere:            "datasphere",
	ServiceDataTransfer:          "datatransfer",
	ServiceDNS:                   "dns",
	ServiceIAM:                   "iam",
	ServiceKMS:                   "kms",
	ServiceKubernetes:            "managed-kubernetes",
	ServiceLockbox:               "lockbox",
	ServiceMDBClickhouse:         "managed-clickhouse",
	ServiceMDBGreenplum:          "managed-greenplum",
	ServiceMDBKafka:              "managed-kafka",
	ServiceMDBMongoDB:            "managed-mongodb",
	ServiceMDBMySQL:              "managed-mysql",
	ServiceMDBOpenSearch:         "managed-opensearch",
	ServiceMDBPostgreSQL:         "managed-postgresql",
	ServiceMDBRedis:              "managed-redis",
	ServiceNLB:                   "load-balancer",
	ServiceOrganizationManager:   "organization-manager",
	ServiceQuotaManager:          "quota-manager",
	ServiceResourceManager:       "resource-manager",
	ServiceServerlessAPIGateway:  "serverless-apigateway",
	ServiceServerlessContainers:  "serverless-containers",
	ServiceServerlessEventRouter: "serverless-eventrouter",
	ServiceServerlessFunctions:   "serverless-functions",
	ServiceServerlessMCPGateway:  "serverless-mcp-gateway",
	ServiceServerlessTriggers:    "serverless-triggers",
	ServiceServerlessWorkflows:   "serverless-workflows",
	ServiceStorage:               "storage-api",
	ServiceTrino:                 "managed-trino",
	ServiceVPC:                   "vpc",
	ServiceYDB:                   "ydb",
}

func TestServiceIDs(t *testing.T) {
	for service, id := range serviceIDs {
		assert.Equal(t, id, string(service))
	}
	// keying by Service collapses any two constants that resolve to the same
	// id, so the count also catches a rename that aliases one onto another
	assert.Len(t, serviceIDs, len(declaredServiceNames(t)), "serviceIDs and the Service constants in service.go disagree")
}

// TestServiceConstantsAreRealEndpoints checks every Service constant against
// the installations themselves. A constant that matches no endpoint id is a
// typo, or an id YC has retired, and would silently skip its tables
// everywhere; nothing offline can catch that now that availability is
// discovered at runtime.
//
// It needs network access (but no credentials), so it is opt-in:
//
//	YC_DISCOVERY_TEST=1 go test ./client/ -run TestServiceConstantsAreRealEndpoints
func TestServiceConstantsAreRealEndpoints(t *testing.T) {
	if os.Getenv("YC_DISCOVERY_TEST") == "" {
		t.Skip("set YC_DISCOVERY_TEST=1 to check the Service constants against live installations")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	offered := make(serviceSet)
	for region, addr := range RegionEndpoints {
		ids, err := listEndpoints(ctx, addr)
		require.NoError(t, err, "discover %s", region)
		t.Logf("%s offers %d services", region, len(ids))
		for _, id := range ids {
			offered[id] = struct{}{}
		}
	}

	for service := range serviceIDs {
		if service == ServiceDataLens {
			continue // an HTTP API, never reported by discovery
		}
		assert.Contains(t, offered, service, "no installation offers %q", service)
	}
}

func listEndpoints(ctx context.Context, addr string) ([]Service, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		return nil, err
	}
	defer func() { _ = conn.Close() }()

	resp, err := endpoint.NewApiEndpointServiceClient(conn).List(ctx, &endpoint.ListApiEndpointsRequest{})
	if err != nil {
		return nil, err
	}

	services := make([]Service, 0, len(resp.GetEndpoints()))
	for _, ep := range resp.GetEndpoints() {
		services = append(services, Service(ep.GetId()))
	}
	return services, nil
}

// declaredServiceNames reads the names of the Service constants out of
// service.go, so that a constant added there cannot escape the checks above.
// It reads names rather than values because most constants are now borrowed
// from the SDK and cannot be evaluated without type-checking the package.
func declaredServiceNames(t *testing.T) []string {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), "service.go", nil, 0)
	require.NoError(t, err)

	var names []string
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valueSpec.Names) != 1 {
				continue
			}
			if name := valueSpec.Names[0].Name; strings.HasPrefix(name, "Service") {
				names = append(names, name)
			}
		}
	}
	require.NotEmpty(t, names)
	return names
}
