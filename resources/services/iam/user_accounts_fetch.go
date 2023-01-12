package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type (
	UserAccount struct {
		*iam.UserAccount
		*iam.SamlUserAccount
		*iam.YandexPassportUserAccount
	}

	accessBindingsClient interface {
		ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	}
)

func fetchUserAccounts(_client accessBindingsClient) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
		c := meta.(*client.Client)
		userClient := c.Services.IAM.UserAccount()
		g := errgroup.Group{}
		ch := make(chan *access.AccessBinding)
		g.Go(func() error {
			defer close(ch)
			return fetchAccessBindings(ctx, _client, c.MultiplexedResourceId, ch)
		})

		g.Go(func() error {
			return fetchUserAccountByID(ctx, userClient, ch, res)
		})

		return g.Wait()
	}
}

func fetchAccessBindings(ctx context.Context, _client accessBindingsClient, resourceID string, res chan<- *access.AccessBinding) error {
	req := &access.ListAccessBindingsRequest{ResourceId: resourceID}
	for {
		resp, err := _client.ListAccessBindings(ctx, req)
		if err != nil {
			return err
		}
		for _, b := range resp.GetAccessBindings() {
			res <- b
		}

		if resp.GetNextPageToken() == "" {
			return nil
		}
		req.PageToken = resp.NextPageToken
	}
}

func fetchUserAccountByID(ctx context.Context, client iam.UserAccountServiceClient, subjects <-chan *access.AccessBinding, res chan<- interface{}) error {
	for subject := range subjects {
		accountType := subject.Subject.Type
		if accountType == "serviceAccount" || accountType == "group" {
			continue
		}
		req := &iam.GetUserAccountRequest{UserAccountId: subject.Subject.Id}
		userAccount, err := client.Get(ctx, req)
		if err != nil {
			return err
		}

		samlUserAccount := userAccount.GetSamlUserAccount()
		if samlUserAccount == nil {
			samlUserAccount = &iam.SamlUserAccount{}
		}

		yandexPassportUserAccount := userAccount.GetYandexPassportUserAccount()
		if yandexPassportUserAccount == nil {
			yandexPassportUserAccount = &iam.YandexPassportUserAccount{}
		}

		res <- UserAccount{
			UserAccount:               userAccount,
			SamlUserAccount:           samlUserAccount,
			YandexPassportUserAccount: yandexPassportUserAccount,
		}
	}
	return nil
}
