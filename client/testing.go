package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		svc, err := createService()
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		var ycSpec Spec
		if err := spec.UnmarshalSpec(&ycSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal yc spec: %w", err)
		}
		c := NewYandexClient(logger, nil, nil,
			[]string{"test-folder-id"},
			[]string{"test-cloud-id"},
			[]string{"test-organization-id"},
			svc,
		)

		return c, nil
	}
	p := plugins.NewSourcePlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"test-destination"},
	})
}
