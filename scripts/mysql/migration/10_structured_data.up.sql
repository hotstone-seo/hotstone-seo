CREATE TABLE structured_datas (
    id          INTEGER AUTO_INCREMENT PRIMARY KEY,
    rule_id     INTEGER NOT NULL,
    `type`      VARCHAR (255) NOT NULL,
    `data`      JSON NOT NULL DEFAULT JSON_TYPE("{}"),
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX structured_datas_rule_id_idx ON structured_datas(rule_id);
