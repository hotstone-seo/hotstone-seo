CREATE TABLE metrics_api_key (
  time           TIMESTAMPTZ       NOT NULL DEFAULT CURRENT_TIMESTAMP,
  api_key_id     INTEGER           NOT NULL
);

SELECT create_hypertable('metrics_api_key', 'time');