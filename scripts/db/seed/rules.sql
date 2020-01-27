INSERT INTO rules (id, data_source_id, "name", url_pattern)
VALUES 
    (1, NULL, 'Airport Rule', '/airports'),
    (2, NULL, 'Hotel Rule', '/hotels'),
    (3, NULL, 'Event Rule', '/events'),
    (4, NULL, 'Event Rule with Path Params', '/events/<eventName>-<eventDate>');

SELECT setval('rules_id_seq', (SELECT MAX(id) FROM rules));