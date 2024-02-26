package client

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func TestResourceHierarchy(t *testing.T) {
	t.Skip("test not mocked")
	ctx := context.Background()
	credentials, err := getCredentials()
	if err != nil {
		t.Error(err)
	}
	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: credentials,
	})
	if err != nil {
		t.Error(err)
	}
	logger := zerolog.New(os.Stdout)

	// Validate that we indeed have hierarchy and don't have dangling resources
	// TODO: mock it
	t.Run("empty_filters", func(t *testing.T) {
		hierarchy, err := NewResourceHierarchy(ctx, logger, sdk, []string{}, []string{}, []string{})
		if err != nil {
			t.Error(err)
		}
		items := hierarchy.All()
		for _, item := range items {
			algebra := []bool{
				item.Organization != "",
				item.Cloud != "",
				item.Folder != "",
			}
			t.Log(item)
			seenOne := false
			seenZeroAfterOne := false
			for _, i := range algebra {
				if i == true {
					if seenZeroAfterOne {
						t.Fatalf("zeros can't be followed by ones: %+v -> %+v", algebra, item)
					}
					seenOne = true
				} else {
					if seenOne {
						seenZeroAfterOne = true
					}
				}
			}
		}

		t.Run("OrganizationRows", func(t *testing.T) {
			rows := hierarchy.OrganizationRows()
			for _, item := range rows {
				if item.Cloud != "" {
					t.Fatalf("OrganizationRows have item with not-null Cloud: %+v", item)
				}
				if item.Folder != "" {
					t.Fatalf("OrganizationRows have item with not-null Folder: %+v", item)
				}
			}
		})

		t.Run("CloudRows", func(t *testing.T) {
			rows := hierarchy.CloudRows()
			for _, item := range rows {
				if item.Organization == "" {
					t.Fatalf("CloudRows have item with null Organization: %+v", item)
				}
				if item.Folder != "" {
					t.Fatalf("CloudRows have item with not-null Folder: %+v", item)
				}
			}
		})

		t.Run("FolderRows", func(t *testing.T) {
			rows := hierarchy.FolderRows()
			for _, item := range rows {
				if item.Folder == "" {
					t.Fatalf("FolderRows have item with null Folder: %+v", item)
				}
			}
		})
	})

}
