-- course sales, anonymous orders, and extended payment_orders

ALTER TABLE payment_orders ALTER COLUMN user_id DROP NOT NULL;

ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS product_ref TEXT;
ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS discount_snapshot JSONB;
ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS channel_code VARCHAR(64);
ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS order_token_hash VARCHAR(64);
ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS is_anonymous BOOLEAN NOT NULL DEFAULT FALSE;

CREATE INDEX IF NOT EXISTS idx_payment_orders_product_type ON payment_orders (product_type);
CREATE INDEX IF NOT EXISTS idx_payment_orders_order_token_hash ON payment_orders (order_token_hash);

CREATE TABLE IF NOT EXISTS courses (
    id          VARCHAR(64)  PRIMARY KEY,
    title       VARCHAR(256) NOT NULL,
    status      VARCHAR(16)  NOT NULL DEFAULT 'on_sale',
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL
);

CREATE TABLE IF NOT EXISTS course_skus (
    id              BIGSERIAL PRIMARY KEY,
    course_id       VARCHAR(64)  NOT NULL REFERENCES courses(id),
    sku_id          VARCHAR(64)  NOT NULL,
    name            VARCHAR(128) NOT NULL,
    base_price_fen  INT          NOT NULL,
    status          VARCHAR(16)  NOT NULL DEFAULT 'on_sale',
    created_at      TIMESTAMPTZ  NOT NULL,
    updated_at      TIMESTAMPTZ  NOT NULL,
    UNIQUE (course_id, sku_id)
);

CREATE INDEX IF NOT EXISTS idx_course_skus_course_id ON course_skus (course_id);

CREATE TABLE IF NOT EXISTS course_discount_rules (
    id              BIGSERIAL PRIMARY KEY,
    course_id       VARCHAR(64)  NOT NULL,
    sku_id          VARCHAR(64)  NOT NULL,
    rule_type       VARCHAR(16)  NOT NULL,
    discount_type   VARCHAR(16)  NOT NULL DEFAULT 'fixed',
    discount_value  INT          NOT NULL,
    channel_code    VARCHAR(64),
    valid_from      TIMESTAMPTZ,
    valid_until     TIMESTAMPTZ,
    status          VARCHAR(16)  NOT NULL DEFAULT 'active',
    created_at      TIMESTAMPTZ  NOT NULL,
    updated_at      TIMESTAMPTZ  NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_course_discount_rules_lookup
    ON course_discount_rules (course_id, sku_id, rule_type, status);

CREATE TABLE IF NOT EXISTS course_enrollments (
    id          BIGSERIAL PRIMARY KEY,
    order_id    BIGINT       NOT NULL UNIQUE REFERENCES payment_orders(id),
    user_id     BIGINT       REFERENCES users(id),
    course_id   VARCHAR(64)  NOT NULL,
    sku_id      VARCHAR(64)  NOT NULL,
    status      VARCHAR(16)  NOT NULL DEFAULT 'active',
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_course_enrollments_user_id ON course_enrollments (user_id);
CREATE INDEX IF NOT EXISTS idx_course_enrollments_course_id ON course_enrollments (course_id);
