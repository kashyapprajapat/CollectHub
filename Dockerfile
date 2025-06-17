# -------- Stage 1: Build --------
FROM golang:1.24-alpine AS builder

# Set environment for static build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the full source code
COPY . .

# Build the Go binary statically
RUN go build -a -installsuffix cgo -o main .

# -------- Stage 2: Run --------
FROM alpine:latest

# Install certs (for HTTPS env & Mongo URI)
RUN apk --no-cache add ca-certificates

# Set working directory inside container
WORKDIR /root/

# Copy built Go binary from builder
COPY --from=builder /app/main .

# Expose your actual app port
EXPOSE 10000

# Run the binary
CMD ["./main"]
