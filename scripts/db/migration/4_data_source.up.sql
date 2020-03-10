CREATE TABLE data_sources (
    id serial PRIMARY KEY,
    "name" VARCHAR (255) NOT NULL,
    "url" VARCHAR (255) NOT NULL,
    is_active CHAR(1) DEFAULT '1',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
