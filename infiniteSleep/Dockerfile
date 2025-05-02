# Use official Go base image
FROM golang:1.21-alpine as builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first for better layer caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app statically
RUN go build -o server .

# Use a minimal base image for production
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./server"]
