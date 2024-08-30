#!/bin/bash

# Set some variables
export APP_NAME="kratom_tracker"
export VITE_API_URL="/api"
export CGO_ENABLED=1 # Required for sqlite3

# Validate pnpm is installed
if ! command -v pnpm &> /dev/null
then
    echo "pnpm could not be found"
    echo "attempting to install pnpm"
    npm install -g pnpm
    if ! command -v pnpm &> /dev/null
    then
        echo "pnpm failed to install"
        exit 1
    fi
fi

# Make the output
mkdir -p output

# Clean the output
rm -rf output/*

# Build the frontend
cd frontend || exit 1

pnpm install
npm run build:prod

# Build the backend
cd ..

# Build for Linux
echo "Building Binary"
go mod tidy
go build -o output/$APP_NAME