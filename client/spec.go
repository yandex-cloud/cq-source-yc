package client

// Spec defines plugin configuration
type Spec struct {
	OrganizationIDs []string `json:"organization_ids"`
	CloudIDs        []string `json:"cloud_ids"`
	FolderIDs       []string `json:"folder_ids"`

	Endpoint     string `json:"endpoint"`
	FolderFilter string `json:"folder_filter"`
}
