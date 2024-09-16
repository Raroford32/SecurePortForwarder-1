#!/bin/bash

set -e

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Install dependencies
go get -u github.com/golang/crypto/...

# Build client and server
go build -o bin/client cmd/client/main.go
go build -o bin/server cmd/server/main.go

echo "Installation completed successfully."
