package integration_tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

var suffix = acctest.RandString(10)

func yandexTestIntegrationHelper(t *testing.T, table *schema.Table, verificationBuilder func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification, tfTmpl string) {
	cfg := client.Config{
		CloudID:   os.Getenv("YC_CLOUD_ID"),
		FolderIDs: []string{os.Getenv("YC_FOLDER_ID")},
	}

	file, err := os.Create(table.Name + ".tf")
	if err != nil {
		t.Fatal(err)
	}

	_, err = file.WriteString(fmt.Sprintf(tfTmpl, suffix))
	if err != nil {
		t.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(table.Name + ".tf")

	providertest.IntegrationTest(t, resources.Provider, providertest.ResourceIntegrationTestData{
		Table:               table,
		Config:              cfg,
		Configure:           client.Configure,
		Suffix:              suffix,
		VerificationBuilder: verificationBuilder,
	})
}

func TestMain(m *testing.M) {
	if _, ok := os.LookupEnv("YC_CLOUD_ID"); !ok {
		fmt.Fprintln(os.Stderr, "YC_CLOUD_ID wasn't specified.")
		return
	}

	if _, ok := os.LookupEnv("YC_FOLDER_ID"); !ok {
		fmt.Fprintln(os.Stderr, "YC_FOLDER_ID wasn't specified.")
		return
	}

	os.Exit(m.Run())
}
