# Stage 1: Build the application
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /device-management-api ./cmd/api/main.go

# Stage 2: Create a lightweight image
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /device-management-api .

EXPOSE 8080
CMD ["./device-management-api"]