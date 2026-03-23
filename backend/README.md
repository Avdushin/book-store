# Bookstore Backend

## Запуск

```bash
make infra     # поднять PostgreSQL (Docker)
make migrate   # применить схему БД
make seed      # заполнить тестовыми данными
make run       # запустить сервер
````

Проверка:

```bash
curl http://localhost:8080/health
```

## API

### Получить все книги

```bash
curl "http://localhost:8080/api/books"
```

### Фильтрация

#### По автору

```bash
curl "http://localhost:8080/api/books?author=Герман%20Гессе"
```

#### По категории

```bash
curl "http://localhost:8080/api/books?category=Философия"
```

#### По году

```bash
curl "http://localhost:8080/api/books?year=1949"
```

### Сортировка

```bash
curl "http://localhost:8080/api/books?sort_by=year&order=desc"
```

Комбинирование:

```bash
curl "http://localhost:8080/api/books?category=Философия&sort_by=year&order=desc"
```

### Получить книгу по ID

```bash
curl "http://localhost:8080/api/books/1"
```

## PostgreSQL

Войти в БД:

```bash
docker exec -it bookstore_postgres psql -U bookstore_user -d bookstore
```

Полезные команды:

```sql
\dt                 -- список таблиц
\d books            -- структура таблицы
SELECT * FROM books LIMIT 5;
```

## Обложки

Доступны по URL:

```text
http://localhost:8080/covers/<filename>
```

Пример:

```bash
http://localhost:8080/covers/master_i_margarita.jpg
```
