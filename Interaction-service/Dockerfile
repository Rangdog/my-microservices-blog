From golang:1.24.1-alpine AS base
WORKDIR /app
COPY interaction-service/go.mod interaction-service/go.sum ./
RUN go mod download
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o interaction-service ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/interaction-service .
EXPOSE 8080
CMD ["./interaction-service"]