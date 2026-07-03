# 支付与课程接口说明（v2.6）

本文档描述 v2.6 新增的**会员优惠券**与**课程售卖**接口。会员续费基础流程见 [API_TEST.md](API_TEST.md) 第七节；本地 curl 见同文档第七、八节。

**Base URL：** `https://api.seedao.xyz`（生产） / `http://127.0.0.1:8080`（本地）

**统一响应：**

```json
{ "code": 0, "message": "ok", "data": { } }
```

**鉴权：** 带 `Authorization: Bearer <JWT>` 表示已登录；标注「JWT 可选」时，不传 Token 也可调用。

**金额单位：** 所有 `*_amount` 字段均为 **分（CNY fen）**，前端展示需自行 `/100`。

---

## 一、会员优惠券

### 1.1 券码试算

`POST /v1/payments/coupons/preview` · **JWT 可选**

| 场景 | 行为 |
|------|------|
| 未登录 | 校验券有效性 + 计算折扣金额；**不**校验「每人限次」 |
| 已登录 | 在上述基础上，额外校验该用户是否超过 `max_uses_per_user` |

**请求：**

```json
{
  "product_type": "membership_renew",
  "days": 365,
  "coupon_code": "SAVE10"
}
```

| 字段 | 必填 | 说明 |
|------|------|------|
| `coupon_code` | 是 | 券码（大小写不敏感，服务端转大写） |
| `product_type` | 否 | 默认 `membership_renew` |
| `days` | 否 | 续费天数，默认 `365`；影响原价计算 |

**成功响应 `data`：**

```json
{
  "coupon_code": "SAVE10",
  "original_amount": 9900,
  "discount_amount": 1000,
  "final_amount": 8900,
  "currency": "CNY"
}
```

**常见错误码：** `40010` 券无效 · `40011` 不适用 · `40012` 次数超限 · `40013` 优惠后金额过低

---

### 1.2 创建会员订单（含券）

`POST /v1/payments/orders` · **JWT 必填** · `[CHANGED]`

**仅接受** `product_type = membership_renew`（可省略，默认即此值）。传其它类型返回 `40001`。

**请求：**

```json
{
  "product_type": "membership_renew",
  "days": 365,
  "pay_channel": "wechat_native",
  "coupon_code": "SAVE10"
}
```

| 字段 | 必填 | 说明 |
|------|------|------|
| `pay_channel` | 否 | `wechat_native`（默认）或 `wechat_h5` |
| `coupon_code` | 否 | 使用时服务端重新验券并**锁定**券额，关单释放、支付成功核销 |

**成功响应 `data`（节选）：**

```json
{
  "order_no": "M202606271430001234",
  "product_type": "membership_renew",
  "product_days": 365,
  "original_amount": 9900,
  "discount_amount": 1000,
  "coupon_code": "SAVE10",
  "amount": 8900,
  "status": "pending",
  "code_url": "weixin://wxpay/bizpayurl?pr=xxxxx",
  "expired_at": "2026-06-27T14:45:00Z"
}
```

其余会员订单接口（列表、查单、关单、微信回调）与 v2.5 相同，见 [API_TEST.md](API_TEST.md)。

---

## 二、课程售卖

### 2.1 业务说明

| 概念 | 说明 |
|------|------|
| 折扣叠加 | **早鸟 + 渠道 + 会员** 可同时生效，服务端计算，前端不传金额 |
| 无主订单 | 未登录下单时不收集个人信息；响应返回 `order_token`，**仅一次**，前端必须保存 |
| 会员折扣 | 仅登录且有效会员可享；未登录 preview 会展示「会员价 xxx，登录后可享」 |
| 履约 | 支付成功后写入 `course_enrollments`；查单返回 `enrollment_status` |

**运营配置：** `course_id`、`sku_id`、折扣规则存于数据库表 `courses` / `course_skus` / `course_discount_rules`（无公开列表 API，前端需与运营约定 ID）。

---

### 2.2 价格试算

`POST /v1/courses/checkout/preview` · **JWT 可选**

**请求：**

```json
{
  "course_id": "course_2026_spring",
  "sku_id": "standard",
  "channel_code": "PARTNER_X",
  "apply_member_discount": true
}
```

| 字段 | 必填 | 说明 |
|------|------|------|
| `course_id` | 是 | 课程 ID |
| `sku_id` | 是 | 票种 ID |
| `channel_code` | 否 | 渠道码；有则尝试匹配 `course_discount_rules` 中 `rule_type=channel` 的规则 |
| `apply_member_discount` | 否 | 默认 `false`；为 `true` 且已登录有效会员时，`final_amount` 含会员价 |

> **与会员券区别：** 课程不使用 `coupons` 表，**没有** `coupon_code`。会员续费券（如 SAVE500）请走 `POST /v1/payments/coupons/preview`；课程渠道优惠传 `channel_code`。

**响应中如何区分折扣类型：** `discounts[].type` 取值 `early_bird`（早鸟）、`channel`（渠道）、`member`（会员）。

#### 三种场景传参示例

| 场景 | `channel_code` | `apply_member_discount` | `Authorization` |
|------|----------------|-------------------------|-----------------|
| 仅早鸟 | 不传 | `false` | 不要 |
| 早鸟 + 渠道 | `"PARTNER_X"` | `false` | 不要 |
| 早鸟 + 渠道 + 会员 | `"PARTNER_X"`（可选） | `true` | **Bearer JWT** |

**1）仅早鸟价**

```bash
curl -s -X POST https://api.seedao.xyz/v1/courses/checkout/preview \
  -H "Content-Type: application/json" \
  -d '{
    "course_id": "course_2026_spring",
    "sku_id": "standard",
    "apply_member_discount": false
  }'
```

早鸟在 `course_discount_rules` 有效期内自动生效，无需传码。预期 `discounts` 含 `type: "early_bird"`。

**2）早鸟 + 渠道价**

```bash
curl -s -X POST https://api.seedao.xyz/v1/courses/checkout/preview \
  -H "Content-Type: application/json" \
  -d '{
    "course_id": "course_2026_spring",
    "sku_id": "standard",
    "channel_code": "PARTNER_X",
    "apply_member_discount": false
  }'
```

`channel_code` 须与库中渠道规则一致。预期 `discounts` 含 `early_bird` 与 `channel`。

**3）早鸟 + 渠道 + 会员价（需登录）**

```bash
curl -s -X POST https://api.seedao.xyz/v1/courses/checkout/preview \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT>" \
  -d '{
    "course_id": "course_2026_spring",
    "sku_id": "standard",
    "channel_code": "PARTNER_X",
    "apply_member_discount": true
  }'
```

须 JWT 且用户会员未过期。预期 `discounts` 可出现 `member` 且 `applied: true`，`final_amount` 为三项叠加价。

创建订单 `POST /v1/courses/orders` 使用相同字段，另加 `pay_channel`（`wechat_native` / `wechat_h5`）。

**成功响应 `data`（未登录示例）：**

```json
{
  "course_id": "course_2026_spring",
  "sku_id": "standard",
  "original_amount": 99900,
  "discounts": [
    { "type": "early_bird", "label": "早鸟价", "amount": 20000, "applied": true },
    { "type": "channel", "label": "渠道 PARTNER_X", "amount": 5000, "applied": true }
  ],
  "final_amount": 74900,
  "amount_if_member": 64900,
  "member_discount": {
    "amount": 10000,
    "applied": false,
    "requires_login": true,
    "hint": "会员价 649.00 元，登录后可享"
  },
  "currency": "CNY"
}
```

| 字段 | 说明 |
|------|------|
| `final_amount` | 当前身份下应付价（未登录 = 早鸟+渠道，不含会员） |
| `amount_if_member` | 若登录且为有效会员时的应付价（三项折扣全叠） |
| `discounts[].applied` | 是否已计入 `final_amount` |

**常见错误码：** `40014` 课程不存在 · `40015` 未在售

---

### 2.3 创建课程订单

`POST /v1/courses/orders` · **JWT 条件必填**

| 场景 | JWT |
|------|-----|
| 仅早鸟 / 渠道价 | 不要 |
| `apply_member_discount: true` | **必须** |

**请求：**

```json
{
  "course_id": "course_2026_spring",
  "sku_id": "standard",
  "channel_code": "PARTNER_X",
  "apply_member_discount": false,
  "pay_channel": "wechat_native"
}
```

**成功响应 `data`（无主订单）：**

```json
{
  "order_no": "C20260627143000a1b2",
  "order_token": "ot_xxxxxxxxxxxxxxxx",
  "product_type": "course_purchase",
  "course_id": "course_2026_spring",
  "sku_id": "standard",
  "original_amount": 99900,
  "discount_amount": 25000,
  "amount": 74900,
  "status": "pending",
  "code_url": "weixin://wxpay/bizpayurl?pr=xxxxx",
  "expired_at": "2026-06-27T14:45:00Z"
}
```

> **`order_token` 只在本响应出现一次。** 建议写入 `localStorage` 或展示「复制订单凭证」。丢失后无法自助查单。

已登录且非无主单：响应无 `order_token`，后续用 JWT 查单/关单。

**常见错误码：** `40100` 会员折扣需登录 · `40016` 同上（业务码）

---

### 2.4 查询课程订单

`GET /v1/courses/orders/:order_no` · **JWT 或 token**

| 身份 | 调用方式 |
|------|----------|
| 已登录且订单已绑定用户 | `Authorization: Bearer <JWT>` |
| 无主订单 / 未认领 | `GET ...?token=<ORDER_TOKEN>` |

**成功响应 `data`（支付后节选）：**

```json
{
  "order_no": "C20260627143000a1b2",
  "status": "paid",
  "amount": 74900,
  "paid_at": "2026-06-27T14:32:00Z",
  "enrollment_status": "active"
}
```

`enrollment_status`：支付成功且履约后为 `active`；未支付为空。

**常见错误码：** `40017` 缺少 token · `40018` token 无效 · `40007` 订单不存在

---

### 2.5 关闭未支付订单

`POST /v1/courses/orders/:order_no/close`

| 身份 | 调用方式 |
|------|----------|
| JWT 拥有者 | Header 带 JWT |
| 无主订单 | Body：`{"order_token":"<ORDER_TOKEN>"}` |

---

### 2.6 认领无主订单

`POST /v1/courses/orders/claim` · **JWT 必填**

用户登录后，将此前无主订单绑定到当前账号。

**请求：**

```json
{
  "order_no": "C20260627143000a1b2",
  "order_token": "ot_xxxxxxxxxxxxxxxx"
}
```

认领成功后，该订单可用 JWT 查单，无需再带 token。

**常见错误码：** `40019` 订单已被他人认领

---

### 2.7 模拟支付（仅 dev_mode）

`POST /v1/courses/dev/orders/:order_no/simulate-pay`

本地 `wechat_pay.dev_mode: true` 时可用。鉴权方式同查单（JWT 或 body/query 中的 `order_token`）。

---

## 三、微信支付回调（共用）

`POST /v1/payments/wechat/notify` · 微信服务器调用 · 无 JWT

| `product_type` | 履约 |
|----------------|------|
| `membership_renew` | 延长 `users.membership_expires_at`；核销优惠券 |
| `course_purchase` | 写入 `course_enrollments` |

生产 Native 回调地址示例：`https://api.seedao.xyz/v1/payments/wechat/notify`

---

## 四、错误码速查（v2.6 新增）

| code | 含义 |
|------|------|
| 40010 | 优惠券无效 |
| 40011 | 优惠券不适用 |
| 40012 | 优惠券使用次数超限 |
| 40013 | 优惠后金额过低（低于 1 分） |
| 40014 | 课程 / 票种不存在 |
| 40015 | 课程未在售 |
| 40016 | 会员折扣需登录 |
| 40017 | 课程订单 token 必填 |
| 40018 | 课程订单 token 无效 |
| 40019 | 课程订单已被认领 |

---

## 五、前端对接建议

### 会员续费页

1. 用户输入券码 → 调 `coupons/preview`（可选带 JWT）
2. 展示 `final_amount` → 调 `payments/orders` 创建订单
3. 展示二维码 / 跳转 H5 → 轮询 `payments/orders/:order_no`

### 课程购买页

1. 进入页 / 切换渠道码 → `courses/checkout/preview`
2. 未登录展示 `member_discount.hint` 与「登录享会员价」
3. 下单 → `courses/orders`；**立即保存 `order_token`**
4. 支付后轮询 `courses/orders/:order_no?token=...`
5. 用户稍后注册登录 → 可选调 `courses/orders/claim` 绑定订单
