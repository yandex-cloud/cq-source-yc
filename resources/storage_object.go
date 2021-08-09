package resources

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func StorageObjects() *schema.Table {
	return &schema.Table{
		Name:        "yandex_storage_objects",
		Resolver:    fetchStorageObjects,
		Multiplex:   client.IdentityMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Description: "name (id) of the bucket.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:     "encryption_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EncryptionEnable"),
			},
		},
	}
}

const (
	endpoint        = "https://storage.yandexcloud.net"
	defaultS3Region = "ru-central1"
)

type StorageObject struct {
	Name             string
	EncryptionEnable bool
}

func fetchStorageObjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	s := session.Must(session.NewSession())
	creds := credentials.NewStaticCredentials("yph13j-Rs7sMMdngIleZ", "iDJo5FInvtfFkJ6yJCOszV5iZeXaJNDM_C2Dvx_S", "")
	c := s3.New(s, aws.NewConfig().
		WithEndpoint(endpoint).
		WithRegion(defaultS3Region).
		WithCredentials(creds),
	)
	resp, err := c.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}
	for _, value := range resp.Buckets {
		_, err := c.GetBucketEncryption(&s3.GetBucketEncryptionInput{
			Bucket: value.Name,
		})
		res <- StorageObject{*value.Name, err == nil}
	}
	return nil
}
