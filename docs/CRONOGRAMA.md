# Cronograma Detalhado - DIM0547

## Desenvolvimento de Sistemas Web II com Go

**Período**: 2026.1 (02/03/2026 a 11/07/2026)  
**Horário**: Terças e Quintas, 13:00 às 14:40 (35T12)

---

## Legenda

| Símbolo | Significado |
|---------|-------------|
| 📺 | Vídeo disponível (assistir antes da aula) |
| 🟢 | Aula presencial (atividades práticas) |
| 🔵 | Acompanhamento online (Google Meet) |
| 📝 | Entrega de exercício |
| 🚀 | Entrega de sprint/projeto |
| 📚 | Prova |
| 🔴 | Feriado/Sem aula |

---

## Calendário Resumido

| Unidade | Semanas | Período | Conteúdo Principal |
|---------|---------|---------|-------------------|
| **U1** | 1-5 | 02/03 - 03/04 | Fundamentos Go + net/http + Chi |
| **U2** | 6-11 | 06/04 - 14/05 | Persistência + Auth + **Prova** |
| **U3** | 12-17 | 18/05 - 26/06 | CI/CD + GraphQL + Apresentações |

---

## UNIDADE 1: Fundamentos (Semanas 1-5)

### Semana 1 (02-05/03): Introdução ao Go e Arquitetura Web

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 01/03 | Dom | 📺 | **Vídeo 1**: Apresentação do curso + Por que Go? |
| 03/03 | Ter | 🟢 | Discussão do vídeo + Arquitetura cliente-servidor |
| 05/03 | Qui | 📺 | Sem aula presencial — assista aos vídeos do **[Curso Básico de Go](https://github.com/classrooms-fmarquesfilho/aprenda-go-com-testes)** se ainda não tiver familiaridade com a linguagem |

**Conteúdo**: Motivação, ecossistema Go, arquitetura web, setup  
**Entregas**: Nenhuma (semana de ambientação)  
**Importante**: Formação dos grupos de projeto até 12/03. Quinta 05/03 sem aula presencial — aproveite para avançar no curso básico de Go.

---

### Semana 2 (09-12/03): HTTP com a Biblioteca Padrão (Parte 1)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 08/03 | Dom | 📺 | **Vídeo 2**: O pacote net/http - Handler e ServeMux |
| 10/03 | Ter | 🟢 | Revisão + Prática: criando handlers |
| 12/03 | Qui | 🔵 | **Acompanhamento Online** - Dúvidas de ambiente + formação de grupos |

**Conteúdo**: http.Handler, HandlerFunc, ServeMux, Request/Response  
**Entregas**: 
- 📝 Lista 1 (net/http básico) - até 13/03 (sex)
- Grupos formados no Discord até 12/03

---

### Semana 3 (16-19/03): HTTP com a Biblioteca Padrão (Parte 2)

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 15/03 | Dom | 📺 | **Vídeo 3**: Roteamento com a stdlib moderna + Middleware |
| 17/03 | Ter | 🟢 | Revisão + Prática: implementando middlewares |
| 19/03 | Qui | 🟢 | Context + Padrões de tratamento de erros |

**Conteúdo**: Roteamento com método+path no ServeMux, wildcards, middleware pattern, context  
**Entregas**: 📝 Lista 2 (net/http + middleware) - até 20/03 (sex)

---

### Semana 4 (23-26/03): Validação e Introdução ao Chi

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 22/03 | Dom | 📺 | **Vídeo 4**: JSON em Go + Introdução ao Chi |
| 24/03 | Ter | 🟢 | Revisão + Prática: validação com go-playground/validator |
| 26/03 | Qui | 🔵 | **Acompanhamento Online** - Sprint 0 (Proposta do projeto) |

**Conteúdo**: encoding/json, struct tags, validator, Chi router  
**Entregas**: 🚀 Sprint 0 - Proposta do projeto (vídeo 5min) - até 26/03

---

### Semana 5 (30/03 - 02/04): Chi Avançado e OpenAPI

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 29/03 | Dom | 📺 | **Vídeo 5**: Chi em profundidade + OpenAPI com swaggo |
| 31/03 | Ter | 🟢 | Revisão + Prática: middleware por grupo e documentação |
| 02/04 | Qui | 🔴 | **Quinta-feira Santa** - Sem aula |

**Conteúdo**: Grupos de rotas, subrouters, swaggo/swag  
**Entregas**: 
- 📝 Lista 3 (Chi + OpenAPI) - até 03/04 (sex)
- 🚀 **Entrega U1** (Listas 1-3 + Sprint 0) - até 03/04 (sex)

---

## UNIDADE 2: Persistência e Arquitetura (Semanas 6-11)

### Semana 6 (06-09/04): PostgreSQL e sqlc

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 05/04 | Dom | 📺 | **Vídeo 6**: PostgreSQL + sqlc - SQL type-safe |
| 07/04 | Ter | 🟢 | Revisão + Prática: primeiras queries com sqlc |
| 09/04 | Qui | 🔵 | **Acompanhamento Online** - Sprint 1 |

**Conteúdo**: SQL vs ORM, sqlc config, queries type-safe  
**Entregas**: 🚀 Sprint 1 (vídeo 8min) - até 10/04 (sex)

---

### Semana 7 (13-16/04): Migrations e Clean Architecture

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 12/04 | Dom | 📺 | **Vídeo 7**: Migrations com Atlas + Clean Architecture |
| 14/04 | Ter | 🟢 | Revisão + Prática: estruturando camadas |
| 16/04 | Qui | 🟢 | Prática: Estruturando o projeto |

**Conteúdo**: Atlas migrations, camadas, inversão de dependência  
**Entregas**: 📝 Lista 4 (sqlc + Clean Architecture) - até 17/04 (sex)

---

### Semana 8 (20-23/04): Testes Automatizados

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 19/04 | Dom | 📺 | **Vídeo 8**: Testes em Go - unitários, integração, e2e |
| 21/04 | Ter | 🔴 | **Tiradentes** - Sem aula (assistir vídeo) |
| 23/04 | Qui | 🔵 | **Acompanhamento Online** - Sprint 2 |

**Conteúdo**: testing package, testify, mocks, dockertest  
**Entregas**: 🚀 Sprint 2 (vídeo 8min) - até 24/04 (sex)

---

### Semana 9 (27-30/04): Autenticação e Autorização

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 26/04 | Dom | 📺 | **Vídeo 9**: JWT + Refresh tokens + OAuth 2.0 |
| 28/04 | Ter | 🟢 | Revisão + Prática: implementando auth |
| 30/04 | Qui | 🟢 | Prática: fluxos de autenticação |

**Conteúdo**: JWT estrutura, refresh token rotation, OAuth 2.0 flows  
**Entregas**: 📝 Lista 5 (Autenticação) - até 01/05 (sex)

---

### Semana 10 (04-07/05): Segurança e Revisão para Prova

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 03/05 | Dom | 📺 | **Vídeo 10**: Segurança - OWASP API Top 10 |
| 05/05 | Ter | 🟢 | Revisão + Prática: corrigindo vulnerabilidades |
| 07/05 | Qui | 🔵 | **Acompanhamento Online** - Sprint 3 + Revisão para prova |

**Conteúdo**: OWASP Top 10, rate limiting, input validation  
**Entregas**: 
- 🚀 Sprint 3 (vídeo 8min) - até 08/05 (sex)
- 🚀 **Entrega U2** (Listas 4-5 + Sprints 1-3) - até 08/05 (sex)

---

### Semana 11 (11-14/05): PROVA PRÁTICA

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 12/05 | Ter | 📚 | **PROVA PRÁTICA - Unidade 2** (laboratório, 1h40min) |
| 14/05 | Qui | 🔴 | Sem aula - Correção da prova (vídeo disponibilizado posteriormente) |

**Conteúdo da Prova**: Todo material das Unidades 1 e 2  
**Formato**: Individual, sem consulta externa, implementação de mini-API

---

## UNIDADE 3: DevOps e Tópicos Avançados (Semanas 12-17)

### Semana 12 (18-21/05): Containerização e CI/CD

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 17/05 | Dom | 📺 | **Vídeo 11**: Docker multi-stage + CI/CD + GOMAXPROCS container-aware (Go 1.25+) |
| 19/05 | Ter | 🟢 | Revisão + Prática: Dockerizando o projeto |
| 21/05 | Qui | 🔵 | **Acompanhamento Online** - Sprint 4 |

**Conteúdo**: Dockerfile, docker-compose, GitHub Actions, GOMAXPROCS, graceful shutdown, health checks  
**Entregas**: 
- 📝 Lista 6 (CI/CD) - até 22/05 (sex)
- 🚀 Sprint 4 (vídeo 8min) - até 22/05 (sex)

---

### Semana 13 (25-28/05): Observabilidade e Frontend

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 24/05 | Dom | 📺 | **Vídeo 12**: Observabilidade - logs, métricas, traces |
| 26/05 | Ter | 🟢 | Revisão + Prática: instrumentando com slog |
| 28/05 | Qui | 🔵 | **Acompanhamento Online** - CORS + Dúvidas de integração |

**Conteúdo**: slog (Go 1.25 GroupAttrs), OpenTelemetry, CORS, integração frontend

---

### Semana 14 (01-04/06): Introdução a GraphQL

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 31/05 | Dom | 📺 | **Vídeo 13**: GraphQL vs REST + gqlgen |
| 02/06 | Ter | 🟢 | Revisão + Prática: Queries e Mutations |
| 04/06 | Qui | 🔴 | **Corpus Christi** - Sem aula |

**Conteúdo**: GraphQL SDL, gqlgen, schema-first development

---

### Semana 15 (08-11/06): Concorrência e Padrões Avançados

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 07/06 | Dom | 📺 | **Vídeo 14**: Concorrência - goroutines, channels + **Go 1.25** |
| 09/06 | Ter | 🟢 | Revisão + Prática: Worker pools |
| 11/06 | Qui | 🔵 | **Acompanhamento Online** - Sprint 5 (final) |

**Conteúdo**: Goroutines, channels, sync.WaitGroup.Go(), testing/synctest, errgroup, worker pools  
**Entregas**: 🚀 Sprint 5 - MVP completo (vídeo 10min) - até 12/06 (sexta)

---

### Semana 16 (15-18/06): Apresentações Finais

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 14/06 | Dom | 📺 | **Vídeo 15**: Próximos passos + Carreira em Go |
| 16/06 | Ter | 🟢 | **Apresentações finais** (Grupos 1-6) |
| 18/06 | Qui | 🟢 | **Apresentações finais** (Grupos 7-12) |

**Formato**: 12 minutos por equipe + 3 minutos de perguntas (15 min total)  
**Entregas**: 
- 🚀 **Entrega U3** (Lista 6 + Projeto Final) - até 26/06 (sexta)

---

### Semana 17 (22-25/06): Correção e Encerramento

| Data | Dia | Tipo | Atividade |
|------|-----|------|-----------|
| 23/06 | Ter | 🔴 | Sem aula - Professor corrigindo trabalhos |
| 25/06 | Qui | 🔵 | **Acompanhamento Online** - Dúvidas sobre correção e notas |

**Prazo final**: Entrega U3 até 26/06 (sexta) às 23:59  
**Consolidação**: Notas no SIGAA até 11/07

---

## Resumo de Entregas

### Listas de Exercícios (GitHub Classroom)

| Lista | Conteúdo | Prazo | Peso |
|-------|----------|-------|------|
| Lista 1 | net/http básico | 13/03 (sex) | 15% |
| Lista 2 | net/http + middleware | 20/03 (sex) | 15% |
| Lista 3 | Chi + OpenAPI | 03/04 (sex) | 20% |
| Lista 4 | sqlc + Clean Architecture | 17/04 (sex) | 20% |
| Lista 5 | Autenticação | 01/05 (sex) | 15% |
| Lista 6 | CI/CD | 22/05 (sex) | 15% |

### Sprints do Projeto

| Sprint | Conteúdo | Prazo | Duração Vídeo |
|--------|----------|-------|---------------|
| Sprint 0 | Proposta do projeto | 26/03 (qui) | 5 min |
| Sprint 1 | Endpoints básicos | 10/04 (sex) | 8 min |
| Sprint 2 | Persistência | 24/04 (sex) | 8 min |
| Sprint 3 | Autenticação | 08/05 (sex) | 8 min |
| Sprint 4 | CI/CD + Deploy | 22/05 (sex) | 8 min |
| Sprint 5 | MVP completo | 12/06 (sex) | 10 min |

### Entregas por Unidade

| Unidade | Prazo | Componentes |
|---------|-------|-------------|
| U1 | 03/04 (sex) | Listas 1-3 + Sprint 0 |
| U2 | 08/05 (sex) | Listas 4-5 + Sprints 1-3 + Prova (12/05) |
| U3 | 26/06 (sex) | Lista 6 + Sprints 4-5 + Apresentação Final |

---

## Feriados e Datas Especiais

| Data | Evento | Impacto |
|------|--------|---------|
| 02/04 | Quinta-feira Santa | Sem aula |
| 21/04 | Tiradentes | Sem aula (assistir vídeo) |
| 04/06 | Corpus Christi | Sem aula |

---

## Datas Importantes

| Data | Evento |
|------|--------|
| 12/03 | Prazo para formação de grupos |
| 03/04 (sex) | Entrega U1 |
| 08/05 (sex) | Entrega U2 |
| **12/05** | **PROVA PRÁTICA** |
| 16 e 18/06 | Apresentações finais |
| 26/06 (sex) | Entrega U3 (final) |
| 11/07 | Término do período |

---

## Acompanhamentos Online (Google Meet)

Os acompanhamentos acontecem às quintas-feiras, alternando com aulas presenciais:

| # | Data | Foco |
|---|------|------|
| 1 | 12/03 | Dúvidas de ambiente + Formação de grupos |
| 2 | 26/03 | Sprint 0 - Proposta do projeto |
| 3 | 09/04 | Sprint 1 - Endpoints básicos |
| 4 | 23/04 | Sprint 2 - Persistência |
| 5 | 07/05 | Sprint 3 - Autenticação + Revisão para prova |
| 6 | 21/05 | Sprint 4 - CI/CD |
| 7 | 28/05 | Dúvidas de integração + CORS |
| 8 | 11/06 | Sprint 5 - MVP final |
| 9 | 25/06 | Dúvidas sobre correção e notas |

**Horário**: 13:00-14:40 (mesmo horário das aulas presenciais)  
**Formato**: ~10 minutos por equipe para tirar dúvidas e revisar progresso.