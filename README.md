# Sorting Service

Сервис предоставляет HTTP API для добавления чисел и получения отсортированного списка.

## Сервис

* Принимает число через HTTP запрос
* Сохраняет его в PostgreSQL
* Выводит отсортированный список чисел
* Автоматически запускает миграции при старте `docker-compose`



## API

| Метод | URL       | Описание        |
| ----- | --------- | --------------- |
| POST  | /numbers  | Добавить число  |

Пример POST-запроса:

```json
{
  "value": 42
}
```

## Docker и запуск

Сервис запускается одной командой:

```bash
docker-compose up --build
```

После запуска сервис доступен на:

```
http://localhost:8080
```

### Переменные окружения (`.env`)

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_PORT=5432

PORT=8080
DATABASE_DSN=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable
```

Пример находится в `.env.example`.

---

## Миграции

Миграция находится в `migrations/init.sql` и выполняется автоматически отдельным сервисом `migrations`:

```sql
CREATE TABLE IF NOT EXISTS numbers (
    id SERIAL PRIMARY KEY,
    value INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

---

## Тестирование

Сервис покрыт unit-тестами.

### Запуск тестов

```bash
go test ./... -v
```

---

## Используемые технологии

* Go 1.24.4
* PostgreSQL 15
* Docker + docker-compose

