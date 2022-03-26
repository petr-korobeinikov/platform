FROM golang

ENV CGO_ENABLED=0 GOOS=linux

RUN \
    mkdir /templates && \
    curl \
      --show-error \
      --silent \
      --location \
      --max-redirs 2 \
      --output /templates/server.go.plush \
      https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/server.go.plush && \
    curl \
      --show-error \
      --silent \
      --location \
      --max-redirs 2 \
      --output /templates/client.js.plush \
      https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/client.js.plush

RUN go install -a -v -ldflags="-w -s" github.com/pacedotdev/oto@latest

COPY entrypoint.sh /bin/entrypoint.sh

WORKDIR /service

ENTRYPOINT ["/bin/entrypoint.sh"]
