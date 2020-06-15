INSERT INTO users (id, email, role_type_id)
VALUES 
    (1, 'iman.tunggono@tiket.com', 1),
    (2, 'fahri.hidayat@tiket.com', 1),
    (3, 'hawari.rahman@tiket.com', 1),
    (4, 'hendri.chia@tiket.com', 1);

SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
