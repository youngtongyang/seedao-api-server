-- Add SIP number column (0 = unassigned). Matches model.Proposal.SIP (int).
-- Index + sequence are created in 009 (also upgrades legacy varchar sip).

ALTER TABLE proposal ADD COLUMN IF NOT EXISTS sip BIGINT NOT NULL DEFAULT 0;
