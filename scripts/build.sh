#!/bin/bash

SCRIPT_DIR="`cd $(dirname $0);pwd`"
PROJECT_DIR="$SCRIPT_DIR/../"

cd ${PROJECT_DIR}

echo "==> Removing old files"
rm -rf bin
mkdir -p bin

CM_ARCH=${CM_ARCH:-"amd64"}
CM_OS=${CM_OS:-"linux"}

echo "==> building..."
GOOS=$CM_OS GOARCH=$CM_ARCH go build -o bin/swan-${CM_OS}-${CM_ARCH} cmd/swan.go