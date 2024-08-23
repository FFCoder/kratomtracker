#!/bin/bash

# Set some variables
APP_NAME="kratom_tracker"
VITE_API_URL="/api"
CGO_ENABLED=1 # Required for sqlite3

# Make the output
mkdir -p output

# Clean the output
rm -rf output/*

# Build the frontend
cd frontend
npm install
npm run build:prod

# Build the backend
cd ..

# Build for Mac M1, M2, etc
echo "Building for Mac M1, M2, etc"
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o output/$APP_NAME.mac.arm64

# Build for Mac Intel
echo "Building for Mac Intel"
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o output/$APP_NAME.mac.amd64

# Build for Linux
echo "Building for Linux"
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o output/$APP_NAME.linux.amd64