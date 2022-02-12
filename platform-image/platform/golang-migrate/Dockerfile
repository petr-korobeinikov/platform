FROM golang:1.16-alpine3.13 AS builder

ARG VERSION="v4.15.1"

RUN apk add --no-cache git gcc musl-dev make

WORKDIR /go/src/github.com/golang-migrate/migrate

RUN set -xe && \
    git clone https://github.com/golang-migrate/migrate . && \
    git checkout "${VERSION}" && \
    make build-docker



FROM alpine:3.13

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/golang-migrate/migrate/build/migrate.linux-386 /usr/local/bin/migrate

RUN ln -s /usr/local/bin/migrate /migrate

ENTRYPOINT ["migrate"]

CMD ["-help"]
