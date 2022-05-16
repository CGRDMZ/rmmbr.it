BEGIN;

CREATE TABLE url_map (
    id serial NOT NULL,
    short_url VARCHAR(100) UNIQUE NOT NULL,
    long_url VARCHAR(5000) NOT NULL,
    visited_count INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE UNIQUE INDEX url_map_short_url_uindex ON url_map (short_url);


COMMIT;