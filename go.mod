module github.com/yandex-cloud/cq-provider-yandex

go 1.19

require (
	github.com/aws/aws-sdk-go v1.44.156
	github.com/cloudquery/plugin-sdk v1.11.2
	github.com/golang/protobuf v1.5.2
	github.com/iancoleman/strcase v0.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rs/zerolog v1.28.0
	github.com/thoas/go-funk v0.9.2
	github.com/yandex-cloud/go-genproto v0.0.0-20221205100932-c2782a87f4d0
	github.com/yandex-cloud/go-sdk v0.0.0-20221205101755-09e71e1b31e4
	golang.org/x/sync v0.1.0
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cloudquery/faker/v3 v3.7.7
	github.com/golang/mock v1.6.0
)

require (
	github.com/getsentry/sentry-go v0.16.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20221208152030-732eee02a75a // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/cloudquery/faker/v3 v3.7.7 => github.com/daniil-ushkov/faker/v3 v3.7.5-0.20210727153506-26b3d57a100b
