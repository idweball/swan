#!/bin/bash

SCRIPT_DIR="`cd $(dirname $0);pwd`"
PROJECT_DIR="$SCRIPT_DIR/../"

cd ${PROJECT_DIR}

echo "==> Running..."
go run cmd/swan.go -cfg test/data/config.toml