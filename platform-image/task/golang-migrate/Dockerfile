FROM golang:1.16 as builder

ARG VERSION="v4.15.1"

WORKDIR /go/src/github.com/golang-migrate/migrate

RUN set -xe && \
    apt-get update -y && \
    apt-get install -y git gcc make && \
    git clone https://github.com/golang-migrate/migrate . && \
    git checkout "${VERSION}" && \
    make build-docker && \
    cp /go/src/github.com/golang-migrate/migrate/build/migrate.linux-386 /bin/migrate



FROM scratch

COPY --from=builder /bin/migrate /bin/migrate

WORKDIR /service

ENTRYPOINT ["migrate"]

CMD ["-help"]
