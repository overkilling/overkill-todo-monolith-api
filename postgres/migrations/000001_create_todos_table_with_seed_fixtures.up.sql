CREATE TABLE IF NOT EXISTS todos(
    id SERIAL PRIMARY KEY,
    todo TEXT
);

INSERT INTO todos (todo) VALUES ('Some todo');
INSERT INTO todos (todo) VALUES ('Another todo');
INSERT INTO todos (todo) VALUES ('Yet another todo');
