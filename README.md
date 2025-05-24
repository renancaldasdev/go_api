# 🛠️ API de Cadastro de Produtos - Go + Docker + PostgreSQL

Esta é uma API REST simples para cadastro e consulta de produtos, desenvolvida em Go e containerizada com Docker. O banco de dados utilizado é o PostgreSQL.

---

## 📦 Tecnologias Utilizadas

- [Go 1.24](https://golang.org/doc/go1.24)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [PostgreSQL 12](https://www.postgresql.org/)
- Docker e Docker Compose

---

## 📁 Estrutura do Projeto

```
├── cmd/
│   └── main.go        # Arquivo principal que inicializa a aplicação
├── controller/        # Controladores da API
├── model/             # Structs e modelos de domínio
├── repository/        # Lógica de persistência
├── usecase/           # Lógica de negócio (serviços)
├── go.mod / go.sum    # Dependências do projeto Go
├── Dockerfile         # Dockerfile da aplicação Go
├── docker-compose.yml # Orquestração com Docker Compose
└── .gitignore
```

---

## 🚀 Como Rodar o Projeto

### Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Passo a passo

1. Clone este repositório:
   ```bash
   git clone https://github.com/renancaldasdev/go_api
   cd go_api
   ```

2. Suba os containers:
   ```bash
   docker-compose up --build
   ```

3. Acesse a API:
   - A API estará disponível em: [http://localhost:8000](http://localhost:8000)

---

## 🔌 Endpoints da API

### `GET /products`
Retorna a lista de produtos cadastrados.

### `GET /product/:id`
Retorna um produto específico por ID.

### `POST /product`
Cria um novo produto.

**Exemplo de payload:**
```json
{
  "name": "Batata Frita",
  "price": 20.0
}
```

---

## 🗃️ Banco de Dados

A API utiliza PostgreSQL 12. Os dados são persistidos no volume `pgdata`, definido no `docker-compose.yml`.

- Host: `go_db`
- Porta: `5432`
- Usuário: `postgres`
- Senha: `1234`
- Banco: `postgres`

---

## 🧼 Parar e limpar containers

```bash
docker-compose down -v
```

---

## 📄 Licença

Este projeto está sob a licença MIT. Sinta-se à vontade para usá-lo como base para seus próprios projetos.
