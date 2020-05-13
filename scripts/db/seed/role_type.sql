INSERT INTO role_type (id, "name", modules)
VALUES 
    (1, 'SUPER ADMIN','{"modules": ["rules","datasources","mismatchrule","analytic","simulation","audittrail","user","roleType"]}'),
    (2, 'ADMIN','{}');

SELECT setval('role_type_id_seq', (SELECT MAX(id) FROM role_type));