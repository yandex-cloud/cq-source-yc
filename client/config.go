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
# Optional. Filter as described https://cloud.yandex.com/docs/resource-manager/grpc/folder_service#List
folder_filter: ""
# Optional. If not specified either using all folders accessible.
folder_ids: [<CHANGE_THIS_TO_YOUR_FOLDER_ID>]
`
}
