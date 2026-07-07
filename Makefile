.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: build
build:
	go build -o yc main.go

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://cli-docs.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf docs/tables
	cloudquery --no-log-file tables --format markdown --output-dir docs test/config.yaml
	mv docs/yc docs/tables

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m

.PHONY: gen-datalens
gen-datalens:
	curl -sSfL https://api.datalens.tech/json/ -o client/yc/datalens/openapi.json
	cd client/yc/datalens && go run ./gen

# All gen targets
.PHONY: gen
gen: gen-docs gen-datalens

.PHONY: update-deps
update-deps:
	go get -u -t ./...