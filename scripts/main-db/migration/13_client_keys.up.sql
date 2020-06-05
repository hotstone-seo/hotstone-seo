CREATE TABLE client_keys (
    id serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    prefix VARCHAR NOT NULL UNIQUE,
    key VARCHAR NOT NULL UNIQUE,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX client_keys_prefix_key_idx ON client_keys(prefix, key);