INSERT INTO client_keys (id, "name", "prefix", "key")
VALUES 
    (1, 'Simulation', 'Otdu1qe', '30c1e2c322bea18f35ad663d0114943a52c2188f71261675d009b5eb8a9b7020');

SELECT setval('client_keys_id_seq', (SELECT MAX(id) FROM client_keys));
