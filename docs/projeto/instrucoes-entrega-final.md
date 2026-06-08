## WEB II (DIM0547) — Entrega Final

**Título da tarefa**: Entrega Final — MVP completo + Auth/Segurança/Testes (Sprint 3) + Vídeo
**Prazo**: 23/06/2026 (terça) às 23:59

**Descrição**:
Entrega que encerra o projeto. A API deve estar completa: login, rotas protegidas, resiliente a ataques comuns, com pipeline CI verde, e o vídeo deve demonstrar o MVP e as decisões técnicas.

**O que entregar via tarefa do SIGAA:**
1. Link do repositório no GitHub (o mesmo das sprints anteriores, agora evoluído)
2. Link do vídeo de **12 minutos** (YouTube/Drive acessível)

**O que precisa estar pronto até o prazo:**
- Endpoint de login retornando JWT
- Rotas protegidas por middleware de autenticação/autorização
- Refresh token funcionando (com rotação)
- Pelo menos **2 correções de segurança OWASP** implementadas (ex.: rate limiting, validação/sanitização de entrada, cabeçalhos de segurança, controle de acesso por objeto)
- Pelo menos **10 testes automatizados** rodando no pipeline CI (verde)
- MVP completo: CRUD persistido em PostgreSQL (sqlc) e ao menos um **relacionamento 1:N** com endpoint de dados aninhados (das sprints anteriores), com tratamento de erros
- *(Opcional — bônus)* pelo menos 1 endpoint **gRPC OU GraphQL**
- Vídeo 12 min mostrando: introdução (projeto/equipe/problema); demonstração do MVP com fluxos completos, tratamento de erros e JWT; arquitetura e código (decisões técnicas e endpoints documentados); conclusão (aprendizados e o que faria diferente)

**Avaliação** — a mesma entrega vale em duas unidades:

*Como Sprint 3 (30% da U2):*
- Funcionalidade implementada (30%)
- Qualidade do código (25%)
- Testes automatizados (20%)
- Demonstração funcional (15%)
- Comunicação técnica (10%)

*Como Vídeo Final (50% da U3):*
- Demonstração do MVP (45%)
- Qualidade técnica — arquitetura e boas práticas (25%)
- Comunicação (15%)
- Organização — tempo, estrutura, edição (10%)
- Reflexão (5%)

> Lembrete: a **Lista 5+6** (auth JWT + autorização + refresh + rate limiting + testes) é entrega **individual** e **separada**, com prazo em **19/06 (sex) às 23:59**. JWT, testes e segurança **não** caem na prova de 09/06, mas são cobrados aqui (projeto) e na Lista 5+6.
