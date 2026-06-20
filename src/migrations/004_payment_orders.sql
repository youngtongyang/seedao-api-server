-- payment orders for WeChat Native membership renew

CREATE TABLE IF NOT EXISTS payment_orders (
    id                  BIGSERIAL PRIMARY KEY,
    order_no            VARCHAR(64)  NOT NULL UNIQUE,
    user_id             BIGINT       NOT NULL REFERENCES users(id),
    product_type        VARCHAR(32)  NOT NULL,
    product_days        INT          NOT NULL DEFAULT 365,
    amount              INT          NOT NULL,
    currency            VARCHAR(8)   NOT NULL DEFAULT 'CNY',
    status              VARCHAR(16)  NOT NULL DEFAULT 'pending',
    pay_channel         VARCHAR(16)  NOT NULL DEFAULT 'wechat_native',
    code_url            TEXT,
    wx_transaction_id   VARCHAR(64),
    paid_at             TIMESTAMPTZ,
    expired_at          TIMESTAMPTZ NOT NULL,
    closed_at           TIMESTAMPTZ,
    notify_payload      JSONB,
    created_at          TIMESTAMPTZ NOT NULL,
    updated_at          TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_payment_orders_user_id ON payment_orders (user_id);
CREATE INDEX IF NOT EXISTS idx_payment_orders_status ON payment_orders (status);
CREATE INDEX IF NOT EXISTS idx_payment_orders_wx_transaction_id ON payment_orders (wx_transaction_id);

CREATE TABLE IF NOT EXISTS payment_refunds (
    id              BIGSERIAL PRIMARY KEY,
    order_id        BIGINT       NOT NULL REFERENCES payment_orders(id),
    refund_no       VARCHAR(64)  NOT NULL UNIQUE,
    wx_refund_id    VARCHAR(64),
    amount          INT          NOT NULL,
    reason          VARCHAR(256),
    status          VARCHAR(16)  NOT NULL DEFAULT 'pending',
    notify_payload  JSONB,
    created_at      TIMESTAMPTZ NOT NULL,
    updated_at      TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_payment_refunds_order_id ON payment_refunds (order_id);
