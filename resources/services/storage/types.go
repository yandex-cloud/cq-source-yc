package storage

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/storage/v1"
)

type Bucket struct {
	*storage.Bucket
	ServerSideEncryption *s3.ServerSideEncryptionConfiguration
}
