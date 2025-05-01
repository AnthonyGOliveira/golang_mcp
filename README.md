# Nome do Projeto

## Sobre o Projeto

Breve descrição do seu projeto e seu propósito principal.

## Histórico de Alterações

O histórico detalhado de alterações pode ser encontrado no arquivo [CHANGELOG.md](./CHANGELOG.md).

## Estrutura do Projeto

A estrutura atual do projeto está organizada da seguinte forma:

```
projeto/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   └── api/
│       ├── handlers/           # Manipuladores de requisições HTTP
│       │   └── health.go       # Handler para verificação de saúde da aplicação
│       └── router/
│           └── router.go       # Configuração de rotas da API
├── .gitignore                  # Arquivos ignorados pelo Git
├── go.mod                      # Definição do módulo Go
└── README.md                   # Este arquivo
```

## Componentes

### Cmd

O diretório `cmd` contém o ponto de entrada principal da aplicação:

- **main.go**: Inicializa o servidor HTTP e configura a aplicação.

### Internal/API

O diretório `internal` contém código específico da aplicação que não é destinado a ser reutilizado por outras aplicações:

#### Handlers

- **health.go**: Implementa o endpoint de verificação de saúde (`/health`) que retorna o status da aplicação.

#### Router

- **router.go**: Define e configura todas as rotas da API, conectando URLs aos manipuladores apropriados.

## Como Executar

Para executar o projeto localmente:

```bash
go run cmd/main.go
```

O servidor será iniciado e estará ouvindo na porta configurada (padrão: 8000).

## Endpoints Disponíveis

Atualmente, o projeto conta com os seguintes endpoints:

- **GET /health**: Verifica se a aplicação está funcionando corretamente.

## Desenvolvimento

Para adicionar novos endpoints à aplicação:

1. Crie um novo arquivo de handler em `internal/api/handlers/`
2. Implemente a lógica do handler
3. Registre o novo endpoint no arquivo `internal/api/router/router.go`