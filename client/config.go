package client

// Config defines Provider Configuration
type Config struct {
	OrganizationIDs []string `yml:"organization_ids,optional"`
	CloudIDs        []string `yml:"cloud_ids,optional"`
	FolderIDs       []string `yml:"folder_ids,optional"`

	FolderFilter string `yml:"folder_filter,optional"`
}

func (c Config) Example() string {
	return `---
Optional. If not specified either using all clouds accessible.
cloud_ids:

Optional. If not specified either using all folders accessible. Might not work without folder specified
folder_ids:

Optional. Filter as described https://github.com/yandex-cloud/cloudapi/blob/master/yandex/cloud/resourcemanager/v1/folder_service.proto
folder_filter: 

list of resources to fetch
`
}
