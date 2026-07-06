# 治理 / 提案 API

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

脚本会先导入 dump（`proposal_*_v2`），再执行 `007_proposals.sql` 重命名为 `proposal` / `proposal_vote` / `proposal_comment`。

## 配置 `governance`（节点白名单，不走链上 SBT）

```yaml
governance:
  quarter_node_total: 20
  hall_proposal_sns: ["baiyu.seedao"]
  node_consensus_proposal_sns: ["node.seedao"]
  hall_voter_sns: ["hall_voter.seedao"]
  node_voter_sns: ["voter1.seedao", "voter2.seedao"]
  default_publicity_seconds: 432000
  default_vote_duration_seconds: 172800
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
| `meta_rule` | 赞成 ≥ 当季节点总数 2/3 |

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
