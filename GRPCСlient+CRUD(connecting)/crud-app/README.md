# CRUD Приложение для Управления Списком Книг
## Пример решения тестового задания для практического проекта #1 курса GOLANG NINJA

### Стэк
- go 1.17
- postgres

### Запуск
Для запуска необходимо указать переменные окружения, например в файле .env

```
$env:DB_HOST="localhost"
$env:DB_PORT="5434"
$env:DB_USERNAME="postgres"
$env:DB_NAME="postgres"
$env:DB_SSLMODE="disable"
$env:DB_PASSWORD="goLANGn1nja"
go run cmd/main.go
```

Сборка и запуск
```
source .env && go build -o app cmd/main.go && ./app
```

Для postgres можно использовать Docker

```
docker run -d --name ninja-db -e POSTGRES_PASSWORD=qwerty123 -v ${HOME}/pgdata/:/var/lib/postgresql/data -p 5432:5432 postgres
```
