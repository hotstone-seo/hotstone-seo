INSERT INTO tags (id, rule_id, locale, "type", attributes, "value")
VALUES 
    (1, 1, 'en-US', 'title', '{}', 'Airport Title'),
    (2, 1, 'en-US', 'meta', '{"name": "description", "content": "Airport Description"}', ''),
    (3, 4, 'en-US', 'title', '{}', 'Event with Path Params Title'),
    (4, 4, 'en-US', 'meta', '{"name": "description", "content": "Event with Path Params Description"}', '');

SELECT setval('tags_id_seq', (SELECT MAX(id) FROM tags));
