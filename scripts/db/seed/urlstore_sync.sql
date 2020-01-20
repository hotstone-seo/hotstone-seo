INSERT INTO urlstore_sync (version, "operation", rule_id, latest_url_pattern)
VALUES
    (1, 'INSERT', 1, '/airports'),
    (2, 'INSERT', 2, '/hotels'),
    (3, 'INSERT', 3, '/events');

SELECT setval('urlstore_sync_version_seq', (SELECT MAX(version) FROM urlstore_sync));