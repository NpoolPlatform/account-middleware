#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

VERSION=v1.51.2
URL_BASE=https://raw.githubusercontent.com/golangci/golangci-lint
URL=$URL_BASE/$VERSION/install.sh

if [[ ! -f .golangci.yml ]]; then
    echo 'ERROR: missing .golangci.yml in repo root' >&2
    exit 1
fi

if ! command -v gofumpt; then
    go install mvdan.cc/gofumpt@v0.3.1
fi

if ! command -v golangci-lint; then
    curl -sfL $URL | sh -s $VERSION
fi

PATH=$PWD/bin:$PATH
whereis golangci-lint

golangci-lint version
golangci-lint linters
golangci-lint run "$@"
