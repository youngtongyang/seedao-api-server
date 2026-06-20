-- user membership expiry

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS membership_expires_at TIMESTAMPTZ;

CREATE INDEX IF NOT EXISTS idx_users_membership_expires_at ON users (membership_expires_at);
