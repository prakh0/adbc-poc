# Multi-stage build to keep the final image lightweight
# Using 'golang' as a base for building, compatible with the project's Go version
FROM golang:bookworm AS builder

# Set working directory
WORKDIR /app

# Copy dependency files and download
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
# CGO_ENABLED=1 is often necessary for ADBC driver management
RUN CGO_ENABLED=1 GOOS=linux go build -o main main.go

# Final runtime stage
FROM debian:bookworm-slim

# Install minimal dependencies for the installer and runtime
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# 2. Install dbc using the provided command
# We set XDG_BIN_HOME to /usr/local/bin so it's globally available
RUN curl -LsSf https://dbc.columnar.tech/install.sh | sh

# 3. Install mysql driver with dbc
RUN $HOME/.local/bin/dbc install mysql

# Copy the built binary from the builder stage
COPY --from=builder /app/main /usr/local/bin/main

# Ensure the binary is executable
RUN chmod +x /usr/local/bin/main

# Run the application
ENTRYPOINT ["/usr/local/bin/main"]
