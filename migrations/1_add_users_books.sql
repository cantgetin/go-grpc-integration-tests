-- +migrate Up

CREATE TABLE books
(
    id     integer not null primary key,
    name   varchar(40),
    author varchar(40)
);

CREATE TABLE users
(
    id            integer     not null primary key,
    name          varchar(40) not null,
    username      varchar(40) not null unique,
    password_hash varchar(40) not null
);

-- +migrate Down

DROP TABLE books;
DROP TABLE users;