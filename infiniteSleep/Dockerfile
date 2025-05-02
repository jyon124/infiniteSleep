# --- Build Stage ---
FROM golang:1.24.2 AS builder

WORKDIR /app

# Copy go.mod first
COPY go.mod ./

# Ensure go.sum exists by running go mod tidy
RUN go mod tidy

# Uncomment only if main.go have third-party dependencies module. If project relies on only standard library, this line is not required.
# COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build statically (safe for small base images if needed)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server .

# --- Runtime Stage ---
FROM golang:1.24.2 AS runner

WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
