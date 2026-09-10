#!/bin/sh
# Install the AURA icon into the user hicolor theme (used by the XFCE panel).
set -e
ROOT="$(CDPATH= cd -- "$(dirname "$0")/../.." && pwd)"
DEST="${HOME}/.local/share/icons/hicolor"
mkdir -p "$DEST/scalable/apps" "$DEST/256x256/apps" "$DEST/48x48/apps"
cp "$ROOT/web/static/aura-icon.svg" "$DEST/scalable/apps/aura.svg"
if [ -f "$ROOT/web/static/aura-icon.png" ]; then
	cp "$ROOT/web/static/aura-icon.png" "$DEST/256x256/apps/aura.png"
fi
if [ -f "$ROOT/web/static/aura-icon-48.png" ]; then
	cp "$ROOT/web/static/aura-icon-48.png" "$DEST/48x48/apps/aura.png"
elif [ -f "$ROOT/web/static/aura-icon.png" ]; then
	cp "$ROOT/web/static/aura-icon.png" "$DEST/48x48/apps/aura.png"
fi
gtk-update-icon-cache -f "$DEST" >/dev/null 2>&1 || true
echo "installed icons in $DEST"
