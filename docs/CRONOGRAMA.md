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
>
> **⚠️ Atualização 25/05/2026**: Aulas de 19/05 e 21/05 não foram realizadas. **Sprint 3 inicia amanhã (26/05)** e tem prazo de entrega prorrogado para **16/06 (ter)**. **Listas 5 e 6 unificadas em uma única Lista 5+6** (auth + segurança + testes). **A Sprint 4 deixa de existir como entrega independente** — é mesclada com a apresentação final, que passa a ser entregue como **vídeo até 30/06 (ter)**.

---

## Legenda

| Símbolo | Significado |
|---------|-------------|
| 📺 | Vídeo disponível (assistir antes da aula) |
| 🟢 | Aula presencial (atividades práticas) |
| 🔵 | Acompanhamento online (Google Meet) |
| 📝 | Entrega de lista de exercícios |
| 🚀 | Entrega de sprint/projeto |
| 🎥 | Entrega de vídeo final (substitui apresentação presencial) |
| 📚 | Prova |
| 🔴 | Feriado/Sem aula |

---

## Visão Geral do Curso por Sprints

| Sprint | Período | Objetivo | Entrega |
|--------|---------|----------|---------|
| **Sprint 0** | 17/03 - 06/04 | Propor o projeto e configurar ambiente Go | 06/04 (dom) ✓ |
| **Sprint 1** | 14/04 - 24/04 | Implementar API REST idiomática com Chi e middleware | 24/04 (sex) ✓ |
| **Sprint 2** | 05/05 - 19/05 | Adicionar persistência com PostgreSQL e relacionamentos 1:N | 19/05 (ter) ✓ |
| **Sprint 3** | 26/05 - 16/06 | Implementar autenticação JWT, segurança OWASP e testes automatizados | **16/06 (ter)** |
| 📚 **PROVA** | 02/06 (ter) | Avaliação prática individual dos conceitos de U1+U2 | — |
| **Conteúdo Final** | 16/06 - 23/06 | gRPC + GraphQL + Concorrência (sem entrega de sprint) | — |
| 🎥 **Vídeo Final** | 16/06 - 30/06 | Defesa final do projeto em vídeo (mescla Sprint 4 + apresentação) | **30/06 (ter)** |
| **Encerramento** | 01-11/07 | Correção, notas, feedback | — |

---

## Mapeamento Unidade ↔ Sprint

| Unidade | Sprints | Conteúdo Principal |
|---------|---------|-------------------|
| **U1** | Sprint 0 + Sprint 1 | Fundamentos Go + net/http + Chi + middleware |
| **U2** | Sprint 2 + Prova | Persistência + Relacionamentos  |
| **U3** | Sprint 3 + Entrega Final do Projeto | Auth + Testes + **gRPC** + GraphQL |

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
- 📝 **Lista 2** (net/http + middleware) — prazo **24/04 (sex)** ✓
- 🚀 **Sprint 0** — Proposta do projeto (vídeo 5min + documento PDF) — prazo estendido até **06/04 (dom)** ✓

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
- 📝 **Lista 3** (Chi + OpenAPI + validator) — 12/05 (ter) ✓
- 🚀 **Sprint 1** (vídeo 8min + API funcionando) — até **30/04** ✓

---

## SPRINT 2 (05/05 - 19/05): Persistência com PostgreSQL e Relacionamentos ✓

### Objetivo do Sprint

A equipe conecta a API a um banco PostgreSQL via **sqlc** (queries type-safe) e modela ao menos um **relacionamento 1:N** entre entidades.

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

### Entregas do Sprint

- 📝 **Lista 3** (Chi + OpenAPI + validator) — 12/05 (ter) ✓
- 📝 **Lista 4** (sqlc + Repository pattern + filtros + JOIN 1:N) — 28/05 (qui)
- 🚀 **Sprint 2** (vídeo 8min + persistência funcionando) — 22/05 (sex) ✓

---

## SPRINT 3 (26/05 - 16/06): Autenticação, Segurança e Testes

### Objetivo do Sprint

A equipe implementa autenticação JWT completa, refresh tokens e middleware de autorização, e introduz testes automatizados nos níveis unitário e de integração. Ao final, a API deve ter login, proteção de rotas, ser resiliente a ataques comuns e ter pipeline CI verde.

> ⚠️ Sprint reorganizada após 25/05: aulas de 19/05 e 21/05 não foram realizadas. O conteúdo da Sprint 3 (JWT + refresh + OWASP + testes) é coberto em 26/05 e 09/06, com fechamento no acompanhamento online de 11/06. **Listas 5 e 6 unificadas** em uma única Lista 5+6 (auth + segurança + testes).

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 26/05 | Ter | 🟢 | **Autenticação**: JWT (estrutura, claims, assinatura HS256, middleware) + Refresh tokens |
| 28/05 | Qui | 🔵 | **Acompanhamento Online** — Sprint 3 + Revisão para a prova de 02/06 |
| 02/06 | Ter | 📚 | **PROVA PRÁTICA** — Unidades 1 e 2 (laboratório, 13:00-14:40) |
| 04/06 | Qui | 🔴 | **Corpus Christi** — Sem aula |
| 09/06 | Ter | 🟢 | OAuth 2.0 flows (visão geral) + Testes em Go (testing, testify, mocks, table-driven) + Segurança API (OWASP API Top 10 + rate limiting) |
| 11/06 | Qui | 🔵 | **Acompanhamento Online** — Sprint 3 (fechamento, dúvidas finais antes da entrega) |

### Vídeos da Sprint

- 📺 **Vídeo 7** (publicado até 27/05): Autenticação JWT + Refresh tokens
- 📺 **Vídeo 8** (publicado até 08/06): OAuth 2.0 + Testes em Go + Segurança API (OWASP Top 10 + rate limiting)

### Entregas do Sprint

- 📝 **Lista 5+6** unificada (auth JWT + autorização + refresh tokens + rate limiting + testes) — até **16/06 (ter) às 23:59**
- 🚀 **Sprint 3** (vídeo 8min + auth + segurança + testes) — até **16/06 (ter)**

**Esperado no vídeo da Sprint 3**:
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

**Conteúdo da prova**: Todo material das Unidades 1 e 2 — fundamentos Go, net/http, Chi, middleware, JSON, validator, sqlc, padrão Repository, autenticação JWT.

**Formato**: Individual, em laboratório. Implementação de uma mini-API com requisitos específicos e testes automáticos de validação.

> Nota: alguns tópicos da Sprint 3 (refresh tokens e testes automatizados) ainda estarão sendo cobertos em 09/06; a prova foca nos fundamentos das duas unidades e na **autenticação JWT**.

---

## Unidade 3

Com a API segura e testada após a Sprint 3, as últimas aulas presenciais ampliam o repertório técnico para protocolos além de REST e concorrência em Go.

> ⚠️ Mudança 25/05: a Sprint 4 deixou de existir como entrega independente. Foi mesclada com a apresentação final, agora entregue como vídeo.

### Aulas

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 16/06 | Ter | 🟢 | **gRPC com Go**: Protocol Buffers, RPC unário e streaming, geração de stubs, comparação com REST |
| 18/06 | Qui | 🟢 | **GraphQL vs REST com gqlgen** + **Concorrência** (goroutines, channels, worker pools, padrões)  |

### Vídeos da Sprint Final

- 📺 **Vídeo 9** (publicado até 14/06): **gRPC com Go** — Protocol Buffers e streaming
- 📺 **Vídeo 10** (publicado até 21/06): GraphQL com gqlgen + Concorrência

### Conteúdo Habilitador

- **gRPC**: definição de serviços em `.proto`, geração de stubs, unary vs server-streaming vs client-streaming vs bidi, trade-offs frente a REST
- **GraphQL**: schema-first com gqlgen, queries vs mutations, N+1 com dataloader, quando faz sentido
- **Concorrência em Go**: goroutines, channels, `select`, padrões (fan-out/fan-in, worker pool, pipeline), `context` para cancelamento

### Entrega Final

🎥 **Vídeo Final** (12 min) — até **30/06 (ter) às 23:59**

**Substitui**: a antiga Sprint 4 (entrega técnica) **e** a antiga apresentação presencial.

**Esperado no vídeo** (12 minutos):
1. **Introdução** (1-2 min) — equipe, projeto, problema que resolve
2. **Demonstração do MVP** (5-6 min) — funcionalidades funcionando, fluxos completos, tratamento de erros
3. **Arquitetura e Código** (3-4 min) — diagrama de arquitetura, decisões técnicas, métricas (cobertura de testes, endpoints documentados), e como tarefa opcional que gera um bonus na nota, uma  **demonstração de um endpoint gRPC ou GraphQL**.
4. **Conclusão** (1 min) — aprendizados principais, o que faria diferente

---

## ENCERRAMENTO (01-11/07)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 30/06 | Ter | 🎥 | **Entrega do vídeo final** (até 23:59) |
| 02/07 | Qui | 🔵 | **Acompanhamento Online** — Dúvidas finais |
| 07/07 | Ter | 🔴 | Sem aula — professor corrigindo trabalhos |
| 09/07 | Qui | 🔵 | **Acompanhamento Online** — Dúvidas sobre correção e notas |

**Consolidação**: Notas no SIGAA até **11/07**

---

## Resumo de Entregas

### Listas de Exercícios (GitHub Classroom)

| Lista | Conteúdo | Prazo | Sprint |
|-------|----------|-------|--------|
| Lista 1 | net/http básico | 27/03 (sex) ✓ | Sprint 0 |
| Lista 2 | net/http + middleware | 24/04 (sex) ✓ | Sprint 0→1 |
| Lista 3 | Chi + OpenAPI + validator | 12/05 (ter) ✓ | Sprint 1→2 |
| Lista 4 | sqlc + Repository + filtros + JOIN 1:N | 19/05 (ter) ✓ | Sprint 2 |
| **Lista 5+6** (unificada) | Auth JWT + autorização + refresh + rate limit + testes | **16/06 (qua)** | Sprint 3 |

> ~~Lista 7 (Deploy + Observabilidade + gRPC)~~ — **removida** após reorganização de 25/05

### Sprints do Projeto

| Sprint | Objetivo | Prazo | Duração Vídeo |
|--------|----------|-------|---------------|
| Sprint 0 | Proposta + ambiente | 06/04 ✓ | 5 min |
| Sprint 1 | API RESTful com Chi | 30/04 ✓ | 8 min |
| Sprint 2 | Persistência com PostgreSQL e relacionamentos | 22/05 ✓ | 8 min |
| Sprint 3 | Auth JWT + Segurança + Testes | **16/06 (ter)** | 8 min |
| 🎥 Vídeo Final | MVP completo + gRPC/GraphQL (substitui Sprint 4 + apresentação) | **30/06 (ter)** | 12 min |

### Entregas por Unidade

| Unidade | Sprints | Prazo Final | Componentes |
|---------|---------|-------------|-------------|
| **U1** | Sprint 0 + Sprint 1 | 24/04 (sex) ✓ | Listas 1-3 + Sprint 0 + Sprint 1 |
| **U2** | Sprint 2 + Sprint 3 + Prova | 02/06 | Listas 4 e 5+6 + Sprints 2-3 + Prova |
| **U3** | Vídeo Final | **30/06 (ter)** | Vídeo Final (mescla Sprint 4 + apresentação) |

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
| **02/06 (ter)** | **PROVA PRÁTICA** |
| **16/06 (ter)** | **Entrega Sprint 3 e Lista 5+6** |
| **30/06 (ter)** | **Entrega do Vídeo Final** (substitui Sprint 4 + apresentação presencial) |
| 11/07 | Término do período (notas no SIGAA) |

---

## Acompanhamentos Online (Google Meet)

Os acompanhamentos acontecem às quintas-feiras, sincronizados com a turma de Processos de Software:

| # | Data | Foco | Sprint |
|---|------|------|--------|
| 1 | 26/03 | Sprint 0 — Proposta do projeto ✓ | Sprint 0 |
| 2 | 14/05 | Sprint 2 — Persistência + relacionamentos ✓ | Sprint 2 |
| 3 | 28/05 | Sprint 3 + Revisão para prova | Sprint 3 |
| 4 | 16/06 | Sprint 3 — fechamento  | Sprint 3 |
| 5 | 25/06 | Dúvidas sobre o vídeo final | Final |
| 6 | 09/07 | Dúvidas sobre correção e notas | — |

**Horário**: 13:00-14:40 (mesmo horário das aulas presenciais)  
**Formato**: ~10 minutos por equipe para tirar dúvidas e revisar progresso
