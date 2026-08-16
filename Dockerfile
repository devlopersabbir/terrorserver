# Stage 1: Build binary using Go toolchain
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o terror ./cmd/terror

# Stage 2: Minimal runtime image
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/terror /usr/local/bin/terror
COPY templates/welcome.html /var/www/terrorserver/index.html
COPY templates/Runtime /etc/terror/Runtime

# Expose HTTP & HTTPS ports
EXPOSE 80 443

# Persistence volumes
VOLUME ["/etc/terror", "/var/lib/terror/certs", "/var/www"]

ENTRYPOINT ["/usr/local/bin/terror"]
CMD ["serve"]
