package client

// Config defines Provider Configuration
type Config struct {
	OrganizationIDs []string `yaml:"organization_ids"`
	CloudIDs        []string `yaml:"cloud_ids"`
	FolderIDs       []string `yaml:"folder_ids"`

	Endpoint     string `yaml:"endpoint"`
	FolderFilter string `yaml:"folder_filter"`
}

func (c Config) Example() string {
	return `
# Specify either organization_ids, cloud_ids or folder_ids
# Optional
# organization_ids:
# Optional 
# cloud_ids:
# Optional. If not specified either using all folders accessible.
# folder_ids:

# Optional.
# endpoint: api.cloud.yandex.net:443

# Optional. Filter as described https:#cloud.yandex.com/docs/resource-manager/grpc/folder_service#List
# folder_filter:
`
}
