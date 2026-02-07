# Verificação do Ambiente - DIM0547

Script para verificar se todas as ferramentas necessárias para o curso estão instaladas corretamente.

## Uso

```bash
cd scripts/check-env

# Verificar requisitos da Unidade 1 (padrão)
go run main.go

# Verificar requisitos da Unidade 2
go run main.go -u 2

# Verificar requisitos da Unidade 3
go run main.go -u 3
```

Ou compile e execute:

```bash
go build -o check-env
./check-env        # U1
./check-env -u 2   # U2
./check-env -u 3   # U3
```

## Requisitos por Unidade

### Unidade 1 (Listas 1-3)

**Obrigatórios:**
- **Go** (1.25+) - Linguagem de programação
- **Git** - Controle de versão
- **golangci-lint** - Linter para Go

**Recomendados:**
- **Docker** - Preparação para U2

### Unidade 2 (Listas 4-5)

**Obrigatórios (além dos da U1):**
- **Docker** - Containerização
- **Docker Compose** - Orquestração de containers
- **sqlc** - Geração de código SQL type-safe

**Recomendados:**
- **Atlas** - Gerenciamento de migrations
- **Air** - Hot reload para desenvolvimento

### Unidade 3 (Lista 6)

**Obrigatórios:** Todos os anteriores

**Recomendados:**
- **swag** - Geração de documentação Swagger/OpenAPI
- **mockery** - Geração de mocks para testes

## Exemplo de Saída (Unidade 1)

```
╔════════════════════════════════════════════════════════════╗
║    DIM0547 - Verificação do Ambiente de Desenvolvimento    ║
╚════════════════════════════════════════════════════════════╝

📚 Verificando requisitos para: Unidade 1

📋 Informações do Sistema
--------------------------------------------------
   Go Version:    go1.25.7
   OS:            linux
   Arquitetura:   amd64
   GOPATH:        /home/user/go
   GOROOT:        /usr/local/go

🔍 Verificação de Ferramentas
--------------------------------------------------
   ✓ go                 (obrigatório) go version go1.25.7 linux/amd64
   ✓ git                (obrigatório) git version 2.43.0
   ✓ golangci-lint      (obrigatório) golangci-lint has version 1.56.0
   ✓ docker             (recomendado) Docker version 25.0.0
   ○ air                (recomendado) - não encontrado
   ○ swag               (recomendado) - não encontrado
   ○ mockery            (recomendado) - não encontrado

📊 Resumo
--------------------------------------------------
   ✓ Todas as ferramentas obrigatórias para U1 estão instaladas!
   ○ Ferramentas recomendadas faltando: air, swag, mockery

   Para instalar as ferramentas recomendadas:

     go install github.com/air-verse/air@latest
     go install github.com/swaggo/swag/cmd/swag@latest
     go install github.com/vektra/mockery/v2@latest

   💡 Dica: Para verificar requisitos da U2, execute:
      go run main.go -u 2
```

## Exemplo de Saída (Unidade 2)

```
📚 Verificando requisitos para: Unidade 2

🔍 Verificação de Ferramentas
--------------------------------------------------
   ✓ go                 (obrigatório) go version go1.25.7 linux/amd64
   ✓ git                (obrigatório) git version 2.43.0
   ✓ golangci-lint      (obrigatório) golangci-lint has version 1.56.0
   ✓ docker             (obrigatório) Docker version 25.0.0
   ✓ docker compose     (obrigatório) Docker Compose version v2.24.0
   ✓ sqlc               (obrigatório) v1.25.0
   ○ atlas              (recomendado) - não encontrado
   ✓ air                (recomendado) v1.51.0
   ...

🐘 Verificação do PostgreSQL
--------------------------------------------------
   ✓ PostgreSQL acessível via Docker
```

## Legenda

- ✓ (verde) - Instalado e funcionando
- ✗ (vermelho) - Obrigatório, não encontrado
- ○ (amarelo) - Recomendado, não encontrado

## Código de Saída

- `0` - Todas as ferramentas obrigatórias para a unidade especificada estão instaladas
- `1` - Alguma ferramenta obrigatória faltando

## Integração com CI

Você pode usar este script em pipelines de CI:

```bash
# Falha se ferramentas obrigatórias da U1 não estiverem instaladas
go run main.go -u 1 || exit 1
```