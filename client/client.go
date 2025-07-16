package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/helpers/grpczerolog"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/pkg/idempotency"
	"github.com/yandex-cloud/go-sdk/pkg/requestid"
	"github.com/yandex-cloud/go-sdk/pkg/retry/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	// Exported on purpose to change via `link -X`
	DefaultUserAgent = "cq-source-yc"
)

type Client struct {
	hierarchy *ResourceHierarchy

	OrganizationId        string
	CloudId               string
	FolderId              string
	MultiplexedResourceId string

	Backend state.Client
	Logger  zerolog.Logger
	SDK     *ycsdk.SDK
}

func (c *Client) ID() string {
	parts := make([]string, 0)
	if c.OrganizationId != "" {
		parts = append(parts, "org:"+c.OrganizationId)
	}
	if c.CloudId != "" {
		parts = append(parts, "cloud:"+c.CloudId)
	}
	if c.FolderId != "" {
		parts = append(parts, "folder:"+c.FolderId)
	}
	return strings.Join(parts, "|")
}

func (c *Client) WithBackend(backend state.Client) *Client {
	newClient := *c
	newClient.Backend = backend
	return &newClient
}

func (c *Client) WithOrganization(id string) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("organization", id).Logger()
	newC.OrganizationId = id
	return &newC
}

func (c *Client) WithCloud(id string) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("cloud", id).Logger()
	newC.CloudId = id
	return &newC
}

func (c *Client) WithFolder(id string) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("folder", id).Logger()
	newC.FolderId = id
	return &newC
}

func (c *Client) WithMultiplexedResourceId(id string) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("multiplexed", id).Logger()
	newC.MultiplexedResourceId = id
	return &newC
}

func New(ctx context.Context, logger zerolog.Logger, spec *Spec) (*Client, error) {
	credentials, err := getCredentials()
	if err != nil {
		return nil, err
	}

	unaryInterceptors := []grpc.UnaryClientInterceptor{
		requestid.Interceptor(),
		idempotency.Interceptor(),
	}
	streamInterceptors := []grpc.StreamClientInterceptor{}

	var dialOpts = []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(unaryInterceptors...),
		grpc.WithChainStreamInterceptor(streamInterceptors...),
	}

	if spec.MaxRetries > 0 {
		o, err := retry.RetryDialOption(
			retry.WithRetries(retry.DefaultNameConfig(), spec.MaxRetries),
			retry.WithRetryableStatusCodes(retry.DefaultNameConfig(), codes.ResourceExhausted, codes.Unavailable),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create retry option: %w", err)
		}
		dialOpts = append(dialOpts, o)
	}

	// debug interceptors are last
	if spec.DebugGRPC {
		unaryInterceptors = append(unaryInterceptors, logging.UnaryClientInterceptor(grpczerolog.InterceptorLogger(logger)))
		streamInterceptors = append(streamInterceptors, logging.StreamClientInterceptor(grpczerolog.InterceptorLogger(logger)))
	}

	sdk, err := ycsdk.Build(ctx,
		ycsdk.Config{
			Credentials: credentials,
			Endpoint:    spec.Endpoint,
		},
		grpc.WithUserAgent(DefaultUserAgent),
		grpc.WithChainUnaryInterceptor(unaryInterceptors...),
		grpc.WithChainStreamInterceptor(streamInterceptors...),
	)
	if err != nil {
		return nil, fmt.Errorf("initialize Yandex Cloud SDK: %w", err)
	}

	client := Client{
		SDK:    sdk,
		Logger: logger,
	}

	hierarchy, err := NewResourceHierarchy(ctx, logger, sdk, spec.OrganizationIDs, spec.CloudIDs, spec.FolderIDs)
	if err != nil {
		return nil, fmt.Errorf("fetch resource hierarchy: %w", err)
	}

	client.hierarchy = hierarchy

	if len(spec.OrganizationIDs) == 0 && len(spec.CloudIDs) == 0 && len(spec.FolderIDs) == 0 {
		client.Logger.Warn().Msg("no organization_ids, cloud_ids, or folder_ids specified – assuming all resources nested in all orgs")
	}
	client.Logger.Debug().
		Interface("organizations", hierarchy.Organizations()).
		Interface("clouds", hierarchy.Clouds()).
		Interface("folders", hierarchy.Folders()).
		Msg("fetched root resources")
	client.Logger.Info().
		Int("organizations", len(hierarchy.Organizations())).
		Int("clouds", len(hierarchy.Clouds())).
		Int("folders", len(hierarchy.Folders())).
		Msg("fetched root resources")

	return &client, nil
}
