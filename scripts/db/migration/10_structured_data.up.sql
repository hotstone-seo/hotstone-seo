CREATE TABLE structured_datas (
    id serial PRIMARY KEY,
    rule_id INTEGER NOT NULL,
    "type" VARCHAR (255) NOT NULL,
    "data" JSONB NOT NULL DEFAULT '{}'::JSONB,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
)