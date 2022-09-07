package client

// Config defines Provider Configuration
type Config struct {
	OrganizationIDs []string `yaml:"organization_ids"`
	CloudIDs        []string `yaml:"cloud_ids"`
	FolderIDs       []string `yaml:"folder_ids"`

	FolderFilter string `yaml:"folder_filter"`
}

func (c Config) Example() string {
	return `---
Optional. If not specified either using all clouds accessible.
cloud_ids:

Optional. If not specified either using all folders accessible. Might not work without folder specified
folder_ids:

Optional. Filter as described https://github.com/yandex-cloud/cloudapi/blob/master/yandex/cloud/resourcemanager/v1/folder_service.proto
folder_filter: 
`
}
