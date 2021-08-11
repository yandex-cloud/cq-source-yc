module github.com/yandex-cloud/cq-provider-yandex

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/aws/aws-sdk-go v1.37.0 // indirect
	github.com/cloudquery/cq-provider-sdk v0.3.1
	github.com/cloudquery/faker/v3 v3.7.4
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-hclog v0.16.1
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
	github.com/iancoleman/strcase v0.1.3
	github.com/jackc/pgx/v4 v4.11.0
	github.com/jhump/protoreflect v1.9.0
	github.com/jinzhu/inflection v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.8.1-0.20210502090430-efae847b30ab
	github.com/yandex-cloud/go-genproto v0.0.0-20210809082946-a97da516c588 //v0.0.0-20210517152439-84c9ad4d8b5f
	github.com/yandex-cloud/go-sdk v0.0.0-20210517154707-ca282b96279e
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	google.golang.org/grpc v1.32.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/cloudquery/cq-provider-sdk v0.3.1 => github.com/daniil-ushkov/cq-provider-sdk v0.3.1-0.20210805130044-aef60fa55baa
	github.com/cloudquery/faker/v3 v3.7.4 => github.com/daniil-ushkov/faker v1.5.1-0.20210727155430-974b577181cb
)
