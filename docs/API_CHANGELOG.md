# API 版本变更一览

**变更标识：** `[NEW]` 本次新增 · `[CHANGED]` 本次变更

完整请求/响应字段以 **[Swagger UI](https://api.seedao.xyz/swagger/index.html)** 为准；本地联调 curl 见 [API_TEST.md](API_TEST.md)。

---

## v2.7

| 接口 | 鉴权 | 标识 | 说明 |
|------|------|------|------|
| `GET /v1/governance/nodes` | 无 | `[NEW]` | 治理节点白名单配置（不在 Swagger 展示） |
| `GET /v1/proposals` | 无 | `[NEW]` | 分页提案列表 |
| `GET /v1/proposals/:id` | 无 | `[NEW]` | 提案详情（可选含投票/评论） |
| `GET /v1/proposals/:id/vote-summary` | 无 | `[NEW]` | 投票计票摘要 |
| `GET /v1/proposals/:id/votes` | 无 | `[NEW]` | 投票列表 |
| `GET /v1/proposals/:id/votes/me` | JWT | `[NEW]` | 当前用户投票 |
| `GET /v1/proposals/:id/comments` | 无 | `[NEW]` | 评论列表 |
| `POST /v1/proposals` | JWT | `[NEW]` | 创建草稿（市政厅/节点共识） |
| `PATCH /v1/proposals/:id` | JWT | `[NEW]` | 修改草稿 |
| `POST /v1/proposals/:id/close` | JWT | `[NEW]` | 关闭草稿 |
| `POST /v1/proposals/:id/publish` | JWT | `[NEW]` | 发布进入公示 |
| `POST /v1/proposals/:id/finalize` | JWT | `[NEW]` | 投票结束后结票 |
| `POST /v1/proposals/:id/votes` | JWT | `[NEW]` | 投票（一 SNS 一票） |
| `POST /v1/proposals/:id/comments` | JWT | `[NEW]` | 发表评论 |

治理业务说明见 [API_GOVERNANCE.md](API_GOVERNANCE.md)。

---

## v2.6

| 接口 | 鉴权 | 标识 | 说明 |
|------|------|------|------|
| `POST /v1/payments/coupons/preview` | JWT 可选 | `[NEW]` | 会员续费券码试算；未登录不做每人限次校验 |
| `POST /v1/payments/orders` | JWT 必填 | `[CHANGED]` | **仅** `membership_renew`；支持 `coupon_code` |
| `POST /v1/courses/checkout/preview` | JWT 可选 | `[NEW]` | 课程价格试算（早鸟+渠道+会员预览） |
| `POST /v1/courses/orders` | JWT 条件必填 | `[NEW]` | 课程下单；无主订单返回 `order_token` |
| `GET /v1/courses/orders/:order_no` | JWT 或 token | `[NEW]` | 查课程订单 |
| `POST /v1/courses/orders/:order_no/close` | JWT 或 token | `[NEW]` | 关闭课程订单 |
| `POST /v1/courses/orders/claim` | JWT 必填 | `[NEW]` | 登录后认领无主订单 |
| `POST /v1/courses/dev/orders/:order_no/simulate-pay` | JWT 或 token | `[NEW]` | 模拟课程支付（仅 `dev_mode`） |

支付与课程完整说明见 [API_PAYMENTS_v2.6.md](API_PAYMENTS_v2.6.md)。

---

## v2.5（历史）

| 接口 | 鉴权 | 标识 | 说明 |
|------|------|------|------|
| `GET /v1/stats/user-count` | 无 | `[NEW]` | 首页展示用户总数 |
| `GET /v1/auth/phone/check` | 无 | `[NEW]` | 注册前检查手机号未占用 |
| `POST /v1/payments/orders` | JWT | `[NEW]` | 创建会员续费订单（Native / H5） |
| `GET /v1/payments/orders` | JWT | `[NEW]` | 订单列表 |
| `GET /v1/payments/orders/:order_no` | JWT | `[NEW]` | 查询订单（前端轮询） |
| `POST /v1/payments/orders/:order_no/close` | JWT | `[NEW]` | 关闭未支付订单 |
| `POST /v1/payments/wechat/notify` | 无 | `[NEW]` | 微信支付异步回调 |
| `POST /v1/payments/dev/simulate-pay` | JWT | `[NEW]` | 模拟支付（仅 `dev_mode`） |
| `GET /v1/user/profile` | JWT | `[CHANGED]` | 响应增加 `membership_expires_at` |
| `POST /v1/auth/login` / `register` | 无 | `[CHANGED]` | `user` 增加 `membership_expires_at` |
