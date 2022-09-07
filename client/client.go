package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/go-homedir"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const defaultEndpoint = "api.cloud.yandex.net:443"
const defaultFolderIdName = "<CHANGE_THIS_TO_YOUR_FOLDER_ID>"

type AccessBindingsLister interface {
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
}

type Client struct {
	organizations []string
	clouds        []string
	folders       []string

	logger hclog.Logger

	// All yandex services initialized by client
	Services *Services
	// S3 client to manage objects storages
	s3Client *s3.S3

	// this is set by table client multiplexer
	MultiplexedResourceId string
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

func (c Client) withResource(id string) *Client {
	return &Client{
		organizations:         c.organizations,
		folders:               c.folders,
		clouds:                c.clouds,
		Services:              c.Services,
		s3Client:              c.s3Client,
		logger:                c.logger.With("id", id),
		MultiplexedResourceId: id,
	}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)
	if providerConfig.Endpoint == "" {
		providerConfig.Endpoint = defaultEndpoint
	}

	clouds := providerConfig.CloudIDs
	folders := providerConfig.FolderIDs

	var err error
	sdk, err := buildSDK(providerConfig.Endpoint)
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}

	extractedClouds, err := getClouds(sdk, providerConfig.OrganizationIDs)
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	clouds = unionStrings(clouds, extractedClouds)

	extractedFolders, err := getFolders(logger, sdk, providerConfig.FolderFilter, clouds)
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	folders = unionStrings(folders, extractedFolders)

	if err = validateFolders(folders); err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}

	services, err := initServices(context.Background(), sdk)
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}

	client := NewYandexClient(logger, folders, clouds, providerConfig.OrganizationIDs, services, nil)
	return client, nil
}

func (c *Client) GetS3Client(ctx context.Context) (*s3.S3, error) {
	if c.s3Client != nil {
		return c.s3Client, nil
	}
	s3Client, err := initS3Clint()
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	c.s3Client = s3Client
	return c.s3Client, nil
}

func buildSDK(endpoint string) (*ycsdk.SDK, error) {
	ctx := context.Background()
	cred, err := getCredentials()
	if err != nil {
		return nil, err
	}
	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: cred,
		Endpoint:    endpoint,
	})
	if err != nil {
		return nil, err
	}
	return sdk, nil
}

func getCredentials() (ycsdk.Credentials, error) {
	if val, ok := os.LookupEnv("YC_SERVICE_ACCOUNT_KEY_FILE"); ok {
		contents, _, err := pathOrContents(val)
		if err != nil {
			return nil, fmt.Errorf("error loading credentials: %s", err)
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
		contents, err := os.ReadFile(path)
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

func getClouds(sdk *ycsdk.SDK, organizationsIds []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	g := errgroup.Group{}
	ch := make(chan string)

	for _, organizationsId := range organizationsIds {
		finalOrganizationId := organizationsId
		g.Go(func() error {
			req := &resourcemanager.ListCloudsRequest{}
			for {
				resp, err := sdk.ResourceManager().Cloud().List(ctx, req)
				if err != nil {
					return err
				}

				for _, cloud := range resp.Clouds {
					if cloud.OrganizationId == finalOrganizationId {
						ch <- cloud.Id
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

func getFolders(logger hclog.Logger, sdk *ycsdk.SDK, filter string, cloudIds []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	g := errgroup.Group{}
	ch := make(chan string)

	for _, cloudId := range cloudIds {
		finalCloudId := cloudId
		g.Go(func() error {
			req := &resourcemanager.ListFoldersRequest{
				CloudId: finalCloudId,
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

func validateFolders(folders []string) error {
	for _, folder := range folders {
		if folder == defaultFolderIdName {
			return fmt.Errorf("please specify a valid folder_id in config.yml instead of %s", defaultFolderIdName)
		}
	}
	return nil
}

func NewYandexClient(log hclog.Logger, folders, clouds, organizations []string, services *Services, s3Client *s3.S3) *Client {
	return &Client{
		logger:        log,
		organizations: organizations,
		folders:       folders,
		clouds:        clouds,
		Services:      services,
		s3Client:      s3Client,
	}
}
