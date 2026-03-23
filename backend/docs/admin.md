Сначала залогинься под admin. В `admin@bookstore.com / admin123`

### Логин admin

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email":"admin@bookstore.com",
    "password":"admin123"
  }'
```

### Создать книгу

```bash
curl -X POST http://localhost:8080/api/admin/books \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -d '{
    "title":"Игра в бисер",
    "description":"Философский роман Германа Гессе",
    "author_id":11,
    "category_id":5,
    "year_written":1943,
    "purchase_price":16.99,
    "rent_price_2_weeks":3.99,
    "rent_price_1_month":6.99,
    "rent_price_3_months":12.99,
    "status":"available",
    "is_available":true,
    "cover_url":"/covers/glass_bead_game.jpg"
  }'
```

### Обновить книгу

```bash
curl -X PUT http://localhost:8080/api/admin/books/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -d '{
    "title":"Мастер и Маргарита",
    "description":"Обновлённое описание",
    "author_id":1,
    "category_id":1,
    "year_written":1967,
    "purchase_price":18.99,
    "rent_price_2_weeks":4.49,
    "rent_price_1_month":7.49,
    "rent_price_3_months":13.49,
    "status":"available",
    "is_available":true,
    "cover_url":"/covers/master_i_margarita.jpg"
  }'
```

### Изменить статус

```bash
curl -X PATCH http://localhost:8080/api/admin/books/1/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -d '{
    "status":"inactive"
  }'
```

### Изменить доступность

```bash
curl -X PATCH http://localhost:8080/api/admin/books/1/availability \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -d '{
    "is_available":false
  }'
```

### Удалить книгу

```bash
curl -X DELETE http://localhost:8080/api/admin/books/16 \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

### Проверка запрета для обычного user

```bash
curl -X POST http://localhost:8080/api/admin/books \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer USER_TOKEN" \
  -d '{
    "title":"Test",
    "description":"Test",
    "author_id":1,
    "category_id":1,
    "year_written":2024,
    "purchase_price":10,
    "rent_price_2_weeks":2,
    "rent_price_1_month":4,
    "rent_price_3_months":8,
    "status":"available",
    "is_available":true,
    "cover_url":"/covers/test.jpg"
  }'
```