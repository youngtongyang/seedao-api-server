-- Feishu / third-party OAuth bindings (bot token flow)

CREATE TABLE IF NOT EXISTS user_oauth_bindings (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT       NOT NULL REFERENCES users (id),
    provider        VARCHAR(32)  NOT NULL,
    provider_app_id VARCHAR(64)  NOT NULL,
    open_id         VARCHAR(128) NOT NULL,
    union_id        VARCHAR(128),
    display_name    VARCHAR(128) NOT NULL DEFAULT '',
    bind_source     VARCHAR(32)  NOT NULL,
    bound_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_user_oauth_bindings_user_id
    ON user_oauth_bindings (user_id);

CREATE INDEX IF NOT EXISTS idx_user_oauth_bindings_deleted_at
    ON user_oauth_bindings (deleted_at);

CREATE UNIQUE INDEX IF NOT EXISTS uq_user_oauth_bindings_user_provider_active
    ON user_oauth_bindings (user_id, provider)
    WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS uq_user_oauth_bindings_provider_openid_active
    ON user_oauth_bindings (provider, provider_app_id, open_id)
    WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_user_oauth_bindings_union_id
    ON user_oauth_bindings (provider, union_id)
    WHERE union_id IS NOT NULL AND deleted_at IS NULL;
