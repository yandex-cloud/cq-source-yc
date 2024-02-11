package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func OrganizationMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	hierarchyItems := client.hierarchy.OrganizationRows()

	var l = make([]schema.ClientMeta, len(hierarchyItems))
	for i, item := range hierarchyItems {
		l[i] = client.WithOrganization(item.Organization).WithMultiplexedResourceId(item.Organization)
	}
	return l
}

func CloudMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	hierarchyItems := client.hierarchy.CloudRows()

	var l = make([]schema.ClientMeta, len(hierarchyItems))
	for i, item := range hierarchyItems {
		l[i] = client.WithOrganization(item.Organization).WithCloud(item.Cloud).WithMultiplexedResourceId(item.Cloud)
	}
	return l
}

func FolderMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	hierarchyItems := client.hierarchy.FolderRows()

	var l = make([]schema.ClientMeta, len(hierarchyItems))
	for i, item := range hierarchyItems {
		l[i] = client.WithOrganization(item.Organization).WithCloud(item.Cloud).WithFolder(item.Folder).WithMultiplexedResourceId(item.Folder)
	}
	return l
}

func PrependEmptyMultiplex(multiplexer schema.Multiplexer) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		return append([]schema.ClientMeta{client}, multiplexer(meta)...)
	}
}
