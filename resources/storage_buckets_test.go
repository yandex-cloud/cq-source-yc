package resources_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func TestStorageBuckets(t *testing.T) {
	cancel, err := createBucketsServer()
	defer cancel()
	if err != nil {
		t.Fatal(err)
	}

	s := session.Must(session.NewSession())
	resource := providertest.ResourceTestData{
		Table:  resources.StorageBuckets(),
		Config: client.Config{},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), nil, nil, nil, nil, s3.New(s, aws.NewConfig().
				WithEndpoint("http://localhost:9000").
				WithRegion("us-east-1").
				WithCredentials(credentials.NewStaticCredentials("user", "12345678", ""))))
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_storage_buckets"),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}

func createBucketsServer() (func(), error) {
	dockerRunCmd := exec.Command("docker", "run",
		"--rm", "-d",
		"--name", "cq_provider_yandex_s3",
		"-p", "9000:9000",
		"-e", "MINIO_ROOT_USER=user",
		"-e", "MINIO_ROOT_PASSWORD=12345678",
		"minio/minio",
		"server", "/data")
	cancelCmd := func() {
		err := exec.Command("docker", "rm", "-f", "cq_provider_yandex_s3").Run()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
	err := dockerRunCmd.Run()
	if err != nil {
		return cancelCmd, err
	}

	awsCmd := exec.Command("aws", "--endpoint=http://localhost:9000", "s3", "mb", "s3://cq-test-bucket")
	awsCmd.Env = append(awsCmd.Env, "AWS_ACCESS_KEY_ID=user", "AWS_SECRET_ACCESS_KEY=12345678")
	err = awsCmd.Run()
	if err != nil {
		return cancelCmd, err
	}

	return cancelCmd, nil
}
