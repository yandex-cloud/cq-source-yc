GOIMPORTS_RESOURCES = find resources -name "*.go" -exec goimports -w {} \;

cloudapi:
	git clone https://github.com/yandex-cloud/cloudapi.git

.PHONY: generate-resources
generate-resources: cloudapi
	go run tools/genresources.go
	$(GOIMPORTS_RESOURCES)

.PHONY: debug
debug:
	env CQ_PROVIDER_DEBUG=1 YC_TOKEN=$(YC_TOKEN) go run main.go
