FROM golang:1.17.2-alpine AS builder

WORKDIR /go/src/service

COPY . .

RUN ls -la && go build -v -o /build/bin/service cmd/service/main.go



FROM alpine:3.14.2

WORKDIR /service

COPY --from=builder /build/bin/service /service/

CMD ["./service"]
