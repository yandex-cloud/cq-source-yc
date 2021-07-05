package client

import (
	"context"
	"fmt"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"os"
	"strings"
	"time"

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
	client := NewYandexClient(logger, folders, services)
	return client, nil
}

func buildSDK() (*ycsdk.SDK, error) {
	ctx := context.Background()
	cred, err := getCredentials()
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

func getCredentials() (ycsdk.Credentials, error) {
	if token, ok := os.LookupEnv("YC_TOKEN"); ok {
		if strings.HasPrefix(token, "t1") {
			return ycsdk.NewIAMTokenCredentials(token), nil
		} else {
			return ycsdk.OAuthToken(token), nil
		}
	}
	if keyFile, ok := os.LookupEnv("YC_SERVICE_ACCOUNT_KEY_FILE"); ok {
		key, err := iamkey.ReadFromJSONFile(keyFile)
		if err != nil {
			return nil, err
		}
		return ycsdk.ServiceAccountKey(key)
	}
	return ycsdk.InstanceServiceAccount(), nil
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

func NewYandexClient(log hclog.Logger, folders []string, services *Services) *Client {
	return &Client{
		logger:   log,
		folders:  folders,
		Services: services,
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
	}
}
