# Flickly

Flickly é uma API em Go que implementa um sistema de gerenciamento de usuários utilizando padrões de arquitetura limpa e CQRS (Command Query Responsibility Segregation).

## Estrutura do Projeto

O projeto segue uma arquitetura em camadas:

- **cmd**: Pontos de entrada da aplicação
- **internal**: Código interno do projeto
  - **api**: Implementação da API REST
  - **domain**: Lógica de negócios e entidades
  - **infra**: Infraestrutura (banco de dados, IoC, etc.)
  - **integration_tests**: Testes de integração

## Requisitos

- Go 1.24 ou superior
- Docker e Docker Compose (para testes de integração)

## Como executar

### Localmente

```bash
# Clonar o repositório
git clone https://github.com/seu-usuario/flickly.git
cd flickly

# Executar a aplicação
go run cmd/flickly/main.go
```

A aplicação estará disponível em `http://localhost:8080`.

### Com Docker

```bash
# Construir e executar com Docker
docker build -t flickly .
docker run -p 8080:8080 flickly
```

## Testes

### Executar testes unitários

```bash
go test -v -short ./...
```

### Executar testes de integração

```bash
go test -v ./internal/integration_tests
```

### Executar testes com Docker Compose

```bash
docker-compose -f docker-compose.test.yml up --build
```

## Endpoints da API

### Saúde da aplicação

```
GET /health
```

Resposta:
```json
{
  "status": "ok",
  "service": "flickly"
}
```

### Versão da API

```
GET /api/flickly/version
```

Resposta:
```json
{
  "version": "1.0.0",
  "api": "flickly"
}
```

### Criar Usuário

```
POST /user
```

Corpo da requisição:
```json
{
  "name": "Nome do Usuário",
  "email": "usuario@example.com",
  "password": "senha123"
}
```

### Autenticar Usuário

```
POST /oauth/token
```

Parâmetros:
- `grant_type`: "password"
- `client_id`: "my_client_id"
- `client_secret`: "my_client_secret"
- `username`: Email do usuário
- `password`: Senha do usuário

## CI/CD

O projeto utiliza GitHub Actions para automação de CI/CD. O pipeline inclui:

1. Verificação de linting
2. Execução de testes unitários
3. Execução de testes de integração
4. Build da aplicação

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## Documentação da API com Swagger

O projeto utiliza Swagger para documentação automática da API.

### Como acessar a documentação

A documentação Swagger está disponível em:
```
http://localhost:8080/swagger/index.html
```
Quando o servidor estiver em execução.

### Como atualizar a documentação

A documentação do Swagger é gerada automaticamente com base nos comentários do código.
Se fizer alterações nas structs ou endpoints, é necessário regenerar a documentação:

1. Execute o script de atualização da documentação:
```
./update_swagger.sh
```

Este script irá:
- Regenerar a documentação Swagger
- Iniciar o servidor automaticamente
- Abrir a interface Swagger UI no navegador

### Estrutura da documentação

A documentação é estruturada com base nos comentários especiais no formato Swagger nos controladores.
Para adicionar novos endpoints à documentação, adicione comentários no formato apropriado:

```go
// @Summary Nome da operação
// @Description Descrição detalhada
// @Tags categoria
// @Accept json
// @Produce json
// @Param parametro body TipoStruct true "Descrição"
// @Success 200 {object} TipoRetorno
// @Failure 400 {object} interface{}
// @Router /caminho [método]
```