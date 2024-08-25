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

# Build for Linux
echo "Building Binary"
go mod tidy
CGO_ENABLED=1 go build -o output/$APP_NAME