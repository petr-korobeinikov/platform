#!/usr/bin/env bash

CFG_FILE=".golangci.yaml"
FALLBACK_CFG_FILE="/etc/golangci.yaml"

USED_CFG_FILE="$CFG_FILE"

[[ -f "$CFG_FILE" ]] || {
  echo "File $CFG_FILE not found in the project."
  echo "Using default $FALLBACK_CFG_FILE instead."

  USED_CFG_FILE="$FALLBACK_CFG_FILE"
}

golangci-lint run --config "$USED_CFG_FILE"
