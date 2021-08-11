package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

// TODO: optimize code

func FolderMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, folderId := range client.folders {
		l = append(l, client.withFolder(folderId))
	}
	return l
}

func CloudMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, cloudId := range client.clouds {
		l = append(l, client.withCloud(cloudId))
	}
	return l
}

func FolderAndCloudMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	clients := FolderMultiplex(meta)
	clients = append(clients, CloudMultiplex(meta)...)
	return clients
}

func IdentityMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	return []schema.ClientMeta{meta}
}
