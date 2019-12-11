CREATE TABLE rules (
    id serial PRIMARY KEY,
    "name" VARCHAR (255) NOT NULL,
    url_pattern TEXT NOT NULL,
    exclusion TEXT,
    id_data_source INTEGER NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);