package client

// Config defines Provider Configuration
type Config struct {
	OrganizationIDs []string `hcl:"organization_ids,optional"`
	CloudIDs        []string `hcl:"cloud_ids,optional"`
	FolderIDs       []string `hcl:"folder_ids,optional"`

	Endpoint     string `hcl:"endpoint,optional"`
	FolderFilter string `hcl:"folder_filter,optional"`
}

func (c Config) Example() string {
	return `configuration {
				// Specify either organization_ids, cloud_ids or folder_ids
				// Optional
				// organization_ids = [<ORGANIZATION_ID>]
				// Optional 
				// cloud_ids = [<CLOUD_ID>]
				// Optional. If not specified either using all folders accessible.
				// folder_ids = [<FOLDER_IDs>]

				// Optional.
				// endpoint = "api.cloud.yandex.net:443"

				// Optional. Filter as described https://cloud.yandex.com/docs/resource-manager/grpc/folder_service#List
				// folder_filter = ""
			}`
}
