-- Run AFTER proposal_v2_tables.sql pg_dump import (do NOT run full 007 before import).

ALTER TABLE proposal_v2 RENAME TO proposal;
ALTER TABLE proposal_vote_v2 RENAME TO proposal_vote;
ALTER TABLE proposal_comment_v2 RENAME TO proposal_comment;

ALTER TABLE proposal ADD COLUMN IF NOT EXISTS resolution_rule VARCHAR(32) NOT NULL DEFAULT 'standard';
ALTER TABLE proposal ADD COLUMN IF NOT EXISTS sip BIGINT NOT NULL DEFAULT 0;

DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_schema = 'public' AND table_name = 'proposal' AND column_name = 'migrated_at'
    ) AND NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_schema = 'public' AND table_name = 'proposal' AND column_name = 'updated_at'
    ) THEN
        ALTER TABLE proposal RENAME COLUMN migrated_at TO updated_at;
    END IF;
END $$;

ALTER TABLE proposal ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

-- Index created/normalized in 009_proposal_sip_alloc.sql after sip is BIGINT.
