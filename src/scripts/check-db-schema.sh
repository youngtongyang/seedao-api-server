#!/usr/bin/env bash
# Read-only remote/local DB schema check.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

PGHOST="${PGHOST:-127.0.0.1}"
PGPORT="${PGPORT:-5432}"
PGUSER="${PGUSER:-seedao}"
PGPASSWORD="${PGPASSWORD:-}"
PGDATABASE="${PGDATABASE:-seedao_api}"

if [[ -z "${PGPASSWORD}" ]]; then
  echo "Set PGPASSWORD (and PGHOST/PGUSER/PGDATABASE for remote)." >&2
  exit 1
fi

export PGPASSWORD

echo "Checking ${PGUSER}@${PGHOST}:${PGPORT}/${PGDATABASE} ..."
psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" \
  -v ON_ERROR_STOP=1 \
  -f "$SCRIPT_DIR/check-db-schema.sql"
