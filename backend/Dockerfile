FROM golang:1.24-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем air (live-reload)
RUN go install github.com/air-verse/air@latest

# Копируем go.mod и go.sum, чтобы перенести зависимости
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем все оставшиеся файлы директории
COPY . .

# Открываем порт
EXPOSE 8080

# Запускаем air с конфигурационным файлом
CMD ["air", "-c", ".air.toml"]

