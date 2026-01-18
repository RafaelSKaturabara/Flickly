# Isolando Variáveis de Ambiente do Go em Projeto Pessoal com direnv

**Data da análise:** 2025-07-06

## Contexto

Em ambientes corporativos, variáveis como `GOPROXY`, `GONOPROXY` e `GOPRIVATE` costumam ser configuradas globalmente (via shell ou arquivo GOENV) para acesso a módulos privados. Para projetos pessoais, é fundamental isolar essas configurações para evitar conflitos e garantir acesso ao proxy público do Go.

## Objetivo

Permitir que um projeto pessoal em Go utilize variáveis de ambiente próprias, sem impactar o ambiente global ou projetos corporativos, usando o [direnv](https://direnv.net/).

## Passo a Passo

### 1. Instale e configure o direnv

- Instale o direnv (caso ainda não tenha):
  ```sh
  brew install direnv
  ```
- Adicione o hook do direnv ao seu shell. No zsh, adicione ao final do `~/.zshrc`:
  ```sh
  eval "$(direnv hook zsh)"
  ```
  **Atenção:** Certifique-se de que está em uma linha separada dos demais comandos. Exemplo correto:
  ```sh
  export PATH="$PATH:$HOME/go/bin"
  eval "$(direnv hook zsh)"
  ```
- Reinicie o terminal ou rode `source ~/.zshrc`.

### 2. Crie o arquivo `.envrc` no diretório do projeto pessoal

Na raiz do projeto, crie um arquivo chamado `.envrc` com o seguinte conteúdo:
```sh
export GOPROXY="https://proxy.golang.org,direct"
unset GONOPROXY
unset GOPRIVATE
```

### 3. Autorize o direnv no projeto

No terminal, dentro da pasta do projeto, execute:
```sh
direnv allow
```

### 4. Valide o isolamento

- Verifique se o direnv está carregando o `.envrc`:
  ```sh
  direnv status
  ```
  Deve aparecer `Loaded RC path .../.envrc`.
- Verifique se as variáveis do Go foram sobrescritas:
  ```sh
  go env | grep -E 'GOPROXY|GONOPROXY|GOPRIVATE'
  ```
  Saída esperada:
  ```
  GOPROXY='https://proxy.golang.org,direct'
  GONOPROXY=''
  GOPRIVATE=''
  ```

## Solução para problemas comuns

- **direnv não carrega o .envrc:**
  - Verifique se o hook está correto no `~/.zshrc`.
  - Certifique-se de abrir um novo terminal dentro da pasta do projeto.
- **Variáveis do Go não mudam:**
  - Confirme se não há erro de sintaxe no `.envrc`.
  - Use sempre `export` para definir variáveis e `unset` para limpar.
- **Ambiente global não deve ser alterado:**
  - Nunca use `go env -w` para variáveis do projeto pessoal.
  - O isolamento é garantido pelo direnv e pelo `.envrc` local.

## Referências
- [direnv - Documentação Oficial](https://direnv.net/docs/)
- [Go - Environment Variables](https://pkg.go.dev/cmd/go#hdr-Environment_variables)

---

**Análise baseada nos comandos, arquivos e ambiente do usuário em 2025-07-06.**

