CREATE TABLE IF NOT EXISTS api_server.users (
    id          UUID            PRIMARY KEY,
    username    VARCHAR(128)    NOT NULL,
    name        VARCHAR(128)    NOT NULL,
    surname     VARCHAR(128)    NOT NULL
);

CREATE TABLE IF NOT EXISTS api_server.user_credentials (
    id      UUID    REFERENCES api_server.users(id) ON DELETE CASCADE   PRIMARY KEY,
    hash    bytea   NOT NULL
);
