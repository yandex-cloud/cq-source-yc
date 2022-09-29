GOIMPORTS_RESOURCES = find resources -name "*.go" -exec goimports -w {} \;
GREEN=\033[0;32m
NC=\033[0m

# Provider generation

.PHONY: generate-resources
generate-resources:
	@go run gen/main.go

# Debug

.PHONY: debug
debug:
	@env CQ_PROVIDER_DEBUG=1 YC_TOKEN=$(YC_TOKEN) go run main.go

.PHONY: postgres
postgres:
	@docker run -d --rm \
	 --name=cq-postgres \
	 -e POSTGRES_PASSWORD=pass \
	 -p 5432:5432 \
	 postgres:latest

# Tests

.PHONY: docker-build-local-tests
docker-build-local-tests:
	@echo "$(GREEN)Building local tests image...$(NC)"
	@test -n "$$(docker image ls -a -q -f reference=cq_provider_yandex_local_tests)" || \
	docker build -t cq_provider_yandex_local_tests -f localtests.Dockerfile .

.PHONY: docker-build-integration-tests
docker-build-integration-tests:
	@echo "$(GREEN)Building integration tests image...$(NC)"
	@test -n "$$(docker image ls -a -q -f reference=cq_provider_yandex_integration_tests)" || \
	docker build -t cq_provider_yandex_integration_tests -f integrationtests.Dockerfile .

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
	@echo "$(GREEN)Waiting for connection to PostgreSQL server...$(NC)"; until pg_isready -q -h localhost -p 5432;do echo -n .;sleep 1;done;echo ""

.PHONY: docker-minio
docker-minio: docker-create-net
	@echo "$(GREEN)Starting MINIO...$(NC)"
	@test -n "$$(docker ps -a -q -f name=cq_provider_yandex_s3)" || \
	docker run --rm -d \
    	--name cq_provider_yandex_s3 \
    	--network=cq_provider_yandex_net \
    	-e MINIO_ROOT_USER=user \
    	-e MINIO_ROOT_PASSWORD=12345678 \
    	-p 9000:9000 \
    	minio/minio server /mnt/data
	@until nc -vz localhost 9000;do echo -n .; sleep 1;done;echo ""
	@env AWS_ACCESS_KEY_ID=user AWS_SECRET_ACCESS_KEY=12345678 aws --endpoint=http://localhost:9000 s3 mb s3://cq-test-bucket

.PHONY: local-tests
local-tests: docker-postgresql docker-build-local-tests docker-minio
	@docker run --rm \
	--name=cq_provider_yandex_local_tests \
	--network=cq_provider_yandex_net \
	-e DATABASE_URL="host=cq_provider_yandex_postgresql user=postgres password=pass DB.name=postgres port=5432" \
	cq_provider_yandex_local_tests

.PHONY: integration-tests
integration-tests: docker-postgresql docker-build-integration-tests
	docker run --rm \
	--name=cq_provider_yandex_integration_test \
	--network=cq_provider_yandex_net \
	-e DATABASE_URL="host=cq_provider_yandex_postgresql user=postgres password=pass DB.name=postgres port=5432" \
	-e YC_CLOUD_ID=$(YC_CLOUD_ID) \
	-e YC_FOLDER_ID=$(YC_FOLDER_ID) \
	-e YC_SERVICE_ACCOUNT_KEY_FILE='$(YC_SERVICE_ACCOUNT_KEY_FILE)' \
	-e YC_STORAGE_ACCESS_KEY=$(YC_STORAGE_ACCESS_KEY) \
	-e YC_STORAGE_SECRET_KEY=$(YC_STORAGE_SECRET_KEY) \
	cq_provider_yandex_integration_tests

.PHONY: clean
clean: clean-docker-postgresql clean-docker-minio clean-docker-net clean-local-tests-image clean-integration-tests-image

.PHONY: clean-local-tests-image
clean-local-tests-image:
	@docker image rm cq_provider_yandex_local_tests || :

.PHONY: clean-integration-tests-image
clean-integration-tests-image:
	@docker image rm cq_provider_yandex_integration_tests || :

.PHONY: clean-docker-net
clean-docker-net:
	@docker network rm cq_provider_yandex_net || :

.PHONY: clean-docker-postgresql
clean-docker-postgresql:
	@docker stop cq_provider_yandex_postgresql || :

.PHONY: clean-docker-minio
clean-docker-minio:
	@docker stop cq_provider_yandex_s3 || :
