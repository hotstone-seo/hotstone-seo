INSERT INTO modules (id, "name", "path", pattern, label, api_path)
VALUES 
    (1, 'rule', 'rules', 'rules*', 'Rules', '{"api_path": [{"path": "/api/rules"},{"path": "/api/center"},{"path": "/api/structured-data"},{"path": "/api/tags"}]}'),
    (2, 'datasources','datasources', 'datasources*','Data Sources', '{"api_path": [{"path": "/api/data_sources"}]}'),
    (3, 'mismatchrule', 'mismatch-rule', 'mismatch*' , 'Mismatch Rule', '{"api_path": [{"path": "/api/metrics/mismatched"}]}'),
    (4, 'analytic', 'analytic', 'analytic*', 'Analytic', '{"api_path": [{"path": "/api/metrics"}]}'),
    (5, 'simulation', 'simulation', 'simulation*', 'Simulation Page', '{"api_path": [{"path": "/p/match"}]}'),
    (6, 'audittrail', 'audit-trail', 'audit*', 'Audit Trail', '{"api_path": [{"path": "/api/audit-trail"}]}'),
    (7, 'user', 'users', 'users*', 'Users', '{"api_path": [{"path": "/api/users"}]}'),
    (8, 'roletype', 'role-type', 'role-type*', 'User Role', '{"api_path": [{"path": "/api/role_types"}]}'),
    (9, 'module', 'modules', 'modules*', 'Modules', '{"api_path": [{"path": "/api/modules"}]}'),
    (10, 'clientkey', 'client-keys', 'client-keys*', 'Client Keys', '{"api_path": [{"path": "/api/client-keys"}]}')
    ;

SELECT setval('modules_id_seq', (SELECT MAX(id) FROM modules));