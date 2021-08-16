GOIMPORTS_RESOURCES = find resources -name "*.go" -exec goimports -w {} \;
GREEN=\033[0;32m
NC=\033[0m

# Provider generation

cloudapi:
	@git clone https://github.com/yandex-cloud/cloudapi.git

.PHONY: generate-resources
generate-resources: cloudapi
	@go run tools/genresources.go
	@$(GOIMPORTS_RESOURCES)

# Debug

.PHONY: debug
debug:
	@env CQ_PROVIDER_DEBUG=1 YC_TOKEN=$(YC_TOKEN) go run main.go

# Tests

.PHONY: docker-build
docker-build:
	@echo "$(GREEN)Building test image...$(NC)"
	@test -n "$$(docker image ls -a -q -f reference=cq_provider_yandex_image)" || docker build -t cq_provider_yandex_image .

.PHONY: docker-create-net
docker-create-net:
	@echo "$(GREEN)Creating network...$(NC)"
	@test -n "$$(docker network ls -q -f name=cq_provider_yandex_net)" || docker network create cq_provider_yandex_net

.PHONY: docker-postgresql
docker-postgresql: docker-create-net
	@echo "$(GREEN)Staring PostgreSQL server...$(NC)"
	@test -n "$$(docker ps -a -q -f name=cq_provider_yandex_postgresql)" || \
	docker run -d --rm \
    --name=cq_provider_yandex_postgresql \
    --network=cq_provider_yandex_net \
    -e POSTGRES_PASSWORD=pass \
    -p 5432:5432 \
    postgres
	@echo "$(GREEN)Waiting for connection to PostgreSQL server...$(NC)"; until pg_isready -q -h localhost -p 5432; do echo -n .;sleep 1;done;echo ""

.PHONY: test
test: docker-postgresql docker-build
	@docker run -it --rm \
	--name=cq_provider_yandex_test \
	--network=cq_provider_yandex_net \
	cq_provider_yandex_image

.PHONY: clean
clean: clean-docker-postgresql clean-docker-net clean-image

.PHONY: clean-image
clean-image:
	@docker image rm cq_provider_yandex_image

.PHONY: clean-docker-net
clean-docker-net:
	@docker network rm cq_provider_yandex_net

.PHONY: clean-docker-postgresql
clean-docker-postgresql:
	@docker stop cq_provider_yandex_postgresql
