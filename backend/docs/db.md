### БД я закинул в докер, по этому чтобы достучаться до неё нужно выполнить данную команду:
```bash
docker exec -it bookstore_postgres psql -U bookstore_user -d bookstore
```

### И там уже работаем с psql. Пример:
```bash
bookstore-# \dt
                List of relations
 Schema |     Name      | Type  |     Owner      
--------+---------------+-------+----------------
 public | authors       | table | bookstore_user
 public | books         | table | bookstore_user
 public | categories    | table | bookstore_user
 public | notifications | table | bookstore_user
 public | purchases     | table | bookstore_user
 public | rentals       | table | bookstore_user
 public | users         | table | bookstore_user
(7 rows)

bookstore-# \d books
                                             Table "public.books"
       Column        |            Type             | Collation | Nullable |              Default              
---------------------+-----------------------------+-----------+----------+-----------------------------------
 id                  | integer                     |           | not null | nextval('books_id_seq'::regclass)
 title               | character varying(255)      |           | not null | 
 description         | text                        |           |          | 
 author_id           | integer                     |           | not null | 
 category_id         | integer                     |           | not null | 
 year_written        | integer                     |           | not null | 
 purchase_price      | numeric(10,2)               |           | not null | 
 rent_price_2_weeks  | numeric(10,2)               |           | not null | 
 rent_price_1_month  | numeric(10,2)               |           | not null | 
 rent_price_3_months | numeric(10,2)               |           | not null | 
 status              | book_status                 |           | not null | 'available'::book_status
 is_available        | boolean                     |           | not null | true
 cover_url           | text                        |           |          | 
 created_at          | timestamp without time zone |           | not null | now()
 updated_at          | timestamp without time zone |           | not null | now()
Indexes:
    "books_pkey" PRIMARY KEY, btree (id)
    "idx_books_author" btree (author_id)
    "idx_books_category" btree (category_id)
    "idx_books_status" btree (status)
    "idx_books_year" btree (year_written)
Check constraints:
    "books_purchase_price_check" CHECK (purchase_price >= 0::numeric)
    "books_rent_price_1_month_check" CHECK (rent_price_1_month >= 0::numeric)
    "books_rent_price_2_weeks_check" CHECK (rent_price_2_weeks >= 0::numeric)
    "books_rent_price_3_months_check" CHECK (rent_price_3_months >= 0::numeric)
    "books_year_written_check" CHECK (year_written > 0)
Foreign-key constraints:
    "books_author_id_fkey" FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE RESTRICT
    "books_category_id_fkey" FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT
Referenced by:
    TABLE "purchases" CONSTRAINT "purchases_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE RESTRICT
    TABLE "rentals" CONSTRAINT "rentals_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE RESTRICT
Triggers:
    trg_books_updated_at BEFORE UPDATE ON books FOR EACH ROW EXECUTE FUNCTION set_updated_at()

bookstore-# 
```
