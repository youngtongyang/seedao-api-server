# API 本地测试手册

本文档汇总 SeeDAO API 本地联调用的 `curl` 命令，便于在 Git Bash 中复制执行。

**前提：**

- API 已启动：`bash src/scripts/start-local.sh` 或 `go run ./cmd/server`
- 默认地址：`http://127.0.0.1:8080`
- 终端：Git Bash（Windows 下可直接用 `curl`）

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

### 2. 发送验证码

```bash
curl -X POST http://127.0.0.1:8080/v1/auth/sms/send \
  -H "Content-Type: application/json" \
  -d '{"phone":"+8613800138000"}'
```

说明：

- 同号 60 秒内不可重复发送
- 同 IP 每小时最多 20 次
- `dev_mode: true` 时验证码打在 API 日志；`false` 时真实发短信

---

## 二、SNS 名称检查（不消耗验证码）

### 3. 查询可用名称

```bash
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=alice_test_123.seedao"
```

预期：`available: true`，响应中 `sns_name` 为 `alice_test_123.seedao`

### 4. 查询已被 registry 占用的名称

```bash
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=baiyu.seedao"
```

预期：`available: false`, `reason: taken_in_registry`

### 5. 查询格式非法的名称

```bash
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=AB.seedao"
```

预期：`available: false`, `reason: invalid_format`

---

## 三、新用户注册（new_sns）

建议顺序：**先 SNS check → 再 register**

### 6. 新用户注册

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
      "register_source": "new_sns"
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

## 四、登录

已注册用户需**重新发送验证码**后再登录。

### 7. 登录

```bash
curl -X POST http://127.0.0.1:8080/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "+8613800138000",
    "code": "301160"
  }'
```

响应格式同注册，返回 `token` 与 `user`。

---

## 五、个人 Profile（需 JWT）

### 8. 查询 Profile

将 `<TOKEN>` 替换为注册/登录返回的 token：

```bash
curl http://127.0.0.1:8080/v1/user/profile \
  -H "Authorization: Bearer <TOKEN>"
```

预期：返回 `id`、`sns_name`、脱敏 `phone`、`register_source` 等。

### 9. 无 Token 访问（应失败）

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

### 10. 获取钱包 Nonce / SIWE Message

```bash
curl -X POST http://127.0.0.1:8080/v1/wallet/nonce \
  -H "Content-Type: application/json" \
  -d '{
    "wallet": "0x78f625a65fc316d32d98d249b698fb509a6d98f2",
    "domain": "app.seedao.top"
  }'
```

返回 `nonce`、`message`、`expires_at`。用 MetaMask 等对 `message` 签名，得到 `signature`。

### 11. 验证 SNS 归属

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

### 12. 老用户注册

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

## 七、推荐测试流程（新用户）

按顺序执行，避免验证码被提前消耗：

```bash
# Step 1: 发验证码
curl -X POST http://127.0.0.1:8080/v1/auth/sms/send \
  -H "Content-Type: application/json" \
  -d '{"phone":"+8613800138000"}'

# Step 2: 查 SNS 是否可用（不消耗验证码）
curl "http://127.0.0.1:8080/v1/sns/check?sns_name=alice_test_123.seedao"

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

## 八、错误码速查

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

## 九、注意事项

1. 手机号必须带 `+86` 前缀，不能写 `13800138000`
2. 验证码 5 分钟有效，验证成功后立即失效
3. 钱包相关接口需真实 SIWE 签名，无法用随意字符串通过
4. 更多接口说明见 [README.md](../README.md)
