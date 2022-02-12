FROM golang

ARG VERSION="v1.44.0"

WORKDIR /service

COPY .golangci.yaml /etc/golangci.yaml

RUN set -xe && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin "${VERSION}"

CMD ["golangci-lint", "run", "--config", "/etc/golangci.yaml"]
