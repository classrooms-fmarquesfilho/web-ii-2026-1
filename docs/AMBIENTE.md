# Configuração do Ambiente de Desenvolvimento

Este guia detalha a instalação e configuração de todas as ferramentas necessárias para o curso.

---

## Curso Básico de Go

Se você está fazendo o **[curso básico "Aprenda Go com Testes"](https://github.com/classrooms-fmarquesfilho/aprenda-go-com-testes)** antes de Web II, você só precisa de:

| Ferramenta | Versão | Obrigatório |
|------------|--------|-------------|
| Go | 1.23+ | Sim |
| Git | 2.40+ | Sim |
| Editor de código | - | Sim |

**Verificação rápida** (no diretório de cada exercício do curso básico):
```bash
go version  # Deve mostrar 1.23 ou superior
git --version
go test     # Deve rodar os testes
```

---

## Curso Web II - Verificação Rápida

Use o script de verificação para saber o que já está instalado:

```bash
cd scripts/check-env

# Verificar requisitos da Unidade 1 (padrão)
go run main.go

# Verificar requisitos da Unidade 2
go run main.go -u 2

# Verificar requisitos da Unidade 3
go run main.go -u 3
```

O script verifica as ferramentas necessárias para cada unidade e indica o que está faltando. Veja mais detalhes em [scripts/check-env/README.md](../scripts/check-env/README.md).

---

## Requisitos por Unidade

### Unidade 1 (Listas 1-3)

| Ferramenta | Versão | Obrigatório |
|------------|--------|-------------|
| Go | 1.25+ | ✅ Sim |
| Git | 2.40+ | ✅ Sim |
| golangci-lint | latest | ✅ Sim |
| VS Code (ou editor) | - | ✅ Sim |

### Unidade 2 (Listas 4-5)

| Ferramenta | Versão | Obrigatório |
|------------|--------|-------------|
| Docker | 24+ | ✅ Sim |
| Docker Compose | 2.x | ✅ Sim |
| PostgreSQL (via Docker) | 18 | ✅ Sim |
| sqlc | latest | ✅ Sim |

### Unidade 3 (Lista 6)

| Ferramenta | Versão | Obrigatório |
|------------|--------|-------------|
| Todas as anteriores | - | ✅ Sim |
| Air (hot reload) | latest | Recomendado |

---

## Instalação do Go

### Linux (Ubuntu/Debian)

```bash
# Remover versão antiga (se existir)
sudo rm -rf /usr/local/go

# Download e instalação (substitua pela versão mais recente)
wget https://go.dev/dl/go1.25.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.25.7.linux-amd64.tar.gz

# Adicionar ao PATH (~/.bashrc ou ~/.zshrc)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc

# Verificar instalação
go version
```

### macOS

```bash
# Via Homebrew 
brew install go

# Ou download direto
# https://go.dev/dl/

# Verificar instalação
go version
```

### Windows

1. Download: https://go.dev/dl/
2. Execute o instalador `.msi`
3. Reinicie o terminal
4. Verifique: `go version`

---

## Instalação do Git

### Linux

```bash
sudo apt update
sudo apt install git

git --version
```

### macOS

```bash
# Git já vem instalado, ou via Homebrew:
brew install git
```

### Windows

Download: https://git-scm.com/download/win

---

## Instalação do Docker

### Linux (Ubuntu)

```bash
# Instalação oficial
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Adicionar usuário ao grupo docker (evita usar sudo)
sudo usermod -aG docker $USER
newgrp docker

# Instalar Docker Compose plugin
sudo apt install docker-compose-plugin

# Verificar
docker --version
docker compose version
```

### macOS

1. Download: https://www.docker.com/products/docker-desktop/
2. Instale o Docker Desktop
3. Inicie o aplicativo
4. Verifique: `docker --version`

### Windows

1. Habilite WSL2 (Windows Subsystem for Linux)
2. Download: https://www.docker.com/products/docker-desktop/
3. Instale o Docker Desktop
4. Verifique: `docker --version`

---

## PostgreSQL via Docker 

O curso utiliza PostgreSQL via Docker para simplificar a configuração. Use o `docker-compose.yml` na raiz do repositório:

```bash
# Iniciar o banco de dados
docker compose up -d

# Verificar se está rodando
docker ps

# Conectar ao banco (para testes)
docker exec -it postgres-dev psql -U dev -d devdb

# Parar o banco
docker compose down
```

### Configuração do docker-compose.yml

O arquivo já está configurado com:
- **Usuário**: dev
- **Senha**: dev
- **Banco**: devdb
- **Porta**: 5432

---

## Ferramentas Go Adicionais 

```bash
# sqlc - Geração de código SQL type-safe 
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Verificar
sqlc version
```

### Ferramentas Recomendadas

```bash
# Air - Hot reload para desenvolvimento
go install github.com/air-verse/air@latest

# swag - Geração de documentação OpenAPI/Swagger
go install github.com/swaggo/swag/cmd/swag@latest

# mockery - Geração de mocks para testes
go install github.com/vektra/mockery/v2@latest

# Verificar
air -v
swag --version
mockery --version
```

### Atlas (Migrations)

```bash
# Linux/macOS
curl -sSf https://atlasgo.sh | sh

# Verificar
atlas version
```

---

## Configuração do Git

### Configuração Básica

```bash
# Identidade
git config --global user.name "Seu Nome"
git config --global user.email "seu.email@ufrn.br"

# Editor padrão
git config --global core.editor "code --wait"

# Branch padrão
git config --global init.defaultBranch main

# Aliases úteis
git config --global alias.st status
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.lg "log --oneline --graph --all"
```

### SSH Key para GitHub

```bash
# Gerar chave
ssh-keygen -t ed25519 -C "seu.email@ufrn.br"

# Iniciar ssh-agent
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519

# Copiar chave pública
cat ~/.ssh/id_ed25519.pub
# Cole em: GitHub → Settings → SSH and GPG keys → New SSH key

# Testar conexão
ssh -T git@github.com
```

---

## Troubleshooting

### "go: command not found"

```bash
# Verifique se Go está no PATH
echo $PATH | grep -q "go/bin" && echo "OK" || echo "Adicione ao PATH"

# Adicione ao ~/.bashrc ou ~/.zshrc
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bashrc
```

### "golangci-lint: command not found"

```bash
# Verifique se GOPATH/bin está no PATH
echo $PATH | grep -q "$(go env GOPATH)/bin" && echo "OK" || echo "Faltando"

# Adicione ao ~/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bashrc

# Reinstale se necessário
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### "permission denied" no Docker

```bash
# Adicione usuário ao grupo docker
sudo usermod -aG docker $USER

# Faça logout e login novamente (ou reinicie)
# Teste novamente
docker ps
```

### PostgreSQL "connection refused"

```bash
# Verifique se o container está rodando
docker ps | grep postgres

# Reinicie se necessário
docker compose down
docker compose up -d

# Verifique logs
docker logs postgres-dev
```
