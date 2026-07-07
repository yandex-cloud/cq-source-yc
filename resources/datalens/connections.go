package datalens

import (
	"context"
	"errors"
	"maps"
	"net/http"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/cq-source-yc/client/yc/datalens"
)

func Connections() *schema.Table {
	t := entriesTable("yc_datalens_connections", "connection")
	t.Description = `DataLens connections, enriched with getConnection details (host, port, ... in the data column). https://yandex.cloud/ru/docs/datalens/operations/api-start`
	t.PostResourceResolver = resolveConnectionDetails
	return t
}

func resolveConnectionDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)

	entry := resource.Item.(*datalens.Entry)
	if entry.IsLocked != nil && bool(*entry.IsLocked) {
		return nil
	}

	conn, err := c.Datalens.GetConnection(ctx, c.OrganizationId, entry.EntryId)
	if err != nil {
		var apiErr *datalens.APIError
		if errors.As(err, &apiErr) && (apiErr.StatusCode == http.StatusForbidden || apiErr.StatusCode == http.StatusNotFound) {
			// Can't fail here because then we wouldn't get all other connections
			c.Logger.Warn().Str("connection", entry.EntryId).Int("status", apiErr.StatusCode).Msg("skipping DataLens connection details")
			return nil
		}
		return err
	}

	data := make(map[string]any, len(conn.Data))
	if entry.Data != nil {
		maps.Copy(data, *entry.Data)
	}
	maps.Copy(data, conn.Data)
	return resource.Set("data", data)
}
