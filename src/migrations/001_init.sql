-- seedao-api-server initial schema (v2.4)

CREATE TABLE IF NOT EXISTS users (
    id              BIGSERIAL PRIMARY KEY,
    register_source VARCHAR(32)  NOT NULL,
    status          VARCHAR(16)  NOT NULL DEFAULT 'active',
    avatar          VARCHAR(512) NOT NULL DEFAULT '',
    bio             TEXT         NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);

CREATE TABLE IF NOT EXISTS user_phones (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT       NOT NULL UNIQUE REFERENCES users (id),
    phone       VARCHAR(20)  NOT NULL UNIQUE,
    phone_hash  VARCHAR(64)  NOT NULL,
    verified_at TIMESTAMPTZ  NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_phones_phone_hash ON user_phones (phone_hash);

CREATE TABLE IF NOT EXISTS sns_names (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(64)  NOT NULL UNIQUE,
    user_id     BIGINT       NOT NULL UNIQUE REFERENCES users (id),
    mint_status VARCHAR(16)  NOT NULL,
    wallet      VARCHAR(42)  UNIQUE,
    reserved_at TIMESTAMPTZ,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_sns_names_lower ON sns_names (LOWER(name));
CREATE UNIQUE INDEX IF NOT EXISTS idx_sns_names_wallet ON sns_names (wallet) WHERE wallet IS NOT NULL;

CREATE TABLE IF NOT EXISTS sns_chain_registry (
    id         BIGSERIAL PRIMARY KEY,
    wallet     VARCHAR(42)  NOT NULL UNIQUE,
    sns_name   VARCHAR(64)  NOT NULL UNIQUE,
    claimed    BOOLEAN      NOT NULL DEFAULT FALSE,
    synced_at  TIMESTAMPTZ  NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_sns_chain_registry_claimed ON sns_chain_registry (claimed) WHERE claimed = FALSE;

CREATE TABLE IF NOT EXISTS sms_send_logs (
    id         BIGSERIAL PRIMARY KEY,
    phone_hash VARCHAR(64)  NOT NULL,
    ip         VARCHAR(45),
    status     VARCHAR(16)  NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_sms_send_logs_phone_hash ON sms_send_logs (phone_hash);
CREATE INDEX IF NOT EXISTS idx_sms_send_logs_created_at ON sms_send_logs (created_at);
