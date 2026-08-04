#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

REGISTRY_IMPORT_SOURCE="${REGISTRY_IMPORT_SOURCE:-}"
PGHOST="${PGHOST:-127.0.0.1}"
PGPORT="${PGPORT:-5432}"
PGUSER="${PGUSER:-seedao}"
PGPASSWORD="${PGPASSWORD:-seedao}"
PGDATABASE="${PGDATABASE:-seedao_api}"
PGSUPERUSER="${PGSUPERUSER:-postgres}"
PGSUPERPASSWORD="${PGSUPERPASSWORD:-}"
PSQL_BIN="${PSQL_BIN:-}"

usage() {
  cat <<'EOF'
Usage:
  ./scripts/setup-db.sh [options]

Options (env vars):
  REGISTRY_IMPORT_SOURCE   Path to sns_chain_registry import SQL
  PGSUPERPASSWORD          postgres superuser password (first-time setup)
  PGHOST PGPORT PGUSER PGPASSWORD PGDATABASE
  PSQL_BIN                 Full path to psql (optional)

Example:
  PGSUPERPASSWORD='your-postgres-password' \
  REGISTRY_IMPORT_SOURCE='/d/path/sns_chain_registry_import.sql' \
  ./scripts/setup-db.sh
EOF
}

find_psql() {
  if [[ -n "$PSQL_BIN" && -x "$PSQL_BIN" ]]; then
    echo "$PSQL_BIN"
    return
  fi

  if command -v psql >/dev/null 2>&1; then
    command -v psql
    return
  fi

  local candidates=(
    "/c/Program Files/PostgreSQL/16/bin/psql.exe"
    "/c/Program Files/PostgreSQL/17/bin/psql.exe"
    "/c/Program Files/PostgreSQL/15/bin/psql.exe"
  )

  for candidate in "${candidates[@]}"; do
    if [[ -x "$candidate" ]]; then
      echo "$candidate"
      return
    fi
  done

  echo "psql not found. Add PostgreSQL bin to PATH or set PSQL_BIN." >&2
  exit 1
}

run_psql() {
  local db="$1"
  local pg_user="$2"
  local pg_password="$3"
  shift 3

  PGPASSWORD="$pg_password" "$PSQL" \
    -h "$PGHOST" \
    -p "$PGPORT" \
    -U "$pg_user" \
    -d "$db" \
    -v ON_ERROR_STOP=1 \
    "$@"
}

if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
  usage
  exit 0
fi

PSQL="$(find_psql)"

echo "==> Using psql: $PSQL"

echo "==> Step 1: ensure role and database"
if PGPASSWORD="$PGPASSWORD" run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -c "SELECT 1;" >/dev/null 2>&1; then
  echo "Database $PGDATABASE is reachable as $PGUSER."
else
  echo "Cannot connect as $PGUSER, trying superuser $PGSUPERUSER ..."
  if [[ -z "$PGSUPERPASSWORD" ]]; then
    echo "Set PGSUPERPASSWORD to your postgres password for first-time setup." >&2
    exit 1
  fi
  run_psql "postgres" "$PGSUPERUSER" "$PGSUPERPASSWORD" -f "$SCRIPT_DIR/init-role.sql"
fi

echo "==> Step 2: run schema migrations"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/001_init.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/002_sns_names_seedao_suffix.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/003_user_membership.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/004_payment_orders.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/005_coupons.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/006_course_payments.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/007_proposals.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/008_proposal_sip.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/009_proposal_sip_alloc.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/010_user_oauth_bindings.sql"
run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$ROOT_DIR/migrations/011_legacy_import.sql"

echo "==> Step 3: import sns_chain_registry"
PREPARED="$ROOT_DIR/migrations/data/sns_chain_registry_import.sql"
if [[ -n "$REGISTRY_IMPORT_SOURCE" ]]; then
  SOURCE_ABS="$(cd "$(dirname "$REGISTRY_IMPORT_SOURCE")" && pwd)/$(basename "$REGISTRY_IMPORT_SOURCE")"
  PREPARED_ABS="$(cd "$(dirname "$PREPARED")" && pwd)/$(basename "$PREPARED")"
  if [[ "$SOURCE_ABS" != "$PREPARED_ABS" ]]; then
    bash "$SCRIPT_DIR/prepare-registry-import.sh" "$REGISTRY_IMPORT_SOURCE" "$PREPARED"
  else
    echo "Using existing import file: $PREPARED"
  fi
fi

if [[ ! -f "$PREPARED" ]]; then
  echo "WARN: No registry import file at $PREPARED"
  echo "Run: ./scripts/prepare-registry-import.sh <your.sql>"
else
  run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -f "$PREPARED"
  COUNT="$(run_psql "$PGDATABASE" "$PGUSER" "$PGPASSWORD" -tAc "SELECT COUNT(*) FROM sns_chain_registry;")"
  echo "Imported sns_chain_registry rows: $(echo "$COUNT" | tr -d '[:space:]')"
fi

echo
echo "Database setup complete."
echo "DSN: postgres://${PGUSER}:${PGPASSWORD}@${PGHOST}:${PGPORT}/${PGDATABASE}?sslmode=disable"
