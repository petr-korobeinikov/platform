FROM golang

ARG VERSION="v1.44.0"

RUN set -xe && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin "${VERSION}"

COPY entrypoint.sh /bin/entrypoint.sh
COPY golangci.yaml /etc/golangci.yaml

WORKDIR /service

ENTRYPOINT ["/bin/entrypoint.sh"]
