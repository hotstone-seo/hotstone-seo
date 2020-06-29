CREATE TABLE audit_trail (
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    time            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    entity_name     VARCHAR (25) NOT NULL,
    entity_id       BIGINT NOT NULL,
    operation       VARCHAR (10) NOT NULL,
    username        VARCHAR (100),
    old_data        JSON NOT NULL DEFAULT JSON_TYPE("{}"),
    new_data        JSON NOT NULL DEFAULT JSON_TYPE("{}")
);
