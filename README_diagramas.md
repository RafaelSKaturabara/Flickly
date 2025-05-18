# Diagramas do Projeto Flickly

Este documento contém diagramas que representam a arquitetura e as dependências do projeto Flickly.

## Como visualizar os diagramas

Os diagramas estão no formato PlantUML e podem ser visualizados de várias maneiras:

1. **Online**: Copie o conteúdo dos arquivos .puml e cole no [PlantUML Web Server](https://www.plantuml.com/plantuml/uml/)

2. **VS Code**: Instale a extensão "PlantUML" para VS Code e use a pré-visualização.

3. **IntelliJ/IDEA**: Instale o plugin "PlantUML integration" e use a pré-visualização.

4. **Linha de comando**: Use a ferramenta `plantuml` para gerar imagens:
   ```
   plantuml package_diagram.puml
   plantuml architecture_diagram.puml
   plantuml cqrs_flow_diagram.puml
   ```

## Descrição dos Diagramas

### 1. Diagrama de Pacotes e Dependências (package_diagram.puml)

Este diagrama mostra a estrutura detalhada de todos os pacotes do projeto e suas dependências. Ele é útil para entender como os diferentes componentes do sistema estão organizados e como eles dependem uns dos outros.

### 2. Diagrama Arquitetural (architecture_diagram.puml)

Este diagrama fornece uma visão de alto nível da arquitetura do sistema, mostrando as principais camadas e suas relações. É uma visão simplificada que facilita a compreensão da estrutura geral do projeto.

### 3. Diagrama de Fluxo CQRS (cqrs_flow_diagram.puml)

Este diagrama de sequência ilustra o fluxo de execução do padrão CQRS (Command Query Responsibility Segregation) implementado no projeto, especificamente para o caso de uso de criação de usuário. Ele mostra como os comandos são processados através do sistema, desde o controlador até o repositório, passando pelo mediador e handlers.

## Estrutura da Aplicação

A aplicação segue uma arquitetura em camadas com os seguintes componentes principais:

1. **Camada de API (api/)**: Contém os controladores e rotas da aplicação.
   - **flickly/**: Endpoints relacionados ao core da aplicação
   - **users/**: Endpoints relacionados aos usuários

2. **Camada de Domínio (domain/)**: Contém a lógica de negócios e entidades.
   - **core/**: Componentes base usados por todo o domínio
   - **users/**: Domínio específico para usuários
   - **flickly/**: Domínio específico para funcionalidades do Flickly

3. **Camada de Infraestrutura (infra/)**: Contém implementações concretas e utilitários.
   - **cross-cutting/**: Componentes utilizados por todas as camadas
   - **data/**: Implementações de repositórios

4. **Testes de Integração (integration_tests/)**: Testes que validam a integração entre os componentes. 