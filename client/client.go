package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/yandex-cloud/go-sdk/iamkey"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

const defaultFolderIdName = "<CHANGE_THIS_TO_YOUR_FOLDER_ID>"

type Client struct {
	folders []string
	logger  hclog.Logger
	// All yandex services initialized by client
	Services *Services
	// this is set by table client multiplexer
	FolderId string

	CloudId string
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	folders := providerConfig.FolderIDs
	var err error
	sdk, err := buildSDK()
	if err != nil {
		return nil, err
	}
	if len(providerConfig.FolderIDs) == 0 {
		folders, err = getFolders(logger, sdk, providerConfig.FolderFilter, providerConfig.CloudID)
		if err != nil {
			return nil, err
		}
		logger.Info("No folder_ids specified in config.yml assuming all active folders", "count", len(folders))
	}
	if err := validateFolders(folders); err != nil {
		return nil, err
	}
	services, err := initServices(context.Background(), sdk)
	if err != nil {
		return nil, err
	}
	client := NewYandexClient(logger, folders, services, providerConfig.CloudID)
	return client, nil
}

func buildSDK() (*ycsdk.SDK, error) {
	ctx := context.Background()
	cred, err := credentials()
	if err != nil {
		return nil, err
	}
	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: cred,
	})
	if err != nil {
		return nil, err
	}
	return sdk, err
}

func credentials() (ycsdk.Credentials, error) {
	if val, ok := os.LookupEnv("YC_SERVICE_ACCOUNT_KEY_FILE"); ok {
		contents, _, err := pathOrContents(val)
		if err != nil {
			return nil, fmt.Errorf("Error loading credentials: %s", err)
		}

		key, err := iamKeyFromJSONContent(contents)
		if err != nil {
			return nil, err
		}
		return ycsdk.ServiceAccountKey(key)
	}

	if val, ok := os.LookupEnv("YC_TOKEN"); ok {
		if strings.HasPrefix(val, "t1.") && strings.Count(val, ".") == 2 {
			return ycsdk.NewIAMTokenCredentials(val), nil
		}
		return ycsdk.OAuthToken(val), nil
	}

	return ycsdk.InstanceServiceAccount(), nil
}

// copy of github.com/hashicorp/terraform-plugin-sdk/helper/pathorcontents.Read()
func pathOrContents(poc string) (string, bool, error) {
	if len(poc) == 0 {
		return poc, false, nil
	}

	path := poc
	if path[0] == '~' {
		var err error
		path, err = homedir.Expand(path)
		if err != nil {
			return path, true, err
		}
	}

	if _, err := os.Stat(path); err == nil {
		contents, err := ioutil.ReadFile(path)
		return string(contents), true, err
	}

	return poc, false, nil
}

func iamKeyFromJSONContent(content string) (*iamkey.Key, error) {
	key := &iamkey.Key{}
	err := json.Unmarshal([]byte(content), key)
	if err != nil {
		return nil, fmt.Errorf("key unmarshal fail: %s", err)
	}
	return key, nil
}

func getFolders(logger hclog.Logger, sdk *ycsdk.SDK, filter string, cloudID string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	req := &resourcemanager.ListFoldersRequest{
		CloudId: cloudID,
		Filter:  filter,
	}

	folders := make([]string, 0)
	for {
		resp, err := sdk.ResourceManager().Folder().List(ctx, req)

		if err != nil {
			return nil, err
		}
		for _, folder := range resp.Folders {
			if folder.GetStatus() == resourcemanager.Folder_ACTIVE {
				folders = append(folders, folder.Id)
			} else {
				logger.Info("Folder state is not active. Folder will be ignored", "folder_id", folder.Id)
			}
		}
		if resp.NextPageToken == "" {
			break
		}
		req.PageToken = resp.NextPageToken
	}
	return folders, nil
}

func validateFolders(folders []string) error {
	for _, folder := range folders {
		if folder == defaultFolderIdName {
			return fmt.Errorf("please specify a valid folder_id in config.yml instead of %s", defaultFolderIdName)
		}
	}
	return nil
}

func NewYandexClient(log hclog.Logger, folders []string, services *Services, cloudId string) *Client {
	return &Client{
		logger:   log,
		folders:  folders,
		Services: services,
		CloudId:  cloudId,
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// withFolder allows multiplexer to create a new client with given subscriptionId
func (c Client) withFolder(folder string) *Client {
	return &Client{
		folders:  c.folders,
		Services: c.Services,
		logger:   c.logger.With("folder_id", folder),
		FolderId: folder,
		CloudId:  c.CloudId,
	}
}
