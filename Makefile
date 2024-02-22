.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: build
build:
	go build -o yc main.go

.PHONY: gen-docs
gen-docs: build
	# TODO: replace with cloudquery cli invocation
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m

# All gen targets
.PHONY: gen
gen: gen-docs
