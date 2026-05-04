# Build stage
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o server .

# Final stage: lightweight image
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]
