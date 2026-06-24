#!/usr/bin/env bash
# 按注册手机号删除用户及其关联数据（users / user_phones / sns_names / payment_* 等）
#
# 默认 dry-run，仅预览；加 --apply 才会真正删除。
#
# Usage:
#   ./scripts/delete-user-by-phone.sh +8618515153726
#   ./scripts/delete-user-by-phone.sh 18515153726 --apply
#
# Env (同 setup-db.sh):
#   PGHOST PGPORT PGUSER PGPASSWORD PGDATABASE
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

PGHOST="${PGHOST:-127.0.0.1}"
PGPORT="${PGPORT:-5432}"
PGUSER="${PGUSER:-seedao}"
PGPASSWORD="${PGPASSWORD:-seedao}"
PGDATABASE="${PGDATABASE:-seedao_api}"

APPLY=0
PHONE_RAW=""

usage() {
  cat <<'EOF'
Usage:
  ./scripts/delete-user-by-phone.sh <phone> [--apply]

Arguments:
  phone     注册手机号，支持 +8618515153726 或 18515153726（自动补 +86）

Options:
  --apply   执行删除（默认仅预览，不改动数据库）

Environment:
  PGHOST PGPORT PGUSER PGPASSWORD PGDATABASE

Example:
  PGPASSWORD='***' ./scripts/delete-user-by-phone.sh 18515153726
  PGPASSWORD='***' ./scripts/delete-user-by-phone.sh +8618515153726 --apply
EOF
}

normalize_phone() {
  local raw="$1"
  raw="$(echo "$raw" | tr -d '[:space:]')"
  if [[ "$raw" =~ ^\+[1-9][0-9]{6,14}$ ]]; then
    echo "$raw"
    return
  fi
  if [[ "$raw" =~ ^1[0-9]{10}$ ]]; then
    echo "+86${raw}"
    return
  fi
  echo "ERROR: invalid phone format: $raw (use +8613... or 11-digit CN mobile)" >&2
  exit 1
}

phone_hash() {
  local phone="$1"
  if command -v sha256sum >/dev/null 2>&1; then
    echo -n "$phone" | sha256sum | awk '{print $1}'
  elif command -v shasum >/dev/null 2>&1; then
    echo -n "$phone" | shasum -a 256 | awk '{print $1}'
  else
    echo ""
  fi
}

run_psql() {
  PGPASSWORD="$PGPASSWORD" psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -v ON_ERROR_STOP=1 "$@"
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    -h|--help)
      usage
      exit 0
      ;;
    --apply)
      APPLY=1
      shift
      ;;
    *)
      if [[ -z "$PHONE_RAW" ]]; then
        PHONE_RAW="$1"
      else
        echo "ERROR: unexpected argument: $1" >&2
        usage
        exit 1
      fi
      shift
      ;;
  esac
done

if [[ -z "$PHONE_RAW" ]]; then
  usage
  exit 1
fi

PHONE="$(normalize_phone "$PHONE_RAW")"
PHONE_HASH="$(phone_hash "$PHONE")"

if [[ -z "$PHONE_HASH" ]]; then
  echo "WARN: sha256sum/shasum not found; sms_send_logs will not be deleted" >&2
fi

SQL_FILE="$(mktemp)"
APPLY_SQL=""
trap 'rm -f "$SQL_FILE" ${APPLY_SQL:+"$APPLY_SQL"}' EXIT

cat >"$SQL_FILE" <<SQL
SELECT u.id AS user_id,
       up.phone,
       u.register_source,
       sn.name AS sns_name,
       sn.wallet AS sns_wallet,
       sn.mint_status,
       (SELECT COUNT(*) FROM payment_orders po WHERE po.user_id = u.id) AS payment_order_count
FROM user_phones up
JOIN users u ON u.id = up.user_id
LEFT JOIN sns_names sn ON sn.user_id = u.id
WHERE up.phone = '$PHONE';
SQL

echo "==> Target phone (E.164): $PHONE"
echo "==> Database: $PGUSER@$PGHOST:$PGPORT/$PGDATABASE"
if [[ -n "$PHONE_HASH" ]]; then
  echo "==> phone_hash (Redis/SMS logs): $PHONE_HASH"
fi
echo ""

USER_ID="$(run_psql -tAc "SELECT user_id FROM user_phones WHERE phone = '$PHONE';" | tr -d '[:space:]')"

if [[ -z "$USER_ID" ]]; then
  echo "No user found for phone: $PHONE"
  exit 0
fi

echo "==> Preview user_id=$USER_ID"
run_psql -f "$SQL_FILE"
echo ""

if [[ "$APPLY" -ne 1 ]]; then
  echo "DRY-RUN only. Re-run with --apply to delete."
  echo ""
  echo "Will delete (in transaction):"
  echo "  - payment_refunds (via user's payment_orders)"
  echo "  - payment_orders"
  echo "  - sns_names"
  echo "  - user_phones"
  echo "  - users"
  echo "  - sms_send_logs (same phone_hash, optional audit rows)"
  if [[ -n "$(run_psql -tAc "SELECT wallet FROM sns_names WHERE user_id = $USER_ID AND wallet IS NOT NULL LIMIT 1;" | tr -d '[:space:]')" ]]; then
    echo "  - reset sns_chain_registry.claimed=false for existing_sns wallet"
  fi
  exit 0
fi

echo "==> Applying delete for user_id=$USER_ID ..."

APPLY_SQL="$(mktemp)"

cat >"$APPLY_SQL" <<SQL
BEGIN;

CREATE TEMP TABLE _del_user ON COMMIT DROP AS
SELECT u.id AS user_id
FROM users u
JOIN user_phones up ON up.user_id = u.id
WHERE up.phone = '$PHONE';

DO \$\$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM _del_user) THEN
    RAISE EXCEPTION 'user not found for phone %', '$PHONE';
  END IF;
END
\$\$;

UPDATE sns_chain_registry r
SET claimed = FALSE,
    updated_at = NOW()
WHERE LOWER(r.wallet) IN (
  SELECT LOWER(sn.wallet)
  FROM sns_names sn
  JOIN _del_user d ON d.user_id = sn.user_id
  WHERE sn.wallet IS NOT NULL
);

DELETE FROM payment_refunds pr
USING payment_orders po, _del_user d
WHERE pr.order_id = po.id
  AND po.user_id = d.user_id;

DELETE FROM payment_orders po
USING _del_user d
WHERE po.user_id = d.user_id;

DELETE FROM sns_names sn
USING _del_user d
WHERE sn.user_id = d.user_id;

DELETE FROM user_phones up
USING _del_user d
WHERE up.user_id = d.user_id;

DELETE FROM users u
USING _del_user d
WHERE u.id = d.user_id;
SQL

if [[ -n "$PHONE_HASH" ]]; then
  cat >>"$APPLY_SQL" <<SQL

DELETE FROM sms_send_logs
WHERE phone_hash = '$PHONE_HASH';
SQL
fi

cat >>"$APPLY_SQL" <<SQL

COMMIT;
SQL

run_psql -f "$APPLY_SQL"

echo "==> Done. Deleted user for phone: $PHONE"
echo ""
echo "Optional Redis cleanup (if same phone still rate-limited or has stale SMS code):"
if [[ -n "$PHONE_HASH" ]]; then
  cat <<EOF
  redis-cli -a '<redis_password>' DEL "sms:code:$PHONE_HASH" "sms:rate:phone:$PHONE_HASH"
EOF
fi
