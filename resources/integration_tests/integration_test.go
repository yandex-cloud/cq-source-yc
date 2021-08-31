package integration_tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

var suffix = acctest.RandString(10)

func testIntegrationHelper(t *testing.T, table *schema.Table, verificationBuilder func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification, tfTmpl string) {
	cfg := client.Config{
		CloudIDs:  []string{os.Getenv("YC_CLOUD_ID")},
		FolderIDs: []string{os.Getenv("YC_FOLDER_ID")},
	}

	tfFilename := table.Name + ".tf"

	createTfFile(t, tfFilename, tfTmpl)
	defer os.Remove(tfFilename)

	createTable(t, table)

	providertest.IntegrationTest(t, resources.Provider, providertest.ResourceIntegrationTestData{
		Table:               table,
		Config:              cfg,
		Configure:           client.Configure,
		VerificationBuilder: verificationBuilder,
	})
}

func createTfFile(t *testing.T, filename, tfTmpl string) {
	file, err := os.Create(filename)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("failed to create %s", filename), err)
	}

	_, err = file.WriteString(tfTmpl)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("failed to write to %s", filename), err)
	}

	err = file.Close()
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("failed to close %s", filename), err)
	}
}

func createTable(t *testing.T, table *schema.Table) {
	pool, err := setupDatabase()
	if err != nil {
		assert.FailNow(t, "failed to connect to database", err)
	}
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		assert.FailNow(t, "failed to acquire connection to database", err)
	}
	defer conn.Release()
	l := logging.New(hclog.DefaultOptions)
	migrator := provider.NewMigrator(l)
	if err = migrator.CreateTable(ctx, conn, table, nil); err != nil {
		assert.FailNow(t, fmt.Sprintf("failed to create tables %s", table.Name), err)
	}
}

func setupDatabase() (*pgxpool.Pool, error) {
	var (
		dbCfg *pgxpool.Config
		err   error
	)
	if config, ok := os.LookupEnv("DATABASE_URL"); ok {
		dbCfg, err = pgxpool.ParseConfig(config)
	} else {
		dbCfg, err = pgxpool.ParseConfig("host=localhost user=postgres password=pass DB.name=postgres port=5432")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to parse config. %w", err)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), dbCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database. %w", err)
	}
	return pool, nil
}

func IdentityFilter(sq squirrel.SelectBuilder, _ *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
	return sq
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
