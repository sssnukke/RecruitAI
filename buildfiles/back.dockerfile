# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make gcc musl-dev

# Копируем весь исходный код
COPY back .
RUN go mod download

# Сборка бинарника для Linux
RUN GOOS=linux GOARCH=amd64 go build -o back ./cmd/back/main.go && ls -l /app

# Stage 2: Run
FROM alpine:3.18

WORKDIR /app

# Устанавливаем бинарник в безопасное место
COPY --from=builder /app/back /usr/local/bin/back

COPY .env /app/.env

CMD ["/usr/local/bin/back"]