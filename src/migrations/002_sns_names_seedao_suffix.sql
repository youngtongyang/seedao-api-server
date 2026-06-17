-- Normalize sns_names to always use .seedao suffix (idempotent)

UPDATE sns_names
SET name = name || '.seedao'
WHERE name NOT LIKE '%.seedao';
