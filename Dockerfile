# Imagen base de Go
FROM golang:tip-alpine3.21

# Variables
WORKDIR /app

# Copiar Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto de la app
COPY . .

# Compilar binario
RUN go build -o main ./cmd/server

# Puerto en el que corre Echo
EXPOSE 8080

# Comando para ejecutar
CMD ["./main"]
