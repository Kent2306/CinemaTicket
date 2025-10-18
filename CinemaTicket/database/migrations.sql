-- Создание таблицы залов
CREATE TABLE IF NOT EXISTS halls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    capacity INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы фильмов
CREATE TABLE IF NOT EXISTS movies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    poster_url TEXT,
    duration INTEGER, -- в минутах
    price DECIMAL(10,2) NOT NULL,
    hall_id INTEGER,
    show_time DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (hall_id) REFERENCES halls (id)
);

-- Создание таблицы бронирований
CREATE TABLE IF NOT EXISTS bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    movie_id INTEGER NOT NULL,
    seat_number INTEGER NOT NULL,
    customer_name TEXT NOT NULL,
    customer_email TEXT NOT NULL,
    booking_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    status TEXT DEFAULT 'active',
    FOREIGN KEY (movie_id) REFERENCES movies (id)
);

-- Вставка тестовых данных залов
INSERT OR IGNORE INTO halls (id, name, capacity) VALUES 
(1, 'Зал 1', 10),
(2, 'Зал 2', 10),
(3, 'Зал 3', 10);

