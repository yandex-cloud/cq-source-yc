package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
)

func OrganizationManagerFederations() *schema.Table {
	return &schema.Table{
		Name:        "yandex_organizationmanager_federations",
		Resolver:    fetchOrganizationManagerFederations,
		Multiplex:   client.MultiplexBy(client.Organizations),
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:            "id",
				Type:            schema.TypeString,
				Description:     "ID of the resource.",
				Resolver:        client.ResolveResourceId,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true, Unique: true},
			},
			{
				Name:        "organization_id",
				Type:        schema.TypeString,
				Description: "ID of the organization that the federation belongs to.",
				Resolver:    schema.PathResolver("OrganizationId"),
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the federation.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the federation.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "cookie_max_age_seconds",
				Type:        schema.TypeBigInt,
				Description: "",
				Resolver:    schema.PathResolver("CookieMaxAge.Seconds"),
			},
			{
				Name:        "cookie_max_age_nanos",
				Type:        schema.TypeInt,
				Description: "",
				Resolver:    schema.PathResolver("CookieMaxAge.Nanos"),
			},
			{
				Name:        "auto_create_account_on_login",
				Type:        schema.TypeBool,
				Description: "Add new users automatically on successful authentication.\n The user becomes member of the organization automatically,\n but you need to grant other roles to them.\n\n If the value is `false`, users who aren't added to the organization\n can't log in, even if they have authenticated on your server.",
				Resolver:    schema.PathResolver("AutoCreateAccountOnLogin"),
			},
			{
				Name:        "issuer",
				Type:        schema.TypeString,
				Description: "ID of the IdP server to be used for authentication.\n The IdP server also responds to IAM with this ID after the user authenticates.",
				Resolver:    schema.PathResolver("Issuer"),
			},
			{
				Name:        "sso_binding",
				Type:        schema.TypeString,
				Description: "Single sign-on endpoint binding type. Most Identity Providers support the `POST` binding type.\n\n SAML Binding is a mapping of a SAML protocol message onto standard messaging\n formats and/or communications protocols.",
				Resolver:    client.EnumPathResolver("SsoBinding"),
			},
			{
				Name:        "sso_url",
				Type:        schema.TypeString,
				Description: "Single sign-on endpoint URL.\n Specify the link to the IdP login page here.",
				Resolver:    schema.PathResolver("SsoUrl"),
			},
			{
				Name:        "security_settings_encrypted_assertions",
				Type:        schema.TypeBool,
				Description: "Enable encrypted assertions.",
				Resolver:    schema.PathResolver("SecuritySettings.EncryptedAssertions"),
			},
			{
				Name:        "case_insensitive_name_ids",
				Type:        schema.TypeBool,
				Description: "Use case insensitive Name IDs.",
				Resolver:    schema.PathResolver("CaseInsensitiveNameIds"),
			},
		},
	}

}

func fetchOrganizationManagerFederations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	req := &saml.ListFederationsRequest{OrganizationId: c.MultiplexedResourceId}
	it := c.Services.OrganizationManagerSAML.Federation().FederationIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}
