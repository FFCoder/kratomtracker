#!/bin/bash

# Set some variables
export APP_NAME="kratom_tracker"
export VITE_API_URL="/api"
export CGO_ENABLED=1 # Required for sqlite3

# Make the output
mkdir -p output

# Clean the output
rm -rf output/*

# Build the frontend
cd frontend || exit 1

# Remove package-lock.json
rm -f package-lock.json
#npm install --legacy-peer-deps
npm run build:prod

# Build the backend
cd ..

# Build for Linux
echo "Building Binary"
go mod tidy
go build -o output/$APP_NAME