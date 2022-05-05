CREATE TABLE IF NOT EXISTS "authors"(
    "id" serial PRIMARY KEY,
    "email" CHAR(30) NOT NULL,
    "name" CHAR(25),
    "gender" CHAR(15)
);

CREATE TABLE IF NOT EXISTS "consumers"(
    "id" serial PRIMARY KEY,
    "name" CHAR(25),
    "gender" CHAR(15)
);

CREATE TABLE IF NOT EXISTS "books"(
    "id" serial PRIMARY KEY,
    "name" CHAR(50),
    "author_id" int,
    FOREIGN KEY("author_id")
    REFERENCES authors("id"),
    "book_id" int,
    FOREIGN KEY("book_id")
    REFERENCES books("id")

);