CREATE TABLE urlstore_sync (
    "version" serial PRIMARY KEY,
    operation TEXT NOT NULL,
    rule_id INTEGER NOT NULL,
    latest_url_pattern TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);