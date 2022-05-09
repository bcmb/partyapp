CREATE TABLE IF NOT EXISTS "user"
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(50)  NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS party
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(50) NOT NULL,
    city      VARCHAR(50) NOT NULL,
    address   VARCHAR(50) NOT NULL,
    date_time TIMESTAMP   NOT NULL
);