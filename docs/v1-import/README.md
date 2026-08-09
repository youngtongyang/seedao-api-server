# 1.0 → 2.0 数据导入

从 SeeDAO 1.0 API 清洗资料与发放记录到 2.0 镜像表（**不回写** `users`）。

## 前置

1. 执行迁移：

```bash
psql "$DSN" -v ON_ERROR_STOP=1 -f src/migrations/011_legacy_import.sql
# 或本地：bash src/scripts/setup-db.sh
```

2. 确保 `sns_chain_registry` 已导入（脚本 1 唯一输入源）。

3. 配置 `config/*.yaml` 中的 `database`、`v1_import`；OSS 后跑时再填 `aliyun_oss`。

## 表

| 表 | 说明 |
|----|------|
| `user_profiles` | bio / email / twitter / mirror / scr_amount，PK=`sns_name` |
| `user_assets` | assets[]（含 WANG/SEE），UNIQUE(`sns_name`,`asset_name`) |
| `user_roles` | sp.roles，PK=(`sns_name`,`role`)，按 sns 先删后插 |
| `user_seeds` | sp.seed + `image_uri_src` / `image_uri_oss` |
| `user_sbts` | sp.sbt + display_name（name→headline→description） |
| `reward_applications` | new_reward 发放记录 |
| `oss_objects` | 图片源 URL → OSS 去重缓存 |

## 命令

在 `src/` 目录：

```bash
# 脚本1：资料/资产/角色/Seed/SBT（默认不传 OSS）
go run ./cmd/v1import profiles -config config/config.local.yaml

# 导入时顺带上传图片（需配置 aliyun_oss）
go run ./cmd/v1import profiles -config config/config.local.yaml -upload-oss

# 脚本2：发放记录
go run ./cmd/v1import rewards -config config/config.local.yaml

# 后补 OSS（扫描 image_uri_oss 为空的 seed/sbt）
# 配置了 v1_import.opensea_api_key 时：直接按 contract+token_id 查 OpenSea 图，不再先走 IPFS
go run ./cmd/v1import oss -config config/config.local.yaml -limit 500
```

## 同步策略

- profiles：UPSERT by `sns_name`
- assets：UPSERT by (`sns_name`,`asset_name`)；SEE 只取 `assets[]`
- roles：按 `sns_name` DELETE → INSERT
- seeds / sbts：UPSERT；OSS URL 仅在新值非空时覆盖
- rewards：UPSERT by `application_id`；wallet→sns 查 `sns_chain_registry`

## 配置片段

```yaml
v1_import:
  base_url: "https://api.seedao.top"
  ipfs_gateway: "https://dweb.link/ipfs/"
  concurrency: 3
  page_size: 100
  opensea_api_key: "YOUR_OPENSEA_KEY"
  opensea_base_url: "https://api.opensea.io"
  opensea_chain: "ethereum"

aliyun_oss:
  endpoint: "oss-cn-heyuan.aliyuncs.com"
  bucket: "your-bucket"
  access_key_id: "..."
  access_key_secret: "..."
  public_base_url: ""  # 可选 CDN
  prefix: "img"
```
