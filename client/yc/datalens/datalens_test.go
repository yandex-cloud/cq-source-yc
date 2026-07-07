package datalens

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestClient(t *testing.T, handler http.Handler) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	c, err := New(Config{
		Endpoint:   srv.URL,
		UserAgent:  "cq-source-yc-test",
		MaxRetries: 2,
		Logger:     zerolog.Nop(),
		Token: func(context.Context) (string, error) {
			return "test-iam-token", nil
		},
	})
	require.NoError(t, err)
	return c
}

func TestGetEntries(t *testing.T) {
	var pageTokens []string
	c := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/rpc/getEntries", r.URL.Path)
		assert.Equal(t, "Bearer test-iam-token", r.Header.Get("Authorization"))
		assert.Equal(t, "org-id", r.Header.Get("x-dl-org-id"))
		assert.Equal(t, "2", r.Header.Get("x-dl-api-version"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var args GetEntriesV2Args
		require.NoError(t, json.NewDecoder(r.Body).Decode(&args))
		require.NotNil(t, args.Scope)
		assert.Equal(t, "connection", *args.Scope)
		token := ""
		if args.PageToken != nil {
			token = *args.PageToken
		}
		pageTokens = append(pageTokens, token)

		w.Header().Set("Content-Type", "application/json")
		if token == "" {
			_, _ = w.Write([]byte(`{
				"nextPageToken": "page2",
				"entries": [
					{"entryId": "e1", "name": "folder-1", "scope": "folder", "type": "", "key": "folder-1", "isLocked": false},
					{"entryId": "e2", "name": "locked", "scope": "connection", "type": "postgres", "isLocked": true}
				]
			}`))
			return
		}
		_, _ = w.Write([]byte(`{"entries": []}`))
	}))

	scope := "connection"
	entries, next, err := c.GetEntries(context.Background(), "org-id", GetEntriesV2Args{Scope: &scope})
	require.NoError(t, err)
	assert.Equal(t, "page2", next)
	require.Len(t, entries, 2)
	assert.Equal(t, "e1", entries[0].EntryId)
	assert.Equal(t, "folder-1", entries[0].Name)
	assert.Equal(t, "e2", entries[1].EntryId)
	assert.Equal(t, "connection", entries[1].Scope)

	entries, next, err = c.GetEntries(context.Background(), "org-id", GetEntriesV2Args{Scope: &scope, PageToken: &next})
	require.NoError(t, err)
	assert.Empty(t, next)
	assert.Empty(t, entries)
	assert.Equal(t, []string{"", "page2"}, pageTokens)
}

func TestGetConnection(t *testing.T) {
	c := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/rpc/getConnection", r.URL.Path)

		var args map[string]string
		require.NoError(t, json.NewDecoder(r.Body).Decode(&args))
		assert.Equal(t, "conn-1", args["connectionId"])

		// The real response has no type and no name, only db_type
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "conn-1",
			"db_type": "trino",
			"key": "Users/user/prod-db",
			"description": "",
			"created_at": "2026-01-01T00:00:00Z",
			"updated_at": "2026-01-02T00:00:00Z",
			"host": "trino.example.com",
			"port": null
		}`))
	}))

	conn, err := c.GetConnection(context.Background(), "org-id", "conn-1")
	require.NoError(t, err)
	assert.Equal(t, "conn-1", conn.Id)
	assert.Equal(t, "trino", conn.Type)
	assert.Equal(t, "trino.example.com", conn.Data["host"])
	require.NotNil(t, conn.Description)
	assert.Empty(t, *conn.Description)
}

func TestRetryOn429(t *testing.T) {
	var requests int
	c := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if requests == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		_, _ = w.Write([]byte(`{"entries": []}`))
	}))

	_, _, err := c.GetEntries(context.Background(), "org-id", GetEntriesV2Args{})
	require.NoError(t, err)
	assert.Equal(t, 2, requests)
}

func TestAPIError(t *testing.T) {
	c := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"code": "ACCESS_DENIED"}`))
	}))

	_, _, err := c.GetEntries(context.Background(), "org-id", GetEntriesV2Args{})
	var apiErr *APIError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, http.StatusForbidden, apiErr.StatusCode)
	assert.Equal(t, "getEntries", apiErr.Method)
	assert.Contains(t, apiErr.Body, "ACCESS_DENIED")
}
