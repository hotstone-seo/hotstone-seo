CREATE TABLE history (
    id bigserial PRIMARY KEY,
    time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    entity_id BIGINT NOT NULL,
    entity_from VARCHAR (25) NOT NULL,
    username VARCHAR (100) NOT NULL,
    data JSONB NOT NULL DEFAULT '{}' :: JSONB
);