# API 本地测试手册

本文档汇总 SeeDAO API 本地联调用的 `curl` 命令，便于在终端（bash/zsh）中复制执行。

**变更标识：** `[NEW]` 本次新增 · `[CHANGED]` 本次变更

> **v2.6 接口字段与鉴权说明（支付/课程）：** 见 [API_PAYMENTS_v2.6.md](API_PAYMENTS_v2.6.md)

### API 变更一览（v2.6）

| 接口 | 标识 | 说明 |
|------|------|------|
| `POST /v1/payments/coupons/preview` | `[NEW]` | 会员券试算，JWT 可选 |
| `POST /v1/payments/orders` | `[CHANGED]` | 仅 membership + 可选 coupon_code |
| `POST /v1/courses/checkout/preview` | `[NEW]` | 课程价格试算 |
| `POST /v1/courses/orders` | `[NEW]` | 课程下单，无主单返回 order_token |
| `GET/POST /v1/courses/orders/*` | `[NEW]` | 查单、关单、认领 |

### API 变更一览（v2.5，历史）

| 接口 | 标识 | 说明 |
|------|------|------|
| `GET /v1/stats/user-count` | `[NEW]` | 首页用户总数，无需鉴权 |
| `GET /v1/auth/phone/check` | `[NEW]` | 注册前检查手机号未占用，`code=40009` |
| `/v1/payments/*` | `[NEW]` | 会员续费微信支付全流程 |
| `GET /v1/user/profile` | `[CHANGED]` | 增加 `membership_expires_at` |
| `POST /v1/auth/login` / `register` 的 `user` | `[CHANGED]` | 增加 `membership_expires_at` |

**前提：**

- API 已启动：`bash src/scripts/start-local.sh` 或 `go run ./cmd/server`
- 默认地址：`http://127.0.0.1:8080`
- 需已安装 `curl`（macOS / Linux 通常自带；Windows 可用 Git Bash 或 `curl.exe`）

**测试变量（执行前替换）：**

| 变量 | 说明 | 示例 |
|------|------|------|
| `PHONE` | 手机号，必须带国际区号 | `+8613800138000` |
| `CODE` | 短信验证码（6 位，5 分钟有效，一次性） | `301160` |
| `TOKEN` | 注册/登录返回的 JWT | `eyJhbG...` |
| `SNS_NAME` | SNS 名，必须为 `{base}.seedao` | `alice_test_123.seedao` |

> **SNS 命名规范：** 所有接口入参/出参字段统一为 `sns_name`，且必须带 `.seedao` 后缀（与 `sns_chain_registry` 一致）。

> 验证码**只用一次**。注册/登录成功后需重新调用「发送验证码」获取新码。

---

## 一、基础接口

### 1. 健康检查

```bash
curl http://127.0.0.1:8080/health
```

预期：`{"code":0,"message":"ok","data":{"status":"ok"}}`

---

### 2. 首页用户总数 `[NEW]`

```bash
curl http://127.0.0.1:8080/v1/stats/user-count
```

预期：

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "total_user_count": 12345
  }
}
```

说明：无需鉴权；`total_user_count` = 链上存量 SNS + `new_sns` 注册用户。

---

### 3. 发送验证码

```bash
curl -X POST http://127.0.0.1:8080/v1/auth/sms/send \
  -H "Content-Type: application/json" \
  -d '{"phone":"+8613800138000"}'
```

说明：

- 同号 60 秒内不可重复发送
- 同 IP 每小时最多 20 次
- 成功响应 `data.retry_after` 为 60；**429 限流时同样返回 `data.retry_after`（剩余秒数）**，前端应展示倒计时而非「发送失败」
- `dev_mode: true` 时验证码打在 API 日志；`false` 时真实发短信

429 示例：

```json
{
  "code": 42901,
  "message": "too many requests",
  "data": { "retry_after": 47 }
}
```

---

### 4. 检查手机号未注册 `[NEW]`

注册前调用，**不消耗验证码**。

```bash
curl "http://127.0.0.1:8080/v1/auth/phone/check?phone=%2B8613800138000"
```

未注册预期：

```json
{
  "code": 40009,
  "message": "phone not registered",
  "data": {
    "phone": "+8613800138000",
    "registered": false
  }
}
```

已注册预期：`code=40006`, `message=phone already registered`

格式非法（未带 `+86`）预期：`40001`

---

## 二、SNS 名称检查（不消耗验证码）

### 4. 查询可用名称

```bash
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=alice_test_123.seedao"
```

预期：`available: true`，响应中 `sns_name` 为 `alice_test_123.seedao`

### 5. 查询已被 registry 占用的名称

```bash
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=baiyu.seedao"
```

预期：`available: false`, `reason: taken_in_registry`

### 6. 查询格式非法的名称

```bash
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=AB.seedao"
```

预期：`available: false`, `reason: invalid_format`

---

## 三、新用户注册（new_sns）

建议顺序：**先 SNS check → 再 register**

### 7. 新用户注册

```bash
curl -X POST http://127.0.0.1:8080/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "+8613800138000",
    "code": "301160",
    "register_type": "new_sns",
    "sns_name": "binbin.seedao"
  }'
```

成功响应示例：

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "token": "eyJhbG...",
    "token_exp": 1718438400,
    "user": {
      "id": 1001,
      "sns_name": "alice_test_123.seedao",
      "register_source": "new_sns",
      "membership_expires_at": "2027-06-15T08:00:00Z"
    }
  }
}
```

常见错误：

| code | 含义 |
|------|------|
| 40002 | 验证码错误/过期/已使用 |
| 40003 | SNS 名称不可用 |
| 40006 | 手机号已注册（改走登录） |

> **注意：** 若 SNS 名不可用（40003），验证码不会被消耗，可换名后重试。  
> `sns_name` 必须带 `.seedao` 后缀，不带后缀会返回 `invalid_format`（40001）。

---

## 四、登录 `[CHANGED]`

已注册用户需**重新发送验证码**后再登录。

### 8. 登录

```bash
curl -X POST http://127.0.0.1:8080/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "+8613800138000",
    "code": "301160"
  }'
```

响应格式同注册，返回 `token` 与 `user`（含 `membership_expires_at`）。`[CHANGED]`

---

## 五、个人 Profile（需 JWT）`[CHANGED]`

### 9. 查询 Profile

将 `<TOKEN>` 替换为注册/登录返回的 token：

```bash
curl http://127.0.0.1:8080/v1/user/profile \
  -H "Authorization: Bearer <TOKEN>"
```

预期：返回 `id`、`sns_name`、脱敏 `phone`、`register_source`、`membership_expires_at` 等。`[CHANGED]`

### 10. 无 Token 访问（应失败）

```bash
curl http://127.0.0.1:8080/v1/user/profile
```

预期：`401 Unauthorized`

---

## 六、老用户注册（existing_sns + 钱包 SIWE）

适用于 `sns_chain_registry` 中已有 SNS、且你**持有对应钱包私钥**的场景。

示例数据（来自导入快照）：

- 钱包：`0x78f625a65fc316d32d98d249b698fb509a6d98f2`
- SNS：`tanghan.seedao`

### 11. 获取钱包 Nonce / SIWE Message

```bash
curl -X POST http://127.0.0.1:8080/v1/wallet/nonce \
  -H "Content-Type: application/json" \
  -d '{
    "wallet": "0x78f625a65fc316d32d98d249b698fb509a6d98f2",
    "domain": "app.seedao.top"
  }'
```

返回 `nonce`、`message`、`expires_at`。用 MetaMask 等对 `message` 签名，得到 `signature`。

### 12. 验证 SNS 归属

```bash
curl -X POST http://127.0.0.1:8080/v1/wallet/verify-sns \
  -H "Content-Type: application/json" \
  -d '{
    "wallet": "0x78f625a65fc316d32d98d249b698fb509a6d98f2",
    "wallet_type": "EOA",
    "domain": "app.seedao.top",
    "message": "上一步返回的完整 message",
    "signature": "0x钱包签名"
  }'
```

成功预期：

```json
{
  "code": 0,
  "data": {
    "wallet": "0x78f625a65fc316d32d98d249b698fb509a6d98f2",
    "sns_name": "tanghan.seedao",
    "claimable": true
  }
}
```

假签名预期：`40004` wallet signature verification failed

### 13. 老用户注册

需：**新手机验证码** + **有效钱包签名**（与 verify-sns 相同 message/signature）

```bash
curl -X POST http://127.0.0.1:8080/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "+8613800138000",
    "code": "301160",
    "register_type": "existing_sns",
    "wallet": "0x78f625a65fc316d32d98d249b698fb509a6d98f2",
    "wallet_type": "EOA",
    "domain": "app.seedao.top",
    "message": "与 verify-sns 相同的 message",
    "signature": "0x与 verify-sns 相同的签名"
  }'
```

---

## 七、会员微信支付

> 接口说明：[API_PAYMENTS_v2.6.md § 一、会员优惠券](API_PAYMENTS_v2.6.md#一会员优惠券)  
> 会员订单创建/查单/关单：`JWT` 必填；券码试算：`JWT` 可选。

支持两种支付渠道，通过创建订单时的 `pay_channel` 区分：

| pay_channel | 适用场景 | 前端行为 |
|-------------|----------|----------|
| `wechat_native`（默认） | PC 网页 | 用 `code_url` 渲染二维码，用户微信扫码 |
| `wechat_h5` | 手机 PWA / 移动浏览器 | 跳转 `h5_url`，唤起微信 App 支付 |

本地默认 `wechat_pay.dev_mode: true`（未配置商户号时自动开启），可用模拟支付接口联调；生产环境需配置真实微信商户参数并将 `dev_mode` 设为 `false`。

**测试变量补充：**

| 变量 | 说明 |
|------|------|
| `ORDER_NO` | 创建订单返回的 `order_no` |
| `ORDER_TOKEN` | 无主课程订单创建时返回，仅一次 |
| `COURSE_ID` / `SKU_ID` | 课程与票种 ID（见库表配置） |

**价格说明：** 默认年费 `9900` 分（99 元），由服务端 `wechat_pay.membership_price_fen` 配置，前端不传金额。

**前端选渠道示例：**

```javascript
const payChannel = /Mobile|Android|iPhone/i.test(navigator.userAgent)
  ? 'wechat_h5'
  : 'wechat_native'
```

### 14. 创建支付订单

用户点击「立即续费」后调用。

#### 14a. PC 扫码（Native，默认）

```bash
curl -X POST http://127.0.0.1:8080/v1/payments/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{
    "product_type": "membership_renew",
    "days": 365,
    "pay_channel": "wechat_native"
  }'
```

`pay_channel` 可省略，默认 `wechat_native`。

成功预期：

```json
{
  "code": 0,
  "data": {
    "order_no": "M202606201430001234",
    "product_type": "membership_renew",
    "product_days": 365,
    "pay_channel": "wechat_native",
    "amount": 9900,
    "currency": "CNY",
    "status": "pending",
    "code_url": "weixin://wxpay/bizpayurl?pr=xxxxx",
    "expired_at": "2026-06-20T14:45:00Z"
  }
}
```

前端用 `code_url` 本地生成二维码图片展示。

#### 14b. 手机 H5（PWA / 移动浏览器）

```bash
curl -X POST http://127.0.0.1:8080/v1/payments/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{
    "product_type": "membership_renew",
    "pay_channel": "wechat_h5"
  }'
```

成功预期：

```json
{
  "code": 0,
  "data": {
    "order_no": "M202606201430001234",
    "pay_channel": "wechat_h5",
    "amount": 9900,
    "status": "pending",
    "h5_url": "https://wx.tenpay.com/cgi-bin/mmpayweb-bin/checkmweb?...",
    "expired_at": "2026-06-20T14:45:00Z"
  }
}
```

前端收到 `h5_url` 后直接 `window.location.href = h5_url` 跳转；支付完成后微信会跳回 `h5_return_url`（服务端配置），同时前端应继续轮询订单状态确认支付结果。

> H5 下单时服务端会自动读取请求来源 IP（`ClientIP`），无需前端传递。

`days` 可省略，默认 `365`。非 365 天时金额按年费比例折算。

### 15. 查询订单状态（前端轮询）

```bash
curl http://127.0.0.1:8080/v1/payments/orders/<ORDER_NO> \
  -H "Authorization: Bearer <TOKEN>"
```

`status=pending` 时响应仍包含 `code_url`（Native）或 `h5_url`（H5），便于用户刷新页面后继续支付。

支付成功后 `status=paid`，并返回 `membership_expires_at`、`paid_at`、`wx_transaction_id`。

### 16. 订单列表

```bash
curl "http://127.0.0.1:8080/v1/payments/orders?limit=20" \
  -H "Authorization: Bearer <TOKEN>"
```

### 17. 关闭未支付订单

用户取消或二维码过期时：

```bash
curl -X POST http://127.0.0.1:8080/v1/payments/orders/<ORDER_NO>/close \
  -H "Authorization: Bearer <TOKEN>"
```

### 18. 微信支付回调（微信服务器调用，无需 JWT）

```
POST https://<你的API域名>/v1/payments/wechat/notify
```

由微信在用户支付成功后异步通知。需公网 HTTPS，并在商户平台配置相同地址。验签通过后自动延长会员。

本地开发可用 ngrok 等暴露该路径；`dev_mode` 下不会收到真实微信回调。

### 19. 模拟支付（仅 dev_mode）

本地联调用，模拟微信回调完成支付：

```bash
curl -X POST http://127.0.0.1:8080/v1/payments/dev/simulate-pay \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"order_no":"<ORDER_NO>"}'
```

成功预期：`status=paid`，并返回更新后的 `membership_expires_at`。

### 会员续费完整联调流程（dev_mode，PC Native）

```bash
# Step 1: 登录获取 TOKEN（见上文注册/登录流程）

# Step 2: 创建 Native 订单
curl -X POST http://127.0.0.1:8080/v1/payments/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"product_type":"membership_renew","pay_channel":"wechat_native"}'

# Step 3: 前端用 code_url 展示二维码；本地可直接模拟支付
curl -X POST http://127.0.0.1:8080/v1/payments/dev/simulate-pay \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"order_no":"<ORDER_NO>"}'

# Step 4: 确认订单与会员到期时间
curl http://127.0.0.1:8080/v1/payments/orders/<ORDER_NO> \
  -H "Authorization: Bearer <TOKEN>"

curl http://127.0.0.1:8080/v1/user/profile \
  -H "Authorization: Bearer <TOKEN>"
```

### 会员续费完整联调流程（dev_mode，手机 H5）

```bash
# Step 1: 登录获取 TOKEN

# Step 2: 创建 H5 订单
curl -X POST http://127.0.0.1:8080/v1/payments/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"product_type":"membership_renew","pay_channel":"wechat_h5"}'

# Step 3: 前端跳转 h5_url；本地 dev_mode 可直接 simulate-pay
curl -X POST http://127.0.0.1:8080/v1/payments/dev/simulate-pay \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"order_no":"<ORDER_NO>"}'

# Step 4: 轮询确认支付结果
curl http://127.0.0.1:8080/v1/payments/orders/<ORDER_NO> \
  -H "Authorization: Bearer <TOKEN>"
```

**生产环境 H5 额外配置（`config.prod.yaml`）：**

```yaml
wechat_pay:
  h5_wap_name: "SeeDAO"
  h5_wap_url: "https://app.seedao.top"          # 须与微信商户平台 H5 支付域名一致
  h5_return_url: "https://app.seedao.top/pay/result"  # 付完跳回的 PWA 页面
```

商户平台还需：**开通 H5 支付** + **配置 H5 支付域名**（需 ICP 备案）。

### 20. 会员券码试算 `[NEW]`（JWT 可选）

```bash
# 未登录：验券 + 试算，不校验每人限次
curl -X POST http://127.0.0.1:8080/v1/payments/coupons/preview \
  -H "Content-Type: application/json" \
  -d '{"coupon_code":"SAVE10","days":365}'

# 已登录：含每人限次校验
curl -X POST http://127.0.0.1:8080/v1/payments/coupons/preview \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"coupon_code":"SAVE10"}'
```

创建会员订单时可带 `coupon_code`：

```bash
curl -X POST http://127.0.0.1:8080/v1/payments/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"product_type":"membership_renew","coupon_code":"SAVE10"}'
```

---

## 八、课程售卖 `[NEW]`

> 接口说明：[API_PAYMENTS_v2.6.md § 二、课程售卖](API_PAYMENTS_v2.6.md#二课程售卖)

**前置：** 库中已有课程配置（`courses` / `course_skus` / `course_discount_rules`）。本地可先执行 `setup-db.sh` 后手动 INSERT，或参考 [OPERATIONS.md § 十二](OPERATIONS.md#十二优惠券与课程数据配置)。

### 21. 价格试算（JWT 可选）

```bash
curl -X POST http://127.0.0.1:8080/v1/courses/checkout/preview \
  -H "Content-Type: application/json" \
  -d '{
    "course_id": "course_2026_spring",
    "sku_id": "standard",
    "channel_code": "PARTNER_X",
    "apply_member_discount": true
  }'
```

未登录成功预期（节选）：

```json
{
  "code": 0,
  "data": {
    "original_amount": 99900,
    "final_amount": 74900,
    "amount_if_member": 64900,
    "member_discount": {
      "amount": 10000,
      "applied": false,
      "requires_login": true,
      "hint": "会员价 649.00 元，登录后可享"
    }
  }
}
```

已登录且有效会员时，若 `apply_member_discount: true`，则 `final_amount` 等于 `amount_if_member`。

### 22. 创建课程订单

**无主（无需 JWT）：**

```bash
curl -X POST http://127.0.0.1:8080/v1/courses/orders \
  -H "Content-Type: application/json" \
  -d '{
    "course_id": "course_2026_spring",
    "sku_id": "standard",
    "channel_code": "PARTNER_X",
    "pay_channel": "wechat_native"
  }'
```

**登录 + 会员价：**

```bash
curl -X POST http://127.0.0.1:8080/v1/courses/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{
    "course_id": "course_2026_spring",
    "sku_id": "standard",
    "channel_code": "PARTNER_X",
    "apply_member_discount": true
  }'
```

成功预期（无主订单）：

```json
{
  "code": 0,
  "data": {
    "order_no": "C20260627143000a1b2",
    "order_token": "ot_xxxxxxxx",
    "amount": 74900,
    "status": "pending",
    "code_url": "weixin://wxpay/bizpayurl?pr=xxxxx"
  }
}
```

**务必保存 `order_token`（仅返回一次）。**

### 23. 查单 / 关单

```bash
# 无主订单
curl "http://127.0.0.1:8080/v1/courses/orders/<ORDER_NO>?token=<ORDER_TOKEN>"

curl -X POST "http://127.0.0.1:8080/v1/courses/orders/<ORDER_NO>/close" \
  -H "Content-Type: application/json" \
  -d '{"order_token":"<ORDER_TOKEN>"}'

# 已登录且已认领 / 登录下单
curl "http://127.0.0.1:8080/v1/courses/orders/<ORDER_NO>" \
  -H "Authorization: Bearer <TOKEN>"
```

支付成功后 `enrollment_status` 为 `active`。

### 24. 模拟支付（dev_mode）

```bash
curl -X POST "http://127.0.0.1:8080/v1/courses/dev/orders/<ORDER_NO>/simulate-pay" \
  -H "Content-Type: application/json" \
  -d '{"order_token":"<ORDER_TOKEN>"}'
```

### 25. 登录后认领无主订单

```bash
curl -X POST http://127.0.0.1:8080/v1/courses/orders/claim \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"order_no":"<ORDER_NO>","order_token":"<ORDER_TOKEN>"}'
```

### 课程购买完整联调（dev_mode）

```bash
# 1. 试算
curl -X POST http://127.0.0.1:8080/v1/courses/checkout/preview \
  -H "Content-Type: application/json" \
  -d '{"course_id":"course_2026_spring","sku_id":"standard","channel_code":"PARTNER_X"}'

# 2. 下单（保存 ORDER_NO、ORDER_TOKEN）
curl -X POST http://127.0.0.1:8080/v1/courses/orders \
  -H "Content-Type: application/json" \
  -d '{"course_id":"course_2026_spring","sku_id":"standard","channel_code":"PARTNER_X"}'

# 3. 模拟支付
curl -X POST "http://127.0.0.1:8080/v1/courses/dev/orders/<ORDER_NO>/simulate-pay" \
  -H "Content-Type: application/json" \
  -d '{"order_token":"<ORDER_TOKEN>"}'

# 4. 确认履约
curl "http://127.0.0.1:8080/v1/courses/orders/<ORDER_NO>?token=<ORDER_TOKEN>"
```

---

## 九、推荐测试流程（新用户）

按顺序执行，避免验证码被提前消耗：

```bash
# Step 1: 发验证码
curl -X POST http://127.0.0.1:8080/v1/auth/sms/send \
  -H "Content-Type: application/json" \
  -d '{"phone":"+8613800138000"}'

# Step 1b: 检查手机号未注册（可选，[NEW]）
curl "http://127.0.0.1:8080/v1/auth/phone/check?phone=%2B8613800138000"

# Step 2: 查 SNS 是否可用（不消耗验证码）
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=alice_test_123.seedao"

# Step 2b: 首页用户总数（可选，[NEW]）
curl http://127.0.0.1:8080/v1/stats/user-count

# Step 3: 注册（消耗验证码，把 CODE 换成短信收到的）
curl -X POST http://127.0.0.1:8080/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "+8613800138000",
    "code": "301160",
    "register_type": "new_sns",
    "sns_name": "alice_test_123.seedao"
  }'

# Step 4: 查 Profile（TOKEN 换成上一步返回值）
curl http://127.0.0.1:8080/v1/user/profile \
  -H "Authorization: Bearer <TOKEN>"

# Step 5: 重新发验证码 → 登录
curl -X POST http://127.0.0.1:8080/v1/auth/sms/send \
  -H "Content-Type: application/json" \
  -d '{"phone":"+8613800138000"}'

curl -X POST http://127.0.0.1:8080/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone":"+8613800138000","code":"新验证码"}'
```

---

## 九、错误码速查

| code | 含义 |
|------|------|
| 0 | 成功 |
| 40001 | 参数错误 |
| 40002 | 验证码错误/过期 |
| 40003 | SNS 名称不可用 |
| 40004 | 钱包验签失败 | 见响应 `detail` |
| 40005 | 钱包无可认领 SNS | `detail`: `wallet_not_in_registry` |

钱包验签失败时 `detail` 取值：

| detail | 含义 |
|--------|------|
| `nonce_expired_or_missing` | 未先调 nonce，或 nonce 已过期（45 秒） |
| `domain_mismatch` | verify 的 domain 与 message 内不一致 |
| `siwe_message_invalid` | message 格式无法解析 |
| `signature_invalid` | 签名无效或与 message 不匹配 |
| `wallet_address_mismatch` | 签名地址与 wallet 不一致 |
| 40005 | 钱包无存量 SNS |
| 40006 | 手机号已注册 |
| 40009 | 手机号未注册（可注册，`/v1/auth/phone/check`） |
| 40007 | 支付订单不存在或无权访问 |
| 40008 | 支付订单状态无效（如已支付、已关闭） |
| 40010 | 优惠券无效 |
| 40011 | 优惠券不适用 |
| 40012 | 优惠券使用次数超限 |
| 40013 | 优惠后金额过低 |
| 40014 | 课程 / 票种不存在 |
| 40015 | 课程未在售 |
| 40016 | 会员折扣需登录 |
| 40017 | 课程订单 token 必填 |
| 40018 | 课程订单 token 无效 |
| 40019 | 课程订单已被认领 |
| 42901 | 发送过于频繁 |
| 50000 | 服务器内部错误 |

---

## 十、注意事项

1. 手机号必须带 `+86` 前缀，不能写 `13800138000`
2. 验证码 5 分钟有效，验证成功后立即失效
3. 钱包相关接口需真实 SIWE 签名，无法用随意字符串通过
4. 会员支付：`dev_mode` 下用 `/v1/payments/dev/simulate-pay` 联调；生产需配置微信商户参数并执行 `004_payment_orders.sql` 迁移
5. 更多接口说明见 [README.md](../README.md)
