INSERT INTO user_roles (id, "name", menus, paths)
VALUES 
    (1, 'Admin','[".*"]','[".*"]'),
    (2, 'Staff','["Rules","Mismatch Rule","Analytic","Simulation"]','["/api/center*", "/api/structured-data*", "/api/tags*", "/api/metrics*", "/p/match*", "/api/rules*", "/api/logout"]');

SELECT setval('user_roles_id_seq', (SELECT MAX(id) FROM user_roles));