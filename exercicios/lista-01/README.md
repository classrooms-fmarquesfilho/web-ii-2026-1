# Lista 1: net/http Básico

## DIM0547 - Desenvolvimento de Sistemas Web II

**Prazo**: 13/03/2026 às 23:59  
**Peso**: 30% da nota de exercícios da U1  
**Entrega**: GitHub Classroom (link no Discord)

---

## Objetivos de Aprendizagem

Ao completar esta lista, você será capaz de:

1. Criar handlers HTTP usando a biblioteca padrão do Go
2. Implementar handlers usando `http.HandleFunc`
3. Trabalhar com `http.Request` (query parameters, headers, body)
4. Retornar respostas JSON com `encoding/json`
5. Usar `http.ServeMux` para roteamento básico

---

## Pré-requisitos

- Assistir ao **Vídeo 2**: O pacote net/http - Handler e ServeMux (disponível no Discord, canal #videos-semanais, até uma semana antes do início do curso)
- Completar o [Tour of Go em Português](https://go-tour-br.appspot.com/) (módulos 1-4)
- Ambiente configurado (Go 1.22+)

---

## Importante: Como os Testes Funcionam

Os testes desta lista usam `httptest`, que **não precisa de servidor rodando**. 

Seu código deve apenas **registrar os handlers** no `main()`. Exemplo:

```go
func main() {
    // ✅ CORRETO: apenas registra o handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
    
    // ❌ NÃO INCLUA http.ListenAndServe() - os testes não precisam disso
}
```

**Por quê?** Os testes chamam `main()` para registrar os handlers e depois usam `httptest.NewRequest()` para simular requisições. Se você incluir `http.ListenAndServe()`, o programa vai bloquear e os testes não vão funcionar.

**Para testar manualmente** (opcional), crie um arquivo separado ou use uma flag:

```go
// Para rodar manualmente: go run main.go -serve
func main() {
    http.HandleFunc("/", handler)
    
    if len(os.Args) > 1 && os.Args[1] == "-serve" {
        log.Fatal(http.ListenAndServe(":8080", nil))
    }
}
```

---

## Exercícios

### Exercício 1: Hello World Server (25 pontos)

Implemente um handler que responda "Hello, World!" na rota `/`.

**Arquivo**: `ex01/main.go`

**Requisitos**:
- Rota `GET /` retorna texto "Hello, World!"
- Content-Type: `text/plain` (opcional, Go define automaticamente)
- Status code: 200

**Teste local**:
```bash
cd ex01
go test -v ./...
```

---

### Exercício 2: Query Parameters (25 pontos)

Implemente um handler que leia query parameters e retorne uma saudação personalizada.

**Arquivo**: `ex02/main.go`

**Requisitos**:
- Rota `GET /greet` aceita parâmetro `name`
- Se `name` não for fornecido ou estiver vazio, usar "World"
- Retornar `Hello, {name}!`

**Exemplos esperados**:
```
GET /greet?name=Maria  →  "Hello, Maria!"
GET /greet?name=       →  "Hello, World!"
GET /greet             →  "Hello, World!"
```

**Teste local**:
```bash
cd ex02
go test -v ./...
```

---

### Exercício 3: Múltiplas Rotas (25 pontos)

Implemente um servidor com múltiplas rotas usando `http.HandleFunc`.

**Arquivo**: `ex03/main.go`

**Requisitos**:
- `GET /` → retorna "Home"
- `GET /about` → retorna "About"
- `GET /health` → retorna JSON `{"status":"ok"}` com Content-Type `application/json`
- Qualquer outra rota → retorna "Not Found" com status 404

**Dica**: Para retornar 404 em rotas não definidas, você precisa verificar `r.URL.Path` no handler de `/`.

**Teste local**:
```bash
cd ex03
go test -v ./...
```

---

### Exercício 4: JSON Echo (25 pontos)

Implemente um handler que receba JSON no body e retorne o mesmo JSON com um campo adicional.

**Arquivo**: `ex04/main.go`

**Requisitos**:
- Rota `POST /echo`
- Recebe JSON com qualquer estrutura
- Adiciona campo `"received": true` ao JSON
- Retorna o JSON modificado com Content-Type `application/json`
- Se o body não for JSON válido, retorna status 400 com `{"error":"invalid JSON"}`

**Exemplo de sucesso**:
```
Entrada:  {"message":"hello","count":42}
Saída:    {"count":42,"message":"hello","received":true}
```

**Exemplo de erro**:
```
Entrada:  "not json"
Saída:    {"error":"invalid JSON"} (status 400)
```

**Teste local**:
```bash
cd ex04
go test -v ./...
```

---

## Como Entregar

### 1. Aceite o Assignment

Clique no link do GitHub Classroom disponível no Discord (canal #exercicios).

### 2. Clone o Repositório

```bash
git clone https://github.com/classrooms-fmarquesfilho/lista-01-seu-usuario
cd lista-01-seu-usuario
```

### 3. Implemente as Soluções

Cada exercício tem sua pasta (`ex01/`, `ex02/`, etc.) com um arquivo `main.go` contendo as instruções.

### 4. Teste Localmente

```bash
# Testar um exercício específico
cd ex01
go test -v ./...

# Ou testar todos da raiz
go test ./...
```

### 5. Envie para Correção

```bash
git add .
git commit -m "Implementa exercícios da lista 1"
git push origin main
```

### 6. Verifique os Resultados

- Acesse a aba **Actions** no seu repositório GitHub
- O workflow mostrará os testes que passaram/falharam
- Você pode fazer múltiplos pushes até o prazo

---

## Critérios de Avaliação

| Critério | Peso |
|----------|------|
| Testes passando | 70% |
| Qualidade do código (golangci-lint) | 20% |
| Entrega no prazo | 10% |

### Política de Atraso

- Até 1 dia: -20%
- Até 3 dias: -50%
- Após 3 dias: Não aceito

---

## Dicas de Implementação

### Registrando um Handler

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello!"))
    })
}
```

### Lendo Query Parameters

```go
func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

### Retornando JSON

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    data := map[string]string{"status": "ok"}
    json.NewEncoder(w).Encode(data)
}
```

### Lendo JSON do Body

```go
func handler(w http.ResponseWriter, r *http.Request) {
    var data map[string]interface{}
    
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"error":"invalid JSON"}`))
        return
    }
    
    // Modificar e retornar data...
}
```

### Retornando 404 para Rotas Não Encontradas

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // O handler "/" captura TODAS as rotas não definidas
        if r.URL.Path != "/" {
            w.WriteHeader(http.StatusNotFound)
            w.Write([]byte("Not Found"))
            return
        }
        w.Write([]byte("Home"))
    })
}
```

---

## Recursos

- **Vídeo 2**: O pacote net/http (estará disponível no Discord, canal #videos-semanais, assim como todos os demais vídeos do curso)
- [Go by Example: HTTP Servers](https://gobyexample.com/http-servers)
- [Go by Example: JSON](https://gobyexample.com/json)
- [Documentação net/http](https://pkg.go.dev/net/http)
- [Documentação httptest](https://pkg.go.dev/net/http/httptest)

---

## Dúvidas?

- **Discord**: Canal #duvidas-u1
- **Aula**: Terças e Quintas, 13:00-14:40