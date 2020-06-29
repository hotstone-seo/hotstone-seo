CREATE TABLE rules (
    id               INTEGER AUTO_INCREMENT PRIMARY KEY,
    `name`           VARCHAR (255) NOT NULL,
    url_pattern      TEXT NOT NULL UNIQUE,
    status           ENUM('start', 'stop', '') DEFAULT '',
    updated_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    change_status_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX rules_status ON rules(status);
