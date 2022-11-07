module github.com/yandex-cloud/cq-provider-yandex

go 1.19

require (
	github.com/aws/aws-sdk-go v1.44.104
	github.com/cloudquery/plugin-sdk v0.13.23
	github.com/golang/protobuf v1.5.2
	github.com/iancoleman/strcase v0.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rs/zerolog v1.28.0
	github.com/thoas/go-funk v0.9.2
	github.com/yandex-cloud/go-genproto v0.0.0-20220916153440-1086fb0ea529
	github.com/yandex-cloud/go-sdk v0.0.0-20220914165927-0dcfb37c9703
	golang.org/x/sync v0.1.0
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cloudquery/faker/v3 v3.7.7
	github.com/golang/mock v1.6.0
)

require (
	github.com/getsentry/sentry-go v0.14.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.2.0.20201002093600-73cf2ae9d891 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20221031165847-c99f073a8326 // indirect
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20220921223823-23cae91e6737 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/cloudquery/faker/v3 v3.7.7 => github.com/daniil-ushkov/faker/v3 v3.7.5-0.20210727153506-26b3d57a100b
