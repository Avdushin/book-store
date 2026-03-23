## 🔔 Модуль уведомлений (Reminder Scheduler)

Фоновый scheduler раз в минуту проверяет таблицу `rentals`:

* если аренда скоро истекает (≤ 24 часа) → создаёт `rent_expiring`
* если аренда уже истекла → помечает её как `expired` и создаёт `rent_expired`

## Проверка

```bash
# уведомления
docker exec -it bookstore_postgres psql -U bookstore_user -d bookstore -c "SELECT * FROM notifications ORDER BY id;"

# аренды
docker exec -it bookstore_postgres psql -U bookstore_user -d bookstore -c "SELECT id, status, end_date FROM rentals;"
```

## Быстрый тест

```bash
# сделать аренду просроченной
docker exec -it bookstore_postgres psql -U bookstore_user -d bookstore -c "
UPDATE rentals SET status='active', end_date=NOW() - INTERVAL '1 hour' WHERE id=3;
"
```

Подожди ~1 минуту → появится `rent_expired` 👍
