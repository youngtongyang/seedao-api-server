-- coupons for membership payment discounts

CREATE TABLE IF NOT EXISTS coupons (
    id                  BIGSERIAL PRIMARY KEY,
    code                VARCHAR(64)  NOT NULL UNIQUE,
    discount_type       VARCHAR(16)  NOT NULL DEFAULT 'fixed',
    discount_value      INT          NOT NULL,
    product_type        VARCHAR(32),
    min_amount_fen      INT          NOT NULL DEFAULT 0,
    max_total_uses      INT,
    used_count          INT          NOT NULL DEFAULT 0,
    max_uses_per_user   INT          NOT NULL DEFAULT 1,
    valid_from          TIMESTAMPTZ  NOT NULL,
    valid_until         TIMESTAMPTZ  NOT NULL,
    status              VARCHAR(16)  NOT NULL DEFAULT 'active',
    created_at          TIMESTAMPTZ  NOT NULL,
    updated_at          TIMESTAMPTZ  NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_coupons_status ON coupons (status);
CREATE INDEX IF NOT EXISTS idx_coupons_valid_until ON coupons (valid_until);

CREATE TABLE IF NOT EXISTS coupon_redemptions (
    id              BIGSERIAL PRIMARY KEY,
    coupon_id       BIGINT       NOT NULL REFERENCES coupons(id),
    user_id         BIGINT       NOT NULL REFERENCES users(id),
    order_id        BIGINT       REFERENCES payment_orders(id),
    order_no        VARCHAR(64)  NOT NULL,
    discount_amount INT          NOT NULL,
    status          VARCHAR(16)  NOT NULL DEFAULT 'reserved',
    created_at      TIMESTAMPTZ  NOT NULL,
    updated_at      TIMESTAMPTZ  NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_coupon_redemptions_coupon_id ON coupon_redemptions (coupon_id);
CREATE INDEX IF NOT EXISTS idx_coupon_redemptions_user_id ON coupon_redemptions (user_id);
CREATE INDEX IF NOT EXISTS idx_coupon_redemptions_order_id ON coupon_redemptions (order_id);
CREATE INDEX IF NOT EXISTS idx_coupon_redemptions_order_no ON coupon_redemptions (order_no);
CREATE INDEX IF NOT EXISTS idx_coupon_redemptions_status ON coupon_redemptions (status);

ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS original_amount INT NOT NULL DEFAULT 0;
ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS discount_amount INT NOT NULL DEFAULT 0;
ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS coupon_code VARCHAR(64);

UPDATE payment_orders SET original_amount = amount WHERE original_amount = 0;
