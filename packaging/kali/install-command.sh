#!/bin/sh
# Install `aura` on PATH and a Kali application menu entry.
set -e
ROOT="$(CDPATH= cd -- "$(dirname "$0")/../.." && pwd)"
cd "$ROOT"

echo "building $ROOT/aura ..."
CGO_ENABLED=1 go build -o "$ROOT/aura" ./cmd/server

WRAP="$(mktemp)"
sed "s|@AURA_ROOT@|$ROOT|g" "$ROOT/packaging/kali/aura-wrapper.sh" > "$WRAP"
chmod 755 "$WRAP"

TARGET="${HOME}/.local/bin/aura"
mkdir -p "$(dirname "$TARGET")"
cp "$WRAP" "$TARGET"
chmod 755 "$TARGET"
rm -f "$WRAP"

mkdir -p "$HOME/.local/share/applications"
cp "$ROOT/packaging/kali/aura.desktop" "$HOME/.local/share/applications/aura.desktop"
sh "$ROOT/packaging/kali/install-icons.sh"
# Old command name still launches AURA.
cat > "${HOME}/.local/bin/meb" << EOF
#!/bin/sh
exec "\$HOME/.local/bin/aura" "\$@"
EOF
chmod 755 "${HOME}/.local/bin/meb"
rm -f "$HOME/.local/share/applications/meb.desktop"

echo "installed: $TARGET"
echo "type:  aura"
