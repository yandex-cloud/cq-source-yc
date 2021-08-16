package client

// Config defines Provider Configuration
type Config struct {
	OrganizationIDs []string `hcl:"organization_ids,optional"`
	CloudIDs        []string `hcl:"cloud_ids,optional"`
	FolderIDs       []string `hcl:"folder_ids,optional"`

	FolderFilter string `hcl:"folder_filter,optional"`
}

func (c Config) Example() string {
	return `configuration {
				// Optional. Filter as described https://cloud.yandex.com/docs/resource-manager/grpc/folder_service#List
				// folder_filter = ""
				// Optional. If not specified either using all folders accessible.
				// folder_ids = [<CHANGE_THIS_TO_YOUR_FOLDER_ID>]
			}`
}
