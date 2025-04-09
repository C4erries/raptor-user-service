# Используем базовый образ Go
FROM golang:1.23.0-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Собираем приложение
RUN go build -o user-service ./cmd/user-service

# alpine -- легковесный образ linux для контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

COPY static/ /app/static/

# Копируем собранное приложение
COPY --from=builder /app/user-service .

# Открываем порт для запросов
EXPOSE ${APP_PORT}

# Запускаем приложение
CMD ["./user-service"]