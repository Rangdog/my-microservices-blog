From golang:1.24.1-alpine AS base
WORKDIR /app
COPY user-service/go.mod ./
COPY user-service/go.sum ./
RUN go mod download
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o user-service ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user-service .
EXPOSE 8080
CMD ["./user-service"]