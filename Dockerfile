FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bot ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bot .
# Копируем необходимые файлы конфигурации
COPY .env .

# Запускаем бота
CMD ["./bot"]
