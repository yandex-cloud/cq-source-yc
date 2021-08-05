package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func DeleteFolderFilter(meta schema.ClientMeta, _ *schema.Resource) []interface{} {
	client := meta.(*Client)
	return []interface{}{"folder_id", client.FolderId}
}
