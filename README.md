# Todoшник

### Запуск проекта

* В корне проекта создать .env file

```
DB_USERNAME=postgres
DB_PASSWORD=qwerty
```

* Запуск сервера:

```
go run cmd/main.go
```

* Запуск бд в docker:

```
docker run --name=todo-app-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
```

* Миграция:

```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

### Архитектура: request -> handler -> service -> repository

* Handlers - обработка запросов, вызов сервисов
* Services - бизнес логика, вызов методов для работы с БД
* Repository - работа с БД


