INSERT INTO role_type (id, "name")
VALUES 
    (1, 'SUPER ADMIN'),
    (2, 'ADMIN');

SELECT setval('role_type_id_seq', (SELECT MAX(id) FROM role_type));