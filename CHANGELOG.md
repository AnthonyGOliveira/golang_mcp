# Changelog

Todas as alterações notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/),
e este projeto adere ao [Versionamento Semântico](https://semver.org/lang/pt-BR/).

## [0.1.0] - 2025-05-01

### Adicionado
- Estrutura inicial do projeto
- Configuração básica do servidor Go
- Organização de código seguindo padrões da comunidade Go
- Diretório `cmd` contendo o ponto de entrada da aplicação
- Diretório `internal/api` com separação entre handlers e router
- Endpoint `/health` para verificação de saúde da aplicação
- Arquivo `.gitignore` para arquivos que não devem ser versionados
- Configuração inicial do `go.mod`
- Documentação básica no README.md

### Notas de Implementação
- Adotada a estrutura padrão de projetos Go com separação entre `cmd` e `internal`
- O router está configurado para expansão futura com novos endpoints