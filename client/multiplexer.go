package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// available reports whether a service can be reached in the installation this
// client is connected to, logging the tables we skip because it cannot.
func (c *Client) available(service Service) bool {
	if c.serviceAvailable(service) {
		return true
	}
	c.Logger.Info().
		Str("service", string(service)).
		Str("region", string(c.Region)).
		Msg("skipping tables: this installation does not offer the service")
	return false
}

// GlobalMultiplex is for tables whose RPC takes no folder, cloud or
// organization: it keeps the single unscoped client the scheduler would use
// anyway, but drops it in regions that do not have the service.
func GlobalMultiplex(service Service) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		if !client.available(service) {
			return nil
		}
		return []schema.ClientMeta{client}
	}
}

func OrganizationMultiplex(service Service) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		if !client.available(service) {
			return nil
		}
		hierarchyItems := client.hierarchy.OrganizationRows()

		var l = make([]schema.ClientMeta, len(hierarchyItems))
		for i, item := range hierarchyItems {
			l[i] = client.WithOrganization(item.Organization).WithMultiplexedResourceId(item.Organization)
		}
		return l
	}
}

func CloudMultiplex(service Service) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		if !client.available(service) {
			return nil
		}
		hierarchyItems := client.hierarchy.CloudRows()

		var l = make([]schema.ClientMeta, len(hierarchyItems))
		for i, item := range hierarchyItems {
			l[i] = client.WithOrganization(item.Organization).WithCloud(item.Cloud).WithMultiplexedResourceId(item.Cloud)
		}
		return l
	}
}

func FolderMultiplex(service Service) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)
		if !client.available(service) {
			return nil
		}
		hierarchyItems := client.hierarchy.FolderRows()

		var l = make([]schema.ClientMeta, len(hierarchyItems))
		for i, item := range hierarchyItems {
			l[i] = client.WithOrganization(item.Organization).WithCloud(item.Cloud).WithFolder(item.Folder).WithMultiplexedResourceId(item.Folder)
		}
		return l
	}
}

func CombineMultiplex(multiplexers ...schema.Multiplexer) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		clients := []schema.ClientMeta{}
		for _, m := range multiplexers {
			clients = append(clients, m(meta)...)
		}
		return clients
	}
}
