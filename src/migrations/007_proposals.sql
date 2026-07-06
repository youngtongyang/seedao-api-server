-- Governance proposals (table names without _v2 suffix)

CREATE TABLE IF NOT EXISTS proposal (
    id                      BIGSERIAL PRIMARY KEY,
    legacy_id               BIGINT,
    legacy_record_id        VARCHAR(128),
    type                    VARCHAR(128) NOT NULL,
    title                   VARCHAR(512) NOT NULL,
    content                 JSONB NOT NULL,
    proposer_sns            VARCHAR(64),
    proposer_wallet         VARCHAR(42) NOT NULL,
    status                  VARCHAR(32) NOT NULL,
    vote_type               SMALLINT NOT NULL DEFAULT 0,
    vote_options            JSONB,
    resolution_rule         VARCHAR(32) NOT NULL DEFAULT 'standard',
    published_at            TIMESTAMPTZ,
    publicity_end_at        TIMESTAMPTZ,
    vote_end_at             TIMESTAMPTZ,
    publicity_seconds       INTEGER,
    vote_duration_seconds   INTEGER,
    created_at              TIMESTAMPTZ NOT NULL,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_proposal_legacy_id ON proposal (legacy_id);
CREATE INDEX IF NOT EXISTS idx_proposal_legacy_rec ON proposal (legacy_record_id);
CREATE INDEX IF NOT EXISTS idx_proposal_proposer_sns ON proposal (proposer_sns);
CREATE INDEX IF NOT EXISTS idx_proposal_pub_end ON proposal (publicity_end_at);
CREATE INDEX IF NOT EXISTS idx_proposal_status ON proposal (status);
CREATE INDEX IF NOT EXISTS idx_proposal_vote_end ON proposal (vote_end_at);
CREATE INDEX IF NOT EXISTS idx_proposal_created_at ON proposal (created_at DESC);

CREATE TABLE IF NOT EXISTS proposal_vote (
    id              BIGSERIAL PRIMARY KEY,
    proposal_id     BIGINT NOT NULL REFERENCES proposal(id) ON DELETE CASCADE,
    sns_name        VARCHAR(64),
    voter_wallet    VARCHAR(42) NOT NULL,
    vote_result     VARCHAR(128) NOT NULL,
    vote_weight     INTEGER NOT NULL DEFAULT 1,
    voted_at        TIMESTAMPTZ NOT NULL,
    legacy_id       BIGINT,
    UNIQUE (legacy_id)
);

CREATE INDEX IF NOT EXISTS idx_proposal_vote_proposal ON proposal_vote (proposal_id);
CREATE INDEX IF NOT EXISTS idx_proposal_vote_sns ON proposal_vote (sns_name);
CREATE INDEX IF NOT EXISTS idx_proposal_vote_wallet ON proposal_vote (voter_wallet);

CREATE UNIQUE INDEX IF NOT EXISTS idx_proposal_vote_proposal_sns_unique
    ON proposal_vote (proposal_id, sns_name)
    WHERE sns_name IS NOT NULL AND sns_name <> '';

CREATE TABLE IF NOT EXISTS proposal_comment (
    id              BIGSERIAL PRIMARY KEY,
    proposal_id     BIGINT NOT NULL REFERENCES proposal(id) ON DELETE CASCADE,
    sns_name        VARCHAR(64),
    author_wallet   VARCHAR(42) NOT NULL,
    parent_id       BIGINT REFERENCES proposal_comment(id),
    content         TEXT NOT NULL,
    is_reject       BOOLEAN NOT NULL DEFAULT FALSE,
    commented_at    TIMESTAMPTZ NOT NULL,
    legacy_id       BIGINT,
    UNIQUE (legacy_id)
);

CREATE INDEX IF NOT EXISTS idx_proposal_comment_proposal ON proposal_comment (proposal_id);
CREATE INDEX IF NOT EXISTS idx_proposal_comment_parent ON proposal_comment (parent_id);

-- Upgrade path: rename legacy *_v2 tables after pg_dump import
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'proposal_v2')
       AND NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'proposal') THEN
        ALTER TABLE proposal_v2 RENAME TO proposal;
    END IF;
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'proposal_vote_v2')
       AND NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'proposal_vote') THEN
        ALTER TABLE proposal_vote_v2 RENAME TO proposal_vote;
    END IF;
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'proposal_comment_v2')
       AND NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'proposal_comment') THEN
        ALTER TABLE proposal_comment_v2 RENAME TO proposal_comment;
    END IF;
END $$;

ALTER TABLE proposal ADD COLUMN IF NOT EXISTS resolution_rule VARCHAR(32) NOT NULL DEFAULT 'standard';
ALTER TABLE proposal ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

-- rename legacy migrated_at if present
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

CREATE UNIQUE INDEX IF NOT EXISTS idx_proposal_vote_proposal_sns_unique
    ON proposal_vote (proposal_id, sns_name)
    WHERE sns_name IS NOT NULL AND sns_name <> '';
