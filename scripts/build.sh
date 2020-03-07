#!/bin/bash

SCRIPT_DIR="`cd $(dirname $0);pwd`"
PROJECT_DIR="$SCRIPT_DIR/../"

cd ${PROJECT_DIR}

echo "==> Removing old files"
rm -f bin
mkdir -p bin

CM_ARCH=${CM_ARCH:-"amd64"}
CM_OS=${CM_OS:-"linux"}

echo "==> building..."
GOOS=$OS GOARCH=$arch go build -o bin/swan-${os}-${arch} cmd/swan.go