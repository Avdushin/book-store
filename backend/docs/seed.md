```bash
➜  backend git:(main) ✗ make seed
docker exec -i bookstore_postgres psql -U bookstore_user -d bookstore < migrations/seed.sql
TRUNCATE TABLE
NOTICE:  truncate cascades to table "notifications"
TRUNCATE TABLE
TRUNCATE TABLE
NOTICE:  truncate cascades to table "rentals"
NOTICE:  truncate cascades to table "purchases"
NOTICE:  truncate cascades to table "notifications"
TRUNCATE TABLE
NOTICE:  truncate cascades to table "books"
NOTICE:  truncate cascades to table "rentals"
NOTICE:  truncate cascades to table "purchases"
NOTICE:  truncate cascades to table "notifications"
TRUNCATE TABLE
NOTICE:  truncate cascades to table "books"
NOTICE:  truncate cascades to table "rentals"
NOTICE:  truncate cascades to table "purchases"
NOTICE:  truncate cascades to table "notifications"
TRUNCATE TABLE
NOTICE:  truncate cascades to table "rentals"
NOTICE:  truncate cascades to table "purchases"
NOTICE:  truncate cascades to table "notifications"
TRUNCATE TABLE
INSERT 0 6
INSERT 0 14
INSERT 0 2
INSERT 0 16
INSERT 0 2
INSERT 0 3
INSERT 0 1
➜  backend git:(main) ✗ curl -X POST http://localhost:8080/api/auth/login \
                              -H "Content-Type: application/json" \
                              -d '{
                            "email":"admin@bookstore.com",
                            "password":"admin123"
                          }'
{"token":"MTphZG1pbjpzdXBlci1zZWNyZXQtZGV2LWtleQ==","user":{"id":1,"full_name":"Admin User","email":"admin@bookstore.com","role":"admin"}}
➜  backend git:(main) ✗ curl -X POST http://localhost:8080/api/auth/login \
                              -H "Content-Type: application/json" \
                              -d '{
                            "email":"user@bookstore.com",
                            "password":"user123"
                          }'
{"token":"Mjp1c2VyOnN1cGVyLXNlY3JldC1kZXYta2V5","user":{"id":2,"full_name":"Test User","email":"user@bookstore.com","role":"user"}}
➜  backend git:(main) ✗ 
```