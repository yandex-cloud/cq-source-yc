.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: build
build:
	go build -o cq-source-yc main.go

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables

# All gen targets
.PHONY: gen
gen: gen-docs


# Integration tests helpers

.PHONY: start-source
start-source: build
	./cq-source-yc serve --log-level debug