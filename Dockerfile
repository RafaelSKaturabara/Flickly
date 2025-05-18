# Estágio de build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instalar dependências
RUN apk add --no-cache git ca-certificates tzdata

# Copiar arquivos de módulo e baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte
COPY . .

# Construir o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main ./cmd/flickly

# Estágio final
FROM alpine:latest

# Adicionar certificados e timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/main

# Porta da aplicação
EXPOSE 8080

# Comando para executar a aplicação
ENTRYPOINT ["/app/main"] 