package storage

import (
	"context"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/storage/v1"
)

func fetchBuckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	s3Client, err := c.GetS3Client()
	if err != nil {
		return err
	}

	it := c.Services.Storage.Bucket().BucketIterator(ctx, &storage.ListBucketsRequest{FolderId: c.MultiplexedResourceId})
	for it.Next() {
		bucket := Bucket{Bucket: it.Value()}
		encryptResp, err := s3Client.GetBucketEncryptionWithContext(ctx, &s3.GetBucketEncryptionInput{Bucket: &bucket.Name})
		if err != nil {
			return err
		}
		if encryptResp != nil {
			bucket.ServerSideEncryption = encryptResp.ServerSideEncryptionConfiguration
		}
		res <- bucket
	}
	return nil
}
