package resources_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"testing"

	iam1 "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-sdk/gen/iam"
	"google.golang.org/grpc"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func TestUserAccountsByFolder(t *testing.T) {
	accessBindingsByFolderServer, serv1, err := createAccessBindingsByFolderServer()
	defer serv1.Stop()
	if err != nil {
		t.Fatal(err)
	}

	userAccountsByFolderServer, serv2, err := createUserAccountsByFolderServer()
	defer serv2.Stop()
	if err != nil {
		t.Fatal(err)
	}

	resource := providertest.ResourceTestData{
		Table:  resources.IAMUserAccountsByFolder(),
		Config: client.Config{FolderIDs: []string{"test"}},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := client.NewYandexClient(
				logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}),
				[]string{"test"},
				nil,
				nil,
				&client.Services{
					IAM:             userAccountsByFolderServer,
					ResourceManager: accessBindingsByFolderServer,
				}, nil)
			return c, nil
		},
		Verifiers: []providertest.Verifier{
			providertest.VerifyAtLeastOneRow("yandex_iam_user_accounts_by_folder"),
			providertest.VerifyNoEmptyColumnsExcept(
				"yandex_iam_user_accounts_by_folder",
				"user_account_yandex_passport_user_account_login",
				"user_account_yandex_passport_user_account_default_email",
				"user_account_saml_user_account_federation_id",
				"user_account_saml_user_account_name_id",
				"user_account_saml_user_account_attributes",
			),
			providertest.VerifyOneOf(
				"yandex_iam_user_accounts_by_folder",
				"user_account_yandex_passport_user_account_login",
				"user_account_saml_user_account_name_id",
			),
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}

type FakeUserAccountByFolderServer struct {
	iam1.UnimplementedUserAccountServiceServer
}

func (f *FakeUserAccountByFolderServer) Get(_ context.Context, req *iam1.GetUserAccountRequest) (*iam1.UserAccount, error) {
	switch req.UserAccountId {
	case "1":
		return &iam1.UserAccount{Id: "1", UserAccount: &iam1.UserAccount_SamlUserAccount{
			SamlUserAccount: &iam1.SamlUserAccount{
				FederationId: "1",
				NameId:       "1",
			},
		}}, nil
	case "2":
		return &iam1.UserAccount{Id: "2", UserAccount: &iam1.UserAccount_YandexPassportUserAccount{
			YandexPassportUserAccount: &iam1.YandexPassportUserAccount{
				Login:        "qwerty",
				DefaultEmail: "qwerty@qwerty.com",
			},
		}}, nil
	default:
		return nil, errors.New("no such user account")
	}
}

func createUserAccountsByFolderServer() (*iam.IAM, *grpc.Server, error) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, nil, err
	}

	serv := grpc.NewServer()
	fakeCloudServiceServer := &FakeUserAccountByFolderServer{}

	iam1.RegisterUserAccountServiceServer(serv, fakeCloudServiceServer)

	go func() {
		err := serv.Serve(lis)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return iam.NewIAM(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), serv, nil
}
