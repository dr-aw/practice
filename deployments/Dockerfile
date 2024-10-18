# Dockerfile
FROM golang:1.23 as builder

WORKDIR /app
COPY go.mod go.sum ./

# Dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app-ololo ./cmd/app

# Alpine-linux
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app-ololo .
COPY --from=builder /app/templates ./templates

# Listening port
EXPOSE 8080

# Default run
CMD ["./app-ololo"]
