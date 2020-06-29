CREATE TABLE url_sync (
    version             INTEGER AUTO_INCREMENT PRIMARY KEY,
    operation           TEXT NOT NULL,
    rule_id             INTEGER NOT NULL,
    latest_url_pattern  TEXT,
    created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
