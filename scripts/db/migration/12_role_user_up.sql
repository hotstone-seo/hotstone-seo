CREATE TABLE role_user (
    id serial PRIMARY KEY,
    email VARCHAR (60) NOT NULL UNIQUE,
    role_type_id INTEGER DEFAULT '0',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL
)