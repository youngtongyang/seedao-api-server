-- Preview/QA test accounts that share production DB seedao_api.
-- users.is_test = true is excluded from GET /v1/stats/user-count.
-- Login on preview: aliyun_sms.dev_mode must be true; verification code is 000000.

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS is_test BOOLEAN NOT NULL DEFAULT FALSE;

CREATE INDEX IF NOT EXISTS idx_users_is_test ON users (is_test) WHERE is_test = TRUE;

-- previewqa01–05 / +8619900000001–0005, all pending (准会员).
-- previewqa03 is under_review so frontend can cover both application states.

WITH ins AS (
    INSERT INTO users (
        register_source, status, avatar, bio,
        membership_tier, application_status, application_submitted_at,
        membership_expires_at, is_test, created_at, updated_at
    )
    SELECT 'new_sns', 'active', '', '',
           'pending', 'not_submitted', NULL,
           NOW() + INTERVAL '365 days', TRUE, NOW(), NOW()
    WHERE NOT EXISTS (SELECT 1 FROM user_phones WHERE phone = '+8619900000001')
    RETURNING id
)
INSERT INTO user_phones (user_id, phone, phone_hash, verified_at, created_at)
SELECT id, '+8619900000001', 'b3a08cfb9d2ab099b0e7fe51449a16543b9808f5d601d9deafaae33c7d6695d6', NOW(), NOW()
FROM ins;

INSERT INTO sns_names (name, user_id, mint_status, reserved_at, created_at)
SELECT 'previewqa01.seedao', user_id, 'reserved', NOW(), NOW()
FROM user_phones
WHERE phone = '+8619900000001'
  AND NOT EXISTS (SELECT 1 FROM sns_names WHERE LOWER(name) = 'previewqa01.seedao');

WITH ins AS (
    INSERT INTO users (
        register_source, status, avatar, bio,
        membership_tier, application_status, application_submitted_at,
        membership_expires_at, is_test, created_at, updated_at
    )
    SELECT 'new_sns', 'active', '', '',
           'pending', 'not_submitted', NULL,
           NOW() + INTERVAL '365 days', TRUE, NOW(), NOW()
    WHERE NOT EXISTS (SELECT 1 FROM user_phones WHERE phone = '+8619900000002')
    RETURNING id
)
INSERT INTO user_phones (user_id, phone, phone_hash, verified_at, created_at)
SELECT id, '+8619900000002', 'e643d5218085d8c77fff0751f5877cd95ba555335f08866a92e203ef76ffe497', NOW(), NOW()
FROM ins;

INSERT INTO sns_names (name, user_id, mint_status, reserved_at, created_at)
SELECT 'previewqa02.seedao', user_id, 'reserved', NOW(), NOW()
FROM user_phones
WHERE phone = '+8619900000002'
  AND NOT EXISTS (SELECT 1 FROM sns_names WHERE LOWER(name) = 'previewqa02.seedao');

WITH ins AS (
    INSERT INTO users (
        register_source, status, avatar, bio,
        membership_tier, application_status, application_submitted_at,
        membership_expires_at, is_test, created_at, updated_at
    )
    SELECT 'new_sns', 'active', '', '',
           'pending', 'under_review', NOW(),
           NOW() + INTERVAL '365 days', TRUE, NOW(), NOW()
    WHERE NOT EXISTS (SELECT 1 FROM user_phones WHERE phone = '+8619900000003')
    RETURNING id
)
INSERT INTO user_phones (user_id, phone, phone_hash, verified_at, created_at)
SELECT id, '+8619900000003', 'fe71c415b0cdfc4da64965dfef47c226af5552b13837af54c9da4ca9ef0d6cf9', NOW(), NOW()
FROM ins;

INSERT INTO sns_names (name, user_id, mint_status, reserved_at, created_at)
SELECT 'previewqa03.seedao', user_id, 'reserved', NOW(), NOW()
FROM user_phones
WHERE phone = '+8619900000003'
  AND NOT EXISTS (SELECT 1 FROM sns_names WHERE LOWER(name) = 'previewqa03.seedao');

WITH ins AS (
    INSERT INTO users (
        register_source, status, avatar, bio,
        membership_tier, application_status, application_submitted_at,
        membership_expires_at, is_test, created_at, updated_at
    )
    SELECT 'new_sns', 'active', '', '',
           'pending', 'not_submitted', NULL,
           NOW() + INTERVAL '365 days', TRUE, NOW(), NOW()
    WHERE NOT EXISTS (SELECT 1 FROM user_phones WHERE phone = '+8619900000004')
    RETURNING id
)
INSERT INTO user_phones (user_id, phone, phone_hash, verified_at, created_at)
SELECT id, '+8619900000004', 'a38cddf61a01a9f562ae81767cf7fdeaad58d24c021ad48e40ab6fc35b1aaa2f', NOW(), NOW()
FROM ins;

INSERT INTO sns_names (name, user_id, mint_status, reserved_at, created_at)
SELECT 'previewqa04.seedao', user_id, 'reserved', NOW(), NOW()
FROM user_phones
WHERE phone = '+8619900000004'
  AND NOT EXISTS (SELECT 1 FROM sns_names WHERE LOWER(name) = 'previewqa04.seedao');

WITH ins AS (
    INSERT INTO users (
        register_source, status, avatar, bio,
        membership_tier, application_status, application_submitted_at,
        membership_expires_at, is_test, created_at, updated_at
    )
    SELECT 'new_sns', 'active', '', '',
           'pending', 'not_submitted', NULL,
           NOW() + INTERVAL '365 days', TRUE, NOW(), NOW()
    WHERE NOT EXISTS (SELECT 1 FROM user_phones WHERE phone = '+8619900000005')
    RETURNING id
)
INSERT INTO user_phones (user_id, phone, phone_hash, verified_at, created_at)
SELECT id, '+8619900000005', 'c96b9303d9c8f7e7c468c2a1a32f3b4c97e9a42c755bcd92abacaba99cde3c25', NOW(), NOW()
FROM ins;

INSERT INTO sns_names (name, user_id, mint_status, reserved_at, created_at)
SELECT 'previewqa05.seedao', user_id, 'reserved', NOW(), NOW()
FROM user_phones
WHERE phone = '+8619900000005'
  AND NOT EXISTS (SELECT 1 FROM sns_names WHERE LOWER(name) = 'previewqa05.seedao');
