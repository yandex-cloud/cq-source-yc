package client

// Spec defines plugin configuration
type Spec struct {
	OrganizationIDs []string `yaml:"organization_ids"`
	CloudIDs        []string `yaml:"cloud_ids"`
	FolderIDs       []string `yaml:"folder_ids"`

	Endpoint     string `yaml:"endpoint"`
	FolderFilter string `yaml:"folder_filter"`
}
