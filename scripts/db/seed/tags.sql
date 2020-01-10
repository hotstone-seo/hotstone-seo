INSERT INTO tags (id, rule_id, "locale", "type", attributes, "value")
VALUES 
    (1, 1, 'EN', 'title', '{}', 'Airport Title'),
    (2, 1, 'EN', 'meta', '{"type": "description", "content": "Airport Description"}', '');

SELECT setval('tags_id_seq', (SELECT MAX(id) FROM tags));