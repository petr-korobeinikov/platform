FROM golang AS builder

ENV CGO_ENABLED=0 GOOS=linux

RUN go install -a -v -ldflags="-w -s" github.com/joho/godotenv/cmd/godotenv@latest



FROM ghcr.io/pkorobeinikov/seeder

COPY --from=builder /go/bin/godotenv /bin/godotenv

WORKDIR /service

ENTRYPOINT ["/bin/godotenv", "-f", ".env,.platform/env/.env", "/bin/seeder"]
