# ex_ms_walletcore

Inicialização da aplicação

go mod init walletcore
Baixar dependencias de pacotes

go mod tidy
executar testes

go test ./...
Criando banco de dados

touch sqlite.db sqlite3 sqlite.db create table products(id string, name string, price float, status string); .tables