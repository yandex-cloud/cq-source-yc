package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"golang.org/x/sync/errgroup"
)

func IAMUserAccountsByClouds() *schema.Table {
	return &schema.Table{
		Name:        "yandex_iam_user_accounts_by_cloud",
		Resolver:    fetchIAMUserAccountsByCloud,
		Multiplex:   client.MultiplexBy(client.Clouds),
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeString,
				Description:     "ID of the user_account.",
				Resolver:        client.ResolveResourceId,
				CreationOptions: schema.ColumnCreationOptions{Nullable: false, Unique: true},
			},
			{
				Name:        "user_account_yandex_passport_user_account_login",
				Type:        schema.TypeString,
				Description: "Login of the Yandex.Passport user account.",
				Resolver:    schema.PathResolver("UserAccount.YandexPassportUserAccount.Login"),
			},
			{
				Name:        "user_account_yandex_passport_user_account_default_email",
				Type:        schema.TypeString,
				Description: "Default email of the Yandex.Passport user account.",
				Resolver:    schema.PathResolver("UserAccount.YandexPassportUserAccount.DefaultEmail"),
			},
			{
				Name:        "user_account_saml_user_account_federation_id",
				Type:        schema.TypeString,
				Description: "ID of the federation that the federation belongs to.",
				Resolver:    schema.PathResolver("UserAccount.SamlUserAccount.FederationId"),
			},
			{
				Name:        "user_account_saml_user_account_name_id",
				Type:        schema.TypeString,
				Description: "Name Id of the SAML federated user.\n The name is unique within the federation. 1-256 characters long.",
				Resolver:    schema.PathResolver("UserAccount.SamlUserAccount.NameId"),
			},
			{
				Name:        "user_account_saml_user_account_attributes",
				Type:        schema.TypeJSON,
				Description: "Additional attributes of the SAML federated user.",
				Resolver:    schema.PathResolver("UserAccount.SamlUserAccount.Attributes"),
			},
		},
	}
}

func fetchIAMUserAccountsByCloud(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	g := errgroup.Group{}
	ch := make(chan interface{})

	g.Go(func() error {
		defer close(ch)
		return fetchAccessBindingsByCloud(ctx, meta, parent, ch)
	})

	g.Go(func() error {
		for value := range ch {
			id := value.(CloudAccessBinding).AccessBinding.Subject.Id
			accountType := value.(CloudAccessBinding).AccessBinding.Subject.Type
			if accountType != "serviceAccount" {
				req := &iam.GetUserAccountRequest{UserAccountId: id}
				userAccount, err := c.Services.IAM.UserAccount().Get(ctx, req)
				if err != nil {
					return err
				}
				res <- userAccount
			}
		}
		return nil
	})

	return g.Wait()
}
