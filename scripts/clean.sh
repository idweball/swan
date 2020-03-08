#!/bin/bash

SCRIPT_DIR="`cd $(dirname $0);pwd`"
PROJECT_DIR="$SCRIPT_DIR/../"

cd ${PROJECT_DIR}

echo "==> cleanning..."
rm -rf bin
rm -rf test/data/tmp/*
