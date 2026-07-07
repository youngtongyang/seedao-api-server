-- Add SIP label column for legacy/imported proposals

ALTER TABLE proposal ADD COLUMN IF NOT EXISTS sip VARCHAR(64);

CREATE INDEX IF NOT EXISTS idx_proposal_sip ON proposal (sip) WHERE sip IS NOT NULL AND sip <> '';
