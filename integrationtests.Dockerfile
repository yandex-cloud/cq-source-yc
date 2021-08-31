# syntax=docker/dockerfile:1

FROM golang:1.16
RUN apt-get update && apt-get install -y curl unzip
RUN curl -o terraform_1.0.5_linux_amd64.zip https://releases.hashicorp.com/terraform/1.0.5/terraform_1.0.5_linux_amd64.zip && \
    unzip terraform_1.0.5_linux_amd64.zip && mv terraform /usr/local/bin

WORKDIR /app

# Copy YC provider
COPY . .

WORKDIR resources/integration_tests
RUN go test -c -o test
CMD ./test -test.v
