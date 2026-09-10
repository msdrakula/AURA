#!/usr/bin/env python3
"""Rasterize the AURA mark to PNG without extra packages."""
from __future__ import annotations

import struct
import zlib
from pathlib import Path

TEAL = (0x3D, 0xBA, 0x9A, 255)
INK = (0x06, 0x20, 0x18, 255)


def point_in_poly(x: float, y: float, poly: list[tuple[float, float]]) -> bool:
    inside = False
    n = len(poly)
    j = n - 1
    for i in range(n):
        xi, yi = poly[i]
        xj, yj = poly[j]
        if ((yi > y) != (yj > y)) and (x < (xj - xi) * (y - yi) / (yj - yi + 0.0) + xi):
            inside = not inside
        j = i
    return inside


def rounded_rect(px: float, py: float, size: int, radius: float) -> bool:
    r = radius
    if px < 0 or py < 0 or px >= size or py >= size:
        return False
    if r <= px <= size - 1 - r or r <= py <= size - 1 - r:
        return True
    corners = (
        (r, r),
        (size - 1 - r, r),
        (r, size - 1 - r),
        (size - 1 - r, size - 1 - r),
    )
    for cx, cy in corners:
        if (px - cx) ** 2 + (py - cy) ** 2 <= r * r:
            return True
    return False


def scale_poly(points: list[tuple[float, float]], size: int) -> list[tuple[float, float]]:
    return [(x / 64.0 * size, y / 64.0 * size) for x, y in points]


def png(size: int, rgba: bytes) -> bytes:
    raw = b"".join(b"\x00" + rgba[i * size * 4 : (i + 1) * size * 4] for i in range(size))

    def chunk(tag: bytes, data: bytes) -> bytes:
        return struct.pack(">I", len(data)) + tag + data + struct.pack(">I", zlib.crc32(tag + data) & 0xFFFFFFFF)

    ihdr = struct.pack(">IIBBBBB", size, size, 8, 6, 0, 0, 0)
    return b"\x89PNG\r\n\x1a\n" + chunk(b"IHDR", ihdr) + chunk(b"IDAT", zlib.compress(raw, 9)) + chunk(b"IEND", b"")


def render(size: int) -> bytes:
    outer = scale_poly([(32, 11), (50, 53), (40.8, 53), (38.2, 46.6), (25.8, 46.6), (23.2, 53), (14, 53)], size)
    hole = scale_poly([(32, 25.2), (27.6, 35.8), (36.4, 35.8)], size)
    radius = size * 14 / 64.0
    out = bytearray(size * size * 4)
    for y in range(size):
        py = y + 0.5
        for x in range(size):
            px = x + 0.5
            i = (y * size + x) * 4
            if not rounded_rect(px, py, size, radius):
                continue
            color = TEAL
            if point_in_poly(px, py, outer) and not point_in_poly(px, py, hole):
                color = INK
            out[i : i + 4] = bytes(color)
    return png(size, bytes(out))


def main() -> None:
    root = Path(__file__).resolve().parents[2]
    static = root / "web" / "static"
    (static / "aura-icon.png").write_bytes(render(256))
    (static / "aura-icon-48.png").write_bytes(render(48))


if __name__ == "__main__":
    main()
