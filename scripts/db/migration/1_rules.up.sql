CREATE TABLE rules (
    id serial PRIMARY KEY,
    data_source_id INTEGER,
    "name" VARCHAR (255) NOT NULL,
    url_pattern TEXT NOT NULL UNIQUE,
    is_active TINYINT DEFAULT '1',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
