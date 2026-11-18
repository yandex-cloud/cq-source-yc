package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestClassifyError(t *testing.T) {
	denied := status.Error(codes.PermissionDenied, "Permission denied")
	notFound := status.Error(codes.NotFound, "Image \"fd8j050mgpmq41lh33pm\" not found")
	unavailable := status.Error(codes.Unavailable, "connection reset")

	for name, tc := range map[string]struct {
		err      error
		suppress bool
		code     string
	}{
		"permission denied":                 {err: denied, suppress: true, code: "PermissionDenied"},
		"wrapped":                           {err: fmt.Errorf("list: %w", notFound), suppress: true, code: "NotFound"},
		"unimplemented":                     {err: status.Error(codes.Unimplemented, ""), suppress: true, code: "Unimplemented"},
		"joined expected":                   {err: fmt.Errorf("quotas: %w", errors.Join(notFound, denied)), suppress: true, code: "NotFound"},
		"joined with an unexpected":         {err: errors.Join(denied, unavailable)},
		"unavailable":                       {err: unavailable},
		"not grpc":                          {err: errors.New("parent is not type of *organizationmanager.Group")},
		"wrapped joined with an unexpected": {err: fmt.Errorf("quotas: %w", errors.Join(denied, unavailable))},
	} {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			c := &Client{Logger: zerolog.New(&buf), FolderId: "b1g"}
			event := schema.ErrorEvent{Table: &schema.Table{Name: "yc_compute_disks_images"}, Client: c}

			require.Equal(t, tc.suppress, ClassifyError(context.Background(), tc.err, event))
			if !tc.suppress {
				assert.Empty(t, buf.String())
				return
			}

			var line map[string]any
			require.NoError(t, json.Unmarshal(buf.Bytes(), &line))
			assert.Equal(t, "warn", line["level"])
			assert.Equal(t, "yc_compute_disks_images", line["table"])
			assert.Equal(t, "folder:b1g", line["client"])
			assert.Equal(t, "table_resolver", line["phase"])
			assert.Equal(t, tc.code, line["grpc_code"])
			assert.Equal(t, tc.err.Error(), line["error"])
			assert.Contains(t, line["message"], "resolver skipped: ")
		})
	}
}
