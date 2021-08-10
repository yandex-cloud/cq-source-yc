package client

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gopkg.in/yaml.v3"
)

const (
	endpoint = "https://storage.yandexcloud.net"
	region   = "ru-central1"
)

func initS3Clint() (*s3.S3, error) {
	s := session.Must(session.NewSession())
	creds, err := getS3StaticCredentials()
	if err != nil {
		return nil, err
	}
	client := s3.New(s, aws.NewConfig().
		WithEndpoint(endpoint).
		WithRegion(region).
		WithCredentials(creds),
	)
	return client, nil
}

func getS3StaticCredentials() (*credentials.Credentials, error) {
	id, idOk := os.LookupEnv("YC_SA_STATIC_KEY_ID")
	secret, secretOk := os.LookupEnv("YC_SA_STATIC_SECRET")
	filename, filenameOk := os.LookupEnv("YC_SA_STATIC_KEY_FILE")

	switch {
	case filenameOk:
		id, secret, err := loadCredentialsFromYaml(filename)
		if err != nil {
			return nil, err
		}
		return credentials.NewStaticCredentials(id, secret, ""), nil
	case idOk && secretOk:
		return credentials.NewStaticCredentials(id, secret, ""), nil
	default:
		return nil, fmt.Errorf("service account static key should be specfied")
	}
}

type staticKey struct {
	AccessKey staticKeyAccessKey `yaml:"access_key"`
	Secret    string             `yaml:"secret"`
}

type staticKeyAccessKey struct {
	KeyId string `yaml:"key_id"`
}

func loadCredentialsFromYaml(filename string) (string, string, error) {
	dataBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", "", err
	}

	var data staticKey
	err = yaml.Unmarshal(dataBytes, &data)
	if err != nil {
		return "", "", err
	}

	return data.AccessKey.KeyId, data.Secret, nil
}
