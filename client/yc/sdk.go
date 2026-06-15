package yc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/helpers/grpczerolog"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/dial"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"github.com/yandex-cloud/go-sdk/pkg/idempotency"
	"github.com/yandex-cloud/go-sdk/pkg/requestid"
	"github.com/yandex-cloud/go-sdk/pkg/retry/v1"
	ycsdkv2 "github.com/yandex-cloud/go-sdk/v2"
	credentialsv2 "github.com/yandex-cloud/go-sdk/v2/credentials"
	iamkeyv2 "github.com/yandex-cloud/go-sdk/v2/pkg/iamkey"
	optionsv2 "github.com/yandex-cloud/go-sdk/v2/pkg/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Config holds the knobs Build needs from the plugin spec, decoupled from the
// client.Spec type so this package stays independent of client.
type Config struct {
	Endpoint   string
	UserAgent  string
	MaxRetries int
	DebugGRPC  bool
}

// Build constructs both the v1 and v2 Yandex Cloud SDKs that back a Client.
func Build(ctx context.Context, logger zerolog.Logger, cfg Config) (*ycsdk.SDK, *ycsdkv2.SDK, error) {
	creds, err := credentials()
	if err != nil {
		return nil, nil, err
	}

	unaryInterceptors := []grpc.UnaryClientInterceptor{
		requestid.Interceptor(),
		idempotency.Interceptor(),
	}
	streamInterceptors := []grpc.StreamClientInterceptor{}

	var dialOpts = []grpc.DialOption{
		grpc.WithUserAgent(dial.UserAgent() + "/" + cfg.UserAgent),
	}

	if cfg.MaxRetries > 0 {
		o, err := retry.RetryDialOption(
			retry.WithRetries(retry.DefaultNameConfig(), cfg.MaxRetries),
			retry.WithRetryableStatusCodes(retry.DefaultNameConfig(), codes.ResourceExhausted, codes.Unavailable),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create retry option: %w", err)
		}
		dialOpts = append(dialOpts, o)
	}

	// debug interceptors are last
	if cfg.DebugGRPC {
		unaryInterceptors = append(unaryInterceptors, logging.UnaryClientInterceptor(grpczerolog.InterceptorLogger(logger)))
		streamInterceptors = append(streamInterceptors, logging.StreamClientInterceptor(grpczerolog.InterceptorLogger(logger)))
	}

	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(unaryInterceptors...), grpc.WithChainStreamInterceptor(streamInterceptors...))

	sdk, err := ycsdk.Build(ctx,
		ycsdk.Config{
			Credentials: creds,
			Endpoint:    cfg.Endpoint,
		},
		dialOpts...,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize Yandex Cloud SDK: %w", err)
	}

	credsv2, err := credentialsV2()
	if err != nil {
		return nil, nil, err
	}
	sdkv2, err := ycsdkv2.Build(ctx,
		optionsv2.WithCredentials(credsv2),
		optionsv2.WithDiscoveryEndpoint(cfg.Endpoint),
		// TODO: use this instead of dialOpts
		// optionsv2.WithDefaultRetryOptions(),
		optionsv2.WithCustomDialOptions(dialOpts...),
	)
	if err != nil {
		return nil, nil, err
	}

	return sdk, sdkv2, nil
}

func iamKeyFromJSONContent(content string) (*iamkey.Key, error) {
	key := &iamkey.Key{}
	err := json.Unmarshal([]byte(content), key)
	if err != nil {
		return nil, fmt.Errorf("key unmarshal: %s", err)
	}
	return key, nil
}

func credentials() (ycsdk.Credentials, error) {
	if val := os.Getenv("YC_SERVICE_ACCOUNT_KEY"); val != "" {
		key, err := iamKeyFromJSONContent(val)
		if err != nil {
			return nil, err
		}
		return ycsdk.ServiceAccountKey(key)
	}

	if val := os.Getenv("YC_TOKEN"); val != "" {
		if strings.HasPrefix(val, "t1.") && strings.Count(val, ".") == 2 {
			return ycsdk.NewIAMTokenCredentials(val), nil
		}
		return ycsdk.OAuthToken(val), nil
	}

	return ycsdk.InstanceServiceAccount(), nil
}

func credentialsV2() (credentialsv2.Credentials, error) {
	if val := os.Getenv("YC_SERVICE_ACCOUNT_KEY"); val != "" {
		key, err := iamkeyv2.ReadFromJSONBytes([]byte(val))
		if err != nil {
			return nil, err
		}
		return credentialsv2.ServiceAccountKey(key)
	}

	if val := os.Getenv("YC_TOKEN"); val != "" {
		if strings.HasPrefix(val, "t1.") && strings.Count(val, ".") == 2 {
			return credentialsv2.IAMToken(val), nil
		}
		return credentialsv2.OAuthToken(val), nil
	}

	return credentialsv2.InstanceServiceAccount(), nil
}
