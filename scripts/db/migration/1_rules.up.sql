CREATE TYPE enum_status AS ENUM ('start', 'stop', '') ; 

CREATE TABLE rules (
    id serial PRIMARY KEY,
    data_source_id INTEGER,
    "name" VARCHAR (255) NOT NULL,
    url_pattern TEXT NOT NULL UNIQUE,
    is_active CHAR(1) DEFAULT '1',
    status enum_status DEFAULT '',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    change_status_at TIMESTAMPTZ DEFAULT NULL
);
