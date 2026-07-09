-- SIP allocation: default 0 (unassigned), sequence for new numbers, associate for close-project inheritance.

ALTER TABLE proposal ADD COLUMN IF NOT EXISTS associate_proposal_id BIGINT REFERENCES proposal(id);

UPDATE proposal SET sip = 0 WHERE sip IS NULL;
ALTER TABLE proposal ALTER COLUMN sip SET DEFAULT 0;

CREATE INDEX IF NOT EXISTS idx_proposal_associate ON proposal (associate_proposal_id);
CREATE INDEX IF NOT EXISTS idx_proposal_sip ON proposal (sip) WHERE sip > 0;

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
