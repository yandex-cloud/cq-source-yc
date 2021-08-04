module github.com/yandex-cloud/cq-provider-yandex

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/cloudquery/cq-provider-sdk v0.2.8
	github.com/cloudquery/faker/v3 v3.7.4
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-hclog v0.16.1
	github.com/iancoleman/strcase v0.1.3
	github.com/jhump/protoreflect v1.9.0
	github.com/jinzhu/inflection v1.0.0
	github.com/thoas/go-funk v0.8.1-0.20210502090430-efae847b30ab
	github.com/yandex-cloud/go-genproto v0.0.0-20210517152439-84c9ad4d8b5f
	github.com/yandex-cloud/go-sdk v0.0.0-20210517154707-ca282b96279e
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/cloudquery/cq-provider-sdk v0.2.8 => github.com/daniil-ushkov/cq-provider-sdk v0.2.9-0.20210803072609-3458aa2c2e18
	github.com/cloudquery/faker/v3 v3.7.4 => github.com/daniil-ushkov/faker v1.5.1-0.20210727155430-974b577181cb
)
