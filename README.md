# API em Go com Autenticação JWT usando brgo (Go em Português)

Uma API RESTful simples construída com Go em sintaxe portuguesa (brgo), framework Gin, SQLite e autenticação JWT.

## Características

- Autenticação de usuário com JWT
- Operações CRUD para usuários, produtos e entregas
- Banco de dados SQLite para persistência de dados
- Armazenamento seguro de senhas com bcrypt
- Sintaxe em português (brgo) para todas as palavras-chave Go

## Requisitos

- Go 1.19 ou superior
- SQLite3
- Wrapper brgo (para converter a sintaxe em português para Go)

## Configuração

1. Clone o repositório
2. Navegue até o diretório do projeto
3. Compile os arquivos brgo para arquivos Go:
   ```bash
   # Script para converter arquivos .brgo para .go
   # (Ferramenta brgo necessária)
   ```
4. Execute a aplicação:
   ```bash
   go run main.go
   ```

A API estará disponível em `http://localhost:8080`.

## Palavras-chave brgo

O projeto usa as seguintes conversões de palavras-chave de português para Go:

| Português | Go       |
|-----------|----------|
| pacote    | package  |
| importa   | import   |
| tipo      | type     |
| estrutura | struct   |
| interface | interface|
| mapa      | map      |
| func      | func     |
| retorna   | return   |
| se        | if       |
| senao     | else     |
| para      | for      |
| vá        | go       |
| adia      | defer    |
| selecione | select   |
| caso      | case     |
| continua  | continue |
| quebra    | break    |
| padrão    | default  |
| nulo      | nil      |
| verdadeiro| true     |
| falso     | false    |
| var       | var      |
| const     | const    |
| anexa     | append   |

## Endpoints da API

### Autenticação

- `POST /api/register`: Registrar um novo usuário
  ```json
  {
    "username": "usuario1",
    "password": "senha123",
    "email": "usuario@exemplo.com"
  }
  ```

- `POST /api/login`: Login e obter token JWT
  ```json
  {
    "username": "usuario1",
    "password": "senha123"
  }
  ```

### Usuários (`/api/usuarios`)

- `GET /api/usuarios`: Obter todos os usuários
- `GET /api/usuarios/:id`: Obter usuário por ID
- `PUT /api/usuarios/:id`: Atualizar usuário
  ```json
  {
    "email": "novoemail@exemplo.com",
    "password": "novasenha"
  }
  ```
- `DELETE /api/usuarios/:id`: Excluir usuário

### Produtos (`/api/produtos`)

- `GET /api/produtos`: Obter todos os produtos
- `GET /api/produtos/:id`: Obter produto por ID
- `POST /api/produtos`: Criar produto
  ```json
  {
    "name": "Nome do Produto",
    "description": "Descrição do Produto",
    "price": 29.99,
    "stock": 100
  }
  ```
- `PUT /api/produtos/:id`: Atualizar produto
- `DELETE /api/produtos/:id`: Excluir produto

### Entregas (`/api/entregas`)

- `GET /api/entregas`: Obter todas as entregas do usuário
- `GET /api/entregas/:id`: Obter entrega por ID
- `POST /api/entregas`: Criar entrega
  ```json
  {
    "product_id": 1,
    "quantity": 2,
    "address": "Rua Principal 123",
    "delivery_at": "2023-06-30T15:00:00Z"
  }
  ```
- `PUT /api/entregas/:id`: Atualizar status da entrega
  ```json
  {
    "status": "shipped"
  }
  ```
- `DELETE /api/entregas/:id`: Excluir entrega

## Autenticação

Todos os endpoints (exceto registro e login) requerem autenticação JWT. Inclua o token JWT no cabeçalho Authorization:

```
Authorization: Bearer <seu_token_jwt>
```

## Notas de Segurança

Para ambientes de produção:
- Altere a chave secreta JWT em middleware/auth.brgo
- Configure HTTPS
- Adicione limitação de taxa
- Implemente validação e tratamento de erros mais abrangentes
