-- Read-only schema check for seedao_api (local vs remote).
-- Usage:
--   export PGHOST=... PGUSER=seedao PGPASSWORD=... PGDATABASE=seedao_api
--   psql -h "$PGHOST" -U "$PGUSER" -d "$PGDATABASE" -f src/scripts/check-db-schema.sql

\pset footer off
\timing off

\echo '========== seedao_api schema check =========='

WITH expected_tables(name) AS (
  VALUES
    ('users'),
    ('user_phones'),
    ('sns_names'),
    ('sns_chain_registry'),
    ('payment_orders'),
    ('payment_refunds'),
    ('coupons'),
    ('coupon_redemptions'),
    ('courses'),
    ('course_skus'),
    ('course_discount_rules'),
    ('course_enrollments'),
    ('proposal'),
    ('proposal_vote'),
    ('proposal_comment')
),
checks AS (
  SELECT
    e.name,
    CASE WHEN t.tablename IS NOT NULL THEN 'OK' ELSE 'MISSING' END AS status
  FROM expected_tables e
  LEFT JOIN pg_tables t
    ON t.schemaname = 'public' AND t.tablename = e.name
)
SELECT * FROM checks ORDER BY status DESC, name;

\echo ''
\echo '---------- proposal columns (governance) ----------'

SELECT
  c.column_name,
  c.data_type,
  c.column_default,
  CASE WHEN c.column_name IS NOT NULL THEN 'OK' ELSE 'MISSING' END AS status
FROM (
  VALUES
    ('sip'),
    ('associate_proposal_id'),
    ('resolution_rule'),
    ('updated_at'),
    ('publicity_seconds'),
    ('vote_duration_seconds')
) AS expected(column_name)
LEFT JOIN information_schema.columns c
  ON c.table_schema = 'public'
 AND c.table_name = 'proposal'
 AND c.column_name = expected.column_name
ORDER BY status DESC, column_name;

\echo ''
\echo '---------- sequences / indexes ----------'

SELECT
  'proposal_sip_seq' AS object,
  CASE WHEN EXISTS (SELECT 1 FROM pg_class WHERE relname = 'proposal_sip_seq') THEN 'OK' ELSE 'MISSING' END AS status
UNION ALL
SELECT
  'idx_proposal_associate',
  CASE WHEN EXISTS (
    SELECT 1 FROM pg_indexes WHERE schemaname = 'public' AND indexname = 'idx_proposal_associate'
  ) THEN 'OK' ELSE 'MISSING' END
UNION ALL
SELECT
  'idx_proposal_sip',
  CASE WHEN EXISTS (
    SELECT 1 FROM pg_indexes WHERE schemaname = 'public' AND indexname = 'idx_proposal_sip'
  ) THEN 'OK' ELSE 'MISSING' END
UNION ALL
SELECT
  'idx_proposal_vote_proposal_sns_unique',
  CASE WHEN EXISTS (
    SELECT 1 FROM pg_indexes WHERE schemaname = 'public' AND indexname = 'idx_proposal_vote_proposal_sns_unique'
  ) THEN 'OK' ELSE 'MISSING' END;

\echo ''
\echo '---------- SIP sanity (if proposal exists) ----------'

SELECT
  COUNT(*) FILTER (WHERE sip IS NULL) AS sip_null_count,
  COUNT(*) FILTER (WHERE sip = 0) AS sip_zero_count,
  COUNT(*) FILTER (WHERE sip > 0) AS sip_assigned_count,
  COALESCE(MAX(sip), 0) AS max_sip
FROM proposal;

SELECT
  CASE WHEN EXISTS (SELECT 1 FROM pg_class WHERE relname = 'proposal_sip_seq')
    THEN (SELECT last_value FROM proposal_sip_seq)
    ELSE NULL
  END AS proposal_sip_seq_last_value;

\echo ''
\echo '---------- migration hints (manual) ----------'
\echo 'If proposal table MISSING        -> run 007_proposals.sql (+ import if needed)'
\echo 'If sip column MISSING            -> run 008_proposal_sip.sql'
\echo 'If associate_proposal_id MISSING -> run 009_proposal_sip_alloc.sql'
\echo 'If proposal_sip_seq MISSING      -> run 009_proposal_sip_alloc.sql'
\echo 'If payment/course tables MISSING -> run 004~006 as needed'
\echo '========== done =========='
