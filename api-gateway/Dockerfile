From golang:1.24.1-alpine AS base
WORKDIR /app
COPY api-gateway/go.mod api-gateway/go.sum ./
COPY go.mod /go.mod
RUN go mod download
FROM base AS builder
COPY api-gateway .
COPY pkg/discovery /pkg/discovery
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o api-gateway ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api-gateway .
EXPOSE 8080
CMD ["./api-gateway"]