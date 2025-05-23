# Документация для проекта raptor-user-service

## Описание проекта

Проект `raptor-user-service` представляет собой сервис для управления пользователями. Сервис предоставляет следующие функции:

- Получение информации о пользователе
- Создание нового пользователя
- Обновление информации о пользователе
- Удаление пользователя

## Структура проекта

Проект состоит из следующих основных компонентов:

- `cmd/user-service/main.go`: Точка входа в приложение.
- `internal/app/app.go`: Основной модуль приложения.
- `internal/app/gprc/app.go`: Модуль для работы с gRPC сервером.
- `internal/config/config.go`: Модуль для загрузки конфигурации приложения.
- `internal/database/postgres/postgres.go`: Модуль для работы с базой данных PostgreSQL.
- `internal/model/user.go`: Модуль для работы с моделью пользователя.
- `internal/server/userinfo/server.go`: Модуль для работы с gRPC сервером.
- `internal/service/userinfo/service.go`: Модуль для работы с информацией о пользователях.

## Запуск приложения

Для запуска приложения выполните следующие шаги:

1. Установите зависимости проекта с помощью `go mod download`.
2. Запустите приложение с помощью команды `go run cmd/user-service/main.go`.

## API

API сервиса предоставляет следующие методы:

- `GetUser`: Получение информации о пользователе.
- `CreateUser`: Создание нового пользователя.
- `UpdateUser`: Обновление информации о пользователе.
- `DeleteUser`: Удаление пользователя.

## База данных

Сервис использует базу данных PostgreSQL для хранения информации о пользователях. Подключение к базе данных осуществляется с помощью модуля `internal/database/postgres/postgres.go`.
