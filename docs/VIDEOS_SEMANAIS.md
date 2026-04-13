# Vídeos Semanais - Sala de Aula Invertida

## DIM0547 - Desenvolvimento de Sistemas Web II

### Sobre a Metodologia

Este curso utiliza a metodologia de **Sala de Aula Invertida (Flipped Classroom)**. Os vídeos são disponibilizados **antes das aulas presenciais** para que você assista e venha à sala preparado para praticar e tirar dúvidas.

> **⚠️ Atualização 09/04/2026**: Os vídeos agora são organizados por **sprint do projeto** (não mais por semana). Cada vídeo habilita uma capacidade técnica que a equipe aplicará no sprint correspondente.

### Pré-requisito: Curso Básico de Go

**Antes de começar os vídeos de Web II**, complete o curso básico:

📚 **[Aprenda Go com Testes](https://github.com/classrooms-fmarquesfilho/aprenda-go-com-testes)**

### Como Aproveitar os Vídeos

1. **Assista com atenção** - Pause e volte quando necessário
2. **Faça anotações** - Especialmente de dúvidas
3. **Teste os exemplos** - Reproduza o código no seu ambiente
4. **Traga dúvidas** - A aula presencial é para esclarecer e praticar

---

## Vídeos por Sprint

### SPRINT 0 (17/03 - 06/04): Proposta e Ambiente Go ✓

**Objetivo da sprint**: Cada equipe propõe seu projeto e configura o ambiente Go.

| Disponível em | Vídeo | Tópicos |
|---------------|-------|---------|
| 15/03 (Dom) ✓ | **Vídeo 1**: Apresentação do curso + Por que Go? | Motivação, ecossistema, cases no Brasil, arquitetura cliente-servidor |
| 15/03 (Dom) ✓ | **Vídeo 2a**: net/http — Handlers, httptest e Query Params | http.Handler, HandlerFunc, httptest, query params, TDD |
| 22/03 (Dom) ✓ | **Vídeo 2b**: net/http — Maps, ServeMux e JSON | maps, for range, interface{}, ServeMux, encoding/json |
| 22/03 (Dom) ✓ | **Vídeo 2c**: GitHub Classroom — Submissão e autograding | init(), DefaultServeMux, classroom.yml, aba Actions |

---

### SPRINT 1 (14/04 - 24/04): API RESTful com Chi

**Objetivo da sprint**: Implementar API REST idiomática com Chi, middleware por grupo e documentação OpenAPI.

| Disponível em | Vídeo | Tópicos |
|---------------|-------|---------|
| 12/04 (Dom) | **Vídeo 3**: Chi Router + JSON + validator | Chi filosofia, grupos de rotas, `chi.URLParam`, encoding/json, go-playground/validator, struct tags |
| 19/04 (Dom) | **Vídeo 4**: Chi avançado + OpenAPI com swaggo | Subrouters, `r.Mount`, middleware por grupo, anotações swaggo, Swagger UI, ponteiros com `new()` (Go 1.26) |

---

### SPRINT 2 (25/04 - 08/05): Persistência + Clean Arch + Testes

**Objetivo da sprint**: Conectar a API a PostgreSQL com sqlc, adotar Clean Architecture e escrever testes automatizados.

| Disponível em | Vídeo | Tópicos |
|---------------|-------|---------|
| 26/04 (Dom) | **Vídeo 5**: PostgreSQL + sqlc + Migrations Atlas | SQL vs ORM, sqlc, queries type-safe, Atlas migrations, transações |
| 03/05 (Dom) | **Vídeo 6**: Clean Architecture + Testes em Go | Camadas, inversão de dependência, testing package, testify, mocks, `dockertest`, `errors.AsType` (Go 1.26), `testing/synctest` (Go 1.25+) |

---

### SPRINT 3 (09/05 - 22/05): Autenticação e Segurança

**Objetivo da sprint**: Implementar autenticação JWT, refresh tokens, OAuth 2.0 e correções OWASP.

| Disponível em | Vídeo | Tópicos |
|---------------|-------|---------|
| 10/05 (Dom) | **Vídeo 7**: Autenticação JWT + OAuth 2.0 | JWT estrutura, claims, refresh tokens, OAuth 2.0 flows, middleware de auth |
| 17/05 (Dom) | **Vídeo 8**: Segurança API — OWASP Top 10 | Vulnerabilidades comuns, rate limiting, input validation, CORS, broken authentication |

---

### PROVA — 26/05 (ter)

Semana de prova. Sem vídeo novo — foco em revisão do conteúdo U1+U2.

---

### SPRINT 4 (27/05 - 12/06): Deploy, Observabilidade e Protocolos Avançados

**Objetivo da sprint**: Containerizar, observar e explorar protocolos alternativos (gRPC e GraphQL).

| Disponível em | Vídeo | Tópicos |
|---------------|-------|---------|
| 24/05 (Dom) | **Vídeo 9**: Docker multi-stage + CI/CD + Deploy | Multi-stage builds, docker-compose, GitHub Actions, GOMAXPROCS container-aware (Go 1.25+), graceful shutdown, health checks |
| 31/05 (Dom) | **Vídeo 10**: Observabilidade — slog + OpenTelemetry | `log/slog` (GroupAttrs Go 1.25), métricas, traces, OpenTelemetry, instrumentação manual |
| 07/06 (Dom) | **🆕 Vídeo 11**: gRPC com Go — Protocol Buffers e streaming | `.proto` files, `protoc-gen-go`, unary calls, server/client/bidirectional streaming, gRPC vs REST |
| 07/06 (Dom) | **Vídeo 12**: GraphQL com gqlgen + Concorrência | SDL, gqlgen, queries, mutations, goroutines, channels, `sync.WaitGroup.Go()` (Go 1.25), errgroup, worker pools |

---

### APRESENTAÇÕES FINAIS (16-18/06)

| Disponível em | Vídeo | Tópicos |
|---------------|-------|---------|
| 14/06 (Dom) | **Vídeo 13**: Próximos passos + Carreira em Go | Carreira, recursos avançados, comunidade |

---

## Resumo: Vídeos por Sprint

| Sprint | Vídeos | Total |
|--------|--------|-------|
| Sprint 0 | 1, 2a, 2b, 2c | 4 vídeos |
| Sprint 1 | 3, 4 | 2 vídeos |
| Sprint 2 | 5, 6 | 2 vídeos |
| Sprint 3 | 7, 8 | 2 vídeos |
| Sprint 4 | 9, 10, **11 (gRPC)**, 12 | 4 vídeos |
| Pós-apresentações | 13 | 1 vídeo |

---

## Onde os Vídeos Ficam Disponíveis

1. **SIGAA** - Anúncio com link de cada vídeo
2. **Discord** - Canal #videos-semanais
3. **YouTube** - Playlist (link no Discord)

---

## FAQ

### E se eu não conseguir assistir antes da aula?

Você pode assistir depois, mas perderá o aproveitamento máximo da aula presencial. As atividades práticas assumem que você viu o conteúdo.

### Os vídeos ficam disponíveis depois do semestre?

Os links permanecem ativos para consulta futura.