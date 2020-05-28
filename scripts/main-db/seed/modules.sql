INSERT INTO modules (id, "name", "path", pattern, label, api_path)
VALUES 
    (1, 'rule', 'rules', 'rules*', 'Rules', 'api/rules'),
    (2, 'datasources','datasources', 'datasources*','Data Sources', 'api/data_sources'),
    (3, 'mismatchrule', 'mismatch-rule', 'mismatch*' , 'Mismatch Rule', 'api/metrics/mismatched'),
    (4, 'analytic', 'analytic', 'analytic*', 'Analytic', 'api/metrics'),
    (5, 'simulation', 'simulation', 'simulation*', 'Simulation Page', 'p/match'),
    (6, 'audittrail', 'audit-trail', 'audit*', 'Audit Trail', 'api/audit-trail'),
    (7, 'user', 'users', 'users*', 'Users', 'api/users'),
    (8, 'roletype', 'role-type', 'role-type*', 'User Role', 'api/role_types'),
    (9, 'module', 'modules', 'modules*', 'Modules', 'api/modules'),
    (10, 'clientkey', 'client-keys', 'client-keys*', 'Client Keys', 'api/client-keys')
    ;

SELECT setval('modules_id_seq', (SELECT MAX(id) FROM modules));