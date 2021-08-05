# syntax=docker/dockerfile:1

FROM golang:1.16 AS cloudquery
WORKDIR /app

# Download CQ
ADD https://github.com/cloudquery/cloudquery/releases/download/v0.13.8/cloudquery_Linux_x86_64 cloudquery
RUN chmod +x cloudquery

# Copy YC provider
COPY . .
RUN go mod tidy

WORKDIR resources
RUN go test -c -o test
ENV DATABASE_URL="host=cq_provider_yandex_postgresql user=postgres password=pass DB.name=postgres port=5432"
CMD ./test -test.v
