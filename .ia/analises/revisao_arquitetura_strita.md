# Revisão Estrita de Arquitetura (DDD & Clean Architecture)

Este documento analisa a estrutura atual sob a ótica estrita dos padrões **Clean Architecture** e **DDD**, confrontando com o diagrama fornecido (Modelo N-Camadas/DDD).

## 1. O Dilema do nome "Services"
Você perguntou o que estaria "errado" no diagrama com base no DDD didático. A resposta é: **Ambiguidade Terminológica**.

### No "DDD Puro" (Eric Evans / Vaughn Vernon):
*   **Presentation Layer**: Onde ficam os Controllers, HTTP Handlers, CLI.
*   **Application Layer**: Onde ficam os Application Services (Orquestração).
*   **Domain Layer**: Onde ficam os Domain Services (Regras de Negócio).

### No seu Diagrama (Modelo Distribuído / N-Tier):
*   **Apresentação**: O Cliente (Frontend/Mobile).
*   **Serviços**: A Interface Remota (Sua API Backend).
*   **Aplicação**: O Backend interno.

**Conclusão**: O diagrama não está "errado", mas usa uma nomenclatura mais comum em arquiteturas .NET/Enterprise distribuídas.
No contexto de **Go** e **Clean Architecture**, chamar a camada de API de `services` é perigoso porque confunde com `Domain Services`.

**Decisão de Design**: Adotaremos o termo **PRESENTATION** para a camada de API.
*   **Nome Escolhido**: `internal/presentation`.
*   **Motivo**: Elimina a confusão com "Domain Services" e alinha com a literatura de Clean Architecture.

---

## 2. Diagnóstico da Estrutura Atual

Independente do nome da pasta (`services` ou `presentation`), o problema estrutural real é o **deslocamento de responsabilidades**.

| Camada no Diagrama | Onde Deveria Estar | Onde Está Hoje | Status |
| :--- | :--- | :--- | :--- |
| **02 - Serviços (API)** | `internal/services` | `internal/services` | ✅ **Correto** (Contém Handlers/Router) |
| **03 - Aplicação** | `internal/application` | ❌ **Inexistente** | 🔴 **Crítico** |
| **04 - Domínio** | `internal/domain` | `internal/domain` | 🟡 **Poluído** |

### O Problema Crítico (Missing Layer)
Seu diagrama mostra explicitamente uma camada **Aplicação** (roxa) entre a API e o Domínio.
No seu código, **essa camada não existe**.

1.  A API (`internal/services`) chama direto o Domínio?
    *   Não, ela chama `command_handlers`.
2.  Onde estão os `command_handlers`?
    *   Estão dentro de `internal/domain/identity/command_handlers`.

**Erro de Conceito**:
*   `Command Handlers` **SÃO** a camada de Aplicação.
*   Ao colocá-los dentro de `domain`, você fundiu as camadas 03 e 04.
*   **Consequência**: Seu Domínio ficou "sujo" com orquestração, dependências e DTOs.

## 3. Plano de Correção (Alinhamento ao Diagrama)

Para seguir o seu diagrama fielmente:

1.  **Renomear a API**:
    *   `internal/services` -> `internal/presentation`.

2.  **Criar a Camada de Aplicação**:
    *   Criar a pasta `internal/application`.
    *   Mover tudo que é "Orquestração" do Domínio para lá:
        *   `internal/domain/identity/command_handlers` -> `internal/application/identity/handlers`
        *   `internal/domain/identity/services` -> `internal/application/identity/services` (Se forem serviços de orquestração como JWT).

3.  **Manter no Domínio** apenas as regras puras:
    *   Entities, Value Objects, Repository Interfaces.

Dessa forma, você terá exatamente as 3 camadas do backend do diagrama:
1.  **Presentation** (API)
2.  **Application** (Handlers/Orchestration)
3.  **Domain** (Core Logic)

---

## 4. Dúvida: Command Handlers vs Clean Architecture UseCases

Você questionou se a estrutura atual (baseada em Commands) foge da Clean Architecture, especialmente ao comparar *Use Cases* com *Command Handlers*.

**Resposta: Não. Elas são conceptualmente idênticas e compatíveis.**

Abaixo, o dicionário de tradução entre os termos que você vê no seu código (estilo CQRS) e os termos da Clean Architecture:

| Termo Clean Architecture | Seu Código (DDD/CQRS) | Função |
| :--- | :--- | :--- |
| **Use Case Interactor** | `CommandHandler` (ex: `CreateUserCommandHandler`) | Contém a regra de orquestração (Application Business Rules). Ele recebe o input e coordena o Domínio. |
| **Request Model** | `Command` (ex: `CreateUserCommand`) | Estrutura de dados simples (DTO) que transporta o input do usuário para o Handler. |
| **Output Port / Response Model** | `mediator.Response` | O resultado devolvido para o Presenter/Controller. |

### Conclusão
Ao usar o padrão **Command**, você está efetivamente implementando **Use Cases**.
*   A diferença é apenas o estilo de agrupamento: em vez de ter um `UserService` "gigante" com métodos `Create`, `Update`, `Delete`, você explode cada método em uma classe separada (`CreateHandler`, `UpdateHandler`).
*   Isso favorece o **SRP (Princípio da Responsabilidade Única)**: cada arquivo tem apenas um motivo para mudar.
*   Portanto, **Command Handler = Use Case**. Seu projeto está correto em relação à arquitetura limpa, apenas usa uma nomenclatura mais específica (CQRS).
