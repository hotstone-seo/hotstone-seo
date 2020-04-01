INSERT INTO data_sources (id, "name", "url")
VALUES 
    (1, 'events', 'http://localhost:3021/events?name=<eventName>');

SELECT setval('data_sources_id_seq', (SELECT MAX(id) FROM data_sources));
