# GoLangApi
Executar:   go run main.go

Criar instância do SqlServer:
docker run --name SqlServerContainer -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=goLang123" -p 1433:1433 -d mcr.microsoft.com/mssql/server:latest 

Conectar com o Db:
docker exec -it SqlServerContainer /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P goLang123