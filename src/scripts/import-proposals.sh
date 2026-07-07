#!/usr/bin/env bash
# Import legacy proposal_v2 pg_dump and normalize table names (strip _v2).
#
# Usage:
#   PGPASSWORD=seedao ./src/scripts/import-proposals.sh "/path/to/proposal_v2_tables.sql"

set -euo pipefail

DUMP="${1:-}"
PGHOST="${PGHOST:-127.0.0.1}"
PGUSER="${PGUSER:-seedao}"
PGDATABASE="${PGDATABASE:-seedao_api}"

if [[ -z "$DUMP" || ! -f "$DUMP" ]]; then
  echo "usage: $0 /path/to/proposal_v2_tables.sql" >&2
  exit 1
fi

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"

echo "==> import legacy dump (creates proposal_*_v2 tables)"
PGPASSWORD="${PGPASSWORD:?set PGPASSWORD}" psql -h "$PGHOST" -U "$PGUSER" -d "$PGDATABASE" -f "$DUMP"

echo "==> normalize schema (rename *_v2, add indexes)"
PGPASSWORD="$PGPASSWORD" psql -h "$PGHOST" -U "$PGUSER" -d "$PGDATABASE" \
  -f "$ROOT/src/scripts/proposal-post-import.sql"

echo "done."
