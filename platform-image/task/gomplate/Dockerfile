FROM golang AS builder

RUN go install github.com/joho/godotenv/cmd/godotenv@latest
RUN go install github.com/hairyhenderson/gomplate/v3/cmd/gomplate@latest

#FROM scratch

#COPY --from=builder /go/bin/godotenv /bin/godotenv
#COPY --from=builder /go/bin/gomplate /bin/gomplate

WORKDIR /service

ENTRYPOINT ["godotenv", "-f", ".env", "-f", ".platform/env/.env", "gomplate"]
