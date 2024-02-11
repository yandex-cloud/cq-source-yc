package storage

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/storage/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name: "yc_storage_buckets",
		Transform: client.TransformWithStruct(
			&storage.Bucket{},
			transformers.WithUnwrapStructFields("Bucket"),
			// Id is always empty ¯\_(ツ)_/¯
			transformers.WithSkipFields("Id"),
			transformers.WithPrimaryKeys("FolderId", "Name"),
		),
		Resolver:  fetchBuckets,
		Multiplex: client.FolderMultiplex,
		Columns: schema.ColumnList{
			client.CloudIdColumn,
		},
	}
}

func fetchBuckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	folderId := c.FolderId

	it := c.SDK.StorageAPI().Bucket().BucketIterator(ctx, &storage.ListBucketsRequest{FolderId: folderId})
	for it.Next() {
		value := it.Value()
		// List doesn't return full info, so we need an additional fetch
		b, err := c.SDK.StorageAPI().Bucket().Get(ctx, &storage.GetBucketRequest{
			Name: value.Name,
			View: storage.GetBucketRequest_VIEW_FULL,
		})
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				return err
			}

			if st.Code() == codes.PermissionDenied {
				// Can't fail here because then we wouldn't get all other buckets
				c.Logger.Warn().Str("bucket", b.Name).Msg("PermissionDenied while fetching bucket")
				continue
			}

			return err
		}

		res <- b
	}

	return it.Error()
}
