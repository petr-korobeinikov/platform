FROM golang:1.17.2-alpine AS builder

WORKDIR /go/src/service

COPY . .

RUN go build -v -o /build/bin/service cmd/service/main.go



FROM alpine:3.14.2

RUN addgroup --gid=1000 service && \
    adduser --ingroup service --uid=1000 --shell=nologin --disabled-password --no-create-home service

USER service:service

WORKDIR /service

COPY --from=builder /build/bin/service /service/

CMD ["./service"]
