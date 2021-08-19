# syntax=docker/dockerfile:1

FROM ubuntu:20.04

ENV DEBIAN_FRONTEND noninteractive

## Install docker
RUN apt-get update
RUN apt-get install -y \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg \
        lsb-release
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
RUN echo \
      "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
      $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN apt-get update
RUN apt-get install -y docker-ce docker-ce-cli containerd.io

# Install utils
RUN apt-get install -y gcc unzip

# Install golang:1.17
RUN curl https://dl.google.com/go/go1.17.linux-amd64.tar.gz -o go1.17.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz
ENV PATH "$PATH:/usr/local/go/bin"

# Install aws
RUN curl https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip -o awscliv2.zip && unzip awscliv2.zip && ./aws/install

WORKDIR /app

# Copy YC provider
COPY . .

WORKDIR resources
RUN go test -c -o test
ENV DATABASE_URL="host=cq_provider_yandex_postgresql user=postgres password=pass DB.name=postgres port=5432"
CMD service docker start && ./test -test.v
