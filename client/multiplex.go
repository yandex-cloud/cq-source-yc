package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func FolderMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, folderId := range client.folders {
		l = append(l, client.withFolder(folderId))
	}
	return l
}

func IdentityMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	return []schema.ClientMeta{meta}
}
