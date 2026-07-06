FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /seedao-api-server ./cmd/server

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /seedao-api-server /app/seedao-api-server

EXPOSE 8080

ENV CONFIG_PATH=/app/config/config.yaml

ENTRYPOINT ["/app/seedao-api-server"]
