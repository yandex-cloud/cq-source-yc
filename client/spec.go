package client

import "github.com/cloudquery/plugin-sdk/v4/scheduler"

const defaultEndpoint = "api.cloud.yandex.net:443"

type Spec struct {
	// List of all organiztions ids to fetch information from
	OrganizationIDs []string `json:"organization_ids"`

	// List of all clouds ids to fetch information from
	CloudIDs []string `json:"cloud_ids"`

	// List of all folder ids to fetch information from
	FolderIDs []string `json:"folder_ids"`

	// If `true`, will log all GRPC calls (currently YC SDK only)
	DebugGRPC bool `json:"debug_grpc"`

	// Defines the maximum number of times an API request will be retried.
	MaxRetries int `json:"max_retries,omitempty" jsonschema:"default=3"`

	// The base URL endpoint the SDK will use
	Endpoint string `json:"endpoint"`

	// The best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency" jsonschema:"minimum=1,default=50000"`

	// The scheduler to use when determining the priority of resources to sync. By default it is set to `shuffle`.
	//
	// Available options: `dfs`, `round-robin`, `shuffle`
	Scheduler scheduler.Strategy `json:"scheduler,omitempty"`
}

func NewDefaultSpec() *Spec {
	return &Spec{
		MaxRetries:  3,
		Endpoint:    defaultEndpoint,
		Concurrency: scheduler.DefaultConcurrency,
		Scheduler:   scheduler.StrategyShuffle,
	}
}

func (spec *Spec) Validate() error {
	return nil
}
