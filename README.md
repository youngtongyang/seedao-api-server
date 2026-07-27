# seedao-api-server

SeeDAO 2.0 用户认证与 SNS 服务 API（v2.7）。

面向 Web / H5 客户端，提供手机号验证码注册/登录、SNS 名称校验与预留、老用户钱包认领链上 SNS、个人 Profile、飞书账号绑定（土地神）、首页用户统计、会员微信支付、课程售卖、治理提案等能力。运行时零链上 RPC 依赖。

## 项目结构

```
seedao-api-server/
├── src/                          # 全部源码与脚本
│   ├── cmd/server/main.go        # 入口
│   ├── internal/
│   │   ├── app/                  # 路由与应用组装
│   │   ├── config/               # 配置加载
│   │   ├── handler/              # HTTP 层 + 单元测试
│   │   ├── middleware/           # JWT / Service Token 鉴权
│   │   ├── service/              # 业务逻辑（含 oauth_binding）
│   │   ├── model/                # GORM 实体
│   │   ├── testutil/             # 测试工具
│   │   └── pkg/                  # jwt / siwe / redis / sms / bindtoken 等
│   ├── migrations/               # SQL 迁移（含 010_user_oauth_bindings）
│   ├── scripts/                  # 建库、导入、本地启动脚本
│   ├── config/
│   │   ├── config.local.example.yaml
│   │   ├── config.dev.example.yaml
│   │   ├── config.prod.example.yaml
│   │   ├── config.local.yaml          # 本地（gitignore）
│   │   ├── config.dev.yaml            # 开发（gitignore）
│   │   └── config.prod.yaml           # 生产（gitignore）
│   ├── docs/                  # Swagger 生成物（swag init）
│   ├── go.mod
│   └── go.sum
├── docker-compose.dev.yml      # 开发环境（postgres + redis + api）
├── docker-compose.prod.yml     # 生产环境（redis + api）
├── docs/
│   ├── DEPLOYMENT.md          # 部署步骤说明
│   ├── OPERATIONS.md          # 运维操作手册（发版、备份、排错、删用户等）
│   ├── API_CHANGELOG.md       # 版本变更一览
│   ├── API_PAYMENTS_v2.6.md   # v2.6 支付/课程业务说明
│   ├── API_GOVERNANCE.md      # v2.7 治理/提案（含内部配置）
│   ├── API_TEST.md            # 本地联调 curl
│   └── feishu-bindings/       # 飞书绑定（土地神 Token）方案与流程图
│       └── README.md
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
| `feishu.*` | 飞书绑定：`app_id`、token/code TTL、`bind_url_template` |
| `internal.service_token` | 土地神调用 `/internal/v1/feishu/*` 的共享密钥 |

## API 文档

**在线 Swagger UI（推荐）：** [https://api.seedao.xyz/swagger/index.html](https://api.seedao.xyz/swagger/index.html)

本地启动后：`http://127.0.0.1:8080/swagger/index.html`

| 文档 | 说明 |
|------|------|
| Swagger UI | 全量接口、请求/响应 schema、在线调试 |
| [docs/API_TEST.md](docs/API_TEST.md) | 本地联调 curl 示例 |
| [docs/API_CHANGELOG.md](docs/API_CHANGELOG.md) | 按版本整理的接口变更一览 |
| [docs/API_PAYMENTS_v2.6.md](docs/API_PAYMENTS_v2.6.md) | 支付/课程业务规则与前端对接流程 |
| [docs/API_GOVERNANCE.md](docs/API_GOVERNANCE.md) | 治理/提案（含白名单配置、导入脚本） |
| [docs/feishu-bindings/README.md](docs/feishu-bindings/README.md) | 飞书绑定方案、流程图与接口说明 |

修改 handler 注释后于 `src/` 重新生成 OpenAPI：

```bash
make swagger
# 或
go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

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
| 40009 | 手机号未注册（可注册） |
| 40010 | 优惠券无效 |
| 40011 | 优惠券不适用 |
| 40012 | 优惠券使用次数超限 |
| 40013 | 优惠后金额过低 |
| 40014 | 课程不存在 |
| 40015 | 课程未在售 |
| 40016 | 会员折扣需登录 |
| 40017 | 订单 token 必填 |
| 40018 | 订单 token 无效 |
| 40019 | 订单已被认领 |
| 40020 | 提案不存在 |
| 40021 | 提案操作无权限 |
| 40022 | 提案状态不允许 |
| 40023 | 不在投票期 |
| 40024 | 已投过票 |
| 40025 | 不在投票人白名单 |
| 40026 | 需绑定 SNS |
| 40027 | 未找到该用户投票 |
| 40040 | 飞书未绑定 |
| 40041 | 当前用户已绑定飞书 |
| 40042 | 该飞书账号已被其他用户绑定 |
| 40043 | 飞书 bind_token / bind_code 无效或过期 |
| 42901 | 发送过于频繁 |
| 50000 | 服务器内部错误 |

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

# 生产环境（推荐 .env 配置 SEEDAO_IMAGE 为阿里云 ACR，见 docs/DEPLOYMENT.md 三-A）
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

- Go 1.23+ / Echo
- PostgreSQL 16 / GORM
- Redis 7
- 阿里云 dysmsapi
- JWT (HS256) / SIWE 本地验签
