CREATE TABLE metrics_unmatched (
    id serial PRIMARY KEY,
    request_path TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_metrics_unmatched_request_path ON metrics_unmatched(request_path);

