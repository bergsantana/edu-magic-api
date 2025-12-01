# Script de Apresentação - EDU Magic API
## Vídeo de 5 Minutos

---

## [0:00 - 0:30] INTRODUÇÃO (30 segundos)

**[Tela: Logo/Título do Projeto]**

Olá! Hoje vou apresentar o **EDU Magic API**, uma API RESTful desenvolvida em Go para gerenciamento de atividades educacionais.

Este projeto foi construído com foco em **performance**, **arquitetura limpa** e **boas práticas de desenvolvimento**.

---

## [0:30 - 1:30] VISÃO GERAL DO PROJETO (1 minuto)

**[Tela: Estrutura de pastas do projeto]**

O EDU Magic API é uma solução backend que permite:
- ✅ **Autenticação de usuários** com JWT
- ✅ **Criação e gerenciamento de atividades educacionais**
- ✅ **Suporte a caça-palavras** (word search) armazenado em JSON

**Stack Tecnológica:**
- **Linguagem:** Go 1.25
- **Framework Web:** Gin (alta performance)
- **ORM:** GORM
- **Banco de Dados:** PostgreSQL
- **Autenticação:** JWT (JSON Web Tokens)
- **Segurança:** bcrypt para hash de senhas

---

## [1:30 - 2:45] ARQUITETURA HEXAGONAL (1 minuto e 15 segundos)

**[Tela: Diagrama da arquitetura ou estrutura de pastas destacada]**

O projeto segue os princípios da **Arquitetura Hexagonal** (Ports & Adapters), garantindo:

### **1. Camada de Domínio (Core)**
```go
// internal/core/domain/
- User: gerenciamento de usuários
- Activity: atividades educacionais com word search
```

### **2. Camada de Serviços**
```go
// internal/core/services/
- AuthService: autenticação e registro
- ActivityService: lógica de negócio das atividades
```

### **3. Camada de Adaptadores**
- **Handlers:** controladores HTTP com Gin
- **Repositories:** acesso ao banco de dados com GORM

### **4. Infraestrutura**
- Configuração com variáveis de ambiente
- Conexão com PostgreSQL
- Middleware de autenticação JWT

**Benefícios dessa arquitetura:**
- 🔄 Facilita testes unitários
- 🔧 Código desacoplado e manutenível
- 📦 Fácil substituição de componentes

---

## [2:45 - 3:45] FUNCIONALIDADES PRINCIPAIS (1 minuto)

**[Tela: Demonstração das requisições ou código dos endpoints]**

### **Autenticação**
```
POST /auth/signup - Registro de novos usuários
POST /auth/login  - Login com email/senha (retorna JWT)
```

### **Gerenciamento de Atividades**
```
POST   /activities        - Criar nova atividade (protegido)
GET    /activities/:ownerID - Listar atividades do usuário
```

**Recursos de Segurança:**
- 🔐 Senhas criptografadas com bcrypt
- 🎫 Autenticação via JWT Bearer Token
- 🛡️ Middleware de proteção de rotas
- 🌐 CORS configurado para frontend

**Exemplo de Activity:**
```json
{
  "title": "Caça-palavras de História",
  "owner_id": 1,
  "word_search": [[1,2], [3,4], [5,6]]
}
```

---

## [3:45 - 4:30] DIFERENCIAIS TÉCNICOS (45 segundos)

**[Tela: Código ou features específicas]**

### **1. Performance**
- Go é até **10x mais rápido** que Node.js
- Gin é um dos frameworks web mais rápidos
- Baixo consumo de memória

### **2. Boas Práticas**
```go
// Logging customizado
log.Println("🚀 Starting EDU Magic API...")

// Tratamento de erros consistente
if err != nil {
    c.JSON(500, gin.H{"error": "Failed to create activity"})
    return
}
```

### **3. Configuração Flexível**
- Variáveis de ambiente via `.env`
- CORS configurado para múltiplos domínios
- Pronto para deploy em produção

### **4. Armazenamento Avançado**
- JSONB no PostgreSQL para dados complexos
- Suporte nativo a estruturas aninhadas

---

## [4:30 - 5:00] CONCLUSÃO E PRÓXIMOS PASSOS (30 segundos)

**[Tela: Repositório GitHub / Contato]**

O **EDU Magic API** demonstra:
- ✅ Arquitetura profissional e escalável
- ✅ Código limpo e bem organizado
- ✅ Segurança e performance

**Próximos passos:**
- 📝 Adicionar testes automatizados
- 📊 Implementar mais tipos de atividades
- 🚀 Deploy em cloud (AWS/GCP/Heroku)
- 📚 Documentação com Swagger

**Obrigado pela atenção!**

GitHub: github.com/bergsantana/edu-magic-api

---

## DICAS PARA GRAVAÇÃO:

### Visual:
- Mostre a estrutura de pastas no VS Code
- Faça requests com Postman/Insomnia
- Destaque trechos de código importantes
- Use um diagrama simples da arquitetura

### Ritmo:
- Fale de forma clara e pausada
- Não tente explicar cada linha de código
- Foque nos conceitos e benefícios
- Pratique antes para caber no tempo

### Edição:
- Use zoom nos códigos importantes
- Adicione legendas com os pontos-chave
- Música de fundo suave (opcional)
- Transições suaves entre seções

**Tempo total:** 5:00 minutos ✅
