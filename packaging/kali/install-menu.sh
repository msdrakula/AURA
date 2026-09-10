#!/bin/sh
set -e
ROOT="$(CDPATH= cd -- "$(dirname "$0")/../.." && pwd)"
BIN="$ROOT/aura"
if [ ! -x "$BIN" ]; then
  echo "build first: go build -o aura ./cmd/server" >&2
  exit 1
fi
mkdir -p "$HOME/.local/share/applications"
sed "s|^Exec=aura|Exec=$BIN|" "$ROOT/packaging/kali/aura.desktop" > "$HOME/.local/share/applications/aura.desktop"
sh "$ROOT/packaging/kali/install-icons.sh"
echo "installed $HOME/.local/share/applications/aura.desktop"
