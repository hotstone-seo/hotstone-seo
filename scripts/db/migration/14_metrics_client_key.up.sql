CREATE TABLE metrics_client_key (
  time              TIMESTAMPTZ       NOT NULL DEFAULT CURRENT_TIMESTAMP,
  client_key_id     INTEGER           NOT NULL
);

SELECT create_hypertable('metrics_client_key', 'time');