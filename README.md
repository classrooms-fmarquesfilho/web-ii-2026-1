# DIM0547 - Desenvolvimento de Sistemas Web II

**Universidade Federal do Rio Grande do Norte**  
**Departamento de Informática e Matemática Aplicada**  
**Prof. Fernando Figueira | 2026.1**

---

## Sobre o Curso

Este repositório contém todo o material didático para a disciplina DIM0547 - Desenvolvimento de Sistemas Web II, com foco em **Go** como linguagem principal e práticas modernas de engenharia de software.

### Metodologia: Sala de Aula Invertida 🎬

Este curso utiliza a metodologia de **Sala de Aula Invertida (Flipped Classroom)**:

```
📺 ANTES DA AULA (em casa)
   └── Assista ao vídeo semanal (liberado até o domingo que antecede as aulas presenciais sobre o conteúdo semanal)
   └── Anote dúvidas

🏫 TERÇA-FEIRA (presencial)
   └── Revisão rápida + Esclarecimento de dúvidas
   └── Atividade prática guiada

🏫 QUINTA-FEIRA (presencial ou online quando for acompanhamento de projeto)
   └── Continuação da prática
   └── Exercícios e code review
   └── Acompanhamento de projeto a cada final de sprint (via Google Meet)
```

**Documentos importantes**

- [Plano de curso](docs/PLANO_DE_CURSO.md)
- [Cronograma detalhado](docs/CRONOGRAMA.md)
- [Programação dos vídeos semanais](docs/VIDEOS_SEMANAIS.md)

**Por que essa metodologia?**
- Aprenda no seu ritmo (pause, volte, reveja)
- Mais tempo para prática em sala
- Dúvidas direcionadas com professor presente
- Preparação para autogestão profissional

### Visão Geral

- **Fundamentos**: 3 semanas de `net/http` puro antes de introduzir Chi
- **Go 1.22+**: Aproveitamento das melhorias de roteamento nativo (métodos e parâmetros no ServeMux)
- **Chi como camada**: Router 100% compatível com net/http, preservando conhecimento transferível
- **Exposição a GraphQL**: 1 semana dedicada ao paradigma alternativo
- **Banco de dados idiomático**: [sqlc](https://sqlc.dev) - SQL real com type-safety em tempo de compilação
- **Alinhamento ACM CS2023**: Estrutura baseada em competências, proporção 60/40 prática-teoria
- **Metodologia PBL**: Project-Based Learning com Scrum adaptado ao contexto acadêmico
---

## Informações do Curso

- **Horário**: Terças e Quintas, 13:00 às 14:40 (35T12)
- **Modalidade**: Presencial com Sala de Aula Invertida
- **Carga horária**: 60 horas (15 semanas)
- **Período**: 02/03/2026 a 11/07/2026

### Discord da Disciplina

Link disponível na turma virtual do SIGAA (logo como um dos primeiros tópicos).

Canais principais:
- `#anuncios` - Avisos importantes do professor
- `#duvidas-u1`, `#duvidas-u2`, `#duvidas-u3` - Dúvidas por unidade
- `#exercicios` - Links do GitHub Classroom
- `#videos-semanais` - Links dos vídeos
- `#projeto` - Discussões sobre o projeto integrador
- `#procurando-grupo` - Para formar equipes

### Outros Canais

- **YouTube**: Playlist de vídeos (link no Discord)
- **Google Meet**: Para acompanhamentos online de projeto
- **SIGAA**: Entregas oficiais e notas
- **GitHub**: Repositórios de projeto e exercícios

---

## Estrutura do Repositório

```
.
├── README.md                    # Este arquivo
├── docker-compose.yml           # PostgreSQL para desenvolvimento
├── docs/
│   ├── PLANO_DE_CURSO.md       # Programa completo (15 semanas)
│   ├── CRONOGRAMA.md           # Cronograma detalhado com datas
│   ├── AVALIACAO.md            # Critérios de avaliação
│   ├── AMBIENTE.md             # Configuração do ambiente de desenvolvimento
│   ├── REFERENCIAS.md          # Bibliografia e recursos
│   └── VIDEOS_SEMANAIS.md      # Lista dos vídeos a serem publicados
├── exercicios/
│   ├── lista-01/               # Lista 1 de exercícios
│   ├── lista-02/               # Lista 2 de exercícios
│   └── ...
└── scripts/
    └── check-env/              # Script de verificação do ambiente
        ├── README.md
        └── main.go
```

---

## Pré-requisitos

### Conhecimentos Esperados

1. Programação em alguma linguagem (C, Java, Python)
2. Familiaridade básica com Git
3. Noções de banco de dados relacional

Embora já possuir esses conhecimentos facilite a aprendizagem neste curso, saiba que iremos abordar cada tópico do início, então não se preocupe!

### Curso Básico de Go (Pré-requisito Recomendado)

**Antes de começar este curso**, é altamente recomendado que você complete o **curso básico de Go**, que cobre os fundamentos da linguagem através da metodologia TDD:

📚 **[Aprenda Go com Testes](https://github.com/classrooms-fmarquesfilho/aprenda-go-com-testes)**

**Conteúdo coberto no curso básico:**
- Sintaxe fundamental de Go (variáveis, funções, tipos)
- Test-Driven Development (TDD) com Go
- Estruturas de dados (arrays, slices, maps, structs)
- Ponteiros e tratamento de erros
- Interfaces e polimorfismo
- Testes unitários e de integração

**Por que fazer o curso básico primeiro?**
- Este curso (Web II) assume familiaridade com a sintaxe de Go
- Começamos direto com `net/http` e desenvolvimento web
- O curso básico estabelece fundamentos de TDD que usaremos extensivamente

**Já tem experiência com Go?** Se você já tem experiência com Go, o curso básico pode não ser necessário.

**Não conseguiu fazer o curso básico ainda?** Os vídeos semanais podem incluir revisões rápidas do conteúdo do curso básico para facilitar o entendimento. No entanto, saiba que você pode precisar dedicar tempo extra para estudar aspectos da linguagem Go para acompanhar as aulas de web.

### Ambiente de Desenvolvimento

```bash
# Go 1.25+
go version  # go1.25.0 ou superior

# PostgreSQL 18
psql --version

# Docker
docker --version

# Verificação completa
cd scripts/check-env
go run main.go
```

Instruções detalhadas em [docs/AMBIENTE.md](docs/AMBIENTE.md).

---

## Metodologia

### Project-Based Learning com Scrum

O curso é centrado no desenvolvimento de um **projeto integrador** ao longo do semestre, utilizando Scrum adaptado ao contexto acadêmico:

- **Equipes de 3-4 estudantes** (formação nas primeiras 2 semanas)
- **Sprints de 2 semanas** alinhados com o conteúdo teórico
- **Acompanhamento online** no último dia de cada sprint via Google Meet
- **Avaliação por vídeo** gravado pela equipe + feedback via Discord e/ou em sala de aula (se der tempo)

### Progressão Pedagógica

```
Semanas 1-4:   Fundamentos Go + net/http puro (HTTP, handlers, middleware, validação)
Semana 5:      Chi avançado + OpenAPI
Semanas 6-8:   Persistência (sqlc + Clean Architecture + Testes)
Semanas 9-10:  Autenticação + Segurança
Semana 11:     Prova Prática
Semanas 12-13: CI/CD + Observabilidade
Semana 14:     GraphQL (paradigma alternativo)
Semana 15:     Concorrência
Semanas 16-17: Apresentações finais + Encerramento
```

### Correção Automática de Exercícios

As listas de exercícios utilizam **GitHub Classroom com autograding**, permitindo:

- Feedback imediato a cada push
- Execução automática de testes
- Visualização de resultados no próprio repositório
- Múltiplas tentativas antes do prazo

---

## Avaliação

A nota final é a **média aritmética das 3 unidades**:

```
Nota Final = (U1 + U2 + U3) / 3
```

### Unidades

| Unidade | Período | Componentes |
|---------|---------|-------------|
| **U1** | Semanas 1-5 | Sprint 0 (60%) + Listas 1-3 (40%) |
| **U2** | Semanas 6-11 | **Prova (50%)** + Sprints 1-3 (35%) + Listas 4-5 (15%) |
| **U3** | Semanas 12-17 | **Prova (50%)** + Projeto Final (45%) + Lista 6 (5%) |

A **Prova Prática** (12/05) vale 50% de U2 e 50% de U3, equivalendo a uma unidade completa.

> **Nota**: A participação individual é avaliada dentro dos componentes de grupo (Sprint 0, Sprints 1-3, Projeto Final).

Detalhes completos em [docs/AVALIACAO.md](docs/AVALIACAO.md).

---

## Referências e Materiais de Apoio

Ver referências completas em [docs/REFERENCIAS.md](docs/REFERENCIAS.md).

---

## Licença

Este material é disponibilizado sob a licença [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

Você pode compartilhar e adaptar este material para fins não comerciais, desde que atribua crédito e distribua sob a mesma licença.
