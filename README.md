# Book Store

Web-приложение книжного магазина с пользовательским и административным интерфейсом.

## Demo
[Скриншоты демо версии](./demo.md)

![demo](./docs/imgs/catalog.jpeg)

## Возможности

### Пользователь:
- Просмотр каталога книг
- Фильтрация:
  - по категории
  - по автору
  - по году
- Сортировка
- Покупка книг
- Аренда книг:
  - 2 недели
  - 1 месяц
  - 3 месяца
- Просмотр своих покупок и аренд

### Администратор:
- Добавление книг
- Редактирование книг
- Удаление книг
- Управление:
  - ценами
  - статусами
  - доступностью
- Добавление авторов и категорий

### Система:
- JWT авторизация
- Роли (user / admin)
- Напоминания об окончании аренды (scheduler)


## Стек

- Backend: Go (chi, PostgreSQL, net/http)
- Frontend: React + Vite
- БД: PostgreSQL
- Docker


## Запуск проекта

### 1. Инфраструктура
```bash
docker-compose -f infra/docker-compose.yml up -d
````

### 2. Backend

```bash
cd backend
make migrate
make seed
make run
```

### 3. Frontend

```bash
cd frontend
pnpm install
pnpm dev
```

Frontend: [http://localhost:5173](http://localhost:5173)
Backend: [http://localhost:8080](http://localhost:8080)


## 🔑 Тестовые пользователи

### Админ

```
email: admin@bookstore.com
password: admin123
```

### Пользователь

```
email: user@bookstore.com
password: user123
```

## API handlers

### Auth

* POST `/api/auth/register`
* POST `/api/auth/login`
* GET `/api/auth/me`

### Books

* GET `/api/books`
* GET `/api/books/{id}`

### Purchases

* POST `/api/purchases`
* GET `/api/purchases/my`

### Rentals

* POST `/api/rentals`
* GET `/api/rentals/my`

### Admin

* POST `/api/admin/books`
* PUT `/api/admin/books/{id}`
* DELETE `/api/admin/books/{id}`

## Дополнительно

* Реализован scheduler напоминаний
* Используется чистая архитектура (repository/service/handler)
* Поддержка ролей и middleware

[Анализ по выполненной задаче](./analys.md) \
[Рекомендации по устранению выявленных ошибок в ходе выполнения задачи](./to_enchance.md)