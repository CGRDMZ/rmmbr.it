BEGIN;

CREATE TABLE users(
    id serial NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);


COMMIT;