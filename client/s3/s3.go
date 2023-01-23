package s3

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/zerolog"
	"github.com/yandex-cloud/cq-provider-yandex/client/utils"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

const (
	endpoint       = "https://storage.yandexcloud.net"
	region         = "ru-central1"
	iamTokenHeader = "X-YaCloud-SubjectToken"
)

func NewS3Client(creds *credentials.Credentials, sdk *ycsdk.SDK, logger zerolog.Logger) (*s3.S3, error) {
	s := session.Must(session.NewSession())
	config := aws.NewConfig().
		WithEndpoint(endpoint).
		WithRegion(region)

	if logger.GetLevel() == zerolog.DebugLevel {
		config.LogLevel = aws.LogLevel(aws.LogDebugWithHTTPBody)
	}
	// TODO: Maybe separeate into two New... func? Or WithIAM/WithCreds?
	if sdk != nil {
		resp, err := sdk.CreateIAMToken(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("cant't get IAM token for S3 client: %w", err)
		}

		config.HTTPClient = http.DefaultClient
		config.HTTPClient.Transport = utils.NewInterceptTransport(nil, func(req *http.Request) error {
			req.Header.Set(iamTokenHeader, resp.IamToken)
			return nil
		})
		config.Credentials = credentials.AnonymousCredentials

	} else if creds != nil {
		config.Credentials = creds
	}
	client := s3.New(s, config)
	return client, nil
}
