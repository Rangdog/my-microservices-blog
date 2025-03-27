From golang:1.24.1-alpine AS base
WORKDIR /app
COPY user-service/go.mod user-service/go.sum ./
COPY go.mod /go.mod
RUN go mod download

FROM base AS builder
WORKDIR /app
COPY user-service .  # Chỉ copy thư mục user-service vào container
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o user-service ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user-service/user-service .
EXPOSE 8000
CMD ["./user-service"]