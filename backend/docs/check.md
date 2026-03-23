## Проверяем логин и покупку

```bash
➜  backend git:(main) ✗ curl -X POST http://localhost:8080/api/auth/login \
                              -H "Content-Type: application/json" \
                              -d '{
                            "email":"user@bookstore.com",
                            "password":"user123"
                          }'
{"token":"Mjp1c2VyOnN1cGVyLXNlY3JldC1kZXYta2V5","user":{"id":2,"full_name":"Test User","email":"user@bookstore.com","role":"user"}}
➜  backend git:(main) ✗ curl -X POST http://localhost:8080/api/purchases \
                              -H "Content-Type: application/json" \
                              -H "Authorization: Bearer Mjp1c2VyOnN1cGVyLXNlY3JldC1kZXYta2V5" \
                              -d '{
                            "book_id": 6
                          }'
{"error":"book is not available for purchase"}
➜  backend git:(main) ✗ curl http://localhost:8080/api/purchases/my \
                              -H "Authorization: Bearer Mjp1c2VyOnN1cGVyLXNlY3JldC1kZXYta2V5"
{"items":[{"id":3,"user_id":2,"book_id":6,"book_title":"Три товарища","price":13.99,"purchased_at":"2026-03-23T12:10:30.843612Z"},{"id":1,"user_id":2,"book_id":1,"book_title":"Мастер и Маргарита","price":15.99,"purchased_at":"2026-03-23T11:59:01.493788Z"},{"id":2,"user_id":2,"book_id":4,"book_title":"1984","price":14.99,"purchased_at":"2026-03-23T11:59:01.493788Z"}],"total":3}
➜  backend git:(main) ✗ curl -X POST http://localhost:8080/api/rentals \
                              -H "Content-Type: application/json" \
                              -H "Authorization: Bearer Mjp1c2VyOnN1cGVyLXNlY3JldC1kZXYta2V5" \
                              -d '{
                            "book_id": 7,
                            "tariff": "2_weeks"
                          }'
{"id":4,"user_id":2,"book_id":7,"book_title":"Убить пересмешника","tariff":"2_weeks","start_date":"2026-03-23T15:12:22.700709Z","end_date":"2026-04-06T15:12:22.700709Z","status":"active","created_at":"2026-03-23T12:12:22.701Z"}
➜  backend git:(main) ✗ curl -X POST http://localhost:8080/api/rentals \
                              -H "Content-Type: application/json" \
                              -H "Authorization: Bearer Mjp1c2VyOnN1cGVyLXNlY3JldC1kZXYta2V5" \
                              -d '{
                            "book_id": 8,
                            "tariff": "1_month"
                          }'
{"id":5,"user_id":2,"book_id":8,"book_title":"Война и мир","tariff":"1_month","start_date":"2026-03-23T15:12:29.512178Z","end_date":"2026-04-23T15:12:29.512178Z","status":"active","created_at":"2026-03-23T12:12:29.512644Z"}
➜  backend git:(main) ✗ curl http://localhost:8080/api/rentals/my \
                              -H "Authorization: Bearer Mjp1c2VyOnN1cGVyLXNlY3JldC1kZXYta2V5"
{"items":[{"id":5,"user_id":2,"book_id":8,"book_title":"Война и мир","tariff":"1_month","start_date":"2026-03-23T15:12:29.512178Z","end_date":"2026-04-23T15:12:29.512178Z","status":"active","created_at":"2026-03-23T12:12:29.512644Z"},{"id":4,"user_id":2,"book_id":7,"book_title":"Убить пересмешника","tariff":"2_weeks","start_date":"2026-03-23T15:12:22.700709Z","end_date":"2026-04-06T15:12:22.700709Z","status":"active","created_at":"2026-03-23T12:12:22.701Z"},{"id":1,"user_id":2,"book_id":2,"book_title":"Преступление и наказание","tariff":"2_weeks","start_date":"2026-03-23T11:59:01.494722Z","end_date":"2026-04-06T11:59:01.494722Z","status":"active","created_at":"2026-03-23T11:59:01.494722Z"},{"id":2,"user_id":2,"book_id":3,"book_title":"Сон смешного человека","tariff":"1_month","start_date":"2026-02-22T11:59:01.494722Z","end_date":"2026-03-24T11:59:01.494722Z","status":"active","created_at":"2026-03-23T11:59:01.494722Z"},{"id":3,"user_id":2,"book_id":5,"book_title":"Маленький принц","tariff":"1_month","start_date":"2026-02-11T11:59:01.494722Z","end_date":"2026-03-13T11:59:01.494722Z","status":"expired","created_at":"2026-03-23T11:59:01.494722Z"}],"total":5}
➜  backend git:(main) ✗ curl http://localhost:8080/api/books/6
                        curl http://localhost:8080/api/books/7
{"id":6,"title":"Три товарища","description":"Роман о дружбе и любви после войны","author_id":5,"author_name":"Эрих Мария Ремарк","category_id":4,"category_name":"Роман","year_written":1936,"purchase_price":13.99,"rent_price_2_weeks":3.49,"rent_price_1_month":6.49,"rent_price_3_months":11.49,"status":"sold_out","is_available":false,"cover_url":"/covers/three_comrades.jpg"}
{"id":7,"title":"Убить пересмешника","description":"Роман о расовой несправедливости","author_id":6,"author_name":"Харпер Ли","category_id":4,"category_name":"Роман","year_written":1960,"purchase_price":11.99,"rent_price_2_weeks":2.99,"rent_price_1_month":5.49,"rent_price_3_months":9.99,"status":"rented","is_available":false,"cover_url":"/covers/to_kill_a_mockingbird.jpg"}
➜  backend git:(main) ✗ 
```