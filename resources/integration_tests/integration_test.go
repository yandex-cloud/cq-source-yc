package integration_tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

// IntegrationTestsEnabledVar is the name of the environment variable that enables integration tests from this package.
// Set it to one of "1", "y", "yes", "true" to enable the tests.
const IntegrationTestsEnabledVar = "INTEGRATION_TESTS"

func yandexTestIntegrationHelper(t *testing.T, table *schema.Table, verificationBuilder func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification) {
	cfg := client.Config{
		CloudID:   os.Getenv("YC_CLOUD_ID"),
		FolderIDs: []string{os.Getenv("YC_FOLDER_ID")},
	}

	providertest.IntegrationTest(t, resources.Provider, providertest.ResourceIntegrationTestData{
		Table:               table,
		Config:              cfg,
		Configure:           client.Configure,
		VerificationBuilder: verificationBuilder,
	})
}

func TestMain(m *testing.M) {
	enabled := os.Getenv(IntegrationTestsEnabledVar)
	enabledValues := map[string]struct{}{
		"1":       {},
		"y":       {},
		"yes":     {},
		"true":    {},
		"enable":  {},
		"enabled": {},
	}

	if _, ok := enabledValues[enabled]; !ok {
		fmt.Fprintln(os.Stderr, "Integration tests are skipped. Set INTEGRATION_TESTS=1 environment variable to enable.")
	}

	if _, ok := os.LookupEnv("YC_CLOUD_ID"); !ok {
		fmt.Fprintln(os.Stderr, "YC_CLOUD_ID wasn't specified.")
	}

	if _, ok := os.LookupEnv("YC_FOLDER_ID"); !ok {
		fmt.Fprintln(os.Stderr, "YC_FOLDER_ID wasn't specified.")
	}

	os.Exit(m.Run())
}
