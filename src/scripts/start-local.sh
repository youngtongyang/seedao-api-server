#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SRC_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
API_PORT="${API_PORT:-8080}"

kill_port_listeners() {
  local port="$1"
  local pids=""

  echo "==> Checking port $port ..."
  if [[ "$(uname -s)" =~ ^(MINGW|MSYS|CYGWIN) ]]; then
    pids="$(netstat -ano 2>/dev/null | grep LISTENING | grep ":${port}" | awk '{print $NF}' | sort -u || true)"
    for pid in $pids; do
      if [[ "$pid" =~ ^[0-9]+$ ]] && [[ "$pid" -gt 0 ]]; then
        echo "Killing process on port $port (PID $pid) ..."
        taskkill //PID "$pid" //F >/dev/null 2>&1 || true
      fi
    done
  elif command -v lsof >/dev/null 2>&1; then
    pids="$(lsof -ti ":${port}" 2>/dev/null || true)"
    for pid in $pids; do
      echo "Killing process on port $port (PID $pid) ..."
      kill -9 "$pid" 2>/dev/null || true
    done
  elif command -v fuser >/dev/null 2>&1; then
    echo "Killing process on port $port ..."
    fuser -k "${port}/tcp" 2>/dev/null || true
  else
    echo "WARN: cannot auto-kill port $port on this platform"
    return 0
  fi

  if [[ "$(uname -s)" =~ ^(MINGW|MSYS|CYGWIN) ]]; then
    if netstat -ano 2>/dev/null | grep LISTENING | grep -q ":${port}"; then
      echo "WARN: port $port is still in use. Stop the process manually and retry." >&2
      exit 1
    fi
  elif command -v lsof >/dev/null 2>&1; then
    if lsof -ti ":${port}" >/dev/null 2>&1; then
      echo "WARN: port $port is still in use. Stop the process manually and retry." >&2
      exit 1
    fi
  fi
  echo "Port $port is free."
}

# 导入文件默认在 migrations/data/，setup-db.sh 会自动使用，无需再设 REGISTRY_IMPORT_SOURCE
echo "==> Step 1: init database + import registry"
bash "$SCRIPT_DIR/setup-db.sh"

echo "==> Step 2: prepare config"
CONFIG_FILE="$SRC_ROOT/config/config.local.yaml"
if [[ ! -f "$CONFIG_FILE" ]]; then
  cp "$SRC_ROOT/config/config.local.example.yaml" "$CONFIG_FILE"
  echo "Created $CONFIG_FILE from example."
fi

cd "$SRC_ROOT"
export CONFIG_PATH="$CONFIG_FILE"

kill_port_listeners "$API_PORT"

echo "==> Step 3: start API server"
if ! command -v go >/dev/null 2>&1; then
  echo "Go not installed. Run: cd src && go run ./cmd/server"
  exit 0
fi

go mod tidy
go run ./cmd/server
