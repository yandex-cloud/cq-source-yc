package s3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"gopkg.in/yaml.v3"
)

// GetStaticCredentials generates [credentials.Credentials] using environtment variables:
// - static key (`YC_STORAGE_ACCESS_KEY` + `YC_STORAGE_SECRET_KEY`)
// - yaml file (`YC_SA_STATIC_KEY_FILE`)
func GetStaticCredentials() (*credentials.Credentials, error) {
	id, idOk := os.LookupEnv("YC_STORAGE_ACCESS_KEY")
	secret, secretOk := os.LookupEnv("YC_STORAGE_SECRET_KEY")
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
		return nil, fmt.Errorf("service account static key should be specified")
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
