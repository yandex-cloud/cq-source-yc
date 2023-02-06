package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	ycs3 "github.com/yandex-cloud/cq-provider-yandex/client/s3"
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
	orgs    []string
	clouds  []string
	folders []string

	logger zerolog.Logger

	// YC SDK
	sdk *ycsdk.SDK

	// S3 client to manage objects storages
	s3 *s3.S3

	// All yandex services initialized by client
	Services *Services

	// this is set by table client multiplexer
	MultiplexedResourceId string
}

func (c *Client) ID() string {
	return c.MultiplexedResourceId
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) S3() *s3.S3 {
	return c.s3
}

func (c *Client) withResource(id string) *Client {
	return &Client{
		orgs:                  c.orgs,
		folders:               c.folders,
		clouds:                c.clouds,
		Services:              c.Services,
		s3:                    c.s3,
		logger:                c.logger.With().Str("id", id).Logger(),
		MultiplexedResourceId: id,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var spec Spec
	err := s.UnmarshalSpec(&spec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Yandex Cloud spec: %w", err)
	}

	if spec.Endpoint == "" {
		spec.Endpoint = defaultEndpoint
	}

	sdk, err := buildSDK(spec.Endpoint)
	if err != nil {
		return nil, err
	}

	extractedClouds, err := getClouds(sdk, spec.OrganizationIDs)
	if err != nil {
		return nil, err
	}
	clouds := unionStrings(spec.CloudIDs, extractedClouds)

	extractedFolders, err := getFolders(logger, sdk, spec.FolderFilter, clouds)
	if err != nil {
		return nil, err
	}
	folders := unionStrings(spec.FolderIDs, extractedFolders)

	if err = validateFolders(folders); err != nil {
		return nil, err
	}

	services, err := initServices(ctx, sdk)
	if err != nil {
		return nil, err
	}

	s3Client, err := initS3(spec.UseIAMForStorage, sdk, logger)
	if err != nil {
		return nil, err
	}

	client := NewYandexClient(logger, sdk, s3Client, folders, clouds, spec.OrganizationIDs, services)
	return client, nil
}

func initS3(useIAM bool, sdk *ycsdk.SDK, logger zerolog.Logger) (*s3.S3, error) {
	if useIAM {
		s3Client, err := ycs3.NewS3Client(nil, sdk, logger)
		if err != nil {
			return nil, err
		}
		return s3Client, nil
	} else {
		creds, err := ycs3.GetStaticCredentials()
		if err != nil {
			return nil, err
		}
		s3Client, err := ycs3.NewS3Client(creds, nil, logger)
		if err != nil {
			return nil, err
		}
		return s3Client, nil
	}
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

func getFolders(logger zerolog.Logger, sdk *ycsdk.SDK, filter string, cloudIds []string) ([]string, error) {
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
						logger.Info().
							Str("folder_id", folder.Id).
							Msg("Folder state is not active. Folder will be ignored")
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

func NewYandexClient(logger zerolog.Logger, sdk *ycsdk.SDK, s3Client *s3.S3, folders, clouds, organizations []string, services *Services) *Client {
	return &Client{
		logger:   logger,
		sdk:      sdk,
		orgs:     organizations,
		folders:  folders,
		clouds:   clouds,
		Services: services,
		s3:       s3Client,
	}
}
