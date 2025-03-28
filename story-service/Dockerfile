From golang:1.24.1-alpine AS base
WORKDIR /app
COPY story-service/go.mod story-service/go.sum ./
COPY go.mod /go.mod
RUN go mod download
FROM base AS builder
COPY story-service .
COPY pkg/discovery /pkg/discovery
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o story-service ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/story-service .
EXPOSE 8100
CMD ["./story-service"]