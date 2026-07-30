// Package datalens is a minimal client for the DataLens API
// (https://yandex.cloud/ru/docs/datalens/operations/api-start).
//
// The API is RPC-style: every call is a POST to /rpc/<method> with a JSON
// body, authorized by an IAM token plus an organization id header.
package datalens

//go:generate go run ./gen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
	"github.com/yandex-cloud/cq-source-yc/internal/util"
	"golang.org/x/time/rate"
)

const (
	DefaultEndpoint = "https://api.datalens.tech"

	headerApiVersion = "2"
	headerAuditMode  = "true"

	// The DataLens API allows 60 requests per minute:
	// https://yandex.cloud/ru/docs/datalens/concepts/limits#datalens-api-limits
	requestsPerMinute = 60

	maxErrorBodySize = 4 * 1024
)

// TokenFunc returns an IAM token to authorize API calls with. It is called on
// every request, so it should cache tokens itself (e.g. be backed by
// ycsdk.IamTokenMiddleware).
type TokenFunc func(ctx context.Context) (string, error)

type logger struct {
	zl zerolog.Logger
}

func (l logger) Error(msg string, keysAndValues ...interface{}) {
	l.zl.Error().Fields(keysAndValues).Msg(msg)
}
func (l logger) Info(msg string, keysAndValues ...interface{}) {
	l.zl.Info().Fields(keysAndValues).Msg(msg)
}
func (l logger) Debug(msg string, keysAndValues ...interface{}) {
	l.zl.Debug().Fields(keysAndValues).Msg(msg)
}
func (l logger) Warn(msg string, keysAndValues ...interface{}) {
	l.zl.Warn().Fields(keysAndValues).Msg(msg)
}

type Config struct {
	// Endpoint of the API, DefaultEndpoint if empty.
	Endpoint string
	// UserAgent is sent with every request if non-empty.
	UserAgent string
	// MaxRetries for 429/5xx responses.
	MaxRetries int
	// Token is required.
	Token  TokenFunc
	Debug  bool
	Logger zerolog.Logger
}

type Client struct {
	endpoint   string
	userAgent  string
	tokenFunc  TokenFunc
	httpClient *retryablehttp.Client
	logger     zerolog.Logger
}

func New(cfg Config) (*Client, error) {
	if cfg.Token == nil {
		return nil, fmt.Errorf("datalens: Config.Token is required")
	}
	endpoint := cfg.Endpoint
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}

	limiter := rate.NewLimiter(rate.Every(time.Minute/requestsPerMinute), 1)

	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = cfg.MaxRetries

	lvl := zerolog.InfoLevel
	// retryablehttp.Client calls Logger.Debug for each request, disable it if user doesn't want to log requests
	if cfg.Debug {
		lvl = zerolog.DebugLevel
	}
	httpClient.Logger = retryablehttp.LeveledLogger(logger{zl: cfg.Logger.Level(lvl)})

	httpClient.HTTPClient.Transport = util.NewInterceptTransport(nil, func(req *http.Request) error {
		return limiter.Wait(req.Context())
	})
	httpClient.RequestLogHook = func(_ retryablehttp.Logger, req *http.Request, attempt int) {
		if attempt > 0 {
			cfg.Logger.Debug().Str("url", req.URL.String()).Int("attempt", attempt).Msg("retrying datalens request")
		}
	}

	return &Client{
		endpoint:   endpoint,
		userAgent:  cfg.UserAgent,
		tokenFunc:  cfg.Token,
		httpClient: httpClient,
		logger:     cfg.Logger,
	}, nil
}

// APIError is a non-200 response from the API.
type APIError struct {
	Method     string
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("datalens: %s returned status %d: %s", e.Method, e.StatusCode, e.Body)
}

func (c *Client) rpc(ctx context.Context, method, orgID string, args, result any) error {
	body, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("datalens: marshal %s args: %w", method, err)
	}

	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, c.endpoint+"/rpc/"+method, bytes.NewReader(body))
	if err != nil {
		return err
	}

	token, err := c.tokenFunc(ctx)
	if err != nil {
		return fmt.Errorf("datalens: obtain IAM token: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("x-dl-org-id", orgID)
	req.Header.Set("x-dl-api-version", headerApiVersion)
	req.Header.Set("x-dl-audit-mode", headerAuditMode)
	req.Header.Set("Content-Type", "application/json")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("datalens: %s: %w", method, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBodySize))
		return &APIError{Method: method, StatusCode: resp.StatusCode, Body: string(errBody)}
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("datalens: decode %s response: %w", method, err)
	}
	return nil
}

// Entry is the full (non-locked) shape of a DataLens entry as returned by
// getEntries. Entries the caller has no access to (isLocked=true) carry only
// EntryId, Name, Scope and Type.
type Entry = GetEntriesV2ResultEntries1

// Connection is a getConnection response. In the spec ConnectionRead is a
// discriminated union of ~30 connection types (postgres, clickhouse, gsheets,
// ...), which cannot be represented as one flat struct, so this type is not
// generated: it carries the fields common to every variant, and Data holds
// the whole response including the type-specific fields (host, port, ...).
type Connection struct {
	Id           string  `json:"id"`
	Type         string  `json:"type"`
	Key          string  `json:"key"`
	Description  *string `json:"description"`
	CollectionId *string `json:"collection_id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`

	// Data is filled by GetConnection with the whole response; there is no
	// "data" key in it, so json.Unmarshal leaves this field alone.
	Data map[string]any `json:"data"`
}

// GetConnection calls the getConnection method.
func (c *Client) GetConnection(ctx context.Context, orgID, connectionID string) (*Connection, error) {
	args := map[string]string{"connectionId": connectionID}
	var raw json.RawMessage
	if err := c.rpc(ctx, "getConnection", orgID, args, &raw); err != nil {
		return nil, err
	}

	var conn Connection
	if err := json.Unmarshal(raw, &conn); err != nil {
		return nil, fmt.Errorf("datalens: decode getConnection response: %w", err)
	}
	if err := json.Unmarshal(raw, &conn.Data); err != nil {
		return nil, fmt.Errorf("datalens: decode getConnection response: %w", err)
	}
	// Unlike the spec, the real API sends the connection type as db_type
	// and does not send name at all.
	if conn.Type == "" {
		if dbType, ok := conn.Data["db_type"].(string); ok {
			conn.Type = dbType
		}
	}
	return &conn, nil
}

// GetEntries calls the getEntries method and returns one page of entries
// together with the token of the next page ("" for the last page).
func (c *Client) GetEntries(ctx context.Context, orgID string, args GetEntriesV2Args) ([]Entry, string, error) {
	var result GetEntriesV2Result
	if err := c.rpc(ctx, "getEntries", orgID, args, &result); err != nil {
		return nil, "", err
	}

	entries := make([]Entry, 0, len(result.Entries))
	for i, item := range result.Entries {
		// Both union variants (locked and regular entries) decode into
		// the full Entry shape: the locked one is a subset of fields.
		entry, err := item.AsGetEntriesV2ResultEntries1()
		if err != nil {
			return nil, "", fmt.Errorf("datalens: decode entry %d: %w", i, err)
		}
		entries = append(entries, entry)
	}

	nextPageToken := ""
	if result.NextPageToken != nil {
		nextPageToken = *result.NextPageToken
	}
	return entries, nextPageToken, nil
}
