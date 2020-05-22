INSERT INTO rules (id, "name", url_pattern)
VALUES 
    (1, 'Airport Rule', '/airports'),
    (2, 'Hotel Rule', '/hotels'),
    (3, 'Event Rule', '/events'),
    (4, 'Event Rule with Path Params', '/events/<eventName>');

SELECT setval('rules_id_seq', (SELECT MAX(id) FROM rules));

INSERT INTO rule_data_sources (rule_id, data_source_id)
VALUES
    (4, 1);
