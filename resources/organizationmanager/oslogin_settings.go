package organizationmanager

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/yandex-cloud/cq-source-yc/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
)

func OsloginSettings() *schema.Table {
	return &schema.Table{
		Name:        "yc_organizationmanager_oslogin_settings",
		Description: `https://yandex.cloud/ru/docs/organization/api-ref/grpc/OsLogin/getSettings#yandex.cloud.organizationmanager.v1.OsLoginSettings`,
		Multiplex:   client.OrganizationMultiplex,
		Resolver:    fetchOsloginSettings,
		Transform:   client.TransformWithStruct(&organizationmanager.OsLoginSettings{}),
		Columns: schema.ColumnList{
			client.TransformColumnPrimaryKey(client.OrganiztionIdColumn),
		},
	}
}

func fetchOsloginSettings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	orgId := c.OrganizationId

	settings, err := c.SDK.OrganizationManager().OsLogin().GetSettings(ctx, &organizationmanager.GetOsLoginSettingsRequest{OrganizationId: orgId})
	if err != nil {
		return err
	}
	res <- settings

	return nil
}
