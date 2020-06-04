CREATE TABLE tags (
    id serial PRIMARY KEY,
    rule_id INTEGER NOT NULL,
    locale VARCHAR (255) NOT NULL,
    "type" VARCHAR (255) NOT NULL,
    attributes JSONB NOT NULL DEFAULT '{}'::JSONB,
    "value" VARCHAR (255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX tags_rule_id_locale_idx ON tags(rule_id, locale);
