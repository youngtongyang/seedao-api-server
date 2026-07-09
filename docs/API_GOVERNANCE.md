# 治理 / 提案 API

## 提案类型与计票规则

| 业务名称 | 创建 API `type` | `resolution_rule` | 通过条件 |
|----------|-----------------|-------------------|----------|
| 节点大会常规提案 | `node_consensus_special` | `standard`（默认） | 参与 **> 1/2** 节点总数；赞成 **> 1/2** 有效票 |
| 节点大会元规则提案 | `node_consensus_rule` | `meta_rule` | 赞成 **≥ 2/3** 节点总数 |
| 市政厅决策提案 | `town_hall_meeting` | `meta_rule`（默认） | 赞成 **≥ 2/3** 市政厅成员（`hall_voter_sns` 人数，当前 4 人即至少 **3** 票赞成） |

元规则提案创建示例：

```json
{
  "type": "node_consensus_rule",
  "resolution_rule": "meta_rule",
  "title": "...",
  "content": { ... }
}
```

---

## 白名单配置（仅 ECS 手动维护）

`governance` 段写在 **`服务器上的生产配置文件（不在本仓库）`**，**勿提交 Git**。

| 配置项 | 作用 |
|--------|------|
| `quarter_node_total` | 当季节点总数（计票分母） |
| `node_consensus_proposal_sns` | 可**发起**节点类提案 |
| `node_voter_sns` | 可**投**节点类提案 |
| `hall_proposal_sns` | 可**发起**市政厅提案 |
| `hall_voter_sns` | 可**投**市政厅提案 |

部署步骤见 [OPERATIONS.md — 治理白名单](OPERATIONS.md#治理白名单仅-ecs-手动配置)。

验证：

```bash
curl -s https://api.seedao.xyz/v1/governance/nodes | jq .
```

---

## 数据表（无 `_v2` 后缀）

| 表名 | 说明 |
|------|------|
| `proposal` | 提案 |
| `proposal_vote` | 投票 |
| `proposal_comment` | 评论 |

## 本地导入历史数据

```bash
export PGPASSWORD='seedao'
chmod +x src/scripts/import-proposals.sh
./src/scripts/import-proposals.sh "/path/to/proposal_v2_tables.sql"
```

脚本会先导入 dump（`proposal_*_v2`），再执行 `proposal-post-import.sql` 重命名为 `proposal` / `proposal_vote` / `proposal_comment`，并确保 `sip`、`resolution_rule` 等字段存在。

导入后若 dump 未含 `sip` 列，可单独执行：

```bash
psql -h "$PGHOST" -U "$PGUSER" -d "$PGDATABASE" -f src/migrations/008_proposal_sip.sql
```

## 接口一览

| 方法 | 路径 | 鉴权 |
|------|------|------|
| GET | `/v1/governance/nodes` | 无 |
| GET | `/v1/proposals` | 无 |
| GET | `/v1/proposals/:id` | 无 |
| GET | `/v1/proposals/:id/vote-summary` | 无 |
| GET | `/v1/proposals/:id/votes` | 无 |
| GET | `/v1/proposals/:id/votes/me` | JWT |
| GET | `/v1/proposals/:id/comments` | 无 |
| POST | `/v1/proposals` | JWT |
| PATCH | `/v1/proposals/:id` | JWT |
| POST | `/v1/proposals/:id/close` | JWT |
| POST | `/v1/proposals/:id/publish` | JWT |
| POST | `/v1/proposals/:id/finalize` | JWT |
| POST | `/v1/proposals/:id/votes` | JWT |
| POST | `/v1/proposals/:id/comments` | JWT |

## 计票规则（`vote_summary`）

| `resolution_rule` | 通过条件 |
|-------------------|----------|
| `standard` | 参与节点 > 当季节点总数 1/2；赞成 > 有效票 1/2 |
| `meta_rule` | 赞成 ≥ 选举人总数 2/3（节点元规则用 `quarter_node_total`；市政厅用 `hall_voter_sns` 人数） |

同一 SNS 对同一提案仅可投一票（应用层 + DB 唯一索引）。

## Swagger

- 启动服务后访问：`http://127.0.0.1:8080/swagger/index.html`
- 按路由分组 Tag：`health` · `auth` · `sns` · `wallet` · `stats` · `user` · `payments` · `courses` · `proposals`（`/v1/governance/nodes` 不对外展示于 Swagger）
- 重新生成文档：

```bash
cd src
go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

## 错误码 40020–40027

见 `internal/pkg/errcode/errcode.go`。
