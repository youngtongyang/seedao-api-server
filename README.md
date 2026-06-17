# seedao-api-server

SeeDAO 2.0 用户认证与 SNS 服务 API（v2.4）。

面向 Web / H5 客户端，提供手机号验证码注册/登录、SNS 名称校验与预留、老用户钱包认领链上 SNS、个人 Profile 等能力。运行时零链上 RPC 依赖。

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
├── docs/DEPLOYMENT.md          # 部署步骤说明
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

## API 接口

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

### 3. 登录

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
      "register_source": "new_sns"
    }
  }
}
```

---

### 4. 注册

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

响应同登录，返回 JWT 与用户信息。

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

### 8. 个人 Profile

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
    "created_at": "2026-06-15T08:00:00Z"
  }
}
```

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

**阿里云 ECS 生产部署完整步骤**见 [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md) 第一章「服务器生产部署」。

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
