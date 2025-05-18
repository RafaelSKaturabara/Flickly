# Testes de Integração - Flickly API

Este diretório contém testes de integração para o projeto Flickly. Os testes de integração verificam se diferentes componentes do sistema funcionam corretamente quando integrados.

## Estrutura dos Testes

1. **api_integration_test.go** - Testa a integração entre os endpoints da API.
2. **data_integration_test.go** - Testa a integração entre o domínio e a camada de dados.
3. **full_flow_integration_test.go** - Testa fluxos completos do aplicativo, desde o registro até a autenticação e acesso a recursos protegidos.

## Como Executar os Testes

### Executar Todos os Testes de Integração

```bash
go test ./internal/integration_tests -v
```

### Executar um Teste Específico

```bash
go test ./internal/integration_tests -run TestRunSuite -v
```

### Executar Testes com Tempo Limitado

```bash
go test ./internal/integration_tests -timeout 30s -v
```

### Ignorar Testes de Integração

Usar o flag `-short` para ignorar os testes de integração (útil durante o desenvolvimento):

```bash
go test ./... -short
```

## Configuração do Ambiente de Teste

Os testes são configurados para usar o ambiente real da aplicação. Se você precisar configurar um ambiente específico para testes, edite as funções `setup()` e `teardown()` no arquivo `integration_test.go`.

## Dependências

Os testes de integração dependem dos mesmos componentes que a aplicação principal, incluindo:

- Framework Gin
- Mediator
- Repositórios
- Serviços de domínio

## Manutenção

Ao adicionar novos componentes ao sistema, lembre-se de adicionar testes de integração correspondentes para garantir que eles funcionem corretamente com os componentes existentes. 