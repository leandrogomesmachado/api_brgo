# API BRGO

API CRUD simples escrita em BRGO (Go com palavras-chave em português do Brasil).

## Estrutura

- `principal.go`: Arquivo principal que inicia o servidor
- `modelos/`: Contém as definições de estruturas para usuários, clientes e produtos
- `armazenamento/`: Implementação do armazenamento em arquivos locais
- `api/`: Handlers da API REST

## Endpoints

### Usuários
- `GET /api/usuarios`: Lista todos os usuários
- `POST /api/usuarios`: Cria um novo usuário
- `GET /api/usuarios/?id=<id>`: Obtém um usuário específico
- `PUT /api/usuarios/?id=<id>`: Atualiza um usuário existente
- `DELETE /api/usuarios/?id=<id>`: Remove um usuário

### Clientes
- `GET /api/clientes`: Lista todos os clientes
- `POST /api/clientes`: Cria um novo cliente
- `GET /api/clientes/?id=<id>`: Obtém um cliente específico
- `PUT /api/clientes/?id=<id>`: Atualiza um cliente existente
- `DELETE /api/clientes/?id=<id>`: Remove um cliente

### Produtos
- `GET /api/produtos`: Lista todos os produtos
- `POST /api/produtos`: Cria um novo produto
- `GET /api/produtos/?id=<id>`: Obtém um produto específico
- `PUT /api/produtos/?id=<id>`: Atualiza um produto existente
- `DELETE /api/produtos/?id=<id>`: Remove um produto
