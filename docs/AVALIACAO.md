# Avaliação - DIM0547

## Desenvolvimento de Sistemas Web II com Go

**Período**: 2026.1  
**Prof. Fernando Figueira**

---

## Estrutura Geral

A avaliação é organizada em **3 unidades** com pesos iguais. A nota final é a média aritmética das três unidades:

```
Nota Final = (U1 + U2 + U3) / 3
```

A **Prova Prática** é compartilhada entre U2 e U3, valendo 50% de cada unidade — ou seja, equivale a **uma unidade completa** do curso.

### Tipos de Avaliação

| Tipo | Componentes | Descrição |
|------|-------------|-----------|
| **Individual** | Listas de exercícios, Prova | Trabalho próprio, sem colaboração |
| **Em grupo** | Sprints, Apresentação Final | Equipes de 3-4 pessoas, participação avaliada |

> **Importante**: As **listas de exercícios são individuais** e qualquer forma de cópia ou plágio resultará em nota zero. O **projeto é em grupo**, e a participação de cada membro é avaliada dentro da nota do projeto.

---

## Calendário de Entregas

> Todas as datas de entrega do curso estão consolidadas aqui.  
> Prazos encerram às **23:59** do dia indicado. Entregas até **sexta-feira** para evitar trabalho no fim de semana.

### Listas de Exercícios (Individual)

| Lista | Conteúdo | Prazo | Unidade |
|-------|----------|-------|---------|
| Lista 1 | net/http básico | 20/03 (sex) | U1 |
| Lista 2 | net/http + middleware | 27/03 (sex) | U1 |
| Lista 3 | Chi + OpenAPI | 03/04 (sex) | U1 |
| Lista 4 | sqlc + Clean Architecture | 17/04 (sex) | U2 |
| Lista 5 | Autenticação | 01/05 (sex) | U2 |
| Lista 6 | CI/CD | 22/05 (sex) | U3 |

### Sprints do Projeto (Grupo)

| Sprint | Conteúdo | Prazo | Vídeo |
|--------|----------|-------|-------|
| Sprint 0 | Proposta do projeto | 26/03 (qui) | 5 min |
| Sprint 1 | Endpoints básicos | 10/04 (sex) | 8 min |
| Sprint 2 | Persistência + testes | 24/04 (sex) | 8 min |
| Sprint 3 | Autenticação + segurança | 08/05 (sex) | 8 min |
| Sprint 4 | CI/CD + Deploy | 22/05 (sex) | 8 min |
| Sprint 5 | MVP completo | 12/06 (sex) | 10 min |

### Datas Especiais

| Evento | Data | Observação |
|--------|------|------------|
| Formação de grupos | 12/03 | Registrar no Discord |
| **Prova Prática** | **12/05 (ter)** | Laboratório, 13:00-14:40 |
| Apresentações Finais | 16 e 18/06 | 12 min + 3 min perguntas |
| Entrega Final U3 | 26/06 (sex) | ZIP no SIGAA |
| Consolidação SIGAA | 11/07 | Término oficial do período |

### Consolidação por Unidade

| Unidade | Período | Entrega Final |
|---------|---------|---------------|
| U1 | 02/03 a 03/04 | 03/04 (sex) |
| U2 | 06/04 a 08/05 | 08/05 (sex) + Prova 12/05 |
| U3 | 18/05 a 26/06 | 26/06 (sex) |

---

## Resumo por Unidade

| Unidade | Componentes | Tipo |
|---------|-------------|------|
| **U1** | Sprint 0 (60%) + Listas 1-3 (40%) | Grupo + Individual |
| **U2** | **Prova (50%)** + Sprints 1-3 (35%) + Listas 4-5 (15%) | Individual + Grupo + Individual |
| **U3** | **Prova (50%)** + Projeto Final (45%) + Lista 6 (5%) | Individual + Grupo + Individual |

---

## Unidade 1 (Semanas 1-5)

### Componentes e Pesos

| Componente | Peso | Tipo | Descrição |
|------------|------|------|-----------|
| Sprint 0 | 60% | Grupo | Proposta do projeto (vídeo 5min) + participação |
| Exercícios (Listas 1-3) | 40% | Individual | Correção automática via GitHub Classroom |

### Fórmula

```
U1 = (Sprint0 × 0.60) + (Exercícios × 0.40)
```

### Peso das Listas na nota de Exercícios U1

| Lista | Peso |
|-------|------|
| Lista 1 | 30% |
| Lista 2 | 30% |
| Lista 3 | 40% |

```
Exercícios_U1 = (Lista1 × 0.30) + (Lista2 × 0.30) + (Lista3 × 0.40)
```

---

## Unidade 2 (Semanas 6-11)

### Componentes e Pesos

| Componente | Peso | Tipo | Descrição |
|------------|------|------|-----------|
| **Prova Prática** | **50%** | Individual | Laboratório (12/05) |
| Sprints 1-3 | 35% | Grupo | Vídeos de acompanhamento (8min cada) + participação |
| Exercícios (Listas 4-5) | 15% | Individual | Correção automática via GitHub Classroom |

### Fórmula

```
U2 = (Prova × 0.50) + (Sprints × 0.35) + (Exercícios × 0.15)
```

### Peso das Listas na nota de Exercícios U2

| Lista | Peso |
|-------|------|
| Lista 4 | 55% |
| Lista 5 | 45% |

```
Exercícios_U2 = (Lista4 × 0.55) + (Lista5 × 0.45)
```

### Peso dos Sprints na nota de Sprints U2

| Sprint | Peso |
|--------|------|
| Sprint 1 | 30% |
| Sprint 2 | 35% |
| Sprint 3 | 35% |

```
Sprints_U2 = (Sprint1 × 0.30) + (Sprint2 × 0.35) + (Sprint3 × 0.35)
```

A nota de cada sprint inclui a avaliação da participação individual de cada membro do grupo.

---

## Unidade 3 (Semanas 12-17)

### Componentes e Pesos

| Componente | Peso | Tipo | Descrição |
|------------|------|------|-----------|
| **Prova Prática** | **50%** | Individual | Mesma prova da U2 (realizada em 12/05) |
| Projeto Final | 45% | Grupo | Sprints 4-5 + Apresentação Final + participação |
| Exercícios (Lista 6) | 5% | Individual | Correção automática via GitHub Classroom |

### Fórmula

```
U3 = (Prova × 0.50) + (Projeto × 0.45) + (Exercícios × 0.05)
```

### Peso do Projeto na U3

| Componente | Peso na U3 |
|------------|------------|
| Sprint 4 | 10% |
| Sprint 5 | 15% |
| Apresentação Final | 20% |

```
Projeto_U3 = Sprint4 (10%) + Sprint5 (15%) + Apresentação (20%) = 45%
```

A nota do projeto inclui a avaliação da participação individual de cada membro do grupo.

---

## Apresentação Final do Projeto (20% da U3)

### Informações Gerais

- **Datas**: 16/06 (terça) e 18/06 (quinta)
- **Horário**: 13:00 às 14:40
- **Local**: Sala de aula com projetor
- **Duração**: 12 minutos de apresentação + 3 minutos de perguntas (15 min total)
- **Grupos por dia**: 6 grupos (total de ~12 grupos)

### Formato

Todos os membros da equipe devem estar presentes e participar da apresentação. A equipe utilizará o projetor e um computador (próprio ou da sala) para demonstrar o MVP.

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

3. Arquitetura e Código (3-4 min)
   - Diagrama de arquitetura (1 slide)
   - Decisões técnicas importantes
   - Métricas: cobertura de testes, endpoints documentados

4. Conclusão (1 min)
   - Aprendizados principais
   - O que faria diferente
```

### Dicas de Preparação

1. **Ensaie com timer** - 12 minutos passam rápido
2. **Prepare backup da demo** - Tenha um vídeo gravado caso algo falhe
3. **Divida as falas** - Todos os membros devem falar
4. **Conheça todo o código** - Qualquer um pode receber perguntas sobre qualquer parte
5. **Tenha a API rodando** - Inicie os containers antes de começar

### Critérios de Avaliação da Apresentação

| Critério | Peso | Descrição |
|----------|------|-----------|
| **Demonstração do MVP** | 35% | Funcionalidades funcionando, fluxos completos, tratamento de erros |
| **Qualidade técnica** | 25% | Arquitetura clara, código organizado, boas práticas |
| **Comunicação** | 20% | Clareza, domínio do conteúdo, participação de todos |
| **Organização** | 10% | Respeito ao tempo, estrutura da apresentação |
| **Perguntas** | 10% | Qualidade das respostas da equipe |

### Ordem das Apresentações

A ordem será definida por sorteio e divulgada no Discord no dia 9/6 (terça da semana anterior à semana das apresentações).

---

## Prova Prática (50% de U2 + 50% de U3 = 1 Unidade)

### Informações Gerais

- **Data**: 12/05/2026 (Terça-feira)
- **Horário**: 13:00 às 14:40
- **Local**: Laboratório A307 (IMD)
- **Duração**: 1h40min
- **Formato**: Individual, sem consulta externa

A prova representa **uma unidade completa** do curso (50% de U2 + 50% de U3). Isso garante que cada estudante demonstre individualmente o domínio dos conceitos fundamentais.

### Conteúdo

Todo o material das **Unidades 1 e 2**:

- Fundamentos de HTTP e arquitetura web
- Go: structs, interfaces, error handling
- net/http: handlers, middleware, context
- Chi: rotas, grupos, middleware
- JSON e validação
- sqlc e persistência
- Clean Architecture
- Testes automatizados
- Autenticação JWT

### Formato da Prova

A prova consiste em **implementar uma mini-API** com requisitos específicos:

1. **Setup inicial fornecido** (10min)
   - Projeto base com estrutura definida
   - Banco de dados configurado
   - Testes de validação prontos

2. **Implementação** (80min)
   - Handlers para CRUD básico
   - Middleware de autenticação
   - Validação de dados
   - Testes unitários

3. **Execução dos testes** (10min)
   - Testes de validação automáticos
   - Correção parcial automática

### Critérios de Avaliação da Prova

| Critério | Peso |
|----------|------|
| Funcionalidade (testes passando) | 50% |
| Qualidade do código | 20% |
| Tratamento de erros | 15% |
| Testes implementados | 15% |

### Preparação

- Revise as listas de exercícios 1-5
- Pratique implementação de APIs do zero
- Familiarize-se com os padrões do curso
- Estude os exemplos das aulas

---

## Projeto Integrador

### Sprint 0: Proposta do Projeto (60% da U1)

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

#### Sprint 5: MVP Completo (10min)
- Funcionalidade completa (25%)
- Qualidade do código e arquitetura (25%)
- Testes e cobertura (20%)
- Deploy funcionando (15%)
- Documentação (15%)

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
- [ ] Deploy em container (Docker)
- [ ] CI/CD com GitHub Actions
- [ ] Graceful shutdown implementado
- [ ] Health checks (/healthz, /readyz)

#### Opcionais (Bônus até +30% na nota do projeto)
- [ ] Endpoint GraphQL com gqlgen (+5%)
- [ ] Rate limiting implementado (+3%)
- [ ] Observabilidade (logs estruturados, métricas) (+5%)
- [ ] Integração com frontend funcional (+5%)
- [ ] Cache (Redis) (+4%)
- [ ] WebSockets (+3%)
- [ ] testing/synctest para testes de concorrência (+3%)
- [ ] `sync.WaitGroup.Go()` no código (+2%)
- [ ] `errors.AsType` para tratamento de erros type-safe (+1%)

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

Cada unidade requer um **arquivo ZIP único** no SIGAA com a estrutura abaixo.

#### U1: Entrega até 03/04

```
equipe-nome-u1.zip
├── README.md                 # Links e informações gerais
├── listas/
│   ├── lista-01/            # Link para repo GitHub Classroom
│   ├── lista-02/            # Link para repo GitHub Classroom
│   └── lista-03/            # Link para repo GitHub Classroom
└── sprint-0/
    ├── video-link.txt       # Link para o vídeo
    └── proposta.pdf         # Documento de proposta (2-3 páginas)
```

#### U2: Entrega até 08/05

```
equipe-nome-u2.zip
├── README.md                    # Links e informações
├── listas/
│   ├── lista-04/               # Link para repo GitHub Classroom
│   └── lista-05/               # Link para repo GitHub Classroom
├── sprints/
│   ├── sprint-1-video.txt      # Link do vídeo Sprint 1
│   ├── sprint-2-video.txt      # Link do vídeo Sprint 2
│   └── sprint-3-video.txt      # Link do vídeo Sprint 3
└── projeto/
    └── repo-link.txt           # Link do repositório do projeto
```

#### U3: Entrega Final até 26/06

```
equipe-nome-final.zip
├── README.md                    # Resumo e links
├── lista/
│   └── lista-06/               # Link para repo GitHub Classroom
├── sprints/
│   ├── sprint-4-video.txt      # Link do vídeo Sprint 4
│   └── sprint-5-video.txt      # Link do vídeo Sprint 5
├── projeto/
│   ├── repo-link.txt           # Link do repositório
│   ├── deploy-link.txt         # URL da aplicação em produção
│   └── documentacao.pdf        # Documentação técnica (5-10 páginas)
└── apresentacao/
    └── slides.pdf              # Slides da apresentação
```

### Documentação Técnica Final (PDF, 5-10 páginas)

1. **Visão Geral** (1 página): Problema, solução, escopo do MVP
2. **Arquitetura** (2-3 páginas): Diagrama de componentes, decisões técnicas, stack
3. **Funcionalidades** (1-2 páginas): Lista de endpoints, principais fluxos
4. **Qualidade** (1 página): Métricas de testes, análise de código
5. **Deploy e Operação** (1 página): Ambiente, graceful shutdown, health checks, monitoramento
6. **Equipe e Contribuições** (1 página): Membros, responsabilidades, distribuição de commits

---

## Participação

### Componentes da Nota de Participação

| Item | Peso | Descrição |
|------|------|-----------|
| Presença | 30% | Frequência nas aulas presenciais |
| Code Reviews | 25% | Revisões de código entre equipes |
| Retrospectivas | 20% | Participação nas discussões de sprint |
| Contribuições | 15% | Ajuda a colegas, perguntas em aula |
| Discord | 10% | Engajamento no servidor da disciplina |

### Code Reviews Entre Equipes

A partir do Sprint 2, cada equipe fará code review de outra equipe:

1. **Semana do review**: Equipe A revisa código da Equipe B
2. **Formato**: Pull Request com comentários construtivos
3. **Prazo**: 48h após a entrega do sprint
4. **Avaliação**: Qualidade e utilidade dos comentários

### Frequência

- **Mínimo**: 75% de presença para aprovação
- **Justificativas**: Via SIGAA até 48h após a falta
- **Aulas online**: Presença verificada via Google Meet

---

## Formação de Grupos

### Requisitos

- **Tamanho**: 3 a 4 pessoas
- **Formação**: Até 12/03/2026
- **Comunicação**: Grupo próprio no Discord

### Processo

1. **Até 10/03**: Encontre sua equipe
2. **Até 12/03**: Registre no formulário (link no Discord)
3. **Até 15/03**: Crie o grupo no Discord da disciplina
4. **Até 26/03**: Defina o tema do projeto (Sprint 0)

### Se Não Encontrar Grupo

1. Poste no canal #procurando-grupo do Discord
2. Participe dos grupos abertos
3. O professor ajudará a formar grupos até 12/03

### Avaliação Individual vs Grupo

Embora o projeto seja em grupo, a avaliação considera:

- **Commits no GitHub**: Distribuição do trabalho
- **Participação nos vídeos**: Todos devem apresentar
- **Code reviews**: Contribuição individual
- **Conhecimento na apresentação**: Todos respondem perguntas

---

## Política de Integridade Acadêmica

### Avaliações Individuais (Listas e Prova)

As **listas de exercícios** e a **prova prática** são avaliações **estritamente individuais**. Cada estudante deve desenvolver suas próprias soluções.

**Permitido:**
- Consultar documentação oficial
- Pesquisar conceitos e técnicas
- Discutir ideias em alto nível com colegas
- Usar snippets de código das aulas como base

**Não Permitido:**
- Copiar código de colegas (total ou parcialmente)
- Compartilhar soluções das listas
- Usar ferramentas de IA para gerar código sem entender
- Plagiar soluções da internet sem adaptação significativa

### Avaliações em Grupo (Projeto)

O **projeto** (sprints e apresentação) é desenvolvido em equipe. A participação individual é avaliada através de:
- Commits no repositório
- Contribuição demonstrada nos vídeos
- Capacidade de responder perguntas sobre qualquer parte do código
- Feedback dos colegas de equipe (quando solicitado)

### Consequências

- **1ª ocorrência**: Nota zero na atividade
- **2ª ocorrência**: Reprovação na disciplina
- **Casos graves**: Encaminhamento à coordenação

---

## Dúvidas Frequentes

### Sobre os Exercícios

**P: Posso fazer mais de um push?**  
R: Sim! Cada push executa os testes novamente. Você pode iterar até o prazo.

**P: E se os testes não passarem no prazo?**  
R: A nota será proporcional aos testes que passaram no último push antes do prazo.

**P: Posso usar bibliotecas externas?**  
R: Apenas as indicadas no enunciado de cada exercício.

**P: A Lista 4 depende de ter um banco rodando?**  
R: Sim, você precisará de PostgreSQL local ou via Docker para testar.

**P: Posso usar ORM ao invés de sqlc?**  
R: Não nos exercícios. No projeto, o padrão é sqlc, mas aceito justificativa para alternativas.

### Sobre os Sprints

**P: O que acontece se não entregar um sprint?**  
R: Zero no sprint específico. Cada sprint tem peso independente.

**P: Posso reenviar o vídeo se ficou ruim?**  
R: Sim, até o prazo. A última versão será avaliada.

**P: Precisa ter código funcionando no Sprint 0?**  
R: Não. O Sprint 0 é apenas a proposta e planejamento.

**P: Posso mudar o tema do projeto depois?**  
R: Sim, mas deve comunicar ao professor com justificativa.

**P: Como gravar o vídeo?**  
R: OBS Studio, Zoom, Google Meet, ou qualquer ferramenta. O importante é a qualidade do conteúdo.

### Sobre a Prova

**P: Posso consultar documentação durante a prova?**  
R: Apenas a documentação offline do Go (`go doc`). Sem internet.

### Sobre os Grupos

**P: Posso fazer o projeto sozinho?**  
R: Não é recomendado. O mínimo é 3 pessoas.

**P: Posso trocar de grupo depois?**  
R: Apenas em casos excepcionais, com autorização do professor.

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
- **Acompanhamento de projeto**: Última quinta-feira de cada sprint (online via Google Meet)

---

## Checklist do Semestre

### Início do Curso
- [ ] Ambiente configurado (Go, Docker, PostgreSQL)
- [ ] Acesso ao Discord da disciplina
- [ ] Acesso ao GitHub Classroom
- [ ] Grupo formado e registrado (até 12/03)

### Durante o Curso
- [ ] Frequência ≥ 75%
- [ ] Exercícios entregues no prazo
- [ ] Sprints entregues com vídeo
- [ ] Participação em code reviews

### Final do Curso
- [ ] Prova realizada (12/05)
- [ ] Projeto completo e funcionando
- [ ] Apresentação final realizada (16 ou 18/06)
- [ ] Todas as entregas consolidadas (até 26/06)