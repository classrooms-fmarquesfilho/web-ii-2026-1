# Cronograma Detalhado - DIM0547

## Desenvolvimento de Sistemas Web II com Go

**Período**: 2026.1 (17/03/2026 a 11/07/2026)  
**Horário**: Terças e Quintas, 13:00 às 14:40 (35T12)

---

> **⚠️ Atualização 09/04/2026**: Cronograma reorganizado por **sprint** (não mais por semana). Cada sprint tem **objetivo claro** alinhado ao conteúdo técnico. Sprints duram **2 semanas**. Prova adiada para **26/05 (ter)**. Total de **4 sprints** (Sprint 1 a Sprint 4). **Aula de gRPC** adicionada na Unidade 3.

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
| **Sprint 1** | 14/04 - 24/04 | Implementar API REST idiomática com Chi e middleware | **24/04 (sex)** |
| **Sprint 2** | 25/04 - 08/05 | Adicionar persistência com PostgreSQL + Clean Architecture + testes | **08/05 (sex)** |
| **Sprint 3** | 09/05 - 22/05 | Implementar autenticação JWT + segurança OWASP | **22/05 (sex)** |
| 📚 **PROVA** | 26/05 (ter) | Avaliação prática individual dos conceitos de U1+U2 | — |
| **Sprint 4** | 27/05 - 12/06 | Containerizar, observar e explorar protocolos avançados (gRPC, GraphQL) | **12/06 (sex)** |
| **Apresentações** | 16-18/06 | Defesa final dos projetos | — |
| **Encerramento** | 23-25/06 | Correção, notas, feedback | 26/06 (sex) |

---

## Mapeamento Unidade ↔ Sprint

| Unidade | Sprints | Conteúdo Principal |
|---------|---------|-------------------|
| **U1** | Sprint 0 + Sprint 1 | Fundamentos Go + net/http + Chi + middleware |
| **U2** | Sprint 2 + Sprint 3 + Prova | Persistência + Clean Arch + Testes + Auth |
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

## SPRINT 1 (14/04 - 24/04): API RESTful Idiomática com Chi

### Objetivo do Sprint

A equipe implementa a primeira versão da API do projeto com rotas via Chi, middleware compartilhado e documentação OpenAPI.

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 14/04 | Ter | 🟢 | Grupos de rotas, subrouters, montagem (r.Mount) |
| 16/04 | Qui | 🟢 | Middleware por grupo + validação com go-playground/validator + OpenAPI (swaggo) |
| 21/04 | Ter | 🔴 | **Tiradentes** — Sem aula |
| 23/04 | Qui | 🔵 | Acompanhamento de projeto |

### Entregas do Sprint

- 📝 **Lista 2** (net/http + middleware) — até 24/04 (sex)
- 📝 **Lista 3** (Chi + OpenAPI + validator) — publicada 21/04, até 5/5 (ter)
- 🚀 **Sprint 1** (vídeo 8min + API funcionando) — até **24/04 (sex)**

**Esperado no vídeo**:
1. Demonstração da API com pelo menos 4-5 endpoints
2. Chi com grupos de rotas e middlewares específicos
3. Documentação OpenAPI gerada com swaggo
4. Validação de entrada com validator

---

## SPRINT 2 (25/04 - 08/05): Persistência, Arquitetura e Testes

### Objetivo do Sprint

A equipe conecta a API a um banco PostgreSQL e introduz **testes automatizados** nos níveis unitário e de integração.

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 28/04 | Ter | 🟢 | PostgreSQL + sqlc: queries type-safe + migrations com Atlas |
| 30/04 | Qui | 🟢 | Clean Architecture em Go: camadas, inversão de dependência, repositories |
| 05/05 | Ter | 🟢 | Testes em Go: testing package, testify, mocks, tabela de testes |
| 07/05 | Qui | 🔵 | **Acompanhamento Online** — Sprint 2 check-in |

### Vídeos da Sprint

- 📺 **Vídeo 5** (publicar até 26/04): PostgreSQL + sqlc + migrations Atlas
- 📺 **Vídeo 6** (publicar até 03/05): Clean Architecture + testes em Go

### Entregas do Sprint

- 📝 **Lista 4** (sqlc + Clean Architecture) — até 02/05 (sex)
- 📝 **Lista 5** (Testes unitários + integração) — publicada 02/05, até 08/05 (sex)
- 🚀 **Sprint 2** (vídeo 8min + persistência + testes) — até **08/05 (sex)**

**Esperado no vídeo**:
1. API conectada a PostgreSQL com sqlc
2. Código organizado em camadas (Clean Architecture)
3. Pelo menos 10 testes automatizados (unitários + integração) passando
4. Pipeline CI rodando os testes automaticamente

---

## SPRINT 3 (09/05 - 22/05): Autenticação e Segurança

### Objetivo do Sprint

A equipe implementa autenticação JWT completa, refresh tokens e middleware de autorização. Ao final, a API deve ter login, proteção de rotas e ser resiliente a ataques comuns.

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 12/05 | Ter | 🟢 | Autenticação com JWT: estrutura, claims, assinatura, middleware de auth |
| 14/05 | Qui | 🟢 | Refresh tokens + OAuth 2.0 flows (authorization code, client credentials) |
| 19/05 | Ter | 🟢 | Segurança: OWASP API Top 10, rate limiting, input validation, CORS |
| 21/05 | Qui | 🔵 | **Acompanhamento Online** — Sprint 3 + Revisão para a prova |

### Vídeos da Sprint

- 📺 **Vídeo 7** (publicar até 10/05): Autenticação JWT + OAuth 2.0 em Go
- 📺 **Vídeo 8** (publicar até 17/05): Segurança API — OWASP Top 10 + rate limiting

### Entregas do Sprint

- 📝 **Lista 6** (Autenticação JWT + OAuth) — publicada 09/05, até 15/05 (sex)
- 🚀 **Sprint 3** (vídeo 8min + auth funcionando + hardening) — até **22/05 (sex)**

**Esperado no vídeo**:
1. Endpoint de login retornando JWT
2. Rotas protegidas por middleware de autenticação
3. Refresh token funcionando
4. Pelo menos 2 correções de segurança OWASP implementadas

---

## PROVA PRÁTICA — 26/05 (Terça)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 26/05 | Ter | 📚 | **PROVA PRÁTICA** — Unidades 1 e 2 (laboratório, 13:00-14:40) |
| 28/05 | Qui | 🟢 | Abertura da U3 (início do Sprint 4) |

**Conteúdo da prova**: Todo material das Unidades 1 e 2 — fundamentos Go, net/http, Chi, middleware, JSON, validator, sqlc, Clean Architecture, testes, autenticação JWT.

**Formato**: Individual, em laboratório. Implementação de uma mini-API com requisitos específicos e testes automáticos de validação.

---

## SPRINT 4 (27/05 - 12/06): Deploy, Observabilidade e Protocolos Avançados

### Objetivo do Sprint

Último sprint. A equipe usa containers e faz deploy do projeto, implementa observabilidade (logs estruturados, métricas, traces) e explora protocolos além de REST: **gRPC** e **GraphQL**. Ao final, o projeto deve estar em produção (mesmo que simulado) com observabilidade funcionando.

### Aulas do Sprint

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 28/05 | Qui | 🟢 | Devolutiva da prova + Docker multi-stage + CI/CD + deploy |
| 02/06 | Ter | 🟢 | Observabilidade: slog, OpenTelemetry, métricas e traces |
| 04/06 | Qui | 🔴 | **Corpus Christi** — Sem aula |
| 09/06 | Ter | 🟢 | **gRPC com Go**: Protocol Buffers, streaming, quando usar vs REST |
| 11/06 | Qui | 🟢 | GraphQL vs REST com gqlgen + Concorrência (goroutines, channels, worker pools) |

### Vídeos da Sprint

- 📺 **Vídeo 9** (publicar até 24/05): Docker multi-stage + CI/CD + deploy
- 📺 **Vídeo 10** (publicar até 31/05): Observabilidade — slog + OpenTelemetry
- 📺 **Vídeo 11** (publicar até 07/06): **gRPC com Go** — Protocol Buffers e streaming
- 📺 **Vídeo 12** (publicar até 07/06): GraphQL com gqlgen + Concorrência

### Entregas do Sprint

- 📝 **Lista 7** (Deploy + Observabilidade + gRPC) — publicada 30/05, até 12/06 (sex)
- 🚀 **Sprint 4** (vídeo 10min + MVP completo + deploy + observabilidade) — até **12/06 (sex)**

**Esperado no vídeo** (sprint final):
1. MVP completo funcionando em container
2. Pipeline CI/CD com deploy automatizado
3. Logs estruturados e pelo menos uma métrica ou trace coletado
4. Um endpoint gRPC ou GraphQL implementado como demonstração de protocolo alternativo
5. Reflexão sobre a arquitetura final

---

## APRESENTAÇÕES FINAIS (16-18/06)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 14/06 | Dom | 📺 | **Vídeo 13**: Próximos passos + Carreira em Go |
| 16/06 | Ter | 🟢 | **Apresentações finais** (Grupos 1-6) |
| 18/06 | Qui | 🟢 | **Apresentações finais** (Grupos 7-12) |

**Formato**: 12 minutos por equipe + 3 minutos de perguntas (15 min total)  
**Foco**: demonstração do MVP, decisões técnicas, aprendizados

---

## ENCERRAMENTO (23-25/06)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 23/06 | Ter | 🔴 | Sem aula — professor corrigindo trabalhos |
| 25/06 | Qui | 🔵 | **Acompanhamento Online** — Dúvidas sobre correção e notas |

**Prazo final**: Entrega U3 até 26/06 (sex) às 23:59  
**Consolidação**: Notas no SIGAA até 11/07

---

## Resumo de Entregas

### Listas de Exercícios (GitHub Classroom)

| Lista | Conteúdo | Prazo | Sprint |
|-------|----------|-------|--------|
| Lista 1 | net/http básico | 27/03 (sex) ✓ | Sprint 0 |
| Lista 2 | net/http + middleware | **17/04 (sex)** | Sprint 0→1 |
| Lista 3 | Chi + OpenAPI + validator | **24/04 (sex)** | Sprint 1 |
| Lista 4 | sqlc + Clean Architecture | **02/05 (sex)** | Sprint 2 |
| Lista 5 | Testes unitários + integração | **08/05 (sex)** | Sprint 2 |
| Lista 6 | Autenticação JWT + OAuth | **15/05 (sex)** | Sprint 3 |
| Lista 7 | Deploy + Observabilidade + gRPC | **12/06 (sex)** | Sprint 4 |

### Sprints do Projeto

| Sprint | Objetivo | Prazo | Duração Vídeo |
|--------|----------|-------|---------------|
| Sprint 0 | Proposta + ambiente | 06/04 (dom) ✓ | 5 min |
| Sprint 1 | API RESTful com Chi | **24/04 (sex)** | 8 min |
| Sprint 2 | Persistência + Clean Arch + Testes | **08/05 (sex)** | 8 min |
| Sprint 3 | Auth JWT + Segurança | **22/05 (sex)** | 8 min |
| Sprint 4 | Deploy + Observabilidade + gRPC/GraphQL | **12/06 (sex)** | 10 min |

### Entregas por Unidade

| Unidade | Sprints | Prazo Final | Componentes |
|---------|---------|-------------|-------------|
| **U1** | Sprint 0 + Sprint 1 | **24/04 (sex)** | Listas 1-3 + Sprint 0 + Sprint 1 |
| **U2** | Sprint 2 + Sprint 3 + Prova | **22/05 (sex)** + 26/05 | Listas 4-6 + Sprints 2-3 + Prova |
| **U3** | Sprint 4 + Apresentação | **26/06 (sex)** | Lista 7 + Sprint 4 + Apresentação Final |

---

## Feriados e Datas Especiais

| Data | Evento | Impacto |
|------|--------|---------|
| 02/04 | Quinta-feira Santa | Sem aula (Sprint 0) |
| 21/04 | Tiradentes | Sem aula (Sprint 1) |
| 01/05 | Dia do Trabalho | Sexta — sem aula regular |
| 04/06 | Corpus Christi | Sem aula (Sprint 4) |

---

## Datas Importantes

| Data | Evento |
|------|--------|
| 19/03 | Prazo para formação de grupos |
| **06/04 (dom)** | Entrega Sprint 0 |
| **24/04 (sex)** | Entrega Sprint 1 (fim da U1) |
| **08/05 (sex)** | Entrega Sprint 2 |
| **22/05 (sex)** | Entrega Sprint 3 |
| **26/05 (ter)** | **PROVA PRÁTICA** |
| **12/06 (sex)** | Entrega Sprint 4 (MVP completo) |
| 16 e 18/06 | Apresentações finais |
| 26/06 (sex) | Entrega U3 (final) |
| 11/07 | Término do período |

---

## Acompanhamentos Online (Google Meet)

Os acompanhamentos acontecem às quintas-feiras, sincronizados com a turma de Processos de Software:

| # | Data | Foco | Sprint |
|---|------|------|--------|
| 1 | 26/03 | Sprint 0 — Proposta do projeto | Sprint 0 |
| 2 | 07/05 | Sprint 2 — Persistência + testes | Sprint 2 |
| 3 | 21/05 | Sprint 3 + Revisão para prova | Sprint 3 |
| 4 | 25/06 | Dúvidas sobre correção e notas | — |

**Horário**: 13:00-14:40 (mesmo horário das aulas presenciais)  
**Formato**: ~10 minutos por equipe para tirar dúvidas e revisar progresso
