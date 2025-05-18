.PHONY: run test coverage swagger clean build clean-swagger lint

# Variáveis
BINARY_NAME=flickly
MAIN_PATH=cmd/flickly/main.go
GO_PATH=$(shell go env GOPATH)
SWAG_PATH=$(shell which swag 2>/dev/null || echo "$(GOPATH)/bin/swag")

# Instalar dependências
deps:
	@echo "Instalando dependências..."
	go mod download
	go mod tidy
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "Instalando golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2; \
	fi

# Rodar a aplicação
run: swagger
	@echo "Iniciando servidor..."
	go run $(MAIN_PATH)

# Rodar testes
test:
	@echo "Executando testes..."
	go test ./... -v

# Executar lint
lint:
	@echo "Executando lint..."
	@golangci-lint run --max-issues-per-linter=0 --max-same-issues=0 || (echo "Lint falhou!" && exit 1)

# Gerar relatório de cobertura
coverage:
	@echo "Gerando relatório de cobertura..."
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# Atualizar e exibir documentação Swagger
swagger:
	@echo "Verificando instalação do swag..."
	@if ! command -v swag &> /dev/null && ! [ -f "$(GO_PATH)/bin/swag" ]; then \
		echo "Instalando swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@echo "Gerando documentação Swagger..."
	@if command -v swag &> /dev/null; then \
		swag init -g $(MAIN_PATH) --parseDependency --parseInternal; \
	elif [ -f "$(GO_PATH)/bin/swag" ]; then \
		$(GO_PATH)/bin/swag init -g $(MAIN_PATH) --parseDependency --parseInternal; \
	else \
		echo "Erro: swag não encontrado no PATH ou em $(GO_PATH)/bin"; \
		exit 1; \
	fi
	@echo "Documentação atualizada. Inicie o servidor com 'make run' para visualizar."

# Limpar arquivos gerados
clean:
	@echo "Limpando binários..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out
	@echo "Limpeza concluída."

# Limpar arquivos do Swagger
clean-swagger:
	@echo "Limpando arquivos do Swagger..."
	@rm -rf docs
	@echo "Arquivos do Swagger removidos."

# Limpar tudo (binários e docs)
clean-all: clean clean-swagger
	@echo "Limpeza completa concluída."

# Construir o binário
build: swagger
	@echo "Construindo aplicação..."
	go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Binário $(BINARY_NAME) criado."

# Construir e executar
build-run: build
	@echo "Iniciando servidor..."
	./$(BINARY_NAME)

# Ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make deps          - Instalar dependências"
	@echo "  make run           - Executar a aplicação"
	@echo "  make test          - Executar testes"
	@echo "  make coverage      - Gerar relatório de cobertura de testes"
	@echo "  make swagger       - Gerar documentação Swagger"
	@echo "  make build         - Construir o binário"
	@echo "  make build-run     - Construir e executar o binário"
	@echo "  make clean         - Limpar arquivos gerados"
	@echo "  make clean-swagger - Limpar apenas arquivos do Swagger"
	@echo "  make clean-all     - Limpar todos os arquivos gerados"

# Comando padrão
default: help 