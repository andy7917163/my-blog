#!/bin/bash
set -e

# 若系統已有 go，直接用
if command -v go &> /dev/null; then
  go run ./cmd/blog build
  exit 0
fi

# 否則下載 Go 1.23
GO_VERSION="1.23.5"
GO_DIR="/tmp/go-${GO_VERSION}"

if [ ! -d "$GO_DIR" ]; then
  echo "下載 Go ${GO_VERSION}..."
  curl -sL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -C /tmp -xz
  mv /tmp/go "$GO_DIR"
fi

"${GO_DIR}/bin/go" run ./cmd/blog build
