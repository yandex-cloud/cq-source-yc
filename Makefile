gen_out := resources

goimports_w_out = find $(gen_out) -name "*.go" -exec goimports -w {} \;

yc_token != cat yc_token

cloudapi:
	git clone https://github.com/yandex-cloud/cloudapi.git

.PHONY: generate-resources
generate-resources: cloudapi
	go run tools/genresources.go
	$(goimports_w_out)

.PHONY: debug
debug:
	env CQ_PROVIDER_DEBUG=1 YC_TOKEN=$(yc_token) go run main.go
