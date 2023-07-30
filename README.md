# Wallet Microservice
Projeto utilizado no módulo de "Arquitetura de Microserviços" e "Arquitetura Baseada em Eventos", do curso Fullcyle 3.0

### Iniciando projeto em GO
go mod init github.com/elieudomaia/arquitetura-microservicos

### Instalando os pacotes do qual o projeto é dependente
go mod tidy

### Rodar todos os testes de todas as pastas
go test ./...

### Instando pacote do sqlite
go get github.com/mattn/go-sqlite3