# Cronograma Detalhado - DIM0547

## Desenvolvimento de Sistemas Web II com Go

**Período**: 2026.1 (17/03/2026 a 11/07/2026)  
**Horário**: Terças e Quintas, 13:00 às 14:40 (35T12)

---

> **⚠️ Atualização 09/04/2026**: Cronograma reorganizado por **sprint** (não mais por semana). Cada sprint tem **objetivo claro** alinhado ao conteúdo técnico. Sprints duram **2 semanas**. Total de **4 sprints** (Sprint 1 a Sprint 4). **Aula de gRPC** adicionada na Unidade 3.
>
> **⚠️ Atualização 05/05/2026**: Cronograma ajustado para reposição de aulas não realizadas (03/03, 05/03, 10/03, 12/03, 31/03, 28/04 e 30/04). **Sprint 2 inicia hoje (05/05)**. Prova adiada para **02/06 (ter)**. Apresentações finais adiadas para **30/06 (ter) e 02/07 (qui)**. Entregas de sprint passam a ocorrer na **terça-feira** que encerra cada sprint.
>
> **⚠️ Lista 3**: prazo de entrega prorrogado para **12/05 (ter) às 23:59**.

---

## Legenda

| Símbolo | Significado |
|---------|-------------|
| 📺 | Vídeo disponível (assistir antes da aula) |
| 🟢 | Aula presencial (atividades práticas) |
| 🔵 | Acompanhamento online (Google Meet) |
| 📝 | Entrega de lista de exercícios |
| 🚀 | Entrega de sprint/projeto |
| 📚 | Prova |
| 🔴 | Feriado/Sem aula |

---

## Visão Geral do Curso por Sprints

| Sprint | Período | Objetivo | Entrega |
|--------|---------|----------|---------|
| **Sprint 0** | 17/03 - 06/04 | Propor o projeto e configurar ambiente Go | 06/04 (dom) ✓ |
| **Sprint 1** | 14/04 - 24/04 | Implementar API REST idiomática com Chi e middleware | 24/04 (sex) ✓ |
| **Sprint 2** | 05/05 - 19/05 | Adicionar persistência com PostgreSQL e relacionamentos 1:N | **19/05 (ter)** |
| **Sprint 3** | 20/05 - 02/06 | Implementar autenticação JWT, segurança OWASP e testes automatizados | **29/05 (sex)** |
| 📚 **PROVA** | 02/06 (ter) | Avaliação prática individual dos conceitos de U1+U2 | — |
| **Sprint 4** | 03/06 - 19/06 | Containerizar, observar e explorar protocolos avançados (gRPC, GraphQL) | **19/06 (sex)** |
| **Apresentações** | 30/06 e 02/07 | Defesa final dos projetos | — |
| **Encerramento** | 07-11/07 | Correção, notas, feedback | 04/07 (sex) |

---

## Mapeamento Unidade ↔ Sprint

| Unidade | Sprints | Conteúdo Principal |
|---------|---------|-------------------|
| **U1** | Sprint 0 + Sprint 1 | Fundamentos Go + net/http + Chi + middleware |
| **U2** | Sprint 2 + Sprint 3 + Prova | Persistência + Relacionamentos + Auth + Testes |
| **U3** | Sprint 4 + Apresentações | Deploy + Observabilidade + **gRPC** + GraphQL + Concorrência |

---

## SPRINT 0 (17/03 - 06/04): Proposta e Ambiente ✓

### Objetivo do Sprint

Cada equipe propõe seu projeto de API em Go, define o MVP, cria o repositório e configura o ambiente de desenvolvimento. Ao final, todo membro deve conseguir rodar um "hello world" em Go usando net/http.

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 15/03 | Dom | 📺 | **Vídeo 1**: Apresentação do curso + Por que Go? |
| 15/03 | Dom | 📺 | **Vídeo 2a**: net/http — Handlers, httptest e Query Params |
| 17/03 | Ter | 🟢 | Apresentação do curso + Início da prática com net/http |
| 19/03 | Qui | 🟢 | Prática: handlers, httptest, query params |
| 22/03 | Dom | 📺 | **Vídeo 2b**: net/http — Maps, ServeMux e JSON |
| 24/03 | Ter | 🟢 | Prática: ServeMux, maps, encoding/json |
| 26/03 | Qui | 🔵 | **Acompanhamento Online** — Sprint 0 (Proposta do projeto) |
| 31/03 | Ter | 🔴 | ~~Sem aula — professor doente~~ |
| 02/04 | Qui | 🔴 | ~~Quinta-feira Santa~~ |
| 07/04 | Ter | 🟢 | Middleware + Roteamento |
| 09/04 | Qui | 🟢 | Acompanhamento presencial + revisão da aula anterior |

### Conteúdo

- Fundamentos de Go aplicados a web: http.Handler, HandlerFunc, httptest
- net/http: ServeMux, roteamento moderno (Go 1.22+), query params, JSON
- Padrão middleware: função `func(http.Handler) http.Handler`
- GitHub Classroom + autograding

### Entregas do Sprint

- 📝 **Lista 1** (net/http básico) — 27/03 (sex) ✓
- 📝 **Lista 2** (net/http + middleware) — prazo **24/04 (sex)**
- 🚀 **Sprint 0** — Proposta do projeto (vídeo 5min + documento PDF) — prazo estendido até **06/04 (dom)**

---

## SPRINT 1 (14/04 - 24/04): API RESTful Idiomática com Chi ✓

### Objetivo do Sprint

A equipe implementa a primeira versão da API do projeto com rotas via Chi, middleware compartilhado e documentação OpenAPI.

### Vídeos da Sprint

- 📺 **Vídeo 3** (publicado 16/04): Chi Router + JSON
- 📺 **Vídeo 4** (publicado 23/04): Chi Avançado, Validação e Middleware

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 14/04 | Ter | 🟢 | Grupos de rotas, subrouters, montagem (r.Mount) |
| 16/04 | Qui | 🟢 | Middleware por grupo + validação com go-playground/validator + OpenAPI (swaggo) |
| 21/04 | Ter | 🔴 | **Tiradentes** — Sem aula |
| 23/04 | Qui | 🔵 | Acompanhamento de projeto |

### Entregas do Sprint

- 📝 **Lista 2** (net/http + middleware) — até 24/04 (sex) ✓
- 📝 **Lista 3** (Chi + OpenAPI + validator) — ⚠️ **prazo prorrogado: 12/05 (ter) às 23:59**
- 🚀 **Sprint 1** (vídeo 8min + API funcionando) — até **24/04 (sex)** ✓

**Esperado no vídeo**:
1. Demonstração da API com pelo menos 4-5 endpoints
2. Chi com grupos de rotas e middlewares específicos
3. Documentação OpenAPI gerada com swaggo
4. Validação de entrada com validator

---

## SPRINT 2 (05/05 - 19/05): Persistência com PostgreSQL e Relacionamentos

### Objetivo do Sprint

A equipe conecta a API a um banco PostgreSQL via **sqlc** (queries type-safe) e modela ao menos um **relacionamento 1:N** entre entidades. Ao final, a aplicação persiste dados de verdade e expõe pelo menos um endpoint que entrega dados aninhados (pai + filhos).

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 05/05 | Ter | 🔴 | Não houve aula |
| 07/05 | Qui | 🟢 | PostgreSQL + sqlc: queries type-safe |
| 12/05 | Ter | 🟢 | JOINs e relacionamentos 1:N: LEFT JOIN, agregação no Go |
| 14/05 | Qui | 🔵 | **Acompanhamento Online** — Sprint 2 check-in |

### Vídeos da Sprint

- 📺 **Vídeo 5**: PostgreSQL + sqlc — queries type-safe e CRUD persistente
- 📺 **Vídeo 6**: JOINs e relacionamentos 1:N em sqlc

> Clean Architecture, migrations com Atlas e testes automatizados foram movidos para sprints posteriores — não são requisitos desta entrega.

### Entregas do Sprint

- 📝 **Lista 3** (Chi + OpenAPI + validator) — ⚠️ **prazo prorrogado: 12/05 (ter) às 23:59**
- 📝 **Lista 4** (sqlc + Repository pattern + filtros + JOIN 1:N) — até **19/05 (ter)**
- 🚀 **Sprint 2** (vídeo 8min + persistência funcionando) — até **19/05 (ter)**

**Esperado no vídeo**:
1. API conectada a PostgreSQL via sqlc
2. CRUD completo em pelo menos uma entidade persistido no banco
3. Pelo menos um endpoint que envolve relacionamento (JOIN, lista paginada, ou busca filtrada)
4. Schema SQL versionado no repositório (em `db/schema/`)

---

## SPRINT 3 (20/05 - 02/06): Autenticação, Segurança e Testes

### Objetivo do Sprint

A equipe implementa autenticação JWT completa, refresh tokens e middleware de autorização, e introduz testes automatizados nos níveis unitário e de integração. Ao final, a API deve ter login, proteção de rotas, ser resiliente a ataques comuns e ter pipeline CI verde.

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 19/05 | Ter | 🟢 | Autenticação com JWT: estrutura, claims, assinatura, middleware de auth |
| 21/05 | Qui | 🟢 | Refresh tokens + OAuth 2.0 flows (authorization code, client credentials) |
| 26/05 | Ter | 🟢 | Testes em Go + Segurança: testing package, testify, mocks, OWASP API Top 10 |
| 28/05 | Qui | 🔵 | **Acompanhamento Online** — Sprint 3 + Revisão para a prova |

### Vídeos da Sprint

- 📺 **Vídeo 7** (publicado até 17/05): Autenticação JWT + OAuth 2.0 em Go
- 📺 **Vídeo 8** (publicado até 24/05): Testes em Go + Segurança API (OWASP Top 10 + rate limiting)

### Entregas do Sprint

- 📝 **Lista 5** (Testes unitários + integração) — até **26/05 (ter)**
- 📝 **Lista 6** (Autenticação JWT + OAuth) — publicada 20/05, até **02/06 (ter)**
- 🚀 **Sprint 3** (vídeo 8min + auth + hardening + testes) — até **29/05 (sex)**

**Esperado no vídeo**:
1. Endpoint de login retornando JWT
2. Rotas protegidas por middleware de autenticação
3. Refresh token funcionando
4. Pelo menos 2 correções de segurança OWASP implementadas
5. Pelo menos 10 testes automatizados rodando no pipeline CI

---

## PROVA PRÁTICA — 02/06 (Terça)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 02/06 | Ter | 📚 | **PROVA PRÁTICA** — Unidades 1 e 2 (laboratório, 13:00-14:40) |
| 04/06 | Qui | 🔴 | **Corpus Christi** — Sem aula |

**Conteúdo da prova**: Todo material das Unidades 1 e 2 — fundamentos Go, net/http, Chi, middleware, JSON, validator, sqlc, Clean Architecture, testes, autenticação JWT.

**Formato**: Individual, em laboratório. Implementação de uma mini-API com requisitos específicos e testes automáticos de validação.

---

## SPRINT 4 (03/06 - 19/06): Deploy, Observabilidade e Protocolos Avançados

### Objetivo do Sprint

Último sprint. A equipe usa containers e faz deploy do projeto, implementa observabilidade (logs estruturados, métricas, traces) e explora protocolos além de REST: **gRPC** e **GraphQL**. Ao final, o projeto deve estar em produção (mesmo que simulado) com observabilidade funcionando.

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 09/06 | Ter | 🟢 | Docker multi-stage + CI/CD + deploy |
| 11/06 | Qui | 🔵 | **Acompanhamento Online** — Sprint 4 check-in |
| 16/06 | Ter | 🟢 | Observabilidade: slog, OpenTelemetry, métricas e traces |
| 18/06 | Qui | 🟢 | **gRPC com Go**: Protocol Buffers, streaming, quando usar vs REST |
| 23/06 | Ter | 🟢 | GraphQL vs REST com gqlgen + Concorrência (goroutines, channels, worker pools) |

### Vídeos da Sprint

- 📺 **Vídeo 9** (publicado até 07/06): Docker multi-stage + CI/CD + deploy
- 📺 **Vídeo 10** (publicado até 14/06): Observabilidade — slog + OpenTelemetry
- 📺 **Vídeo 11** (publicado até 14/06): **gRPC com Go** — Protocol Buffers e streaming
- 📺 **Vídeo 12** (publicado até 21/06): GraphQL com gqlgen + Concorrência

### Entregas do Sprint

- 📝 **Lista 7** (Deploy + Observabilidade + gRPC) — publicada 03/06, até **19/06 (sex)**
- 🚀 **Sprint 4** (vídeo 10min + MVP completo + deploy + observabilidade) — até **19/06 (sex)**

**Esperado no vídeo** (sprint final):
1. MVP completo funcionando em container
2. Pipeline CI/CD com deploy automatizado
3. Logs estruturados e pelo menos uma métrica ou trace coletado
4. Um endpoint gRPC ou GraphQL implementado como demonstração de protocolo alternativo
5. Reflexão sobre a arquitetura final

---

## APRESENTAÇÕES FINAIS (30/06 e 02/07)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 28/06 | Dom | 📺 | **Vídeo 13**: Próximos passos + Carreira em Go |
| 30/06 | Ter | 🟢 | **Apresentações finais** (Grupos 1-6) |
| 02/07 | Qui | 🟢 | **Apresentações finais** (Grupos 7-12) |

**Formato**: 12 minutos por equipe + 3 minutos de perguntas (15 min total)  
**Foco**: demonstração do MVP, decisões técnicas, aprendizados

---

## ENCERRAMENTO (07-11/07)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 07/07 | Ter | 🔴 | Sem aula — professor corrigindo trabalhos |
| 09/07 | Qui | 🔵 | **Acompanhamento Online** — Dúvidas sobre correção e notas |

**Prazo final**: Entrega U3 até **04/07 (sex) às 23:59**  
**Consolidação**: Notas no SIGAA até 11/07

---

## Resumo de Entregas

### Listas de Exercícios (GitHub Classroom)

| Lista | Conteúdo | Prazo | Sprint |
|-------|----------|-------|--------|
| Lista 1 | net/http básico | 27/03 (sex) ✓ | Sprint 0 |
| Lista 2 | net/http + middleware | 24/04 (sex) ✓ | Sprint 0→1 |
| Lista 3 | Chi + OpenAPI + validator | ⚠️ **12/05 (ter) às 23:59** | Sprint 1→2 |
| Lista 4 | sqlc + Repository + filtros + JOIN 1:N | **19/05 (ter)** | Sprint 2 |
| Lista 5 | Testes unitários + integração | **26/05 (ter)** | Sprint 3 |
| Lista 6 | Autenticação JWT + OAuth | **02/06 (ter)** | Sprint 3 |
| Lista 7 | Deploy + Observabilidade + gRPC | **19/06 (sex)** | Sprint 4 |

### Sprints do Projeto

| Sprint | Objetivo | Prazo | Duração Vídeo |
|--------|----------|-------|---------------|
| Sprint 0 | Proposta + ambiente | 06/04 (dom) ✓ | 5 min |
| Sprint 1 | API RESTful com Chi | 24/04 (sex) ✓ | 8 min |
| Sprint 2 | Persistência com PostgreSQL e relacionamentos | **19/05 (ter)** | 8 min |
| Sprint 3 | Auth JWT + Segurança + Testes | **29/05 (sex)** | 8 min |
| Sprint 4 | Deploy + Observabilidade + gRPC/GraphQL | **19/06 (sex)** | 10 min |

### Entregas por Unidade

| Unidade | Sprints | Prazo Final | Componentes |
|---------|---------|-------------|-------------|
| **U1** | Sprint 0 + Sprint 1 | 24/04 (sex) ✓ | Listas 1-3 + Sprint 0 + Sprint 1 |
| **U2** | Sprint 2 + Sprint 3 + Prova | **29/05 (sex)** + 02/06 | Listas 4-6 + Sprints 2-3 + Prova |
| **U3** | Sprint 4 + Apresentação | **04/07 (sex)** | Lista 7 + Sprint 4 + Apresentação Final |

---

## Feriados e Datas Especiais

| Data | Evento | Impacto |
|------|--------|---------|
| 02/04 | Quinta-feira Santa | Sem aula (Sprint 0) |
| 21/04 | Tiradentes | Sem aula (Sprint 1) |
| 01/05 | Dia do Trabalho | Sexta — sem aula regular |
| 04/06 | Corpus Christi | Sem aula — dia após a prova |

---

## Datas Importantes

| Data | Evento |
|------|--------|
| 19/03 | Prazo para formação de grupos |
| 06/04 (dom) | Entrega Sprint 0 ✓ |
| 24/04 (sex) | Entrega Sprint 1 (fim da U1) ✓ |
| ⚠️ **12/05 (ter)** | **Lista 3 — prazo prorrogado (23:59)** |
| **19/05 (ter)** | Entrega Sprint 2 |
| **29/05 (sex)** | Entrega Sprint 3 |
| **02/06 (ter)** | **PROVA PRÁTICA** |
| **19/06 (sex)** | Entrega Sprint 4 (MVP completo) |
| 30/06 e 02/07 | Apresentações finais |
| 04/07 (sex) | Entrega U3 (final) |
| 11/07 | Término do período |

---

## Acompanhamentos Online (Google Meet)

Os acompanhamentos acontecem às quintas-feiras, sincronizados com a turma de Processos de Software:

| # | Data | Foco | Sprint |
|---|------|------|--------|
| 1 | 26/03 | Sprint 0 — Proposta do projeto ✓ | Sprint 0 |
| 2 | 14/05 | Sprint 2 — Persistência + relacionamentos | Sprint 2 |
| 3 | 28/05 | Sprint 3 + Revisão para prova | Sprint 3 |
| 4 | 11/06 | Sprint 4 — Deploy + observabilidade | Sprint 4 |
| 5 | 09/07 | Dúvidas sobre correção e notas | — |

**Horário**: 13:00-14:40 (mesmo horário das aulas presenciais)  
**Formato**: ~10 minutos por equipe para tirar dúvidas e revisar progresso