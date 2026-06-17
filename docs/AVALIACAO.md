# Avaliação - DIM0547

## Desenvolvimento de Sistemas Web II com Go

**Período**: 2026.1  
**Prof. Fernando Figueira**

> **⚠️ Atualização 05/05/2026**: Datas ajustadas para reposição de aulas não realizadas. Prova adiada para **02/06 (ter)**. Apresentações finais adiadas para **30/06 (ter) e 02/07 (qui)**. Entrega U3 passa a ser **04/07 (sex)**.
>
> **⚠️ Lista 3**: prazo de entrega prorrogado para **12/05 (ter) às 23:59**.
>
> **⚠️ Atualização 25/05/2026**:
> - **Listas 5 e 6 unificadas** em uma única **Lista 5+6** (auth + segurança + testes), prazo **03/06 (qua)**.
> - **Lista 7 removida** do cronograma.
> - **Sprint 3** com prazo prorrogado para **16/06 (ter)**.
> - **Sprint 4** mesclada com a Apresentação Final em um único entregável: o **Vídeo Final**, com prazo **30/06 (ter) às 23:59** (substitui a apresentação presencial).
> - **Pesos redistribuídos** nas três unidades, com cada componente explicitado individualmente:
>   - **U1**: Sprint 0 (30%) + Sprint 1 (40%) + Lista 1 (10%) + Lista 2 (10%) + Lista 3 (10%)
>   - **U2**: Prova (20%) + Sprint 2 (30%) + Sprint 3 (30%) + Lista 4 (10%) + Lista 5+6 (10%)
>   - **U3**: Prova (50%) + Vídeo Final (50%)
> - A **Sprint 1** passa a contar em U1 (deixa de pesar em U2). A **Prova** continua sendo uma única prova realizada em 02/06, com pesos diferentes em U2 (20%) e U3 (50%).
>
> **⚠️ Atualização 06/06/2026**:
> - A **Prova foi adiada para 11/06 (ter)**. Pesos mantidos: 20% da U2 e 50% da U3.
> - A **aula de 02/06 foi cancelada**; o conteúdo de testes + segurança fica no Vídeo 8 e não é foco da prova.
> - A **Lista 5+6** passa a ser entregue em **23/06 (ter)**.
> - A **Sprint 3 é fundida na Entrega Final do Projeto**, entregue em **23/06 (ter)** como **vídeo único**. Não há mais vídeo separado de Sprint 3: a mesma entrega é avaliada como **Sprint 3 (30% da U2)** e como **Vídeo Final (50% da U3)** — os pesos das duas unidades recaem sobre o mesmo entregável.
>
> **⚠️ Atualização 17/06/2026**:
> - **Listas 1–4 reabertas valendo 70%**: quem não submeteu pode entregar até **25/06 (qui)** por uma **tarefa de reposição** específica. Quem já entregou no prazo mantém a nota.
> - **Lista 5+6 não será publicada** (sem tempo hábil). Será **resolvida no Vídeo 8** (Autenticação e Segurança — Parte 2). Os conceitos seguem sendo cobrados na **Entrega Final** (Sprint 3 / Vídeo Final).
> - **Peso da Lista 5+6 redistribuído**: o peso que a Lista 5+6 tinha na U2 foi somado à **Lista 4** (de peso 1 para **peso 2**).
> - **Notas da prova e das Listas 1–4 já estão disponíveis no SIGAA.**
> - **Sem aula presencial nesta semana**; professor no Meet no horário das aulas (16 e 18/06) para dúvidas. Vídeos 8 e 9 a serem publicados nos próximos dias.

---

## Estrutura Geral

A avaliação é organizada em **3 unidades** com pesos iguais. A nota final é a média aritmética das três unidades:

```
Nota Final = (U1 + U2 + U3) / 3
```

A **Prova** (aplicada no Multiprova) é compartilhada entre U2 e U3, valendo 20% de U2 e 50% de U3.

### Tipos de Avaliação

| Tipo | Componentes | Descrição |
|------|-------------|-----------|
| **Individual** | Listas de exercícios, Prova | Trabalho próprio, sem colaboração |
| **Em grupo** | Sprints, Vídeo Final | Equipes de 3-4 pessoas |

> **Importante**: As **listas de exercícios são individuais** e qualquer forma de cópia ou plágio resultará em nota zero. O **projeto é em grupo**, e a contribuição individual de cada membro é avaliada através de commits, presença nos vídeos e capacidade de responder perguntas sobre o código.

---

## Calendário de Entregas

> Todas as datas de entrega do curso estão consolidadas aqui.  
> Prazos encerram às **23:59** do dia indicado.

### Listas de Exercícios (Individual)

| Lista | Conteúdo | Prazo | Unidade |
|-------|----------|-------|---------|
| Lista 1 | net/http básico | 27/03 (sex) ✓ | U1 |
| Lista 2 | net/http + middleware | 24/04 (sex) ✓ | U1 |
| Lista 3 | Chi + OpenAPI + validator | 12/05 (ter) ✓ | U1 |
| Lista 4 | sqlc + Repository + filtros + JOIN 1:N | 28/05 (qui) | U2 |
| ~~**Lista 5+6**~~ | ~~Auth JWT + autorização + refresh + rate limit + testes~~ | **Não publicada** (resolvida no Vídeo 8) | — |
| **Reabertura Listas 1–4** | Reposição (vale **70%**) para quem não submeteu | **25/06 (qui)** | U1/U2 |

> ~~Lista 7 (Deploy + Observabilidade + gRPC)~~ — **removida** após reorganização de 25/05

### Sprints do Projeto (Grupo)

| Sprint | Conteúdo | Prazo | Vídeo |
|--------|----------|-------|-------|
| Sprint 0 | Proposta do projeto | 06/04 (dom) ✓ | 5 min |
| Sprint 1 | API RESTful com Chi + middleware | 24/04 (sex) ✓ | 8 min |
| Sprint 2 | Persistência com PostgreSQL e relacionamentos 1:N | 19/05 (ter) ✓ | 8 min |
| Sprint 3 | Autenticação JWT + Segurança + Testes | **23/06 (ter)** | — |
| **🎥 Vídeo Final** | MVP completo | **23/06 (ter)** | 12 min |

### Datas Especiais

| Evento | Data | Observação |
|--------|------|------------|
| Formação de grupos | 12/03 | Registrar no Discord |
| **Prova (Multiprova)** | **11/06 (qui)** ✓ | Laboratório, 13:00-14:40 |
| **Reabertura Listas 1–4** | **25/06 (qui)** | Tarefa de reposição, vale 70% |
| **Vídeo Final** | **23/06 (ter)** | Substitui apresentação presencial |

---

## Resumo por Unidade

| Unidade | Componentes | Tipo |
|---------|-------------|------|
| **U1** | Sprint 0 (30%) + Sprint 1 (40%) + Listas 1, 2 e 3 (10% cada) | Grupo + Individual |
| **U2** | Prova (20%) + Sprint 2 (30%) + Sprint 3 (30%) + Lista 4 (20%) | Individual + Grupo + Individual |
| **U3** | **Prova (50%)** + Vídeo Final (50%) | Individual + Grupo |

---

## Unidade 1 (Sprint 0 + Sprint 1 + Listas 1-3)

### Componentes e Pesos

| Componente | Peso | Tipo | Descrição |
|------------|------|------|-----------|
| Sprint 0 | 30% | Grupo | Proposta do projeto (vídeo 5min + documento) |
| Sprint 1 | 40% | Grupo | API RESTful idiomática com Chi + middleware (vídeo 8min) |
| Lista 1 | 10% | Individual | net/http básico — entregue ✓ |
| Lista 2 | 10% | Individual | net/http + middleware — entregue ✓ |
| Lista 3 | 10% | Individual | Chi + OpenAPI + validator — entregue ✓ |

### Fórmula

```
U1 = (Sprint0 × 0.30) + (Sprint1 × 0.40)
   + (Lista1 × 0.10) + (Lista2 × 0.10) + (Lista3 × 0.10)
```

> Mudança 25/05: a Sprint 1 passa a contar **em U1** (antes pesava apenas dentro do bloco "Sprints 1-3" em U2). As três listas de U1 têm agora pesos explícitos e iguais (10% cada).

---

## Unidade 2 (Sprint 2 + Sprint 3 + Lista 4 + Prova)

### Componentes e Pesos

| Componente | Peso | Tipo | Descrição |
|------------|------|------|-----------|
| Prova | 20% | Individual | Prova individual no Multiprova (11/06) |
| Sprint 2 | 30% | Grupo | Persistência com PostgreSQL e relacionamentos 1:N (vídeo 8min) |
| Sprint 3 | 30% | Grupo | Autenticação JWT + segurança + testes — entregue na Entrega Final (23/06) |
| Lista 4 | 20% | Individual | sqlc + Repository + filtros + JOIN 1:N (absorve o peso da Lista 5+6) — entregue ✓ |

### Fórmula

```
U2 = (Prova × 0.20)
   + (Sprint2 × 0.30) + (Sprint3 × 0.30)
   + (Lista4 × 0.20)
```

> Pesos da U2: Prova 20%, Sprint 2 30%, Sprint 3 30%, Lista 4 20%. A Lista 4 absorveu o peso que era da Lista 5+6 (que não será publicada — atualização de 17/06).
>
> Mudança 06/06: a **Sprint 3 deixa de ter vídeo próprio** e é entregue dentro da **Entrega Final do Projeto (23/06)**. Os 30% de Sprint 3 (U2) e os 50% de Vídeo Final (U3) recaem a mesma entrega do último vídeo de projeto, avaliado segundo os critérios publicados na tarefa.

---

## Unidade 3 (Vídeo Final)

### Componentes e Pesos

| Componente | Peso | Tipo | Descrição |
|------------|------|------|-----------|
| **Prova** | **50%** | Individual | Mesma prova da U2 (a ser realizada em 11/06) |
| **Vídeo Final** | 50% | Grupo | MVP completo |

### Fórmula

```
U3 = (Prova × 0.50) + (VídeoFinal × 0.50)
```



## Vídeo Final (50% da U3)

### Descrição

O Vídeo Final é a **apresentação do projeto** da equipe. Ele substitui a **Sprint 3** (entrega técnica de auth/segurança/testes), a antiga Sprint 4 e a antiga apresentação presencial — todas unificadas em um único vídeo de 12 minutos, entregue em **23/06**. Como absorve a Sprint 3, o vídeo deve **também** evidenciar: login retornando JWT, rotas protegidas por middleware, refresh token, ao menos 2 correções de segurança OWASP e ao menos 10 testes no CI.

### Informações Gerais

- **Prazo**: **23/06/2026 (terça) às 23:59**
- **Formato**: Vídeo gravado (link YouTube/Drive acessível)
- **Duração**: **12 minutos**

### Estrutura Sugerida (12 minutos)

```
1. Introdução (1-2 min)
   - Nome do projeto e equipe
   - Problema que resolve
   - Público-alvo

2. Demonstração do MVP (5-6 min)
   - Funcionalidades principais funcionando
   - Fluxos de usuário completos
   - Tratamento de erros (mostrar pelo menos um caso)
   - Autenticação JWT em ação

3. Arquitetura e Código (3-4 min)
   - Diagrama de arquitetura (1 slide)
   - Decisões técnicas importantes
   - Métricas: cobertura de testes, endpoints documentados
   - DEMONSTRAÇÃO DE PROTOCOLO ALTERNATIVO: pelo menos 1 endpoint gRPC OU GraphQL (bônus na nota)

4. Conclusão (1 min)
   - Aprendizados principais
   - O que faria diferente
```

### Dicas de Preparação

1. **Ensaie antes** — 12 minutos passam rápido
2. **Edite o vídeo** — evite pausas longas, mantenha o ritmo
3. **Divida as falas** — é considerado positivo que todos membros participem
4. **Áudio** — teste a gravação da voz para verificar se dá pra entender
5. **Teste antes de enviar** — verifique que o arquivo abre e tem áudio

### Critérios de Avaliação do Vídeo Final

| Critério | Peso | Descrição |
|----------|------|-----------|
| **Demonstração do MVP** | 30% | Funcionalidades funcionando, fluxos completos, tratamento de erros |
| **Qualidade técnica** | 25% | Arquitetura clara, código organizado, boas práticas |
| **Exploração de protocolo alternativo** | 15% | Demonstração de endpoint gRPC OU GraphQL implementado |
| **Comunicação** | 15% | Clareza, domínio do conteúdo, participação de todos |
| **Organização** | 10% | Respeito ao tempo, estrutura do vídeo, edição |
| **Reflexão** | 5% | Aprendizados, o que faria diferente |

---

## Prova (20% de U2 + 50% de U3)

### Informações Gerais

- **Data**: 11/06/2026 (Terça-feira)
- **Horário**: 13:00 às 14:40
- **Local**: Laboratório A307 (IMD)
- **Duração**: 1h40min
- **Formato**: Individual, aplicada no **Multiprova** (sem consulta externa)

A prova é uma única avaliação realizada em 11/06, com pesos diferentes em cada unidade: vale 20% da U2 e 50% da U3.

### Conteúdo

Todo o material das **Unidades 1 e 2**:

- Fundamentos de HTTP e arquitetura web
- Go: structs, interfaces, error handling
- net/http: handlers, middleware, context
- Chi: rotas, grupos, middleware
- JSON e validação
- sqlc e persistência
- Relacionamentos 1:N e JOINs
- Repository pattern (desacoplamento de camadas)

> Nota: **autenticação JWT, testes e segurança/OWASP estão excluídos da prova** — os alunos não tiveram tempo de praticar esses conceitos. Ficam por conta do **Vídeo 8** e são cobrados no projeto e na Lista 5+6. A prova concentra-se nos fundamentos das unidades 1 e 2 (Go, net/http, Chi, JSON/validação, sqlc, Repository).

### Formato da Prova

A prova é aplicada no **Multiprova**, combinando questões de múltipla escolha, verdadeiro/falso, resposta curta e leitura/análise de código Go. A composição e a quantidade de questões serão definidas na própria prova.

### Critérios de Avaliação da Prova

A nota é a pontuação obtida no conjunto de questões do Multiprova, distribuída entre:

| Critério | Peso |
|----------|------|
| Conceitos (HTTP, Go, Chi, middleware, JSON, validação) | 50% |
| Persistência (sqlc, Repository, relacionamentos 1:N) | 30% |
| Leitura/análise de código | 20% |

### Preparação

- Revise as listas de exercícios 1-4
- Revise os conceitos e padrões do curso
- Pratique leitura e depuração de código Go
- Estude os exemplos das aulas

---

## Projeto Integrador

### Sprint 0: Proposta do Projeto (30% da U1)

O Sprint 0 é a **fase de planejamento** onde a equipe define o que será construído ao longo do semestre.

#### Entregáveis do Sprint 0

1. **Vídeo de apresentação** (5 minutos)
2. **Documento de proposta** (PDF, 2-3 páginas)
3. **Repositório do projeto** (criado no GitHub)

---

#### Visão do Produto

A visão do produto é uma declaração estratégica que responde: "Por que este produto existe?" e "Qual problema ele resolve?"

**Template obrigatório:**

```
Para [usuários-alvo]
Que [problema/necessidade]
O [nome do produto] é um [categoria do produto]
Que [benefício principal/capacidade]
Diferente de [alternativa existente]
Nosso produto [diferencial único]
```

**Exemplo — API de Gestão de Biblioteca:**

```
Para bibliotecários e estudantes universitários
Que perdem tempo com processos manuais de empréstimo e busca
O BiblioAPI é uma API REST de gestão de biblioteca
Que automatiza empréstimos, devoluções e busca de acervo
Diferente de sistemas legados como Pergamum
Nosso produto oferece API moderna, documentada e fácil de integrar
```

**Checklist da Visão:**
- [ ] Define claramente o usuário-alvo
- [ ] Identifica o problema específico a ser resolvido
- [ ] Explicita o valor único oferecido
- [ ] É realista para o escopo do curso (1 semestre, equipe de 3-4)

---

#### Definição do MVP

O MVP (Minimum Viable Product) define o **escopo mínimo** que entrega valor ao usuário.

**Framework de definição:**

1. **Problema Central**: Qual o principal problema que seu produto resolve?
2. **Hipótese de Valor**: "Acreditamos que [usuários] vão [comportamento] porque [benefício]"
3. **Funcionalidades Essenciais**: Lista priorizada do que DEVE estar no MVP
4. **Fora do Escopo**: O que explicitamente NÃO será implementado

**Exemplo — BiblioAPI:**

| No MVP | Fora do MVP |
|--------|-------------|
| CRUD de livros | Recomendações por IA |
| CRUD de usuários | Integração com e-commerce |
| Sistema de empréstimo/devolução | Multas automáticas |
| Busca por título/autor | Reserva de livros |
| Autenticação JWT | Notificações por email |

---

#### Product Backlog

O backlog é a **lista priorizada** de tudo que será desenvolvido, organizada por sprints.

**Estrutura recomendada:**

| Prioridade | User Story | Critérios de Aceitação | Sprint |
|------------|------------|------------------------|--------|
| P1 | Como bibliotecário, quero cadastrar livros para manter o acervo atualizado | Campos: título, autor, ISBN, quantidade. Validação de ISBN único. | 1 |
| P1 | Como usuário, quero me autenticar para acessar o sistema | Login com email/senha. Token JWT retornado. Refresh token implementado. | 3 |
| P2 | Como estudante, quero buscar livros por título para encontrar o que preciso | Busca parcial (LIKE). Retorna lista paginada. Ordena por relevância. | 2 |
| P2 | Como bibliotecário, quero registrar empréstimos para controlar o acervo | Vincula usuário + livro. Data de devolução calculada. Atualiza quantidade disponível. | 2 |

**Dicas para o backlog:**
- Use formato de User Story: "Como [papel], quero [ação] para [benefício]"
- Priorize: P1 = essencial para MVP, P2 = importante, P3 = desejável
- Estime: Use horas ou pontos de complexidade
- Distribua nos sprints de forma realista

---

#### Estrutura do Vídeo (5 minutos)

```
1. Apresentação da Equipe (30s)
   - Nome da equipe e do projeto
   - Integrantes e papéis (quem faz o quê)

2. Visão do Produto (1min30s)
   - Problema que resolve
   - Público-alvo
   - Proposta de valor (usar o template)

3. Definição do MVP (1min30s)
   - Funcionalidades essenciais (o que entra)
   - O que NÃO está no MVP (o que fica de fora)
   - Critérios de sucesso

4. Stack Tecnológico (1min)
   - Go + Chi + PostgreSQL + sqlc (obrigatório)
   - Outras escolhas e justificativas

5. Planejamento (30s)
   - Visão geral do backlog
   - Riscos identificados
```

---

#### Documento de Proposta (PDF, 2-3 páginas)

O documento deve conter:

1. **Visão do Produto** (usando o template)
2. **Definição do MVP** (funcionalidades dentro e fora do escopo)
3. **Product Backlog** (tabela com user stories priorizadas)
4. **Stack Tecnológico** (com justificativas)
5. **Equipe** (nome, matrícula, papel de cada integrante)
6. **Link do repositório** (já criado no GitHub)

---

#### Exemplos de Projetos Aceitos

**Nível de complexidade esperado:**

1. **API de Gerenciamento de Tarefas**
   - CRUD de tarefas e projetos
   - Categorias e tags
   - Autenticação de usuários
   - Filtros e busca

2. **API de E-commerce Simples**
   - Catálogo de produtos
   - Carrinho de compras
   - Checkout (simulado)
   - Histórico de pedidos

3. **API de Agendamento**
   - Cadastro de serviços
   - Disponibilidade de horários
   - Reservas
   - Confirmação

4. **API de Rede Social Básica**
   - Perfis de usuário
   - Posts e comentários
   - Likes e seguidores
   - Feed simples

---

#### Critérios de Avaliação do Sprint 0

| Critério | Peso | Descrição |
|----------|------|-----------|
| Clareza da visão | 25% | Problema e solução bem definidos, template preenchido corretamente |
| Escopo do MVP | 25% | Funcionalidades realistas, priorizadas, com critérios de aceitação |
| Viabilidade técnica | 20% | Stack adequado, riscos mapeados, escopo factível |
| Organização da equipe | 15% | Papéis claros, backlog bem estruturado |
| Qualidade do vídeo | 15% | Comunicação clara, todos participam, respeita o tempo |

---

#### Sprints 1-4: Desenvolvimento (8min cada)
- Funcionalidade implementada (30%)
- Qualidade do código (25%)
- Testes automatizados (20%)
- Demonstração funcional (15%)
- Comunicação técnica (10%)

> A Sprint 4 deixou de existir como entrega independente após a reorganização de 25/05 — agora é parte do Vídeo Final.

### Estrutura do Vídeo de Sprint

```
1. Introdução (30s)
   - Nome da equipe e do projeto
   - Objetivo do sprint

2. Demonstração (50% do tempo)
   - Funcionalidades implementadas funcionando
   - Testes passando
   - Deploy (quando aplicável)

3. Código (30% do tempo)
   - Principais decisões técnicas
   - Trechos relevantes de código
   - Arquitetura

4. Desafios e Próximos Passos (20% do tempo)
   - Dificuldades encontradas
   - Como foram superadas
   - Planejamento do próximo sprint
```

### Requisitos Técnicos do Projeto

#### Obrigatórios
- [ ] API REST com Go + Chi
- [ ] Persistência com PostgreSQL + sqlc
- [ ] Autenticação JWT
- [ ] Testes automatizados (cobertura ≥ 70%)
- [ ] Documentação OpenAPI
- [ ] CI/CD com GitHub Actions
- [ ] Graceful shutdown implementado
- [ ] Health checks (/healthz, /readyz)
- [ ] **Exploração de protocolo alternativo no Vídeo Final**: pelo menos 1 endpoint gRPC ou GraphQL implementado

#### Opcionais (Bônus até +30% na nota do projeto)
- [ ] Endpoint gRPC E GraphQL implementados (+5%)
- [ ] Deploy em produção funcionando (+5%)
- [ ] Observabilidade (logs estruturados, métricas) (+5%)
- [ ] Docker multi-stage com imagem otimizada (+3%)
- [ ] Rate limiting implementado (+3%)
- [ ] Integração com frontend funcional (+5%)
- [ ] Cache (Redis) (+4%)
- [ ] WebSockets (+3%)
- [ ] testing/synctest para testes de concorrência (+3%)

---

## Exercícios Práticos

### Sistema de Correção Automática

Os exercícios utilizam **GitHub Classroom com autograding**, oferecendo:

- Feedback imediato a cada push
- Execução automática de testes Go
- Múltiplas tentativas antes do prazo
- Visualização de resultados no próprio repositório

### Como Funciona o Autograding

1. **Aceite a tarefa** no GitHub Classroom (link no Discord e no SIGAA)
2. **Clone o repositório** criado automaticamente
3. **Implemente as soluções** seguindo as instruções
4. **Faça push** para o GitHub
5. **Verifique os resultados** na aba "Actions" do repositório

### Critérios de Avaliação dos Exercícios

- **Testes passando** (70%): Quantidade de testes automatizados que passam
- **Qualidade do código** (20%): golangci-lint sem erros/warnings
- **Prazo** (10%): Entregas após o prazo perdem pontos conforme política abaixo

### Política de Atraso

- Até 1 dia: -20%
- Até 3 dias: -50%
- Após 3 dias: Não aceito

---

## Estrutura de Entrega

### Entregas de Unidade (SIGAA)

#### U1: Entregue ✓

#### U2: Entrega até 23/06 (Sprint 3 + Listas) + Prova em 11/06

```
equipe-nome-u2.zip
├── README.md                    # Links e informações
├── listas/
│   ├── lista-04/               # Link para repo GitHub Classroom
│   └── lista-05-06/            # Link para repo GitHub Classroom (Lista 5+6 unificada)
├── sprints/
│   ├── sprint-1-video.txt      # Link do vídeo Sprint 1
│   ├── sprint-2-video.txt      # Link do vídeo Sprint 2
│   └── sprint-3-video.txt      # Link do vídeo Sprint 3
└── projeto/
    └── repo-link.txt           # Link do repositório do projeto
```

#### U3: Entrega Final até 23/06

```
equipe-nome-final.zip
├── README.md                    # Resumo e links
├── video-final/
│   └── video-link.txt          # Link do Vídeo Final (12min)
├── projeto/
│   ├── repo-link.txt           # Link do repositório
│   └── deploy-link.txt         # URL da aplicação (se aplicável)
└── documentacao/
    └── documentacao.pdf        # Documentação técnica (5-10 páginas)
```

### Documentação Técnica Final (PDF, 5-10 páginas)

1. **Visão Geral** (1 página): Problema, solução, escopo do MVP
2. **Arquitetura** (2-3 páginas): Diagrama de componentes, decisões técnicas, stack
3. **Funcionalidades** (1-2 páginas): Lista de endpoints, principais fluxos
4. **Qualidade** (1 página): Métricas de testes, análise de código
5. **Exploração técnica** (1 página): gRPC ou GraphQL implementado — escolha e justificativa
6. **Equipe e Contribuições** (1 página): Membros, responsabilidades, distribuição de commits

---

## Frequência

- **Mínimo**: 75% de presença para aprovação
- **Justificativas**: Via SIGAA até 48h após a falta
- **Aulas online**: Presença verificada via Google Meet

---

## Formação de Grupos

- **Tamanho**: 3 a 4 pessoas
- **Formação**: Até 12/03/2026 ✓
- **Comunicação**: Grupo próprio no Discord

---

## Política de Integridade Acadêmica

### Avaliações Individuais (Listas e Prova)

As **listas de exercícios** e a **prova prática** são avaliações **estritamente individuais**.

**Permitido**: consultar documentação oficial, pesquisar conceitos, discutir ideias em alto nível, usar snippets das aulas como base.

**Não Permitido**: copiar código de colegas, compartilhar soluções das listas, usar ferramentas de IA para gerar código sem entender, plagiar soluções da internet sem adaptação significativa.

### Avaliações em Grupo (Projeto e Vídeo Final)

O **projeto** (sprints e Vídeo Final) é desenvolvido em equipe. A contribuição individual é avaliada através de:
- Commits no repositório
- Presença demonstrada nos vídeos (todos devem aparecer no Vídeo Final)
- Capacidade de responder perguntas sobre qualquer parte do código

---

## Dúvidas Frequentes

### Sobre o Vídeo Final (23/06)

**P: Por que substituiu a apresentação presencial?**  
R: Reorganização de cronograma após reposições. Vídeo dá mais flexibilidade para a equipe gravar com calma e permite uma janela maior para correção antes do fim do período.

**P: Todos os membros precisam aparecer no vídeo?**  
R: Sim, a presença de todos os membros entra no critério "Comunicação" na avaliação do Vídeo Final.

**P: Posso usar vídeo gravado de sprint anterior?**  
R: Não. O Vídeo Final tem estrutura própria (focada no MVP completo + exploração técnica). Pode reaproveitar slides ou recortes, mas o vídeo final deve ser único.

### Sobre os Exercícios

**P: Posso fazer mais de um push?**  
R: Sim! Cada push executa os testes novamente. Você pode iterar até o prazo.

**P: Não entreguei a Lista 1, 2, 3 ou 4. Ainda dá?**  
R: Sim. As Listas 1–4 foram **reabertas valendo 70%**. Submeta pela **tarefa de reposição** específica até **25/06 (qui)**. Quem já entregou no prazo não precisa fazer nada — a nota original é mantida.

**P: A Lista 5+6 vai ser publicada?**  
R: Não. Por falta de tempo hábil, a Lista 5+6 **não será publicada como tarefa**. Ela é **resolvida no Vídeo 8** (Autenticação e Segurança — Parte 2). Os conceitos (JWT, autorização/ownership, refresh tokens, rate limiting, testes) seguem sendo cobrados na **Entrega Final**. O peso que ela tinha foi somado à **Lista 4**, que passa a valer 20% da U2 (peso 2 na planilha).

**P: Cadê a Lista 7?**  
R: Removida na reorganização de 25/05. O conteúdo (gRPC) aparece no Vídeo Final como requisito de exploração técnica.

### Sobre os Sprints

**P: Cadê a Sprint 4?**  
R: Mesclada com a apresentação final no Vídeo Final.

**P: Posso reenviar o vídeo se ficou ruim?**  
R: Sim, até o prazo. A última versão será avaliada.

### Sobre a Prova

**P: Posso consultar documentação durante a prova?**  
R: Apenas a documentação oficial do Go e uma folha A4 manuscrita frente e verso.

**P: A prova cobra refresh tokens e OWASP?**  
R: Não. Auth JWT, testes e OWASP estão **excluídos da prova** (os alunos não tiveram tempo de praticar). Ficam no Vídeo 8 e são exercitados no projeto e na Lista 5+6. A prova foca nos fundamentos das unidades 1 e 2 (Go, net/http, Chi, JSON, sqlc, Repository).

---

## Dúvidas e Suporte

### Canais de Atendimento

| Canal | Uso | Resposta |
|-------|-----|----------|
| Discord | Dúvidas técnicas, discussões | Até 24h |
| Aula presencial | Dúvidas conceituais | Imediato |
| Acompanhamento online | Revisão de projeto | Agendado |

### Horários de Atendimento

- **Presencial**: Após as aulas (14:40-15:10)
- **Online**: Discord, resposta até 24h em dias úteis
- **Acompanhamentos**: ver cronograma

---

## Checklist do Semestre

### Início do Curso ✓
- [x] Ambiente configurado (Go, Docker, PostgreSQL)
- [x] Acesso ao Discord da disciplina
- [x] Acesso ao GitHub Classroom
- [x] Grupo formado e registrado

### Durante o Curso
- [ ] Frequência ≥ 75%
- [ ] Exercícios entregues no prazo (Lista 5+6 até 23/06)
- [ ] Entrega Final (Sprint 3 + Vídeo) até 23/06
- [ ] Prova realizada (11/06)

### Final do Curso
- [ ] Vídeo Final entregue até 23/06
- [ ] Todas as entregas consolidadas no SIGAA