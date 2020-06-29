CREATE TYPE enum_status AS ENUM ('start', 'stop', '') ; 

CREATE TABLE rules (
    id serial PRIMARY KEY,
    "name" VARCHAR (255) NOT NULL,
    url_pattern TEXT NOT NULL UNIQUE,
    status enum_status DEFAULT '',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    change_status_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
create index rules_status ON rules(status);