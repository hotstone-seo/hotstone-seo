INSERT INTO role_module (id, module)
VALUES 
    (1, 'Rules'),
    (2, 'Data Source'),
    (3, 'Mismatched Rule'),
    (4, 'Analytic'),
    (5, 'Simulation'),
    (6, 'Audit Trail');

SELECT setval('role_module_id_seq', (SELECT MAX(id) FROM role_module));