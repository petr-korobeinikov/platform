#!/usr/bin/env bash

set -euo pipefail

echo "Generating Go source"
oto \
  -template /templates/server.go.plush \
  -out /service/oto/generated/oto.gen.go \
  -ignore Ignorer \
  -pkg generated \
  /service/oto/definition

echo "Running gofmt on generated Go source"
gofmt -w /service/oto/generated/oto.gen.go

echo "Generating JS source"
oto \
  -template /templates/client.js.plush \
  -out /service/oto/generated/oto.gen.js \
  /service/oto/definition
