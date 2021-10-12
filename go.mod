module github.com/yandex-cloud/cq-provider-yandex

go 1.17

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/aws/aws-sdk-go v1.37.0
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
	github.com/yandex-cloud/go-genproto v0.0.0-20210811160424-06c241452aa5 //v0.0.0-20210517152439-84c9ad4d8b5f
	github.com/yandex-cloud/go-sdk v0.0.0-20210811160850-f28151fc0a62
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.32.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/creasty/defaults v1.5.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/doug-martin/goqu/v9 v9.13.0 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/georgysavva/scany v0.2.8 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-test/deep v1.0.7 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/hashicorp/go-plugin v1.4.1 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/hashicorp/terraform-exec v0.13.3 // indirect
	github.com/hashicorp/terraform-json v0.10.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20210316155119-a95892c5f864 // indirect
	github.com/huandu/go-sqlbuilder v1.12.1 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.8.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.7.0 // indirect
	github.com/jackc/puddle v1.1.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/tmccombs/hcl2json v0.3.3 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.4 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/zclconf/go-cty v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/cloudquery/cq-provider-sdk v0.3.1 => github.com/daniil-ushkov/cq-provider-sdk v0.3.1-0.20210817100343-43126bbee7d2
	github.com/cloudquery/faker/v3 v3.7.4 => github.com/daniil-ushkov/faker v1.5.1-0.20210727155430-974b577181cb
)
