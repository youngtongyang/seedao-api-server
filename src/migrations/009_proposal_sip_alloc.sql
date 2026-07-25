-- SIP allocation: default 0 (unassigned), sequence for new numbers, associate for close-project inheritance.
-- Also upgrades legacy varchar sip (from older 008) to BIGINT.

ALTER TABLE proposal ADD COLUMN IF NOT EXISTS associate_proposal_id BIGINT REFERENCES proposal(id);

DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_schema = 'public'
          AND table_name = 'proposal'
          AND column_name = 'sip'
          AND data_type IN ('character varying', 'text')
    ) THEN
        DROP INDEX IF EXISTS idx_proposal_sip;
        ALTER TABLE proposal ALTER COLUMN sip DROP DEFAULT;
        ALTER TABLE proposal
            ALTER COLUMN sip TYPE BIGINT
            USING (
                CASE
                    WHEN sip IS NULL OR btrim(sip) = '' THEN 0
                    WHEN sip ~ '^[0-9]+$' THEN sip::bigint
                    WHEN sip ~* '^SIP-?[0-9]+$' THEN NULLIF(regexp_replace(sip, '[^0-9]', '', 'g'), '')::bigint
                    ELSE 0
                END
            );
    ELSIF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_schema = 'public'
          AND table_name = 'proposal'
          AND column_name = 'sip'
    ) THEN
        ALTER TABLE proposal ADD COLUMN sip BIGINT NOT NULL DEFAULT 0;
    END IF;
END $$;

UPDATE proposal SET sip = 0 WHERE sip IS NULL;
ALTER TABLE proposal ALTER COLUMN sip SET DEFAULT 0;
ALTER TABLE proposal ALTER COLUMN sip SET NOT NULL;

CREATE INDEX IF NOT EXISTS idx_proposal_associate ON proposal (associate_proposal_id);
DROP INDEX IF EXISTS idx_proposal_sip;
CREATE INDEX idx_proposal_sip ON proposal (sip) WHERE sip > 0;

-- Continue from imported legacy data (run once; safe if sequence already exists).
DO $$
DECLARE
    max_sip BIGINT;
BEGIN
    SELECT COALESCE(MAX(sip), 0) INTO max_sip FROM proposal WHERE sip > 0;
    IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'proposal_sip_seq') THEN
        EXECUTE format('CREATE SEQUENCE proposal_sip_seq START WITH %s', GREATEST(max_sip + 1, 1));
    ELSE
        PERFORM setval('proposal_sip_seq', GREATEST(max_sip, 1));
    END IF;
END $$;
