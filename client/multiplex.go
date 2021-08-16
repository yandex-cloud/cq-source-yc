package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func MultiplexBy(resourcesGetter func(client *Client) []string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for _, id := range resourcesGetter(client) {
			l = append(l, client.withResource(id))
		}
		return l
	}
}

func Organizations(client *Client) []string {
	return client.organizations
}

func Clouds(client *Client) []string {
	return client.clouds
}

func Folders(client *Client) []string {
	return client.folders
}

func EmptyMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	return []schema.ClientMeta{meta}
}
