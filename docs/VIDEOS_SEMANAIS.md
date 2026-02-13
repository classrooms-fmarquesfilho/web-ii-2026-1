# Vídeos Semanais - Sala de Aula Invertida

## DIM0547 - Desenvolvimento de Sistemas Web II

### Sobre a Metodologia

Este curso utiliza a metodologia de **Sala de Aula Invertida (Flipped Classroom)**. Cada semana, um vídeo é disponibilizado **até domingo às 20h** para que você assista **antes das aulas presenciais** de terça e quinta-feira.

### Pré-requisito: Curso Básico de Go

**Antes de começar os vídeos de Web II**, complete o curso básico:

📚 **[Aprenda Go com Testes](https://github.com/classrooms-fmarquesfilho/aprenda-go-com-testes)**

**Tópicos essenciais do curso básico para Web II:**

| Capítulo | Tópico | Por que é importante para Web II |
|----------|--------|----------------------------------|
| 1-2 | Hello World + Integers | Sintaxe básica, funções, testes |
| 3-5 | Iteration + Arrays/Slices | Manipulação de coleções (comum em APIs) |
| 6 | Structs, Methods, Interfaces | base para handlers HTTP |
| 7 | Pointers & Errors | tratamento de erros em APIs |
| 8 | Maps | Dados JSON, configurações |
| 9-10 | Dependency Injection + Mocking | arquitetura testável |

**Tempo estimado**: 2-3 semanas de estudo (2-3h/dia)

### Como Aproveitar os Vídeos

1. **Assista com atenção** - Pause e volte quando necessário
2. **Faça anotações** - Especialmente de dúvidas
3. **Teste os exemplos** - Reproduza o código no seu ambiente
4. **Traga dúvidas** - A aula presencial é para esclarecer e praticar

---

## Lista de Vídeos

### UNIDADE 1: Fundamentos (Semanas 1-5)

| Semana | Disponível em | Vídeo | Duração | Tópicos |
|--------|---------------|-------|---------|---------|
| 1 | 01/03 (Dom) | **Vídeo 1**: Apresentação do curso + Por que Go? | ~30 min | Motivação, ecossistema, cases no Brasil, arquitetura cliente-servidor |
| 2 | 08/03 (Dom) | **Vídeo 2**: O pacote net/http - Handler e ServeMux | ~35 min | http.Handler, HandlerFunc, ServeMux, Request/Response |
| 3 | 15/03 (Dom) | **Vídeo 3**: Go 1.22+ roteamento + Middleware | ~40 min | Novidades do Go 1.22, padrão middleware, cadeia de handlers |
| 4 | 22/03 (Dom) | **Vídeo 4**: JSON em Go + Introdução ao Chi | ~35 min | encoding/json, struct tags, validator, Chi router |
| 5 | 29/03 (Dom) | **Vídeo 5**: Chi em profundidade + OpenAPI | ~40 min | Grupos de rotas, middleware por grupo, swaggo |

### UNIDADE 2: Persistência e Arquitetura (Semanas 6-10)

| Semana | Disponível em | Vídeo | Duração | Tópicos |
|--------|---------------|-------|---------|---------|
| 6 | 05/04 (Dom) | **Vídeo 6**: PostgreSQL + sqlc | ~40 min | Por que SQL vs ORM, sqlc, queries type-safe |
| 7 | 12/04 (Dom) | **Vídeo 7**: Migrations + Clean Architecture | ~45 min | Atlas, estrutura de camadas, inversão de dependência |
| 8 | 19/04 (Dom) | **Vídeo 8**: Testes em Go | ~40 min | testing package, testify, mocks, testes de integração |
| 9 | 26/04 (Dom) | **Vídeo 9**: Autenticação JWT + OAuth 2.0 | ~45 min | JWT, refresh tokens, OAuth 2.0, middleware de auth |
| 10 | 03/05 (Dom) | **Vídeo 10**: Segurança - OWASP API Top 10 | ~35 min | Vulnerabilidades comuns, rate limiting, input validation |

### UNIDADE 3: DevOps e Conceitos Avançados (Semanas 11-16)

| Semana | Disponível em | Vídeo | Duração | Tópicos |
|--------|---------------|-------|---------|---------|
| 11 | — | *Semana de prova - sem vídeo novo* | — | Revisão do conteúdo U1-U2 |
| 12 | 17/05 (Dom) | **Vídeo 11**: Docker + CI/CD | ~50 min | Multi-stage builds, docker-compose, GitHub Actions, GOMAXPROCS container-aware, graceful shutdown, health checks |
| 13 | 24/05 (Dom) | **Vídeo 12**: Observabilidade | ~40 min | slog (GroupAttrs Go 1.25), métricas, traces, OpenTelemetry |
| 14 | 31/05 (Dom) | **Vídeo 13**: GraphQL vs REST | ~40 min | Quando usar GraphQL, gqlgen, queries e mutations |
| 15 | 07/06 (Dom) | **Vídeo 14**: Concorrência em Go | ~50 min | Goroutines, channels, sync.WaitGroup.Go(), testing/synctest, errgroup, worker pools |
| 16 | 14/06 (Dom) | **Vídeo 15**: Próximos passos | ~30 min | Carreira em Go, recursos avançados, comunidade |

---

## Onde os Vídeos Ficam Disponíveis

1. **SIGAA** - Anúncio com link toda semana
2. **Discord** - Canal #videos-semanais
3. **YouTube** - Playlist (link no Discord)

---

## FAQ

### E se eu não conseguir assistir antes da aula?

Você pode assistir depois, mas perderá o aproveitamento máximo da aula presencial. As atividades práticas assumem que você viu o conteúdo.

### Os vídeos ficam disponíveis depois do semestre?

Os links permanecem ativos para consulta futura.