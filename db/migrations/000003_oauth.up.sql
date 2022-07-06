BEGIN;

ALTER TABLE url_map 
    ADD user_id VARCHAR(100);

ALTER TABLE users
    DROP COLUMN password;

COMMIT;