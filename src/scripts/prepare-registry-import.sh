#!/usr/bin/env bash
set -euo pipefail

SOURCE_FILE="${1:-}"
OUTPUT_FILE="${2:-}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

if [[ -z "$OUTPUT_FILE" ]]; then
  OUTPUT_FILE="$ROOT_DIR/migrations/data/sns_chain_registry_import.sql"
fi

if [[ -z "$SOURCE_FILE" ]]; then
  echo "Usage: $0 <source.sql> [output.sql]" >&2
  exit 1
fi

if [[ ! -f "$SOURCE_FILE" ]]; then
  echo "Source file not found: $SOURCE_FILE" >&2
  exit 1
fi

mkdir -p "$(dirname "$OUTPUT_FILE")"

SOURCE_ABS="$(cd "$(dirname "$SOURCE_FILE")" && pwd)/$(basename "$SOURCE_FILE")"
OUTPUT_ABS="$(cd "$(dirname "$OUTPUT_FILE")" && pwd)/$(basename "$OUTPUT_FILE")"
if [[ "$SOURCE_ABS" == "$OUTPUT_ABS" ]]; then
  echo "Import file already in place: $OUTPUT_ABS"
  exit 0
fi

cp -f "$SOURCE_FILE" "$OUTPUT_FILE"

echo "Prepared import file: $OUTPUT_FILE"
echo "Kept original sns_name values (including .seedao suffix)."
