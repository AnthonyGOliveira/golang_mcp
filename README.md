# Golang mcp server
## Histórico de Alterações
O histórico detalhado de alterações pode ser encontrado no arquivo [CHANGELOG.md](./CHANGELOG.md).

## Estrutura do Projeto
A estrutura atual do projeto está organizada da seguinte forma:
```
projeto/
├── cmd/
│ └── main.go # Ponto de entrada da aplicação
├── internal/
│ └── api/
│ ├── handlers/ # Manipuladores de requisições HTTP
│ │ └── health.go # Handler para verificação de saúde da aplicação
│ │ └── health_test.go # Testes para o handler de saúde
│ └── router/
│ └── router.go # Configuração de rotas da API
├── .gitignore # Arquivos ignorados pelo Git
├── go.mod # Definição do módulo Go
├── Makefile # Automação de tarefas comuns do projeto
└── README.md # Este arquivo
```

## Componentes
### Cmd
O diretório `cmd` contém o ponto de entrada principal da aplicação:
- **main.go**: Inicializa o servidor HTTP e configura a aplicação.

### Internal/API
O diretório `internal` contém código específico da aplicação que não é destinado a ser reutilizado por outras aplicações:

#### Handlers
- **health.go**: Implementa o endpoint de verificação de saúde (`/health`) que retorna o status da aplicação.
- **health_test.go**: Contém testes unitários para validar o funcionamento do HealthCheckHandler.

#### Router
- **router.go**: Define e configura todas as rotas da API, conectando URLs aos manipuladores apropriados.

## Como Executar
Para executar o projeto localmente:

```bash
# Usando go diretamente
go run cmd/main.go

# Ou usando o Makefile
make run
```

O servidor será iniciado e estará ouvindo na porta configurada (padrão: 8000).

## Testes
Para executar os testes do projeto:

```bash
# Executar todos os testes
make test

# Executar testes com cobertura e gerar relatório
make cover
```

## Automação com Makefile
O projeto inclui um Makefile para facilitar tarefas comuns:

- `make run`: Executa a aplicação
- `make test`: Executa todos os testes
- `make cover`: Executa testes com cobertura e gera relatório HTML
- `make cover-test`: Gera arquivo de cobertura de testes
- `make cover-report`: Visualização de cobertura de testes
- `make clean`: Remove arquivos gerados (ex: relatórios de cobertura)
- `make help`: Exibe ajuda sobre os comandos disponíveis

## Endpoints Disponíveis
Atualmente, o projeto conta com os seguintes endpoints:
- **GET /health**: Verifica se a aplicação está funcionando corretamente.

## Desenvolvimento
Para adicionar novos endpoints à aplicação:
1. Crie um novo arquivo de handler em `internal/api/handlers/`
2. Implemente a lógica do handler
3. Registre o novo endpoint no arquivo `internal/api/router/router.go`
4. Adicione testes unitários para validar o comportamento esperado