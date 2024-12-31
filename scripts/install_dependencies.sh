#!/bin/bash
# Exit immediately if a command exits with a non-zero status
set -e

# Print commands and arguments as they are executed
set -x

cd /home/ubuntu/
# Update the package list and install Go if it’s not installed
# if ! command -v go &> /dev/null; then
#     echo "Go is not installed. Installing Go..."
#     sudo apt update
#     sudo apt install -y golang-go
# else
#     echo "Go is already installed"
# fi

# Verify Go installation
go version

rm -rf vendor
go clean -modcache

# Download project dependencies
echo "Tidying and downloading dependencies..."
sudo go mod tidy

# Build the project binary (optional but helpful for deployment)
echo "Building the Go project..."
go build -o app .

echo "Dependencies installed and project built successfully."