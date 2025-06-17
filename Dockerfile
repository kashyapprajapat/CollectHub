# Use official Go 1.24 image
FROM golang:1.24 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the app
RUN go build -o main .

# Use a minimal image to run the app
FROM debian:bullseye-slim
WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/main .

# Copy .env if needed
# COPY --from=builder /app/.env .

# Set environment variable to production
ENV FIBER_ENV=production

# Expose port
EXPOSE 10000

# Run the app
CMD ["./main"]
