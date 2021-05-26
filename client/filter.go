package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func DeleteFolderFilter(meta schema.ClientMeta) []interface{} {
	client := meta.(*Client)
	return []interface{}{"folder_id", client.FolderId}
}
