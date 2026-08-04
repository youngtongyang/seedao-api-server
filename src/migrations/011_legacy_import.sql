-- Legacy 1.0 data mirror tables (profiles / assets / roles / seeds / sbts / rewards)
-- Import via: go run ./cmd/v1import <profiles|rewards|oss>

CREATE TABLE IF NOT EXISTS user_profiles (
    sns_name         VARCHAR(64) PRIMARY KEY,
    bio              TEXT         NOT NULL DEFAULT '',
    email            VARCHAR(255) NOT NULL DEFAULT '',
    twitter_profile  VARCHAR(512) NOT NULL DEFAULT '',
    mirror           VARCHAR(512) NOT NULL DEFAULT '',
    scr_amount       NUMERIC(78, 30),
    created_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_assets (
    id                  BIGSERIAL PRIMARY KEY,
    sns_name            VARCHAR(64)    NOT NULL,
    asset_name          VARCHAR(32)    NOT NULL,
    dealt_amount        NUMERIC(78, 30) NOT NULL DEFAULT 0,
    processing_amount   NUMERIC(78, 30) NOT NULL DEFAULT 0,
    claimed             BOOLEAN        NOT NULL DEFAULT FALSE,
    source_create_ts    BIGINT,
    created_at          TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    UNIQUE (sns_name, asset_name)
);

CREATE INDEX IF NOT EXISTS idx_user_assets_sns_name ON user_assets (sns_name);

CREATE TABLE IF NOT EXISTS user_roles (
    sns_name    VARCHAR(64)  NOT NULL,
    role        VARCHAR(64)  NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    PRIMARY KEY (sns_name, role)
);

CREATE TABLE IF NOT EXISTS user_seeds (
    id              BIGSERIAL PRIMARY KEY,
    sns_name        VARCHAR(64)  NOT NULL,
    token_id        VARCHAR(64)  NOT NULL,
    contract_addr   VARCHAR(42)  NOT NULL,
    contract_type   VARCHAR(16)  NOT NULL DEFAULT '',
    token_amount    VARCHAR(32)  NOT NULL DEFAULT '',
    image_uri_src   TEXT         NOT NULL DEFAULT '',
    image_uri_oss   VARCHAR(1024) NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    UNIQUE (sns_name, contract_addr, token_id)
);

CREATE INDEX IF NOT EXISTS idx_user_seeds_sns_name ON user_seeds (sns_name);
CREATE INDEX IF NOT EXISTS idx_user_seeds_oss_pending
    ON user_seeds (id)
    WHERE image_uri_src <> '' AND image_uri_oss = '';

CREATE TABLE IF NOT EXISTS user_sbts (
    id               BIGSERIAL PRIMARY KEY,
    sns_name         VARCHAR(64)   NOT NULL,
    token_id         VARCHAR(64)   NOT NULL,
    contract_addr    VARCHAR(42)   NOT NULL,
    contract_type    VARCHAR(16)   NOT NULL DEFAULT '',
    token_amount     VARCHAR(32)   NOT NULL DEFAULT '',
    collection_name  VARCHAR(255)  NOT NULL DEFAULT '',
    display_name     VARCHAR(512)  NOT NULL DEFAULT '',
    description      TEXT          NOT NULL DEFAULT '',
    image_uri_src    TEXT          NOT NULL DEFAULT '',
    image_uri_oss    VARCHAR(1024) NOT NULL DEFAULT '',
    created_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (sns_name, contract_addr, token_id)
);

CREATE INDEX IF NOT EXISTS idx_user_sbts_sns_name ON user_sbts (sns_name);
CREATE INDEX IF NOT EXISTS idx_user_sbts_oss_pending
    ON user_sbts (id)
    WHERE image_uri_src <> '' AND image_uri_oss = '';

CREATE TABLE IF NOT EXISTS reward_applications (
    application_id       BIGINT PRIMARY KEY,
    season_name          VARCHAR(32)   NOT NULL DEFAULT '',
    entity_name          VARCHAR(512)  NOT NULL DEFAULT '',
    entity_type          VARCHAR(64)   NOT NULL DEFAULT '',
    asset_name           VARCHAR(32)   NOT NULL DEFAULT '',
    amount               NUMERIC(78, 30),
    budget_source        VARCHAR(512)  NOT NULL DEFAULT '',
    status               VARCHAR(32)   NOT NULL DEFAULT '',
    target_user_wallet   VARCHAR(42)   NOT NULL DEFAULT '',
    target_sns_name      VARCHAR(64),
    applicant_wallet     VARCHAR(42)   NOT NULL DEFAULT '',
    applicant_sns_name   VARCHAR(64),
    source_create_ts     BIGINT,
    source_update_ts     BIGINT,
    created_at           TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_reward_applications_target_sns
    ON reward_applications (target_sns_name);
CREATE INDEX IF NOT EXISTS idx_reward_applications_applicant_sns
    ON reward_applications (applicant_sns_name);
CREATE INDEX IF NOT EXISTS idx_reward_applications_season_asset
    ON reward_applications (season_name, asset_name);
CREATE INDEX IF NOT EXISTS idx_reward_applications_status
    ON reward_applications (status);

CREATE TABLE IF NOT EXISTS oss_objects (
    src_url_hash  CHAR(64) PRIMARY KEY,
    src_url       TEXT         NOT NULL,
    oss_url       VARCHAR(1024) NOT NULL,
    content_type  VARCHAR(128) NOT NULL DEFAULT '',
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
