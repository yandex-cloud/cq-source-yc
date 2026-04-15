package client

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
	credentialsv2 "github.com/yandex-cloud/go-sdk/v2/credentials"
	iamkeyv2 "github.com/yandex-cloud/go-sdk/v2/pkg/iamkey"
)

func iamKeyFromJSONContent(content string) (*iamkey.Key, error) {
	key := &iamkey.Key{}
	err := json.Unmarshal([]byte(content), key)
	if err != nil {
		return nil, fmt.Errorf("key unmarshal: %s", err)
	}
	return key, nil
}

func getCredentials() (ycsdk.Credentials, error) {
	if val := os.Getenv("YC_SERVICE_ACCOUNT_KEY"); val != "" {
		key, err := iamKeyFromJSONContent(val)
		if err != nil {
			return nil, err
		}
		return ycsdk.ServiceAccountKey(key)
	}

	if val := os.Getenv("YC_TOKEN"); val != "" {
		if strings.HasPrefix(val, "t1.") && strings.Count(val, ".") == 2 {
			return ycsdk.NewIAMTokenCredentials(val), nil
		}
		return ycsdk.OAuthToken(val), nil
	}

	return ycsdk.InstanceServiceAccount(), nil
}

func getCredentialsV2() (credentialsv2.Credentials, error) {
	if val := os.Getenv("YC_SERVICE_ACCOUNT_KEY"); val != "" {
		key, err := iamkeyv2.ReadFromJSONBytes([]byte(val))
		if err != nil {
			return nil, err
		}
		return credentialsv2.ServiceAccountKey(key)
	}

	if val := os.Getenv("YC_TOKEN"); val != "" {
		if strings.HasPrefix(val, "t1.") && strings.Count(val, ".") == 2 {
			return credentialsv2.IAMToken(val), nil
		}
		return credentialsv2.OAuthToken(val), nil
	}

	return credentialsv2.InstanceServiceAccount(), nil
}
