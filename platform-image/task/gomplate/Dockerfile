FROM golang AS builder

ENV CGO_ENABLED=0 GOOS=linux

RUN go install -a -v -ldflags="-w -s" github.com/joho/godotenv/cmd/godotenv@latest
RUN go install -a -v -ldflags="-w -s" github.com/hairyhenderson/gomplate/v3/cmd/gomplate@latest



FROM scratch

COPY --from=builder /go/bin/godotenv /bin/godotenv
COPY --from=builder /go/bin/gomplate /bin/gomplate

WORKDIR /service

ENTRYPOINT ["/bin/godotenv", "-f", ".env,.platform/env/.env", "/bin/gomplate"]
