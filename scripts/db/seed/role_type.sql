INSERT INTO role_type (id, "type")
VALUES 
    (1, 'ADMIN'),
    (2, 'USER');

SELECT setval('role_type_id_seq', (SELECT MAX(id) FROM role_type));