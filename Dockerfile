# syntax=docker/dockerfile:1

FROM golang:1.16
WORKDIR /app

# Copy YC provider
COPY . .

WORKDIR resources
RUN go test -c -o test
ENV DATABASE_URL="host=cq_provider_yandex_postgresql user=postgres password=pass DB.name=postgres port=5432"
CMD ./test -test.v
