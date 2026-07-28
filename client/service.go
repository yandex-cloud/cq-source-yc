package client

import (
	ycsdk "github.com/yandex-cloud/go-sdk"
)

// Service identifies a Yandex Cloud API service.
//
// Values must be the endpoint ids reported by the discovery API
// (yandex.cloud.endpoint.ApiEndpointService.List), because that is what we
// match them against at runtime: an id no installation serves means the tables
// using it are skipped everywhere.
//
// Wherever the SDK we call already names the id, take it from there rather
// than repeating the string, so that gating and routing cannot drift apart.
// The rest are spelled out because no SDK constant covers them. Note that the
// ids in go-sdk/v2 (services/endpoints/dynamic_endpoints.go) are not always
// the ones discovery serves; where they differ, discovery wins.
// TestServiceIDs pins the resulting values and
// TestServiceConstantsAreRealEndpoints checks them against a live installation.
type Service string

const (
	ServiceALB                   = Service(ycsdk.ApplicationLoadBalancerServiceID)
	ServiceAuditTrails           = Service(ycsdk.AuditTrailsServiceID)
	ServiceBaremetal             = Service(ycsdk.BaremetalServiceID)
	ServiceCDN                   = Service(ycsdk.CDNID)
	ServiceCloudRegistry         = Service(ycsdk.CloudRegistryServiceID)
	ServiceCompute               = Service(ycsdk.ComputeServiceID)
	ServiceContainerRegistry     = Service(ycsdk.ContainerRegistryServiceID)
	ServiceDataSphere            = Service(ycsdk.DatasphereServiceID)
	ServiceDataTransfer          = Service(ycsdk.DataTransferServiceID)
	ServiceDNS                   = Service(ycsdk.DNSServiceID)
	ServiceIAM                   = Service(ycsdk.IAMServiceID)
	ServiceKMS                   = Service(ycsdk.KMSServiceID)
	ServiceKubernetes            = Service(ycsdk.KubernetesServiceID)
	ServiceLockbox               = Service(ycsdk.LockboxSecretServiceID)
	ServiceMDBClickhouse         = Service(ycsdk.MDBClickhouseServiceID)
	ServiceMDBGreenplum          = Service(ycsdk.MDBGreenplumServiceID)
	ServiceMDBKafka              = Service(ycsdk.MDBKafkaServiceID)
	ServiceMDBMongoDB            = Service(ycsdk.MDBMongoDBServiceID)
	ServiceMDBMySQL              = Service(ycsdk.MDBMySQLServiceID)
	ServiceMDBOpenSearch         = Service(ycsdk.MDBOpenSearchID)
	ServiceMDBPostgreSQL         = Service(ycsdk.MDBPostgreSQLServiceID)
	ServiceMDBRedis              = Service(ycsdk.MDBRedisServiceID)
	ServiceNLB                   = Service(ycsdk.LoadBalancerServiceID)
	ServiceOrganizationManager   = Service(ycsdk.OrganizationManagementServiceID)
	ServiceQuotaManager          = Service(ycsdk.QuotaManagementServiceID)
	ServiceResourceManager       = Service(ycsdk.ResourceManagementServiceID)
	ServiceServerlessAPIGateway  = Service(ycsdk.APIGatewayServiceID)
	ServiceServerlessContainers  = Service(ycsdk.ServerlessContainersServiceID)
	ServiceServerlessEventRouter = Service(ycsdk.EventrouterServiceID)
	ServiceServerlessFunctions   = Service(ycsdk.FunctionServiceID)
	ServiceServerlessMCPGateway  = Service(ycsdk.McpGatewayServiceID)
	ServiceServerlessTriggers    = Service(ycsdk.TriggerServiceID)
	ServiceServerlessWorkflows   = Service(ycsdk.WorkflowServiceID)
	ServiceStorage               = Service(ycsdk.StorageAPIServiceID)
	ServiceTrino                 = Service(ycsdk.TrinoServiceID)
	ServiceVPC                   = Service(ycsdk.VpcServiceID)
	ServiceYDB                   = Service(ycsdk.YDBServiceID)

	// ServiceAIBatchInference has no endpoint id of its own: batch inference is
	// served by the Foundation Models installation, so key it on that endpoint.
	ServiceAIBatchInference = Service(ycsdk.AIFM)

	// Interconnect and these AI services are reached over the v2 SDK, which
	// resolves endpoints by proto package rather than by named constant, so
	// there is nothing to borrow.
	ServiceAIAssistants Service = "ai-assistants"
	ServiceAIDatasets   Service = "fomo-dataset"
	ServiceAIFiles      Service = "ai-files"
	ServiceAITuning     Service = "fomo-tuning"
	ServiceCIC          Service = "cic"

	// ServiceDataLens is not a gRPC service and so never appears in discovery.
	// Its availability is pinned to the region instead, see serviceAvailable.
	ServiceDataLens Service = "datalens"
)

// serviceSet holds the services an installation offers.
type serviceSet map[Service]struct{}

func newServiceSet(ids []string) serviceSet {
	services := make(serviceSet, len(ids))
	for _, id := range ids {
		services[Service(id)] = struct{}{}
	}
	return services
}

// serviceAvailable reports whether a service can be reached in the
// installation this client is connected to.
//
// The answer comes from the endpoints the installation itself listed when the
// SDK started up, so it needs no maintenance and holds for installations we
// know nothing else about. A client whose services were never discovered – an
// unconfigured one in a test – filters nothing rather than filtering
// everything.
func (c *Client) serviceAvailable(service Service) bool {
	if service == ServiceDataLens {
		// not a gRPC service, so discovery has nothing to say about it
		return c.Region == RegionRU || c.Region == ""
	}
	if c.services == nil {
		return true
	}
	_, ok := c.services[service]
	return ok
}
