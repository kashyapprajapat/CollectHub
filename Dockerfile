# 1. Use Go base image
FROM golang:1.21-alpine

# 2. Set working directory
WORKDIR /app

# 3. Copy go.mod and go.sum
COPY go.mod go.sum ./

# 4. Download dependencies
RUN go mod download

# 5. Copy the entire project
COPY . .

# 6. Build the Go app
RUN go build -o main .

# 7. Command to run the executable
CMD ["/app/main"]
