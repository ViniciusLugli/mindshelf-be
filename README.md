# Mindshelf BE

Backend da aplicacao Mindshelf, construido em Go com Gin, GORM, PostgreSQL, JWT e WebSocket.

O projeto oferece:

- autenticacao com JWT
- gerenciamento de usuarios
- gerenciamento de grupos
- gerenciamento de tarefas
- amizade entre usuarios
- chat em tempo real via WebSocket
- compartilhamento e importacao de tarefas pelo chat
- documentacao Swagger

## Visao geral

Este repositorio expoe uma API HTTP para operacoes principais da aplicacao e um canal WebSocket autenticado para eventos em tempo real.

Em termos praticos, o fluxo principal e:

1. o usuario se registra ou faz login
2. recebe um token JWT no corpo da resposta
3. o backend tambem grava esse token em cookie HTTP-only
4. com autenticacao ativa, o usuario acessa rotas protegidas em `/api/*`
5. o mesmo usuario pode abrir uma conexao WebSocket em `/api/ws`
6. a partir do socket, pode conversar, compartilhar tarefas e gerenciar amizades

### Links para o repositório

- [Github]("https://github.com/ViniciusLugli/mindshelf-be")
- [Bitbucket]("https://bitbucket.org/viniciuslugli/mindshelf-be/src/main/")

## Funcionalidades

### HTTP / REST

- cadastro e login
- consulta e atualizacao do usuario autenticado
- listagem e consulta de usuarios
- CRUD de grupos
- CRUD de tarefas
- importacao de tarefas compartilhadas via chat

### WebSocket

- envio de mensagens em tempo real
- listagem de conversas
- listagem de chats recentes
- marcacao de mensagens como lidas
- envio, aceite, rejeicao e remocao de amizades
- compartilhamento de tarefa com snapshot dos dados

## Stack usada

- Go
- Gin
- GORM
- PostgreSQL
- golang-migrate
- JWT
- Gorilla WebSocket
- Swagger / Swaggo

## Estrutura do projeto

```text
.
|- cmd/
|  |- api/            # inicializacao do servidor HTTP/WebSocket
|  |- up_migrate/     # executa migracoes para cima
|  `- down_migrate/   # rollback de migracoes
|- docs/              # arquivos gerados do Swagger
|- internal/
|  |- dtos/           # contratos de entrada e saida
|  |- handlers/       # camada HTTP e WebSocket
|  |- middlewares/    # autenticacao, recovery, logging, request id
|  |- models/         # entidades do dominio
|  |- repositories/   # acesso a dados e migracoes
|  |- services/       # regras de negocio
|  `- utils/          # JWT, logger e utilitarios de WS
|- migrations/        # arquivos SQL versionados
|- Dockerfile
`- docker-compose.yml
```

## Arquitetura em camadas

- `handlers`: recebem requests, validam entrada, retornam status HTTP e respostas JSON
- `services`: concentram a regra de negocio
- `repositories`: falam com o banco usando GORM ou migracoes SQL
- `models`: representam as entidades persistidas
- `dtos`: padronizam payloads de request e response

Essa separacao facilita manutencao, testes e evolucao do projeto.

## Requisitos

- Go `1.26.1` ou compativel com o `go.mod`
- PostgreSQL
- variaveis de ambiente configuradas

Observacoes importantes:

- o `go.mod` declara `Go 1.26.1`
- o `Dockerfile` atual usa imagem `golang:1.24-alpine`; se houver erro de compatibilidade, alinhe a versao da imagem com o `go.mod`
- a aplicacao roda migracoes no startup

## Variaveis de ambiente

Exemplo recomendado:

```env
DATABASE_URL="postgres://my_user:my_password@localhost:5433/my_db?sslmode=disable"
DSN="host=localhost user=my_user password=my_password dbname=my_db port=5433 sslmode=disable"
JWT_SECRET="change-me"
LOG_LEVEL="debug"
ALLOWED_ORIGINS="http://localhost:3000"
PORT="8080"
COOKIE_SECURE="false"
```

### O que cada variavel faz

| Variavel          | Obrigatoria | Descricao                                                                    |
| ----------------- | ----------- | ---------------------------------------------------------------------------- |
| `DATABASE_URL`    | Sim         | URL usada nas migracoes e tambem como opcao principal de conexao com o banco |
| `DSN`             | Opcional    | String alternativa de conexao usada como fallback em `ConnectDB()`           |
| `JWT_SECRET`      | Sim         | Segredo para gerar e validar tokens JWT                                      |
| `LOG_LEVEL`       | Nao         | Nivel de log: `debug`, `info`, `warn`, `error`                               |
| `ALLOWED_ORIGINS` | Nao         | Lista separada por virgula de origens aceitas no upgrade WebSocket           |
| `PORT`            | Nao         | Porta do servidor HTTP; padrao `8080`                                        |
| `COOKIE_SECURE`   | Nao         | Se `true`, marca o cookie de autenticacao como seguro                        |

Importante:

- para subir a API sem surpresa, configure `DATABASE_URL`
- `JWT_SECRET` e obrigatorio; sem ele, login e registro falham

## Como rodar localmente

### 1. Suba o banco de dados

O repositorio ja possui um `docker-compose.yml` simples para PostgreSQL:

```bash
docker compose up -d
```

Isso sobe um PostgreSQL com:

- banco: `my_db`
- usuario: `my_user`
- senha: `my_password`
- porta externa: `5433`

### 2. Configure o `.env`

Se quiser, use como base o arquivo `.env` existente e ajuste os valores para seu ambiente.

### 3. Baixe as dependencias

```bash
go mod download
```

### 4. Inicie a API

```bash
go run ./cmd/api
```

Servidor padrao:

- HTTP: `http://localhost:8080`
- Swagger: `http://localhost:8080/swagger/index.html`
- WebSocket autenticado: `ws://localhost:8080/api/ws`

## Rodando com Docker

Para build da aplicacao:

```bash
docker build -t mindshelf-be .
```

Depois execute a imagem com as variaveis de ambiente necessarias:

```bash
docker run --rm -p 8080:8080 --env-file .env mindshelf-be
```

## Migracoes

As migracoes ficam em `migrations/` e sao executadas com `golang-migrate`.

### Subir migracoes

```bash
go run ./cmd/up_migrate
```

### Reverter 1 migracao

```bash
go run ./cmd/down_migrate 1
```

### Reverter todas

```bash
go run ./cmd/down_migrate all
```

### Forcar uma versao de migracao

```bash
go run ./cmd/down_migrate force 9
```

### Limpar completamente o schema publico

```bash
go run ./cmd/down_migrate wipe yes
```

Atencao:

- `wipe yes` apaga todos os dados do schema `public`
- a aplicacao chama `RunUpMigrations()` no startup da API

## Autenticacao

O projeto aceita autenticacao de duas formas nas rotas protegidas:

- header `Authorization: Bearer <token>`
- cookie `mindshelf_token`

Quando o usuario faz registro ou login:

- o token JWT volta no JSON da resposta
- o backend tambem grava um cookie HTTP-only com esse token

## Endpoints principais

As rotas abaixo representam o uso principal da API.

### Publicas

| Metodo | Rota                  | Descricao                         |
| ------ | --------------------- | --------------------------------- |
| `POST` | `/register`           | Cria usuario e retorna token      |
| `POST` | `/login`              | Autentica usuario e retorna token |
| `GET`  | `/swagger`            | Redireciona para a UI do Swagger  |
| `GET`  | `/swagger/index.html` | Documentacao interativa           |

### Usuarios

| Metodo   | Rota             | Descricao                     |
| -------- | ---------------- | ----------------------------- |
| `GET`    | `/api/users`     | Lista usuarios com paginacao  |
| `GET`    | `/api/users/me`  | Retorna o usuario autenticado |
| `GET`    | `/api/users/:id` | Busca usuario por ID          |
| `PATCH`  | `/api/users/me`  | Atualiza o proprio usuario    |
| `DELETE` | `/api/users/me`  | Remove o proprio usuario      |

### Grupos

| Metodo   | Rota              | Descricao                           |
| -------- | ----------------- | ----------------------------------- |
| `GET`    | `/api/groups`     | Lista grupos do usuario autenticado |
| `GET`    | `/api/groups/:id` | Busca grupo por ID                  |
| `POST`   | `/api/groups`     | Cria grupo                          |
| `PATCH`  | `/api/groups/:id` | Atualiza grupo                      |
| `DELETE` | `/api/groups/:id` | Remove grupo                        |

### Tarefas

| Metodo   | Rota             | Descricao                             |
| -------- | ---------------- | ------------------------------------- |
| `GET`    | `/api/tasks`     | Lista tarefas com filtros e paginacao |
| `GET`    | `/api/tasks/:id` | Busca tarefa por ID                   |
| `POST`   | `/api/tasks`     | Cria tarefa                           |
| `PATCH`  | `/api/tasks/:id` | Atualiza tarefa                       |
| `DELETE` | `/api/tasks/:id` | Remove tarefa                         |

### Tarefas compartilhadas

| Metodo | Rota                       | Descricao                             |
| ------ | -------------------------- | ------------------------------------- |
| `POST` | `/api/shared-tasks/import` | Importa uma tarefa recebida pelo chat |

## Parametros comuns de listagem

As rotas de listagem usam pagina e limite.

Exemplos:

- `GET /api/users?page=1&limit=10`
- `GET /api/groups?page=1&limit=10&name=work`
- `GET /api/tasks?page=1&limit=20&title=deploy`
- `GET /api/tasks?page=1&limit=20&group_id=<uuid>`

## Exemplos de uso HTTP

### Registrar usuario

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Joao",
    "email": "joao@example.com",
    "password": "123456"
  }'
```

Resposta esperada:

```json
{
  "token": "<jwt>",
  "user": {
    "id": "<uuid>",
    "name": "Joao",
    "email": "joao@example.com",
    "avatar_url": ""
  }
}
```

### Criar grupo

```bash
curl -X POST http://localhost:8080/api/groups \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Trabalho",
    "color": "#2563eb"
  }'
```

### Criar tarefa

```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Preparar apresentacao",
    "notes": "Finalizar ate sexta",
    "group_id": "<group-uuid>"
  }'
```

## WebSocket

### Endpoint

Conexao autenticada:

```text
GET /api/ws
```

O WebSocket usa a mesma autenticacao das rotas HTTP protegidas.

### Formato das mensagens

Cliente -> servidor:

```json
{
  "action": "send_message",
  "payload": {
    "to_user_id": "<uuid>",
    "content": "Ola!"
  }
}
```

Servidor -> cliente:

```json
{
  "action": "message_sent",
  "success": true,
  "data": {},
  "error": ""
}
```

### Acoes disponiveis

| Action                        | Objetivo                            | Resposta principal                                                  |
| ----------------------------- | ----------------------------------- | ------------------------------------------------------------------- |
| `send_message`                | Envia mensagem para outro usuario   | `message_sent` para o remetente e `message_received` para o destino |
| `share_task`                  | Compartilha snapshot de uma tarefa  | `message_sent` / `message_received`                                 |
| `get_conversation`            | Retorna historico com outro usuario | `get_conversation`                                                  |
| `get_chats`                   | Lista chats recentes                | `get_chats`                                                         |
| `mark_messages_read`          | Marca mensagens como lidas          | `mark_messages_read` e possivel broadcast `messages_read`           |
| `send_friend_request`         | Envia pedido de amizade             | `send_friend_request`                                               |
| `accept_friend_request`       | Aceita amizade pendente             | `accept_friend_request`                                             |
| `reject_friend_request`       | Rejeita amizade pendente            | `reject_friend_request`                                             |
| `remove_friend`               | Remove amizade                      | `remove_friend`                                                     |
| `get_friends`                 | Lista amigos                        | `get_friends`                                                       |
| `get_pending_friend_requests` | Lista convites pendentes recebidos  | `get_pending_friend_requests`                                       |

### Exemplo: enviar mensagem

Payload enviado:

```json
{
  "action": "send_message",
  "payload": {
    "to_user_id": "1f7c2a36-0000-0000-0000-000000000000",
    "content": "Oi, tudo bem?"
  }
}
```

Eventos recebidos:

- remetente: `message_sent`
- destinatario: `message_received`

### Exemplo: compartilhar tarefa

Payload enviado:

```json
{
  "action": "share_task",
  "payload": {
    "to_user_id": "1f7c2a36-0000-0000-0000-000000000000",
    "task_id": "2c7f4e55-0000-0000-0000-000000000000",
    "content": "Da uma olhada nessa tarefa"
  }
}
```

Quando a mensagem e do tipo `shared_task`, a resposta inclui um snapshot com:

- tarefa original
- titulo
- notas
- nome do grupo
- cor do grupo
- `imported_task_id` quando a tarefa ja foi importada

## Swagger

Documentacao interativa:

- `http://localhost:8080/swagger`
- `http://localhost:8080/swagger/index.html`

Se voce alterar anotacoes dos handlers e precisar regenerar os arquivos do Swagger, um comando comum e:

```bash
swag init -g cmd/api/main.go -o docs
```

## Testes

Para rodar todos os testes:

```bash
go test ./...
```

Para verificar cobertura:

```bash
go test ./... -cover
```

O pipeline do repositorio tambem executa:

- `go fmt ./...`
- validacao de arquivos nao formatados com `gofmt`
- `go vet ./...`
- `go test ./... -coverprofile=coverage.out`
- build do projeto

## Build de producao

Para gerar o binario localmente:

```bash
go build -o server ./cmd/api
```

## Fluxo sugerido para desenvolvimento

1. suba o PostgreSQL com `docker compose up -d`
2. configure `DATABASE_URL` e `JWT_SECRET`
3. rode `go run ./cmd/api`
4. teste as rotas no Swagger
5. conecte o frontend ou um cliente WebSocket em `/api/ws`
6. rode `go test ./...` antes de finalizar alteracoes

## Troubleshooting

### Erro de banco ao iniciar

Verifique:

- se o PostgreSQL esta rodando
- se `DATABASE_URL` aponta para a porta correta
- se usuario, senha e nome do banco estao corretos

### Erro de JWT

Se aparecer erro relacionado a token ou configuracao JWT:

- confira se `JWT_SECRET` foi definido
- confira se o token enviado nao expirou nem foi alterado

### Erro 401 nas rotas protegidas

Confirme se voce esta enviando:

- `Authorization: Bearer <token>`
- ou o cookie `mindshelf_token`

### Falha no WebSocket por origem bloqueada

Se o upgrade falhar no navegador, confira `ALLOWED_ORIGINS`.

Exemplo:

```env
ALLOWED_ORIGINS="http://localhost:3000,http://127.0.0.1:3000"
```

### Observacao sobre migracoes no startup

O servidor executa migracoes automaticamente ao iniciar. Se voce encontrar problema relacionado a estado de migracao, valide a tabela de versoes do banco e, se necessario, use os comandos de `down_migrate` com cuidado.

## Resumo rapido

- backend em Go para usuarios, grupos, tarefas, amizade e chat
- autenticacao por JWT via header ou cookie
- WebSocket autenticado para eventos em tempo real
- PostgreSQL com migracoes SQL versionadas
- Swagger para exploracao e testes da API
