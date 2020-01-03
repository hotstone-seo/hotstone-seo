INSERT INTO locales (id, lang_code, country_code)
VALUES 
    (1, 'en', 'US'),
    (2, 'id', 'ID');

SELECT setval('locales_id_seq', (SELECT MAX(id) FROM locales));