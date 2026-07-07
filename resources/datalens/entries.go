package datalens

import (
	"context"
	"errors"
	"net/http"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/client/yc/datalens"
)

func entriesTable(name, scope string) *schema.Table {
	return &schema.Table{
		Name:        name,
		Description: `DataLens entries with scope=` + scope + `. https://yandex.cloud/ru/docs/datalens/operations/api-start`,
		Multiplex:   client.OrganizationMultiplex,
		Resolver:    fetchEntries(scope),
		Transform:   client.TransformWithStruct(&datalens.Entry{}, transformers.WithPrimaryKeys("EntryId")),
		Columns: schema.ColumnList{
			client.OrganiztionIdColumn,
		},
	}
}

func fetchEntries(scope string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)

		pageSize := 200 // API maximum
		includeAll := true
		args := datalens.GetEntriesV2Args{
			Scope:                  &scope,
			PageSize:               &pageSize,
			IncludeData:            &includeAll,
			IncludeLinks:           &includeAll,
			IncludePermissionsInfo: &includeAll,
		}
		for {
			entries, nextPageToken, err := c.Datalens.GetEntries(ctx, c.OrganizationId, args)
			if err != nil {
				var apiErr *datalens.APIError
				if errors.As(err, &apiErr) && (apiErr.StatusCode == http.StatusForbidden || apiErr.StatusCode == http.StatusNotFound) {
					// The organization has no DataLens instance or we have no access to it
					c.Logger.Warn().Int("status", apiErr.StatusCode).Str("scope", scope).Msg("skipping DataLens entries")
					return nil
				}
				return err
			}

			for i := range entries {
				res <- &entries[i]
			}

			if nextPageToken == "" {
				return nil
			}
			args.PageToken = &nextPageToken
		}
	}
}

func Folders() *schema.Table {
	return entriesTable("yc_datalens_folders", "folder")
}

func Datasets() *schema.Table {
	return entriesTable("yc_datalens_datasets", "dataset")
}

func Dashboards() *schema.Table {
	return entriesTable("yc_datalens_dashboards", "dash")
}

func Widgets() *schema.Table {
	return entriesTable("yc_datalens_widgets", "widget")
}

func Reports() *schema.Table {
	return entriesTable("yc_datalens_reports", "report")
}
