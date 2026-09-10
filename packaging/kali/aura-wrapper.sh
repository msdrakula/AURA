#!/bin/sh
# Installed as ~/.local/bin/aura
ROOT="@AURA_ROOT@"
if [ ! -x "$ROOT/aura" ]; then
	echo "AURA binary not found: $ROOT/aura" >&2
	echo "Build it with: (cd \"$ROOT\" && go build -o aura ./cmd/server)" >&2
	exit 1
fi
export AURA_ROOT="$ROOT"
cd "$ROOT" || exit 1
exec "$ROOT/aura" "$@"
