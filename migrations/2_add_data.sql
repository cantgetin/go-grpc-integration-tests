-- +migrate Up

INSERT INTO books (id, name, author)
VALUES (1, 'Martin Eden', 'Jack London'),
       (2, 'Ulysses', 'James Joyce');

INSERT INTO users (id, name, username, password_hash)
VALUES (1, 'John', 'john12', '2fbe00f6a2f5ca35a3d49adbdc33ce23'),
       (2, 'Ross', 'frosty', '11340a040bf55d21ab9e0cb1323ffa7c');

-- +migrate Down