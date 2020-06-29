CREATE TABLE rule_data_sources (
    rule_id INTEGER NOT NULL,
    data_source_id INTEGER NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(rule_id, data_source_id)
);

CREATE INDEX rule_data_sources_rule_id_idx ON rule_data_sources(rule_id);
