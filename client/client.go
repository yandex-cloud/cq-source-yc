package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"github.com/yandex-cloud/cq-source-yc/client/yc"
	ycsdk "github.com/yandex-cloud/go-sdk"
	ycsdkv2 "github.com/yandex-cloud/go-sdk/v2"
)

const (
	// Exported on purpose to change via `link -X`
	DefaultUserAgent = "cq-source-yc"
)

type Client struct {
	hierarchy *yc.ResourceHierarchy

	OrganizationId          string
	CloudId                 string
	FolderId                string
	MultiplexedResourceId   string
	MultiplexedResourceType ResourceType

	Backend state.Client
	Logger  zerolog.Logger
	SDK     *ycsdk.SDK
	SDKv2   *ycsdkv2.SDK
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
	newC.MultiplexedResourceType = ResourceTypeOrganization
	return &newC
}

func (c *Client) WithCloud(id string) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("cloud", id).Logger()
	newC.CloudId = id
	newC.MultiplexedResourceType = ResourceTypeCloud
	return &newC
}

func (c *Client) WithFolder(id string) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("folder", id).Logger()
	newC.FolderId = id
	newC.MultiplexedResourceType = ResourceTypeFolder
	return &newC
}

func (c *Client) WithMultiplexedResourceId(id string) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("multiplexed", id).Logger()
	newC.MultiplexedResourceId = id
	return &newC
}

func New(ctx context.Context, logger zerolog.Logger, spec *Spec) (*Client, error) {
	sdk, sdkv2, err := yc.Build(ctx, logger, yc.Config{
		Endpoint:   spec.Endpoint,
		UserAgent:  DefaultUserAgent,
		MaxRetries: spec.MaxRetries,
		DebugGRPC:  spec.DebugGRPC,
	})
	if err != nil {
		return nil, err
	}

	client := Client{
		SDK:    sdk,
		SDKv2:  sdkv2,
		Logger: logger,
	}

	hierarchy, err := yc.NewResourceHierarchy(ctx, logger, sdk, spec.OrganizationIDs, spec.CloudIDs, spec.FolderIDs)
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
