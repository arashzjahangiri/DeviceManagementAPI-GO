# Use the official Go image to build the application
FROM golang:1.24.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /device-management-api ./cmd/api/main.go

# Use a minimal base image for the final container
FROM alpine:latest

WORKDIR /root/
# Copy the compiled binary from the builder stage
COPY --from=builder /device-management-api /device-management-api

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/device-management-api"]