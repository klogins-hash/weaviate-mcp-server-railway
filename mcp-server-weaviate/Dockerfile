# Use official Go image as builder
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcp-server .

# Use minimal alpine image for final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN adduser -D -s /bin/sh appuser

WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/mcp-server .

# Change ownership to non-root user
RUN chown appuser:appuser mcp-server

# Switch to non-root user
USER appuser

# Expose port (Railway will override this with PORT env var)
EXPOSE 8675

# Run the binary
CMD ["./mcp-server"]
