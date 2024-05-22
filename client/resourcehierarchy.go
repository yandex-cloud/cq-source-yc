package client

import (
	"context"
	"fmt"
	"slices"
	"sort"

	"github.com/rs/zerolog"
	"github.com/yandex-cloud/cq-source-yc/internal/util"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"google.golang.org/grpc"
)

// Self-made name for struct holding info about resource model hierarchy.
// Each field could be nil.
type ResourceHierarchyItem struct {
	Organization string
	Cloud        string
	Folder       string
}

type ResourceHierarchy struct {
	items []ResourceHierarchyItem
}

func (h *ResourceHierarchy) All() []ResourceHierarchyItem {
	return h.items
}

func (h *ResourceHierarchy) Organizations() []string {
	result := make(map[string]bool)
	for _, item := range h.items {
		if item.Organization != "" {
			result[item.Organization] = true
		}
	}
	return util.SetToSlice(result)
}

func (h *ResourceHierarchy) Clouds() []string {
	result := make(map[string]bool)
	for _, item := range h.items {
		if item.Cloud != "" {
			result[item.Cloud] = true
		}
	}
	return util.SetToSlice(result)
}

func (h *ResourceHierarchy) Folders() []string {
	result := make(map[string]bool)
	for _, item := range h.items {
		if item.Folder != "" {
			result[item.Folder] = true
		}
	}
	return util.SetToSlice(result)
}

// OrganizationRows returns all ResourceHierarchyItems with Organization != ""
func (h *ResourceHierarchy) OrganizationRows() []ResourceHierarchyItem {
	return util.Filter(h.items, func(item ResourceHierarchyItem) bool {
		return item.Organization != "" && item.Cloud == "" && item.Folder == ""
	})
}

// CloudRows returns all ResourceHierarchyItems with Cloud != ""
func (h *ResourceHierarchy) CloudRows() []ResourceHierarchyItem {
	return util.Filter(h.items, func(item ResourceHierarchyItem) bool {
		return item.Cloud != "" && item.Folder == ""
	})
}

// FolderRows returns all ResourceHierarchyItems with Folder != ""
func (h *ResourceHierarchy) FolderRows() []ResourceHierarchyItem {
	return util.Filter(h.items, func(item ResourceHierarchyItem) bool {
		return item.Folder != ""
	})
}

const (
	serviceResourceManager     = "resource-manager"
	serviceOrganizationManager = "organization-manager"
)

// discover the hierarchy using Breadth-first search
func bfs(ctx context.Context, sdk *ycsdk.SDK, logger zerolog.Logger, init []ResourceHierarchyItem, organizations []string, clouds []string, folders []string, opts ...grpc.CallOption) ([]ResourceHierarchyItem, error) {
	// call ApiEndpoint.List to fill internal endpoint list in ycsdk.SDK
	_, err := sdk.ApiEndpoint().ApiEndpoint().List(ctx, &endpoint.ListApiEndpointsRequest{})
	if err != nil {
		return nil, fmt.Errorf("discover endpoints: %w", err)
	}
	services := util.SliceToSet(sdk.KnownServices())
	logger.Debug().Interface("services", services).Msg("available services")

	items := make([]ResourceHierarchyItem, 0)

	organizationsSet := util.SliceToSet(organizations)
	cloudsSet := util.SliceToSet(clouds)
	foldersSet := util.SliceToSet(folders)

	queue := []ResourceHierarchyItem{}
	if len(init) == 0 {
		queue = append(queue, ResourceHierarchyItem{})
	} else {
		queue = append(queue, init...)
	}

	for len(queue) > 0 {
		item := queue[0]
		logger.Debug().Interface("item", item).Msg("bfs iteration")
		// always append items
		items = append(items, item)
		queue = queue[1:]

		// for the sake of readability
		if item.Folder != "" {
			// just append to items
			continue
		} else if item.Cloud != "" && services[serviceResourceManager] {
			// discover folders
			it := sdk.ResourceManager().Folder().FolderIterator(ctx, &resourcemanager.ListFoldersRequest{CloudId: item.Cloud}, opts...)
			for it.Next() {
				folder := it.Value()
				if len(folders) == 0 || foldersSet[folder.Id] {
					// org -> cloud -> folder
					queue = append(queue, ResourceHierarchyItem{Organization: item.Organization, Cloud: item.Cloud, Folder: folder.Id})
				} else {
					logger.Warn().Str("id", folder.Id).Msg("Skipping folder with id")
				}
			}
			if err := it.Error(); err != nil {
				return nil, err
			}
		} else if item.Organization != "" && services[serviceResourceManager] {
			// discover clouds
			it := sdk.ResourceManager().Cloud().CloudIterator(ctx, &resourcemanager.ListCloudsRequest{OrganizationId: item.Organization}, opts...)
			for it.Next() {
				cloud := it.Value()
				if len(clouds) == 0 || cloudsSet[cloud.Id] {
					// org -> cloud -> nil
					queue = append(queue, ResourceHierarchyItem{Organization: item.Organization, Cloud: cloud.Id})
				} else {
					logger.Warn().Str("id", cloud.Id).Msg("Skipping cloud with id")
				}
			}
			if err := it.Error(); err != nil {
				return nil, err
			}
		} else if services[serviceOrganizationManager] {
			// discover organizations
			it := sdk.OrganizationManager().Organization().OrganizationIterator(ctx, &organizationmanager.ListOrganizationsRequest{}, opts...)
			for it.Next() {
				org := it.Value()
				if len(organizations) == 0 || organizationsSet[org.Id] {
					// org -> nil -> nil
					queue = append(queue, ResourceHierarchyItem{Organization: org.Id})
				} else {
					logger.Warn().Str("id", org.Id).Msg("Skipping organization with id")
				}
			}
			if err := it.Error(); err != nil {
				return nil, err
			}
		} else {
			// query Clouds without parent Organization
			logger.Warn().Msg("organization-manager unavailable, resolving clouds")
			it := sdk.ResourceManager().Cloud().CloudIterator(ctx, &resourcemanager.ListCloudsRequest{}, opts...)
			for it.Next() {
				cloud := it.Value()
				if len(clouds) == 0 || cloudsSet[cloud.Id] {
					// org -> nil -> nil
					queue = append(queue, ResourceHierarchyItem{Cloud: cloud.Id})
				} else {
					logger.Warn().Str("id", cloud.Id).Msg("Skipping cloud with id")
				}
			}
			if err := it.Error(); err != nil {
				return nil, err
			}
		}
	}

	return items, nil
}

// NewResourceHirarchy fetches all reachable resourse model entities and creates ResourceHirarchy struct
// while holding the hierarchy (i.e. organization -> cloud -> folder).
// Some properties of the hierarchy:
//   - the hierarchy could be partially incomplete, e.g. nil -> cloud_id1 -> folder_id1.
//   - likewise both entries must exist for scoped resolution (e.g. Cloud Access Bindings):
//     cloud_id1 -> nil        (resolve tables for Cloud)
//     cloud_id1 -> folder_id1 (resolve tables for Folder)
//
// Discover algorithm is:
//  1. Recursively discover resources using Breadth-first search on resources returned by calling `List` method on Organizations, Clouds, Folders
//  2. Recursively discover resources using Breadth-first search on resources provided in arguments (`organizations`, `clouds`, `folders`)
//  3. Merge and deduplicate results from (1) and (2)
//  4. Apply filters from arguments (`organizations`, `clouds`, `folders`)
//
// Step 2 is needed because API does not return all Organizations, which caller account has access to, but only Organization, which caller account belong to
func NewResourceHierarchy(ctx context.Context, logger zerolog.Logger, sdk *ycsdk.SDK, organizations []string, clouds []string, folders []string, opts ...grpc.CallOption) (*ResourceHierarchy, error) {
	var items []ResourceHierarchyItem

	// discover resources based on API knowledge
	itemsA, err := bfs(ctx, sdk, logger, nil, organizations, clouds, folders, opts...)
	if err != nil {
		return nil, err
	}
	logger.Debug().Int("count", len(itemsA)).Msg("discovered resources using List calls")

	// discover resources using provided ids
	init := make([]ResourceHierarchyItem, 0)
	for _, org := range organizations {
		init = append(init, ResourceHierarchyItem{Organization: org})
	}
	for _, cloud := range clouds {
		init = append(init, ResourceHierarchyItem{Cloud: cloud})
	}
	for _, folder := range folders {
		init = append(init, ResourceHierarchyItem{Folder: folder})
	}
	itemsB, err := bfs(ctx, sdk, logger, init, organizations, clouds, folders, opts...)
	if err != nil {
		return nil, err
	}
	logger.Debug().Int("count", len(itemsB)).Msg("discovered resources using List calls on provided ids")

	// merge & deduplicate
	items = append(itemsA, itemsB...)
	sort.Slice(items, func(i, j int) bool {
		return (items[i].Organization < items[j].Organization) &&
			(items[i].Cloud < items[j].Cloud) &&
			(items[i].Folder < items[j].Folder)
	})
	items = slices.Compact(items)

	return &ResourceHierarchy{
		items: items,
	}, nil
}
