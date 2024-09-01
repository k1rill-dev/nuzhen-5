# Используем официальный образ Golang версии 1.22 для сборки
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Перемещаемся в директорию с исходным кодом
WORKDIR /app/cmd/main

# Сборка приложения
RUN go build -o /app/myapp

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем нужные зависимости для запуска
RUN apk --no-cache add ca-certificates

# Копируем собранное приложение из образа builder
COPY --from=builder /app/myapp /app/myapp

# Задаем рабочую директорию
WORKDIR /app

# Указываем команду запуска приложения
CMD ["./myapp"]
