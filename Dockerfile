# Используем официальный образ Golang для сборки
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum файлы для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь проект в контейнер
COPY . .

# Сборка исполняемого файла
RUN go build -o main ./cmd/main/main.go

# Минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /root/

# Копируем скомпилированное приложение и файл конфигурации
COPY --from=builder /app/main .
COPY --from=builder /app/config/config.yaml ./config/

# Открываем порт, который будет использоваться приложением
EXPOSE 8080

# Команда запуска приложения
CMD ["./main"]
