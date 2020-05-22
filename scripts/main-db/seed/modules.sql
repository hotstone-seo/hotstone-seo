INSERT INTO modules (id, "name", "path", pattern, label)
VALUES 
    (1, 'rule', 'rules', 'rules*', 'Rules'),
    (2, 'datasources','datasources', 'datasources*','Data Sources'),
    (3, 'mismatchrule', 'mismatch-rule', 'mismatch*' , 'Mismatch Rule'),
    (4, 'analytic', 'analytic', 'analytic*', 'Analytic'),
    (5, 'simulation', 'simulation', 'simulation*', 'Simulation Page'),
    (6, 'audittrail', 'audit-trail', 'audit*', 'Audit Trail'),
    (7, 'user', 'users', 'users*', 'Users'),
    (8, 'roletype', 'role-type', 'role-types*', 'User Role'),
    (9, 'module', 'modules', 'modules*', 'Modules'),
    (10, 'clientkey', 'client-keys', 'client-keys*', 'Client Keys')
    ;

SELECT setval('modules_id_seq', (SELECT MAX(id) FROM modules));