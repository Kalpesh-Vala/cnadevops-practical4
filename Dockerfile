# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy only the go.mod and go.sum files to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:3.19

# Add basic security patches
RUN apk --no-cache add ca-certificates && \
    adduser -D -H -h /app appuser

WORKDIR /app

# Copy only the compiled binary from the build stage
COPY --from=builder /app/main .

# Use non-root user for security
USER appuser

# Expose the application port
EXPOSE 8081

# Run the application
CMD ["./main"]
