-- =========================================
-- CLEAN
-- =========================================
TRUNCATE TABLE notifications RESTART IDENTITY CASCADE;
TRUNCATE TABLE rentals RESTART IDENTITY CASCADE;
TRUNCATE TABLE purchases RESTART IDENTITY CASCADE;
TRUNCATE TABLE books RESTART IDENTITY CASCADE;
TRUNCATE TABLE authors RESTART IDENTITY CASCADE;
TRUNCATE TABLE categories RESTART IDENTITY CASCADE;
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

-- =========================================
-- CATEGORIES
-- =========================================
INSERT INTO categories (name) VALUES
('Классика'),
('Антиутопия'),
('Психология'),
('Роман'),
('Философия'),
('Современная литература');

-- =========================================
-- AUTHORS
-- =========================================
INSERT INTO authors (full_name) VALUES
('Михаил Булгаков'),
('Фёдор Достоевский'),
('Джордж Оруэлл'),
('Антуан де Сент-Экзюпери'),
('Эрих Мария Ремарк'),
('Харпер Ли'),
('Лев Толстой'),
('Оскар Уайльд'),
('Дэниел Киз'),
('Джером Сэлинджер'),
('Герман Гессе'),
('Томас Манн'),
('Кристофер Бакли'),
('Лилиан Войнич');

-- =========================================
-- USERS
-- =========================================
INSERT INTO users (full_name, email, password_hash, role) VALUES
('Admin User', 'admin@bookstore.com', 'admin123', 'admin'),
('Test User', 'user@bookstore.com', 'user123', 'user');

-- =========================================
-- BOOKS
-- =========================================
INSERT INTO books (
    title, description, author_id, category_id,
    year_written,
    purchase_price,
    rent_price_2_weeks,
    rent_price_1_month,
    rent_price_3_months,
    status,
    is_available,
    cover_url
) VALUES

-- Булгаков
('Мастер и Маргарита', 'Роман о добре и зле, любви и мистике',
 1, 1, 1967, 15.99, 3.99, 6.99, 12.99, 'available', true, '/covers/master_i_margarita.jpg'),

-- Достоевский
('Преступление и наказание', 'Психологический роман о морали и вине',
 2, 3, 1866, 12.99, 2.99, 5.99, 10.99, 'available', true, '/covers/crime_and_punishment.jpg'),

('Сон смешного человека', 'Философский рассказ о смысле жизни',
 2, 5, 1877, 7.99, 1.99, 3.99, 6.99, 'available', true, '/covers/dream_of_ridiculous_man.jpg'),

-- Оруэлл
('1984', 'Антиутопия о тоталитарном обществе',
 3, 2, 1949, 14.99, 3.49, 6.49, 11.99, 'available', true, '/covers/1984.jpg'),

-- Экзюпери
('Маленький принц', 'Философская сказка о жизни и любви',
 4, 5, 1943, 9.99, 1.99, 3.99, 7.99, 'available', true, '/covers/little_prince.jpg'),

-- Ремарк
('Три товарища', 'Роман о дружбе и любви после войны',
 5, 4, 1936, 13.99, 3.49, 6.49, 11.49, 'available', true, '/covers/three_comrades.jpg'),

-- Харпер Ли
('Убить пересмешника', 'Роман о расовой несправедливости',
 6, 4, 1960, 11.99, 2.99, 5.49, 9.99, 'available', true, '/covers/to_kill_a_mockingbird.jpg'),

-- Толстой
('Война и мир', 'Эпопея о войне и судьбах людей',
 7, 1, 1869, 19.99, 4.99, 8.99, 15.99, 'available', true, '/covers/war_and_peace.jpg'),

-- Уайльд
('Портрет Дориана Грея', 'Философский роман о красоте и морали',
 8, 5, 1890, 10.99, 2.49, 4.99, 8.99, 'available', true, '/covers/dorian_gray.jpg'),

-- Киз
('Цветы для Элджернона', 'История о развитии интеллекта',
 9, 3, 1966, 12.49, 2.99, 5.49, 9.49, 'available', true, '/covers/flowers_for_algernon.jpg'),

-- Сэлинджер
('Над пропастью во ржи', 'Роман о подростковом кризисе',
 10, 4, 1951, 11.49, 2.99, 5.49, 9.49, 'available', true, '/covers/catcher_in_the_rye.jpg'),

-- Гессе
('Степной волк', 'Философский роман о раздвоении личности',
 11, 5, 1927, 13.49, 3.49, 6.49, 11.49, 'available', true, '/covers/steppenwolf.jpg'),

('Демиан', 'История взросления и поиска себя',
 11, 5, 1919, 11.49, 2.99, 5.49, 9.49, 'available', true, '/covers/demian.jpg'),

-- Манн
('Волшебная гора', 'Философский роман о времени и жизни',
 12, 5, 1924, 14.99, 3.99, 6.99, 12.49, 'available', true, '/covers/magic_mountain.jpg'),

-- Бакли
('Здесь курят', 'Сатирический роман о PR и лоббизме',
 13, 6, 1994, 10.99, 2.99, 5.49, 9.49, 'available', true, '/covers/thank_you_for_smoking.jpg'),

-- Войнич
('Овод', 'Роман о революции и личной драме',
 14, 4, 1897, 9.99, 2.49, 4.99, 8.99, 'available', true, '/covers/gadfly.jpg');

-- =========================================
-- PURCHASES
-- =========================================
INSERT INTO purchases (user_id, book_id, price)
VALUES
(2, 1, 15.99),
(2, 4, 14.99);

-- =========================================
-- RENTALS
-- =========================================
INSERT INTO rentals (
    user_id,
    book_id,
    tariff,
    start_date,
    end_date,
    status
)
VALUES
(2, 2, '2_weeks', NOW(), NOW() + INTERVAL '14 days', 'active'),
(2, 3, '1_month', NOW() - INTERVAL '29 days', NOW() + INTERVAL '1 day', 'active'),
(2, 5, '1_month', NOW() - INTERVAL '40 days', NOW() - INTERVAL '10 days', 'expired');

-- =========================================
-- NOTIFICATIONS
-- =========================================
INSERT INTO notifications (
    user_id,
    rental_id,
    type,
    message,
    status
)
VALUES
(2, 2, 'rent_expiring', 'Срок аренды скоро истекает', 'pending');
