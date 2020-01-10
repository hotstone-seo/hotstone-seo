CREATE TABLE metrics_rule_matching (
  time        TIMESTAMPTZ       NOT NULL DEFAULT CURRENT_TIMESTAMP,
  is_matched  INTEGER           NOT NULL,
  url_mismatched                TEXT
);

SELECT create_hypertable('metrics_rule_matching', 'time');