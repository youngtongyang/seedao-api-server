# seedao-api-server

SeeDAO 2.0 用户认证与 SNS 服务 API（v2.5）。

面向 Web / H5 客户端，提供手机号验证码注册/登录、SNS 名称校验与预留、老用户钱包认领链上 SNS、个人 Profile、首页用户统计、会员微信支付等能力。运行时零链上 RPC 依赖。

## 项目结构

```
seedao-api-server/
├── src/                          # 全部源码与脚本
│   ├── cmd/server/main.go        # 入口
│   ├── internal/
│   │   ├── app/                  # 路由与应用组装
│   │   ├── config/               # 配置加载
│   │   ├── handler/              # HTTP 层 + 单元测试
│   │   ├── middleware/           # JWT 鉴权
│   │   ├── service/              # 业务逻辑
│   │   ├── model/                # GORM 实体
│   │   ├── testutil/             # 测试工具
│   │   └── pkg/                  # jwt / siwe / redis / sms 等
│   ├── migrations/               # SQL 迁移
│   ├── scripts/                  # 建库、导入、本地启动脚本
│   ├── config/
│   │   ├── config.local.example.yaml
│   │   ├── config.dev.example.yaml
│   │   ├── config.prod.example.yaml
│   │   ├── config.local.yaml          # 本地（gitignore）
│   │   ├── config.dev.yaml            # 开发（gitignore）
│   │   └── config.prod.yaml           # 生产（gitignore）
│   ├── go.mod
│   └── go.sum
├── docker-compose.dev.yml      # 开发环境（postgres + redis + api）
├── docker-compose.prod.yml     # 生产环境（redis + api）
├── docs/
│   ├── DEPLOYMENT.md          # 部署步骤说明
│   ├── OPERATIONS.md          # 运维操作手册（发版、备份、排错、删用户等）
│   └── API_TEST.md            # 本地联调 curl（含 [NEW]/[CHANGED] 标注）
└── README.md
```

## 配置说明

三套环境配置模板（复制后填写密钥，**切勿提交实际 yaml**）：

| 环境 | 模板 | 实际配置 |
|------|------|----------|
| local | `config.local.example.yaml` | `config.local.yaml` |
| dev | `config.dev.example.yaml` | `config.dev.yaml` |
| prod | `config.prod.example.yaml` | `config.prod.yaml` |

```bash
cd src/config
cp config.local.example.yaml config.local.yaml
```

本地联调 curl 命令见 **[docs/API_TEST.md](docs/API_TEST.md)**。

| 配置项 | 说明 |
|--------|------|
| `database.dsn` | PostgreSQL 连接串 |
| `redis.addr` | Redis 地址 |
| `redis.password` | Redis 密码（必填） |
| `jwt.secret` | JWT 签名密钥（生产环境使用强随机值） |
| `aliyun_sms.*` | 阿里云短信 AccessKey、签名、模板 |
| `aliyun_sms.dev_mode` | local 可设 `true`（日志验证码）；dev/prod 设 `false`（真实发短信） |
| `siwe.chain_id` | SIWE 验签链 ID |
| `wechat_pay.*` | 微信支付（会员续费）；生产 `dev_mode: false`，需挂载商户证书 |

## API 接口

**变更标识：** `[NEW]` 本次新增 · `[CHANGED]` 本次变更（对接时请优先关注）

### API 变更一览（v2.5）

| 接口 | 鉴权 | 标识 | 说明 |
|------|------|------|------|
| `GET /v1/stats/user-count` | 无 | `[NEW]` | 首页展示用户总数 |
| `POST /v1/payments/orders` | JWT | `[NEW]` | 创建会员续费订单（Native / H5） |
| `GET /v1/payments/orders` | JWT | `[NEW]` | 订单列表 |
| `GET /v1/payments/orders/:order_no` | JWT | `[NEW]` | 查询订单（前端轮询） |
| `POST /v1/payments/orders/:order_no/close` | JWT | `[NEW]` | 关闭未支付订单 |
| `POST /v1/payments/wechat/notify` | 无 | `[NEW]` | 微信支付异步回调 |
| `POST /v1/payments/dev/simulate-pay` | JWT | `[NEW]` | 模拟支付（仅 `dev_mode`） |
| `GET /v1/user/profile` | JWT | `[CHANGED]` | 响应增加 `membership_expires_at` |
| `POST /v1/auth/login` / `register` | 无 | `[CHANGED]` | `user` 增加 `membership_expires_at` |

本地联调 curl 见 **[docs/API_TEST.md](docs/API_TEST.md)**（含 `[NEW]` / `[CHANGED]` 标注）。

统一响应格式：

```json
{ "code": 0, "message": "ok", "data": {} }
```

### 错误码

| code | 含义 |
|------|------|
| 0 | 成功 |
| 40001 | 参数错误 |
| 40002 | 验证码错误/过期 |
| 40003 | SNS 名称不可用 |
| 40004 | 钱包验签失败 |
| 40005 | 钱包无存量 SNS |
| 40006 | 手机号已注册 |
| 40007 | 支付订单不存在或无权访问 |
| 40008 | 支付订单状态无效 |
| 42901 | 发送过于频繁 |
| 50000 | 服务器内部错误 |

---

### 1. 健康检查

`GET /health` · 无需鉴权

```bash
curl http://127.0.0.1:8080/health
```

响应：

```json
{ "code": 0, "message": "ok", "data": { "status": "ok" } }
```

---

### 2. 发送验证码

`POST /v1/auth/sms/send` · 无需鉴权

请求：

```json
{ "phone": "+8613800138000" }
```

响应：

```json
{ "code": 0, "message": "ok", "data": { "retry_after": 60 } }
```

说明：验证码 6 位、5 分钟有效；同号 60s、同 IP 20 次/小时限流。

---

### 3. 登录 `[CHANGED]`

`POST /v1/auth/login` · 无需鉴权

请求：

```json
{ "phone": "+8613800138000", "code": "123456" }
```

响应：

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "token": "eyJhbG...",
    "token_exp": 1718438400,
    "user": {
      "id": 1001,
      "sns_name": "alice.seedao",
      "register_source": "new_sns",
      "membership_expires_at": "2027-06-15T08:00:00Z"
    }
  }
}
```

---

### 4. 注册 `[CHANGED]`

`POST /v1/auth/register` · 无需鉴权

**新用户 `new_sns`：**

```json
{
  "phone": "+8613800138000",
  "code": "123456",
  "register_type": "new_sns",
  "sns_name": "alice_seed.seedao"
}
```

`sns_name` 必须为 `{base}.seedao` 格式（与 `sns_chain_registry` 一致）。

**老用户 `existing_sns`（需先完成钱包验签）：**

```json
{
  "phone": "+8613800138000",
  "code": "123456",
  "register_type": "existing_sns",
  "wallet": "0xAbCd...",
  "wallet_type": "EOA",
  "domain": "app.seedao.top",
  "message": "...",
  "signature": "0x..."
}
```

响应同登录，返回 JWT 与用户信息。`user.membership_expires_at` 为会员到期时间（ISO 8601 UTC）。`[CHANGED]`

---

### 5. SNS 名称可用性

`GET /v1/sns/check?sns_name=baiyu.seedao` · 无需鉴权

可用：

```json
{ "code": 0, "data": { "sns_name": "alice_test.seedao", "available": true } }
```

不可用：

```json
{
  "code": 0,
  "data": {
    "sns_name": "baiyu.seedao",
    "available": false,
    "reason": "taken_in_registry"
  }
}
```

`reason`：`taken_in_registry` | `taken_in_app` | `invalid_format` | `reserved_word`

---

### 6. 钱包 Nonce（SIWE）

`POST /v1/wallet/nonce` · 无需鉴权

请求：

```json
{ "wallet": "0xAbCd...", "domain": "app.seedao.top" }
```

响应：

```json
{
  "code": 0,
  "data": {
    "nonce": "8f3c2a...",
    "message": "app.seedao.top wants you to sign in...",
    "expires_at": 1718438445
  }
}
```

---

### 7. 验证 SNS 归属

`POST /v1/wallet/verify-sns` · 无需鉴权

请求：

```json
{
  "wallet": "0xAbCd...",
  "wallet_type": "EOA",
  "domain": "app.seedao.top",
  "message": "...",
  "signature": "0x..."
}
```

响应：

```json
{
  "code": 0,
  "data": {
    "wallet": "0xabcd...",
    "sns_name": "baiyu.seedao",
    "claimable": true
  }
}
```

---

### 8. 个人 Profile `[CHANGED]`

`GET /v1/user/profile` · **需要 JWT**

```bash
curl -H "Authorization: Bearer <token>" http://127.0.0.1:8080/v1/user/profile
```

响应：

```json
{
  "code": 0,
  "data": {
    "id": 1001,
    "sns_name": "alice.seedao",
    "phone": "+86138****8000",
    "register_source": "new_sns",
    "avatar": "",
    "bio": "",
    "membership_expires_at": "2027-06-15T08:00:00Z",
    "created_at": "2026-06-15T08:00:00Z"
  }
}
```

---

### 9. 首页用户总数 `[NEW]`

`GET /v1/stats/user-count` · 无需鉴权

```bash
curl https://api.seedao.xyz/v1/stats/user-count
```

响应：

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "total_user_count": 12345
  }
}
```

说明：`total_user_count` = 链上存量 SNS（`sns_chain_registry`）+ 新注册用户（`register_source=new_sns`）。

---

### 10. 会员微信支付 `[NEW]`

以下接口均需 JWT（微信回调除外）。`pay_channel` 支持 `wechat_native`（PC 扫码）与 `wechat_h5`（手机浏览器）。

| 方法 | 路径 | 鉴权 |
|------|------|------|
| POST | `/v1/payments/orders` | JWT |
| GET | `/v1/payments/orders` | JWT |
| GET | `/v1/payments/orders/:order_no` | JWT |
| POST | `/v1/payments/orders/:order_no/close` | JWT |
| POST | `/v1/payments/wechat/notify` | 无（微信服务器） |
| POST | `/v1/payments/dev/simulate-pay` | JWT（仅 `dev_mode`） |

**创建订单** `POST /v1/payments/orders`：

```json
{
  "product_type": "membership_renew",
  "days": 365,
  "pay_channel": "wechat_native"
}
```

- `pay_channel` 可省略，默认 `wechat_native`
- `days` 可省略，默认 `365`；金额由服务端 `wechat_pay.membership_price_fen` 决定
- Native 返回 `code_url`（生成二维码）；H5 返回 `h5_url`（页面跳转）

**查询订单** `GET /v1/payments/orders/:order_no`：支付成功后 `status=paid`，并返回 `membership_expires_at`。

本地联调与完整 curl 见 **[docs/API_TEST.md](docs/API_TEST.md) 第七节**。

---

## 本地开发（local）

### 1. 配置

```bash
cd src/config
cp config.local.example.yaml config.local.yaml
```

### 2. 建库与导入（脚本见 `src/scripts/`）

```bash
PGSUPERPASSWORD='postgres密码' \
bash src/scripts/setup-db.sh
```

### 3. 启动

```bash
cd src
export CONFIG_PATH="$(pwd)/config/config.local.yaml"
go run ./cmd/server
```

或一键：`bash src/scripts/start-local.sh`

### 4. Docker 部署

```bash
# 开发环境（postgres + redis + api）
cp src/config/config.dev.example.yaml src/config/config.dev.yaml
# 编辑 config.dev.yaml，填写 redis.password 等配置
docker compose -f docker-compose.dev.yml up -d --build

# 生产环境（从 GHCR 拉镜像，需先 docker login ghcr.io）
cp src/config/config.prod.example.yaml src/config/config.prod.yaml
docker compose -f docker-compose.prod.yml up -d
# 更新：bash deploy/pull-deploy.sh
```

**阿里云 ECS 生产部署完整步骤**见 [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md) 第一章；**日常运维**见 [docs/OPERATIONS.md](docs/OPERATIONS.md)。

---

## 单元测试

测试使用内存 SQLite + miniredis，无需外部 PostgreSQL/Redis：

```bash
cd src
go test ./...
```

仅跑 HTTP 接口测试：

```bash
go test ./internal/handler/... -v
```

---

## 技术栈

- Go 1.22+ / Echo
- PostgreSQL 16 / GORM
- Redis 7
- 阿里云 dysmsapi
- JWT (HS256) / SIWE 本地验签
