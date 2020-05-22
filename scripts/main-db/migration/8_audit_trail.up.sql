CREATE TABLE audit_trail (
    id bigserial PRIMARY KEY,
    time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    entity_name VARCHAR (25) NOT NULL,
    entity_id BIGINT NOT NULL,
    operation VARCHAR (10) NOT NULL,
    username VARCHAR (100),
    old_data JSONB NOT NULL DEFAULT '{}' :: JSONB,
    new_data JSONB NOT NULL DEFAULT '{}' :: JSONB
);