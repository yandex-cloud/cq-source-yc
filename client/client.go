package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/go-homedir"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"golang.org/x/sync/errgroup"
)

const defaultFolderIdName = "<CHANGE_THIS_TO_YOUR_FOLDER_ID>"

type Client struct {
	folders []string
	clouds  []string

	logger hclog.Logger

	// All yandex services initialized by client
	Services *Services
	// S3 client to manage objects storages
	S3Client *s3.S3

	// this is set by table client multiplexer
	FolderId string
	CloudId  string
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	folders := providerConfig.FolderIDs

	var err error
	sdk, err := buildSDK()
	if err != nil {
		return nil, err
	}

	extractedFolders, err := getFolders(logger, sdk, providerConfig.FolderFilter, providerConfig.CloudIDs)
	if err != nil {
		return nil, err
	}
	folders = unionStrings(folders, extractedFolders)

	if err = validateFolders(folders); err != nil {
		return nil, err
	}

	services, err := initServices(context.Background(), sdk)
	if err != nil {
		return nil, err
	}

	s3Client, err := initS3Clint()
	if err != nil {
		return nil, err
	}

	client := NewYandexClient(logger, folders, providerConfig.CloudIDs, services, s3Client)
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

func unionStrings(strs1, strs2 []string) (res []string) {
	m := map[string]struct{}{}
	for _, s := range strs1 {
		m[s] = struct{}{}
	}
	for _, s := range strs2 {
		m[s] = struct{}{}
	}
	for k := range m {
		res = append(res, k)
	}
	return
}

func getFolders(logger hclog.Logger, sdk *ycsdk.SDK, filter string, cloudIDs []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	ch := make(chan string)

	for _, cloudId := range cloudIDs {
		cloudId := cloudId
		g.Go(func() error {
			req := &resourcemanager.ListFoldersRequest{
				CloudId: cloudId,
				Filter:  filter,
			}
			for {
				resp, err := sdk.ResourceManager().Folder().List(ctx, req)
				if err != nil {
					return err
				}

				for _, folder := range resp.Folders {
					if folder.GetStatus() == resourcemanager.Folder_ACTIVE {
						ch <- folder.Id
					} else {
						logger.Info("Folder state is not active. Folder will be ignored", "folder_id", folder.Id)
					}
				}

				if resp.NextPageToken == "" {
					break
				}

				req.PageToken = resp.NextPageToken
			}
			return nil
		})
	}

	folders := make([]string, 0)
	go func() {
		for folder := range ch {
			folders = append(folders, folder)
		}
	}()

	err := g.Wait()
	close(ch)
	if err != nil {
		return nil, err
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

func NewYandexClient(log hclog.Logger, folders, clouds []string, services *Services, s3Client *s3.S3) *Client {
	return &Client{
		logger:   log,
		folders:  folders,
		clouds:   clouds,
		Services: services,
		S3Client: s3Client,
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// withFolder allows multiplexer to create a new client with given subscriptionId
func (c Client) withFolder(folder string) *Client {
	return &Client{
		folders:  c.folders,
		clouds:   c.clouds,
		Services: c.Services,
		S3Client: c.S3Client,
		logger:   c.logger.With("folder_id", folder),
		FolderId: folder,
	}
}

// withCloud allows multiplexer to create a new client with given subscriptionId
func (c Client) withCloud(cloud string) *Client {
	return &Client{
		folders:  c.folders,
		clouds:   c.clouds,
		Services: c.Services,
		S3Client: c.S3Client,
		logger:   c.logger.With("cloud_id", cloud),
		CloudId:  cloud,
	}
}
