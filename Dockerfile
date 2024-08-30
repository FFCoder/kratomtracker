# Build stage
FROM golang:1.22-bullseye AS build

LABEL authors="Jonathon Chambers"
LABEL description="This is the Dockerfile for setting up the Kratom Tracker Application"

# Set the Current Working Directory inside the container
WORKDIR /app

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - && \
    apt-get install -y nodejs

# Validate Node and NPM installation
RUN node -v && npm -v

# Copy the source
COPY . .

# Build the application
RUN chmod +x build_prod.sh && \
    ./build_prod.sh

# Final stage
FROM debian:bullseye-slim

WORKDIR /app

# Copy the build output from the build stage
COPY --from=build /app/output/kratom_tracker /app/kratom_tracker

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["/app/kratom_tracker"]