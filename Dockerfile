# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git ca-certificates tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -extldflags '-static'" -o /bin/jawad cmd/jawad/main.go

# Final minimal production stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -g 10001 -S jawad \
    && adduser -u 10001 -S jawad -G jawad

WORKDIR /app

RUN mkdir -p /app/.ssh /app/configs && chown -R jawad:jawad /app

COPY --from=builder --chown=jawad:jawad /bin/jawad /app/jawad
COPY --from=builder --chown=jawad:jawad /src/configs/portfolio.yaml /app/configs/portfolio.yaml

USER jawad

# Persistent volume for SSH host keys and custom configuration
VOLUME ["/app/.ssh", "/app/configs"]

EXPOSE 2222

ENV PORT=2222 \
    SSH_HOST=0.0.0.0 \
    SSH_KEY_PATH=/app/.ssh/id_ed25519

ENTRYPOINT ["/app/jawad", "-ssh"]
